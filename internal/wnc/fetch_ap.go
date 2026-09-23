package wnc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// The fields expression names the nodes this view renders and no others. It keeps the
// certificate leaves, the external-module serial numbers and proxy-info on the controller, and
// an unpruned read returns proxy-info's username and password. The serial number survives
// under device-detail because the Serial column renders it.
const apViewFields = "name;wtp-mac;ip-addr;num-radio-slots;country-code;" +
	"device-detail;ap-mode-data;ap-state;ap-time-info"

// AP is one access point's identity, state, power, uplink neighbor and position. BootTime is the
// instant the access point itself came up, and JoinTime the instant the current CAPWAP
// association began. A view carrying only one of them reads a controller switchover as a reboot
// of every access point.
//
// A zero value means the controller reported no instant.
type AP struct {
	Name        string
	WtpMAC      string
	EthernetMAC string
	Model       string
	Serial      string
	IPAddr      string
	SWVersion   string
	Slots       uint8
	Country     string
	Mode        string
	SubMode     string
	AdminState  string
	OperState   string
	BootTime    time.Time
	JoinTime    time.Time
	PowerType   string
	PowerMode   string
	Neighbors   []string

	// Longitude and Latitude are WGS 84 degrees, both set or both nil.
	Longitude *float64
	Latitude  *float64
}

// APReads reports which secondary read failed.
type APReads struct {
	Power       error
	LLDP        error
	Geolocation error
}

// APs reads the access point view. The CAPWAP collection drives the rows, so its failure is
// returned as the error and costs the controller its rows. The power pair, the LLDP neighbors
// and the position are secondary, and their failures cost only those cells.
func (c *Client) APs(ctx context.Context) ([]AP, APReads, error) {
	resp, err := c.sdk.AP().ListCAPWAPData(ctx, sdk.WithFields(apViewFields))
	if err != nil {
		return nil, APReads{}, fmt.Errorf("reading capwap-data: %w", err)
	}

	if resp == nil {
		return nil, APReads{}, nil
	}

	power, powerErr := c.apPower(ctx)
	neighbors, lldpErr := c.apNeighbors(ctx)
	positions, geoErr := c.apPositions(ctx)
	reads := APReads{Power: powerErr, LLDP: lldpErr, Geolocation: geoErr}

	aps := make([]AP, 0, len(resp.CAPWAPData))

	for _, ap := range resp.CAPWAPData {
		static := ap.DeviceDetail.StaticInfo
		row := AP{
			Name:        ap.Name,
			WtpMAC:      ap.WtpMAC,
			EthernetMAC: static.BoardData.WtpEnetMAC,
			Model:       static.ApModels.Model,
			Serial:      static.BoardData.WtpSerialNum,
			IPAddr:      ap.IPAddr,
			SWVersion:   ap.DeviceDetail.WtpVersion.SwVersion,
			Slots:       ap.NumRadioSlots,
			Country:     ap.CountryCode,
			Mode:        string(ap.ApModeData.WtpMode),
			SubMode:     ap.ApModeData.ApSubMode,
			AdminState:  string(ap.ApState.ApAdminState),
			OperState:   ap.ApState.ApOperationState,
			BootTime:    parseInstant(ap.ApTimeInfo.BootTime),
			JoinTime:    parseInstant(ap.ApTimeInfo.JoinTime),
			Neighbors:   neighbors[ap.WtpMAC],
		}

		if p, ok := power[ap.WtpMAC]; ok {
			row.PowerType, row.PowerMode = p.kind, p.mode
		}

		if pos, ok := positions[ap.WtpMAC]; ok {
			row.Longitude, row.Latitude = ptrTo(pos.longitude), ptrTo(pos.latitude)
		}

		aps = append(aps, row)
	}

	return aps, reads, nil
}

type powerPair struct {
	kind string
	mode string
}

