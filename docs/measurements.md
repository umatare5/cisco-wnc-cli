# Measurements

This document records every reading taken on a live Catalyst 9800, each one only once.

`All` is `17.12.8`, `17.15.6` and `17.18.4a` – `Unrecorded` is a reading whose release was not recorded.
`Not release-bound` is a reading that no controller release changes.

## Device Heading Map

Columns whose device heading differs, with the command that prints them.

| Column                           | Device heading                                       | Device command                                         |
| :------------------------------- | :--------------------------------------------------- | :----------------------------------------------------- |
| `admin_state`                    | `Administrative State`                               | `show ap config general`                               |
| `ap_join_profile`                | `AP Profile`                                         | `show wireless tag site detailed`                      |
| `assoc_uptime_seconds`           | `Association Up Time`                                | `show ap uptime`                                       |
| `band` and `protocol`            | `11n(2.4)`, the pair in one field                    | `show wireless client summary`                         |
| `channel`                        | `(64,60)`, a pair per radio                          | `show ap dot11 5ghz summary`                           |
| `connected_seconds`              | `Connected For`, in seconds                          | `show wireless client mac-address ... detail`          |
| `device`                         | `Device Name`                                        | `show wireless client mac-address ... detail`          |
| `last_config_failure`            | `Reason for last unsuccessful configuration attempt` | `show wireless stats ap mac-address ... join detailed` |
| `last_discovery_failure`         | `Reason for last unsuccessful discovery attempt`     | `show wireless stats ap mac-address ... join detailed` |
| `last_join_failure`              | `Reason for last unsuccessful join attempt`          | `show wireless stats ap mac-address ... join detailed` |
| `last_reboot_reason`             | `Reboot reason from AP`                              | `show wireless stats ap mac-address ... join detailed` |
| `misconfiguration_reason`        | none at all on the device                            | `show ap tag summary`                                  |
| `mode`                           | `AP Mode : FlexConnect`                              | `show ap ... config general`                           |
| `p2p_blocking`                   | `Peer-to-Peer Blocking Action`                       | `show wlan id`                                         |
| `power_type`                     | `PoE`, both mechanisms collapsed                     | `show ap config general`                               |
| `profile_24ghz`                  | `2.4ghz RF Policy`                                   | `show wireless tag rf detailed`                        |
| `profile_5ghz`                   | `5ghz RF Policy`                                     | `show wireless tag rf detailed`                        |
| `profile_6ghz`                   | `6ghz RF Policy`                                     | `show wireless tag rf detailed`                        |
| `radio_mac`                      | `Base MAC`                                           | `show wireless stats ap join summary`                  |
| `radio_mac`                      | `AP Mac`                                             | `show ap tag summary`                                  |
| `radio_mac`                      | `Mac Address`                                        | `show ap dot11 5ghz summary`                           |
| `rate`                           | `Current Rate`, an MCS index and stream count        | `show wireless client mac-address ... detail`          |
| `rx_bytes`                       | `Number of Bytes Received from Client`               | `show wireless client mac-address ... detail`          |
| `tx_bytes`                       | `Number of Bytes Sent to Client`                     | `show wireless client mac-address ... detail`          |
| `txpower`                        | `Txpwr`, level and dBm in one field                  | `show ap dot11 5ghz summary`                           |
| `uptime_seconds`                 | `AP Up Time`                                         | `show ap uptime`                                       |
| `version`                        | `Software Version`                                   | `show ap config general`                               |
| `wlan_profile`, `policy_profile` | `WLAN Profile Name`, `Policy Name`                   | `show wireless tag policy detailed`                    |

## Access Point

The access-point collections and the RPCs that name one.

