# Show Commands

These commands read a controller and print a table or JSON, and none of them changes anything on it.

## wnc show overview

One row per access point radio, with the RF settings and the load on it.

### Format

```bash
wnc show overview
```

### Expected result

```text
AP Name    AP MAC             Slot  Mode         Band  Admin    Oper  Channel  Width  TxPower  Clients    ChUtil  RF Profile         Controller
TEST-AP01  00:00:5e:00:53:01  0     FlexConnect  2.4   Enabled  Up    11ch     20MHz  20dBm    1clients   23%     test-rf-profile01  WNC1
TEST-AP01  00:00:5e:00:53:01  1     FlexConnect  5     Enabled  Up    64ch     40MHz  18dBm    1clients   1%      test-rf-profile03  WNC1
TEST-AP02  00:00:5e:00:53:02  0     FlexConnect  2.4   Enabled  Up    1ch      20MHz  19dBm    2clients   16%     test-rf-profile01  WNC1
TEST-AP02  00:00:5e:00:53:02  1     FlexConnect  5     Enabled  Up    48ch     40MHz  17dBm    0clients   1%      test-rf-profile04  WNC1
TEST-AP03  00:00:5e:00:53:03  0     FlexConnect  2.4   Enabled  Up    6ch      20MHz  22dBm    14clients  10%     test-rf-profile01  WNC1
TEST-AP03  00:00:5e:00:53:03  1     FlexConnect  5     Enabled  Up    116ch    40MHz  22dBm    2clients   2%      test-rf-profile03  WNC1
TEST-AP03  00:00:5e:00:53:03  2     FlexConnect  6     Enabled  Up    5ch      40MHz  18dBm    1clients   2%      test-rf-profile05  WNC1
```

### Use cases

<details><summary>Case 1: List the radios above 50% utilization, worst first</summary><p>

```bash
wnc show overview -f json | jq -r '.[] | select(.ch_util_percent > 50) | "\(.ap_name)/\(.slot) \(.ch_util_percent)%"'
```

</p></details>

<details><summary>Case 2: List the 6 GHz radios only</summary><p>

```bash
wnc show overview -r 6
```

</p></details>

<details><summary>Case 3: List the radios carrying the most clients</summary><p>

```bash
wnc show overview -b clients --sort-order desc
```

</p></details>

## wnc show ap

One row per access point: what it is, how it is powered, what it is plugged into and how long it has been up.

### Format

```bash
wnc show ap
```

### Expected result

```text
AP Name    Model             Serial       Ethernet MAC       Radio MAC          IP Address    SW Version  Slots  Country  Mode         Admin    State       LLDP Neighbor                     Power Type    Power Mode  Uptime  Assoc  Controller
TEST-AP01  AIR-AP1815I-Q-K9  TST0000AP01  00:00:5e:00:53:11  00:00:5e:00:53:01  192.168.0.11  17.12.7.13  2      J4       FlexConnect  Enabled  Registered  test-sw01.example.internal:Gi0/2  PoE+          Full Power  1d14h   1d14h  WNC1
TEST-AP02  AIR-AP2802I-Q-K9  TST0000AP02  00:00:5e:00:53:12  00:00:5e:00:53:02  192.168.0.12  17.12.7.13  2      J4       FlexConnect  Enabled  Registered  test-sw01.example.internal:Gi0/4  PoE+          Full Power  1d1h    1d1h   WNC1
TEST-AP03  CW9166I-Q         TST0000AP03  00:00:5e:00:53:13  00:00:5e:00:53:03  192.168.0.13  17.12.7.13  3      J4       FlexConnect  Enabled  Registered  test-sw01.example.internal:Gi0/3  PoE (legacy)  Full Power  5d20h   5d20h  WNC1
```

### Use cases

<details><summary>Case 1: List the access points that are not registered</summary><p>

```bash
wnc show ap -f json | jq -r '.[] | select(.state != "Registered") | "\(.ap_name) \(.state)"'
```