// apPower indexes the power pair by access point. ap-pow is a pointer, so an access point missing
// from the map is reported as unreported rather than as unpowered.
func (c *Client) apPower(ctx context.Context) (map[string]powerPair, error) {
	resp, err := c.sdk.AP().ListApOperData(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading oper-data: %w", err)
	}

	if resp == nil {
		return nil, nil
	}

	out := make(map[string]powerPair, len(resp.OperData))

	for _, op := range resp.OperData {
		if op.ApPow == nil {
			continue
		}

		out[op.WtpMAC] = powerPair{kind: op.ApPow.PowerType, mode: op.ApPow.PowerMode}
	}

	return out, nil
}

// apNeighbors indexes the LLDP neighbors by access point. The list is keyed on the access point
// and the neighbor together, so reading the whole list and grouping it keeps an access
// point's row from being duplicated or dropped.
// Keying the read by access point alone answers 404,
// because that is a partial list key.
func (c *Client) apNeighbors(ctx context.Context) (map[string][]string, error) {
	resp, err := c.sdk.AP().ListLldpNeigh(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading lldp-neigh: %w", err)
	}

	if resp == nil {
		return nil, nil
	}

	out := make(map[string][]string, len(resp.LldpNeigh))

	for _, n := range resp.LldpNeigh {
		if label := neighborLabel(n.SystemName, n.PortID); label != "" {
			out[n.WtpMAC] = append(out[n.WtpMAC], label)
		}
	}

	return out, nil
}

// Bounds of a WGS 84 position in degrees. Both leaves are decimal64 with no range statement, so a
// value past these is well formed on the wire.
const (
	longitudeBound = 180
	latitudeBound  = 90
)

type position struct {
	longitude float64
	latitude  float64
}

// apPositions indexes the position by access point. Measured on 17.15.6, ap-mac is the base radio
// address capwap-data keys on, and not the Ethernet address. A record yields a position only when
// both halves parse within their bounds, because one half alone names no place. The invalid case
// of the location choice carries no ellipse, so the nil chain covers it.
func (c *Client) apPositions(ctx context.Context) (map[string]position, error) {
	resp, err := c.sdk.Geolocation().ListAPGeolocationData(ctx)
	if err != nil {
		return nil, fmt.Errorf("reading ap-geo-loc-data: %w", err)
	}

	if resp == nil {
		return nil, nil
	}

	out := make(map[string]position, len(resp.ApGeoLocData))

	for _, g := range resp.ApGeoLocData {
		if g.Loc == nil || g.Loc.Ellipse == nil || g.Loc.Ellipse.Center == nil {
			continue
		}

		lon, lonOK := parseDegrees(g.Loc.Ellipse.Center.Longitude, longitudeBound)
		lat, latOK := parseDegrees(g.Loc.Ellipse.Center.Latitude, latitudeBound)

		if lonOK && latOK {
			out[g.ApMAC] = position{longitude: lon, latitude: lat}
		}
	}

	return out, nil
}

// parseDegrees reads one coordinate, which RFC 7951 writes as a JSON string because it is a
// decimal64. The range test is affirmative: ParseFloat accepts "NaN" and "Inf", and json/v2 fails
// the whole output on either.
func parseDegrees(leaf *string, bound float64) (float64, bool) {
	if leaf == nil {
		return 0, false
	}

	v, err := strconv.ParseFloat(*leaf, 64)
	if err == nil && v >= -bound && v <= bound {
		return v, true
	}

	return 0, false
}

// neighborLabel names one neighbor. The port is appended because a system name alone does not say
// which switchport the access point is on.
func neighborLabel(system, port string) string {
	switch {
	case system == "" && port == "":
		return ""
	case port == "":
		return system
	case system == "":
		return port
	default:
		return system + ":" + port
	}
}

// parseInstant decodes a controller timestamp. The fraction is variable width, five digits on one
// access point and six on another in the same read, and an unparseable value yields the zero time,
// which the renderer treats as unreported.
func parseInstant(s string) time.Time {
	if s == "" {
		return time.Time{}
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}
	}

	return t
}