| Fact                                                                | Release           | Condition                                          |
| :------------------------------------------------------------------ | :---------------- | :------------------------------------------------- |
| `ap-name-mac-map=<name>` returns name, base radio and Ethernet MAC  | All               | `404` for an unheld name, a rowless `200` alike    |
| `ap-geo-loc-data` keys on the base radio MAC, not the Ethernet MAC  | 17.15.6           | 2 of 2 records, both set by hand                   |
| An invalid location still carries its above-ground height           | 17.15.6           | 1 record, its coordinates deleted by hand          |
| `ap-location/floor-id` holds the floor set by hand, `floor` reads 0 | 17.15.6           | 2 of 2 records, `floor` 0 on both                  |
| `fields=ap-location/floor-id` prunes to that one leaf               | 17.15.6           | `ap-location(floor-id)` read the same              |
| The keyed read `404`s a long, spaced, slashed or multi-byte name    | 17.18.4a          | Four spellings probed, one 256 characters long     |
| `reset ap`: out of `capwap-data` in 16s, rejoin by 285s             | 17.12.8, 17.18.4a | AIR-AP1815I, clients down throughout               |
| `Not Joined` with `Wtp reset config cmd sent`, then reboot-cmd      | 17.12.8           | Through a `reset ap`, in `wnc show ap-join` alone  |
| Name-arm `reset ap`: new boot time, no other AP moved               | 17.18.4a          | Through the `ap-name` arm, not the address arm     |
| `reset capwap`: rejoin in 10s, uptime held, assoc restarted         | 17.12.8, 17.18.4a | `reset ap` restarts both quantities                |
| `num-join-req-recvd` and its siblings accumulate across CAPWAP      | 17.12.8, 17.18.4a | On `ap-join-stats`, so a reset shows as a delta    |
| An AP-level disable takes `ap-admin-state` to `adminstate-disabled` | 17.15.6, 17.18.4a | Radios held `enabled`, `oper-state` `radio-down`   |
| A disabled AP stays `Registered`, `uptime_seconds` climbing         | 17.15.6, 17.18.4a | No reboot, the difference from a `reset ap`        |
| `set-ap-admin-state`: the name arm read as the address arm did      | 17.18.4a, 17.15.6 | Name arm at 17.18.4a, address arm at 17.15.6       |
| An unjoined AP leaves `capwap-data` and is still counted            | Unrecorded        | `show ap summary` read 2, the join summary 3       |
| Six of eleven join-record instants read `1970-01-01T00:00:00+00:00` | All               | Every lab record, the sentinel for no event        |
| Nineteen join-record counters, with no clear time declared          | All               | On `ap-join-stats`, read by `wnc show ap-join`     |
| Join, config, discovery and reboot domains: 42, 14, 17 and 59       | All               | No display string for any, where phases print      |
| The disconnect reason enum read `unkown` on 5 of 7 records          | Unrecorded        | Controller's spelling, so the column is free text  |
| A second slot count includes the remote-LAN port, one radio high    | Unrecorded        | Only on a model carrying one, in `show ap summary` |

## Radio

`radio-oper-data`, the band number the slot RPC takes, and the RRM measurement beside it.

| Fact                                                                    | Release           | Condition                                          |
| :---------------------------------------------------------------------- | :---------------- | :------------------------------------------------- |
| The band number follows `radio-type`, not the band served               | All               | The served band's number `400`s, dual band takes 3 |
| XOR slot 2 on 6 GHz, band 4: "does not have a dedicated radio"          | 17.15.6           | `TEST-AP03`, reporting `dot11-6-ghz-band`          |
| Band 3 on that radio `204`s and takes slot 2 down, 0 and 1 up           | 17.15.6           | The other direction of the pair, one radio         |
| Band 3 in slot 0: "AP does not support the specified radio type"        | Unrecorded        | A dedicated 2.4 GHz layout, which takes band 1     |
| Band 3 in slot 0 is accepted where band 1, the band served, `400`s      | 17.12.8           | A 2.4-or-5 GHz XOR layout, the other AP            |
| `must` accepts band 1 on slot 0 and no other slot                       | All               | A dedicated 2.4 GHz radio, `radio-80211bg`         |
| `must` accepts band 2 on slot 1 or slot 2                               | All               | A dedicated 5 GHz radio, `radio-80211a`            |
| `must` accepts band 3 on slot 0 or slot 2                               | All               | `radio-80211abgn` and `radio-80211-xor-5-6ghz`     |
| `must` accepts band 4 on slot 2 or slot 3                               | All               | `radio-80211-6ghz`, and never once written         |
| Band 1 with slot 1 `400`s, the controller naming that `must` clause     | 17.15.6           | The one forbidden pair probed on the wire          |
| `--slot 3` is refused only after the controller answers, so exit 1      | 17.12.8, 17.15.6  | An AP holding slots 0 and 1, so never exit 2       |
| `enm-radio-type`: eight members at 17.12.8 and 17.15.6, nine later      | All               | The ninth, `radio-80211-xor-24-6ghz`, fits 3 and 4 |
| `radio-80211-xor-24-6ghz`, `radio-uwb` and `radio-invalid` take no band | All               | No leaf says whether such a radio wants 3 or 4     |
| Radio disable then enable: slot 1 `Enabled/Up` to `Disabled/Down`       | 17.15.6, 17.18.4a | Slot 0 untouched, and `--slot 0` mirrored it       |
| An XOR radio sends a power entry per band, the first 22 dBm             | Unrecorded        | The device summary says 18 on the selected band    |
| `curr-freq` is guarded on radio mode, absent on Monitor and Sniffer     | All               | Width and power unguarded, device prints all `N/A` |
| A remote-LAN port has no mode, band, state, channel or power            | All               | Listed among the radios, so the slot can be named  |
| The RRM measurement list is shorter than the radio list                 | Unrecorded        | So a radio can hold no channel-utilization row     |
| `cca-util-percentage` is a `uint16` with no `range` statement declared  | All               | 100 is no ceiling, so spare is not 100 minus it    |
| `units "dBm"` on transmit power, `units "percentage"` on utilization    | All               | Width a bare `uint8`, channel and Hz no unit       |