</p></details>

<details><summary>Case 2: Find the access points that rejoined without rebooting</summary><p>

```bash
wnc show ap -f json | jq -r '.[] | select(.uptime_seconds - .assoc_uptime_seconds > 600)? | .ap_name'
```

</p></details>

<details><summary>Case 3: List the most recently rebooted access points first</summary><p>

```bash
wnc show ap -b uptime_seconds
```

</p></details>

## wnc show ap-join

One row per access point the controller remembers, joined or not.

### Format

```bash
wnc show ap-join
```

### Expected result

```text
AP Name    Radio MAC          Ethernet MAC       IP Address    Status  Last Failure Phase  Last Join Failure        Last Config Failure  Last Discovery Failure  Last Disconnect Reason      Reboot Reason                 Last Join  Last Config  Last Discovery  Last Error  Controller
TEST-AP01  00:00:5e:00:53:01  00:00:5e:00:53:11  192.168.0.11  Joined  Join                jf-dtls-alert-from-peer  None                 None                    DTLS close alert from peer  ap-reboot-reason-img-upgrade  3h3m       3h3m         3h3m            3h7m        WNC1
TEST-AP02  00:00:5e:00:53:02  00:00:5e:00:53:12  192.168.0.12  Joined  Image-Download      None                     None                 None                    Image Download Success      ap-reboot-reason-img-upgrade  3h3m       3h3m         3h3m            3h9m        WNC1
TEST-AP03  00:00:5e:00:53:03  00:00:5e:00:53:13  192.168.0.13  Joined  Join                jf-dtls-alert-from-peer  None                 None                    DTLS close alert from peer  ap-reboot-reason-img-upgrade  3h3m       3h3m         3h3m            3h7m        WNC1
```

### Use cases

<details><summary>Case 1: List the access points the controller remembers but is not serving</summary><p>

```bash
wnc show ap-join -f json | jq -r '.[] | select(.status != "Joined") | "\(.ap_name) \(.disconnect_reason)"'
```

</p></details>

<details><summary>Case 2: Find the access points that joined but were never configured</summary><p>

```bash
wnc show ap-join -f json | jq -r '.[] | select(.last_join_seconds and (.last_config_seconds | not)) | .ap_name'
```

</p></details>

<details><summary>Case 3: List the most recently joined first</summary><p>

```bash
wnc show ap-join -b last_join_seconds
```

</p></details>

## wnc show ap-tag

One row per access point: which tags are in effect on it, and where they came from.

### Format

```bash
wnc show ap-tag
```

### Expected result

```text
AP Name    AP MAC             Misconfigured  Misconfig Reason  Tag Source  Filter Name  Policy Tag      Site Tag        RF Tag        AP Profile         Flex Profile         Controller
TEST-AP01  00:00:5e:00:53:01  No             -                 Static      -            test-wlan-flex  test-site-flex  test-inside   test-ap-profile01  test-flex-profile01  WNC1
TEST-AP02  00:00:5e:00:53:02  No             -                 Static      -            test-wlan-flex  test-site-flex  test-outside  test-ap-profile01  test-flex-profile01  WNC1
TEST-AP03  00:00:5e:00:53:03  No             -                 Static      -            test-wlan-flex  test-site-flex  test-inside   test-ap-profile01  test-flex-profile01  WNC1
```

### Use cases

<details><summary>Case 1: List the access points the controller considers misconfigured</summary><p>

```bash
wnc show ap-tag -f json | jq -r '.[] | select(.misconfigured == true) | .ap_name'
```

</p></details>

<details><summary>Case 2: Find the non-static tag sources, so the profile columns can disagree</summary><p>

```bash
wnc show ap-tag -f json | jq -r '.[] | select(.tag_source != "Static") | "\(.ap_name) \(.tag_source)"'
```

</p></details>

<details><summary>Case 3: Group the access points by RF tag</summary><p>

```bash
wnc show ap-tag -b rf_tag
```

</p></details>

