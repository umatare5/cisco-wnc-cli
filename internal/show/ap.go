package show

import (
	"context"
	"strings"
	"time"

	"github.com/umatare5/cisco-wnc-cli/internal/config"
	"github.com/umatare5/cisco-wnc-cli/internal/render"
	"github.com/umatare5/cisco-wnc-cli/internal/wnc"
)

// APRow is one row of show ap.
type APRow struct {
	APName       *string  `json:"ap_name,omitzero"`
	Model        *string  `json:"model,omitzero"`
	Serial       *string  `json:"serial,omitzero"`
	EthernetMAC  *string  `json:"ethernet_mac,omitzero"`
	RadioMAC     *string  `json:"radio_mac,omitzero"`
	IPAddress    *string  `json:"ip_address,omitzero"`
	SWVersion    *string  `json:"version,omitzero"`
	Slots        *uint8   `json:"slots,omitzero"`
	Country      *string  `json:"country,omitzero"`
	Mode         *string  `json:"mode,omitzero"`
	Admin        *string  `json:"admin_state,omitzero"`
	State        *string  `json:"state,omitzero"`
	LLDPNeighbor *string  `json:"lldp_neighbor,omitzero"`
	Longitude    *float64 `json:"longitude,omitzero"`
	Latitude     *float64 `json:"latitude,omitzero"`
	Height       *int16   `json:"height,omitzero"`
	Floor        *int     `json:"floor,omitzero"`
	PowerType    *string  `json:"power_type,omitzero"`
	PowerMode    *string  `json:"power_mode,omitzero"`
	Uptime       *int64   `json:"uptime_seconds,omitzero"`
	AssocUptime  *int64   `json:"assoc_uptime_seconds,omitzero"`
	Controller   string   `json:"controller"`
}

// APColumns describes the access point view. Uptime is the access point's own age from boot-time
// and Assoc Uptime is the age of the current CAPWAP association from join-time. A controller
// switchover renews only the second, so one column for both would report it as a reboot of
// every access point.
func APColumns() []render.Column[APRow] {
	return []render.Column[APRow]{
		{Key: keyAPName, Header: headAPName, Cell: func(r APRow) string { return render.StrPtr(r.APName) }},
		{Key: "model", Header: "Model", Cell: func(r APRow) string { return render.StrPtr(r.Model) }},
		{
			Key: "serial", Header: "Serial", Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.Serial) },
		},
		{
			Key: keyEthernetMAC, Header: headEthernetMAC, Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.EthernetMAC) },
		},
		{
			Key: keyRadioMAC, Header: headRadioMAC, Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.RadioMAC) },
		},
		{
			Key: keyIPAddress, Header: headIPAddress, Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.IPAddress) },
		},
		{Key: "version", Header: "Version", Cell: func(r APRow) string { return render.StrPtr(r.SWVersion) }},
		{
			Key: "slots", Header: "Slots", Hidden: true,
			Cell: func(r APRow) string { return render.IntPtr(r.Slots) },
			Sort: func(r APRow) any { return render.SortValue(r.Slots) },
		},
		{
			Key: "country", Header: "Country", Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.Country) },
		},
		{Key: keyMode, Header: "Mode", Cell: func(r APRow) string { return render.StrPtr(r.Mode) }},
		{
			Key: keyAdminState, Header: headAdminState,
			Cell:   func(r APRow) string { return render.StrPtr(r.Admin) },
			Pretty: func(r APRow) string { return prettyState(r.Admin, dispEnabled, dispDisabled, glyphOff) },
		},
		{
			Key: keyState, Header: "State",
			Cell:   func(r APRow) string { return render.StrPtr(r.State) },
			Pretty: func(r APRow) string { return prettyOtherwise(r.State, dispRegistered, glyphWarn) },
		},
		{
			Key: "lldp_neighbor", Header: "LLDP Neighbor", Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.LLDPNeighbor) },
		},
		{
			Key: "longitude", Header: "Longitude", Hidden: true,
			Cell: func(r APRow) string { return render.FloatPtr(r.Longitude) },
			Sort: func(r APRow) any { return render.SortValue(r.Longitude) },
		},
		{
			Key: "latitude", Header: "Latitude", Hidden: true,
			Cell: func(r APRow) string { return render.FloatPtr(r.Latitude) },
			Sort: func(r APRow) any { return render.SortValue(r.Latitude) },
		},
		{
			Key: "height", Header: "Height", Hidden: true,
			Cell: func(r APRow) string { return render.UnitPtr(r.Height, "m") },
			Sort: func(r APRow) any { return render.SortValue(r.Height) },
		},
		{
			Key: "floor", Header: "Floor", Hidden: true,
			Cell: func(r APRow) string { return render.IntPtr(r.Floor) },
			Sort: func(r APRow) any { return render.SortValue(r.Floor) },
		},
		{
			Key: "power_type", Header: "Power Type", Hidden: true,
			Cell: func(r APRow) string { return render.StrPtr(r.PowerType) },
		},
		{Key: "power_mode", Header: "Power Mode", Cell: func(r APRow) string { return render.StrPtr(r.PowerMode) }},
		{
			Key: "uptime_seconds", Header: "Uptime",
			Cell: func(r APRow) string { return render.Duration(r.Uptime) },
			Sort: func(r APRow) any { return render.SortValue(r.Uptime) },
		},
		{
			Key: "assoc_uptime_seconds", Header: "Assoc Uptime",
			Cell: func(r APRow) string { return render.Duration(r.AssocUptime) },
			Sort: func(r APRow) any { return render.SortValue(r.AssocUptime) },
		},
		{Key: keyController, Header: headController, Cell: func(r APRow) string { return render.Str(r.Controller) }},
	}
}