## Client

The client collections and `apf-ms-delete-all`, whose scope no schema states.

| Fact                                                                 | Release           | Condition                                            |
| :------------------------------------------------------------------- | :---------------- | :--------------------------------------------------- |
| Every `client-mac` a controller serves is already lowercase          | 17.12.8, 17.15.6  | The form the SDK normalizes to, so none is folded    |
| The client RPC module at `2023-03-01`: two `clear-sisf-binding`      | 17.12.8           | Read 2026-08-29, with no client delete there         |
| The same module at `2024-03-01` adds `apf-ms-delete-all`             | 17.15.6, 17.18.4a | Read 2026-08-29, so `deauth` needs 17.15.6           |
| Its description reads "... client MAC or IP address or username"     | Unrecorded        | The name reads as a purge its model contradicts      |
| A post to the release without it `400`s `malformed-message`          | 17.12.8           | This CLI re-words that status rather than passing it |
| `204` for a target with no session as for one it dropped             | 17.18.4a          | Both under 330ms, so neither confirms a client       |
| `--mac` post: `connected_seconds` 6102 to 13, `Run` to `IP Learning` | 17.18.4a          | Recovery to `ipv4` by 210s, three clients untouched  |
| Two `--mac` posts each dropped two of eighteen, not one              | 17.15.6           | A 25s no-post window moved none, so no accident      |
| The extra station shared the AP, BSSID, radio and WLAN               | 17.15.6           | Another vendor, and thirteen on that BSS stayed      |
| `--username` post: the row gone in 1.9s, deleted not reset in place  | 17.15.6           | Target stable 82 minutes at `connected_seconds` 4943 |
| The station re-associated as a new record in 1.5s, `Run` inside 42s  | 17.15.6           | So a later read sees a young association             |
| That post moved none of the other 17 clients on the controller       | 17.15.6           | Two 35s control windows moved 1 and 0 of eighteen    |
| Sixteen of eighteen clients carried an empty username                | 17.15.6           | So an empty `--username` selects nearly every one    |
| The RPC's `ip-addr` arm answers `204` and is not refused             | 17.15.6           | Every SISF binding measured carries zone 0           |
| The username key arrives with an empty value, not omitted            | Unrecorded        | So an empty username reads as unreported             |
| Up to eight IPv6 addresses per client, some compressed               | Unrecorded        | A textual compare differs between two polls          |
| Sibling instant leaves on the client read return the Unix epoch      | Unrecorded        | An age off one would read as fifty-six years         |
| The list carrying a client hostname answers with no content          | Unrecorded        | So the device-classification label is all there is   |

## Tag

The tag configuration lists and the operational container an AP resolves to.