## wnc show client

One row per associated client, joined across the collections that describe one.

### Format

```bash
wnc show client
```

### Expected result

```text
MAC                IPv4          IPv6          Device          Username   SSID          AP Name    Slot  Band  Protocol  Channel  State           RSSI    SNR   Rate     Streams  Assoc  Rx        Tx       Controller
00:00:5e:00:53:a1  192.168.0.21  2001:db8::11  Example Vendor  -          test-essid01  TEST-AP03  0     2.4   11ax      6ch      Run             -21dBm  78dB  143Mbps  1ss      2h15m  17.0MiB   19.6MiB  WNC1
00:00:5e:00:53:a2  192.168.0.22  -             Example Phone   test-user  test-essid02  TEST-AP01  1     5     11ac      64ch     Run             -43dBm  56dB  866Mbps  2ss      17m    107.9KiB  30.7KiB  WNC1
00:00:5e:00:53:a3  192.168.0.23  -             Example Sensor  -          test-essid03  TEST-AP03  2     6     11be      5ch      Authenticating  -55dBm  40dB  -        -        42s    -         -        WNC1
```

### Use cases

<details><summary>Case 1: List the clients with a weak signal</summary><p>

```bash
wnc show client -f json | jq -r '.[] | select(.rssi_dbm != null and .rssi_dbm < -70) | "\(.mac) \(.ap_name) \(.rssi_dbm)"'
```

</p></details>

<details><summary>Case 2: Find the clients stuck short of the run state</summary><p>

```bash
wnc show client -f json | jq -r '.[] | select(.state != "Run") | "\(.mac) \(.state) \(.assoc_seconds)s"'
```

</p></details>

<details><summary>Case 3: List the 6 GHz clients on one SSID</summary><p>

```bash
wnc show client -r 6 -s test-essid03
```

</p></details>

<details><summary>Case 4: Rank the clients by transmitted bytes, largest first</summary><p>

```bash
wnc show client -b tx_bytes --sort-order desc
```

</p></details>

## wnc show wlan

One row per WLAN and the policy profile bound to it.

### Format

```bash
wnc show wlan
```

### Expected result

```text
ID  Profile              SSID          Status   Security            Bands  Broadcast  P2P Block  Policy Status  Switching  Interface      Session TO  DHCP Required  Policy Profile         Tags            Controller
5   test-wlan-profile01  test-essid01  Enabled  WPA2 PSK            2.4    Enabled    Disabled   Active         Local      TEST-INTERNAL  43200       Yes            test-policy-profile01  test-wlan-flex  WNC1
6   test-wlan-profile02  test-essid02  Enabled  WPA2 PSK            5      Enabled    Disabled   Active         Local      TEST-INTERNAL  43200       Yes            test-policy-profile01  test-wlan-flex  WNC1
7   test-wlan-profile03  test-essid03  Enabled  WPA3 802.1X-SHA256  5/6    Enabled    Disabled   Active         Local      TEST-INTERNAL  43200       Yes            test-policy-profile01  test-wlan-flex  WNC1
```

### Use cases

<details><summary>Case 1: Find the enabled WLANs whose policy profile is shut, so they are off air</summary><p>

```bash
wnc show wlan -f json | jq -r '.[] | select(.status == "Enabled" and .policy_status == "Shutdown") | .profile'
```

</p></details>

<details><summary>Case 2: List the WLANs with no encryption</summary><p>

```bash
wnc show wlan -f json | jq -r '.[] | select(.security == "Open") | "\(.wlan_id) \(.ssid)"'
```

</p></details>

<details><summary>Case 3: List the WLANs on 6 GHz</summary><p>

```bash
wnc show wlan -f json | jq -r '.[] | select(.bands | test("6")) | .ssid'
```

</p></details>

## wnc show policy-tag

One row per WLAN binding a policy tag carries: what exists, and what each tag binds.

### Format

```bash
wnc show policy-tag
```

### Expected result