func FetchAPs(ctx context.Context, c *wnc.Client, t config.Target, rep *Reporter) ([]APRow, error) {
	aps, reads, err := c.APs(ctx)
	if err != nil {
		return nil, err
	}

	rep.Degraded("oper-data", reads.Power)
	rep.Degraded("lldp-neigh", reads.LLDP)
	rep.Degraded("ap-geo-loc-data", reads.Geolocation)

	return apRows(aps, t), nil
}

func apRows(aps []wnc.AP, t config.Target) []APRow {
	now := time.Now()
	rows := make([]APRow, 0, len(aps))

	for _, ap := range aps {
		rows = append(rows, APRow{
			APName:      optional(ap.Name),
			Model:       optional(ap.Model),
			Serial:      optional(ap.Serial),
			EthernetMAC: optional(ap.EthernetMAC),
			RadioMAC:    optional(ap.WtpMAC),
			IPAddress:   optional(ap.IPAddr),
			SWVersion:   optional(ap.SWVersion),
			// No joined access point has zero radios, so zero is an omitted leaf. The
			// count comes from num-radio-slots and not from the static-info slot count,
			// which also counts a remote-LAN port and reads one high on such a model.
			Slots: zeroAbsent(ap.Slots),
			// The controller pads this to a fixed width, so "J4 " and "J4" are the same
			// regulatory domain.
			Country:      optional(strings.TrimSpace(ap.Country)),
			Mode:         optional(showAPModeWithSub(ap.Mode, ap.SubMode)),
			Admin:        optional(showAPAdmin(ap.AdminState)),
			State:        optional(showAPState(ap.OperState)),
			LLDPNeighbor: optional(strings.Join(ap.Neighbors, ", ")),
			Longitude:    ap.Longitude,
			Latitude:     ap.Latitude,
			Height:       ap.Height,
			Floor:        ap.Floor,
			PowerType:    optional(showPowerType(ap.PowerType)),
			PowerMode:    optional(showPowerMode(ap.PowerMode)),
			Uptime:       render.SecondsSince(now, ap.BootTime),
			AssocUptime:  render.SecondsSince(now, ap.JoinTime),
			Controller:   t.Name,
		})
	}

	return rows
}