| Fact                                                              | Release           | Condition                                                  |
| :---------------------------------------------------------------- | :---------------- | :--------------------------------------------------------- |
| 32 characters accepted, 33 answering `400 ... exceed 32`          | 17.12.8           | Per kind, the key leaves declaring a `pattern`             |
| No `leafref` and no `require-instance` in the tag modules         | All               | Names are plain strings, so a dangling ref persists        |
| A merge `PATCH` leaves an omitted leaf at the value held          | 17.12.8, 17.15.6  | A second write kept the description and other profile      |
| A flex profile without `is-local-site` false `400`s on `when`     | 17.12.8           | `when "../is-local-site = 'false'"`, defaulting to true    |
| `rf-tag-radio-profiles` with a null list `400`s `invalid value`   | 17.12.8, 17.15.6  | So the per-slot radio profile list is out of reach         |
| A plain read of `rf-tags` omits `default-rf-tag`'s profiles       | All               | `report-all` returns them, so no view reads plain          |
| An omitted description and RF profile name arrive absent          | Unrecorded        | The mirror of filter-name, which arrives empty             |
| `is-ap-misconfigured` arrives as explicit `false` on a healthy AP | 17.12.8           | `No` is a reading, not a substitute for silence            |
| `ap-misconfig` on 3 of 3 records, its domain going three to four  | 17.15.6, 17.18.4a | Undeclared at 17.12.8, with only `apmgr-no-misconfig` seen |
| The filter-name leaf arrives empty, on 3 of 3 records             | 17.12.8           | With no `ap-filter-configs` on that controller             |
| No resolved counterpart of the AP join or flex profile            | All               | Both columns come from the configured site tag             |

## Configuration

What a write reaches, what a save persists, and the dial timeout the SDK pins.

| Fact                                                                 | Release           | Condition                                            |
| :------------------------------------------------------------------- | :---------------- | :--------------------------------------------------- |
| `writable-running` advertised, `:startup` and `:candidate` by none   | All               | So the startup configuration is unreachable          |
| One RF tag: running only, both after a save, startup after delete    | All               | So an unsaved delete restores it on a reload         |
| `cisco-ia:save-config` returns `Save running-config successful`      | All               | A 78-byte body, identical over six posts             |
| A save takes at most 3.7s against 0.13s for a container read         | All               | The one request a read-sized `--timeout` refuses     |
| `ap-cfg` holds no admin-state leaf, and the name filter finds none   | 17.15.6           | Per-AP configuration keys on the dotted MAC          |
| A plain read omits `wpa2-enabled` and `auth-key-mgmt-dot1x` where on | Unrecorded        | Present only where explicitly switched off           |
| A controller that rejects `?with-defaults=report-all` answers `400`  | Unrecorded        | So a configuration read cannot fall back to plain    |
| `report-all` also materializes the PSK and the WEP key material      | Unrecorded        | So the raw struct is an allow-list at decode         |
| The legacy WLAN band setting is obsolete and reads all bands         | All               | Gone at 17.18.4a, so `bands` reads the per-band list |
| The SDK pins the dialer at 30s, `--timeout 60s` giving up at 30      | Not release-bound | Against an unroutable address, so no route at all    |

## Not Measured

Known gaps, each with its reason and what an operator carries for it.

| Gap                                          | Why                                          | Consequence                                         |
| :------------------------------------------- | :------------------------------------------- | :-------------------------------------------------- |
| Deleting a tag an AP resolves to             | It would delete a tag in use in the lab      | The controller keeps a dangling reference           |
| A sea-level elevation on any access point    | The lab records above-ground heights only    | The height column reads the above-ground case alone |
| The floor of an access point never given one | Both lab access points carry a floor id      | `floor` prints what the leaf holds, 0 included      |
| A dedicated 6 GHz radio on any slot          | The lab holds no AP carrying one             | Band 4 rests on the clause, never on a write        |
| A username on more than one session          | No lab controller holds two under one name   | The prompt's count is the only warning given        |
| The 17.12.8 `deauth` `400` in this CLI       | Every client there carries an empty username | Pinned on fixtures under `internal/wnc` instead     |
| Client impact of a `reset capwap`            | The lab AP carried no clients at the time    | A control teardown is not a radio reset             |
| The `reset capwap` name arm on the wire      | No write through it recorded                 | So neither arm settles anything about the other     |
| Whether the `ip-addr` arm drops a client     | This CLI posts through no such arm           | `--mac` already selects a client by address         |
| Site-tag leaves this CLI never renders       | None verified: fabric pair, ARP, DHCP, load  | The view renders the profiles and the local flag    |
| An independent source for AP admin state     | `show running-config all` by name finds none | Read it back with `wnc show ap` and `show overview` |
| The age of the channel-utilization read      | No module declares a timestamp for it        | RRM's cycle is invisible, so stale looks fresh      |