```text
Policy Tag          Description                       WLAN                 Policy Profile         Controller
default-policy-tag  Preconfigured default policy-tag  -                    -                      WNC3
test-wlan-flex      -                                 test-wlan-profile01  test-policy-profile01  WNC3
test-wlan-flex      -                                 test-wlan-profile02  test-policy-profile01  WNC3
test-wlan-flex      -                                 test-wlan-profile03  test-policy-profile01  WNC3
```

### Use cases

<details><summary>Case 1: List every tag with its bindings, before deciding what to delete</summary><p>

```bash
wnc show policy-tag --controller 192.168.0.1
```

</p></details>

<details><summary>Case 2: Find the tags that bind nothing, the usual target of a delete</summary><p>

```bash
wnc show policy-tag --format json | jq -r '.[] | select(.wlan == null) | .policy_tag'
```

</p></details>

<details><summary>Case 3: Group the tags by the policy profile the bindings point at</summary><p>

```bash
wnc show policy-tag --sort-by policy_profile
```

</p></details>

<details><summary>Case 4: Find which tags bind one WLAN profile</summary><p>

```bash
wnc show policy-tag --format json | jq -r '.[] | select(.wlan == "test-wlan-profile01") | .policy_tag'
```

</p></details>

## wnc show site-tag

One row per site tag: which profiles it names, and whether the site is local.

### Format

```bash
wnc show site-tag
```

### Expected result

```text
Site Tag          Description                     AP Join Profile     Flex Profile          Local Site  Controller
default-site-tag  Preconfigured default site tag  default-ap-profile  default-flex-profile  No          WNC3
test-site-flex    -                               test-ap-profile01   test-flex-profile01   No          WNC3
```

### Use cases

<details><summary>Case 1: List every site tag before deciding what to delete</summary><p>

```bash
wnc show site-tag --controller 192.168.0.1
```

</p></details>

<details><summary>Case 2: List the non-local sites, the ones a flex profile can be bound to</summary><p>

```bash
wnc show site-tag --format json | jq -r '.[] | select(.local_site == false) | .site_tag'
```

</p></details>

<details><summary>Case 3: Find the site tags naming no AP join profile</summary><p>

```bash
wnc show site-tag --format json | jq -r '.[] | select(.ap_join_profile == null) | .site_tag'
```

</p></details>

<details><summary>Case 4: Group the site tags by the AP join profile they name</summary><p>

```bash
wnc show site-tag --sort-by ap_join_profile
```

</p></details>

## wnc show rf-tag

One row per RF tag: which RF profile it names on each band.

### Format

```bash
wnc show rf-tag
```

### Expected result

```text
RF Tag          Description                   2.4 GHz Profile    5 GHz Profile      6 GHz Profile            Controller
default-rf-tag  Preconfigured default RF tag  default_rf_24gh    default_rf_5gh     default-rf-profile-6ghz  WNC3
test-inside     -                             test-rf-profile01  test-rf-profile03  test-rf-profile05        WNC3
test-outside    -                             test-rf-profile01  test-rf-profile04  test-rf-profile05        WNC3
```

### Use cases

<details><summary>Case 1: List every RF tag before deciding what to delete</summary><p>

```bash
wnc show rf-tag --controller 192.168.0.1
```

</p></details>

<details><summary>Case 2: Find the tags with no 6 GHz profile, so 6E falls back to default</summary><p>

```bash
wnc show rf-tag --format json | jq -r '.[] | select(.profile_6ghz == null) | .rf_tag'
```

</p></details>

<details><summary>Case 3: Find which tags name one profile on any band</summary><p>

```bash
wnc show rf-tag --format json \
  | jq -r '.[] | select([.profile_24ghz, .profile_5ghz, .profile_6ghz] | index("test-rf-profile02")) | .rf_tag'
```

</p></details>

<details><summary>Case 4: Group the tags by the 5 GHz profile</summary><p>

```bash
wnc show rf-tag --sort-by profile_5ghz
```

</p></details>
