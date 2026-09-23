# Action Commands

These commands act on a controller, and each one runs the sequence of guards that [Architecture](architecture.md#acting-on-a-controller) sets out.

## wnc reset ap

Restart one access point.

### Format

```bash
wnc reset ap --ap-name <ap-name>
```

### Expected result

```text
Reset TEST-AP01 on WNC1? Its clients disconnect for about four minutes. [y/N]: y
TEST-AP01 on WNC1: reset sent
```

### Use cases

<details><summary>Case 1: Restart the access point a weak-signal client is on</summary><p>

```bash
wnc reset ap --ap-name TEST-AP01
```

</p></details>

<details><summary>Case 2: Check the target without acting</summary><p>

```bash
wnc --dry-run reset ap --ap-name TEST-AP01
```

Expected result:

```text
TEST-AP01 on WNC1: would reset
```

</p></details>

<details><summary>Case 3: Watch it leave and come back</summary><p>

```bash
wnc reset ap --ap-name TEST-AP01 --yes
watch -n 10 'wnc show ap-join -f json | jq -r ".[] | select(.ap_name==\"TEST-AP01\") | .status"'
```

</p></details>

## wnc reset capwap

Reset one access point's controller session.

### Format

```bash
wnc reset capwap --ap-name <ap-name>
```

### Expected result

```text
Reset the CAPWAP session of TEST-AP01 on WNC1? It rejoins within about ten seconds and does not reboot. [y/N]: y
TEST-AP01 on WNC1: CAPWAP reset sent
```

### Use cases

<details><summary>Case 1: Reset the session of the access point a stuck join is suspected on</summary><p>

```bash
wnc reset capwap --ap-name TEST-AP01
```

</p></details>

<details><summary>Case 2: Check the target without acting</summary><p>

```bash
wnc --dry-run reset capwap --ap-name TEST-AP01
```

Expected result:

```text
TEST-AP01 on WNC1: would reset the CAPWAP session
```

</p></details>

<details><summary>Case 3: Watch the association age reset while the uptime keeps climbing</summary><p>

```bash
wnc reset capwap --ap-name TEST-AP01 --yes
watch -n 5 'wnc show ap -f json | jq -r ".[] | select(.ap_name==\"TEST-AP01\") | \"up \(.uptime_seconds)s assoc \(.assoc_uptime_seconds)s\""'
```

</p></details>

## wnc enable, wnc disable

Set the administrative state of one access point, or of one of its radios.

### Format

```bash
wnc enable ap --ap-name <ap-name>
wnc disable ap --ap-name <ap-name>
wnc enable radio --ap-name <ap-name> --slot <n>
wnc disable radio --ap-name <ap-name> --slot <n>
```

### Expected result

```text
Disable TEST-AP01 on WNC1? This sets the access point's admin state, not one radio's. [y/N]: y
TEST-AP01 on WNC1: disable sent
```

### Use cases

<details><summary>Case 1: Take one radio out of service</summary><p>

```bash
wnc disable radio --ap-name TEST-AP01 --slot 1
```

</p></details>

<details><summary>Case 2: Check the target and the band without acting</summary><p>

```bash
wnc --dry-run disable radio --ap-name TEST-AP01 --slot 1
```

Expected result:

```text
slot 1 (5 GHz) of TEST-AP01 on WNC1: would disable
```

</p></details>

<details><summary>Case 3: Bring an access point back and read both views</summary><p>

```bash
wnc enable ap --ap-name TEST-AP01 --yes
wnc show ap -f json | jq -r '.[] | select(.ap_name=="TEST-AP01") | .admin'
wnc show overview -f json | jq -r '.[] | select(.ap_name=="TEST-AP01") | "\(.slot) \(.admin)/\(.oper)"'
```

</p></details>

## wnc set policy-tag, site-tag, rf-tag

Create or update one tag on a controller.

### Format

```bash
wnc set rf-tag --name <name> [--profile-5ghz <profile>]
wnc set site-tag --name <name> [--ap-join-profile <profile>]
wnc set policy-tag --name <name> [--wlan <profile> --policy-profile <profile>]
```

### Expected result

```text
Create RF tag test-rf-inside on WNC1? [y/N]: y
RF tag test-rf-inside on WNC1: created
```

### Use cases

<details><summary>Case 1: Create an RF tag bound to two band profiles</summary><p>

```bash
wnc set rf-tag --name test-rf-inside \
  --description "inside coverage" \
  --profile-24ghz test-rf-profile01 \
  --profile-5ghz test-rf-profile03
```

</p></details>

<details><summary>Case 2: Add a band to it later, leaving the rest as it is</summary><p>

```bash
wnc set rf-tag --name test-rf-inside --profile-6ghz test-rf-profile05
```

</p></details>

<details><summary>Case 3: Bind a WLAN to a policy profile</summary><p>

```bash
wnc set policy-tag --name test-wlan --wlan test-wlan-profile01 --policy-profile test-policy-profile01
```

</p></details>

<details><summary>Case 4: Check what would happen without acting</summary><p>

```bash
wnc --dry-run set site-tag --name test-site --ap-join-profile test-ap-profile01
```

Expected result:

```text
site tag test-site on WNC1: would update
```

</p></details>

## wnc delete policy-tag, site-tag, rf-tag

Delete one tag from a controller.

### Format

```bash
wnc delete policy-tag --name <name>
wnc delete site-tag --name <name>
wnc delete rf-tag --name <name>
```

### Expected result

```text
Delete RF tag test-rf-inside on WNC1? [y/N]: y
RF tag test-rf-inside on WNC1: deleted
```

### Use cases

<details><summary>Case 1: Delete a tag nothing resolves to</summary><p>

```bash
wnc show ap-tag -f json | jq -r '.[] | .rf_tag' | sort -u
wnc delete rf-tag --name test-rf-retired
```

</p></details>

<details><summary>Case 2: Check the target without acting</summary><p>

```bash
wnc --dry-run delete policy-tag --name test-wlan
```

Expected result:

```text
policy tag test-wlan on WNC1: would delete
```

</p></details>

## wnc deauth

Deauthenticate a client on a controller, by address or by username.

### Format

```bash
wnc deauth --mac <mac>
wnc deauth --username <username>
```

### Expected result

```text
Deauthenticate 00:00:5e:00:53:a1 on WNC3? It is dropped and reconnects on its own. [y/N]: y
00:00:5e:00:53:a1 on WNC3: deauthenticate sent
```

### Use cases

<details><summary>Case 1: Drop a client stuck short of `Run`</summary><p>

```bash
wnc show client -f json | jq -r '.[] | select(.state != "Run") | .mac'
wnc deauth --mac 00:00:5e:00:53:a1
```

</p></details>

<details><summary>Case 2: Drop every session a user holds</summary><p>

```bash
wnc show client -f json | jq -r '.[] | select(.username) | .username' | sort -u
wnc deauth --username test-user
```

</p></details>

<details><summary>Case 3: Check the target without acting, on either arm</summary><p>

```bash
wnc --dry-run deauth --mac 00:00:5e:00:53:a1
wnc --dry-run deauth --username test-user
```

Expected result:

```text
00:00:5e:00:53:a1 on WNC3: would deauthenticate
2 clients authenticated as test-user on WNC3: would deauthenticate
```

</p></details>

<details><summary>Case 4: Drop a client and watch it come back</summary><p>

```bash
wnc deauth --mac 00:00:5e:00:53:a1 --yes
watch -n 5 'wnc show client -f json | jq -r ".[] | select(.mac==\"00:00:5e:00:53:a1\") | \"\(.state) assoc \(.assoc_seconds)s\""'
```

</p></details>
