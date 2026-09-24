# Help

This page transcribes the help text of every command except `completion`, as the binary prints it.

## wnc

```text
NAME:
   wnc - Operate Cisco Catalyst 9800 Wireless Network Controllers

USAGE:
   wnc [global options] [command [command options]]

VERSION:
   dev

COMMANDS:
   deauth             Deauthenticate a client on a controller
   delete             Delete a tag from a controller
   disable            Disable an access point or one of its radios
   enable             Enable an access point or one of its radios
   generate-token, g  Print the Basic auth token for a controller account
   reset              Restart an access point or its controller session
   save-config        Save the running configuration to the startup configuration
   set                Create or update a tag on a controller
   show, s            Display controller state

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
   --help, -h          show help
   --version, -v       print the version
```

## wnc deauth

```text
NAME:
   wnc deauth - Deauthenticate a client on a controller

USAGE:
   wnc deauth (--mac <mac> | --username <username>) [options]

DESCRIPTION:
   --mac and --username are the mac and username columns of wnc show client,
   and they select different clients, so give only one. The controller resolves
   the target first, so a value it holds no client at is refused before the RPC,
   which answers the same whether or not a client was there. A username may hold
   more than one session, and the prompt says how many. The client is dropped and
   re-associates on its own, on a timer its supplicant sets rather than the
   controller, so allow about four minutes. The operation is absent before 17.15.6.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --mac string                                                       client MAC address, as wnc show client --columns mac prints it
   --username string                                                  client username, as wnc show client --columns username prints it
   --yes                                                              act without the confirmation prompt
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc delete

```text
NAME:
   wnc delete - Delete a tag from a controller

USAGE:
   wnc delete [command [command options]]

COMMANDS:
   policy-tag  Delete one policy tag
   site-tag    Delete one site tag
   rf-tag      Delete one RF tag

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc delete policy-tag

```text
NAME:
   wnc delete policy-tag - Delete one policy tag

USAGE:
   wnc delete policy-tag --name <name> [options]

DESCRIPTION:
   The name --name gives is read on the controller first, so a name it does not
   hold is a failure rather than a silent success. Pass --dry-run to report and
   change nothing.

OPTIONS:
   --name string  policy tag name, at most 32 characters
   --yes          act without the confirmation prompt
   --help, -h     show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc delete site-tag

```text
NAME:
   wnc delete site-tag - Delete one site tag

USAGE:
   wnc delete site-tag --name <name> [options]

DESCRIPTION:
   The name --name gives is read on the controller first, so a name it does not
   hold is a failure rather than a silent success. Pass --dry-run to report and
   change nothing.

OPTIONS:
   --name string  site tag name, at most 32 characters
   --yes          act without the confirmation prompt
   --help, -h     show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc delete rf-tag

```text
NAME:
   wnc delete rf-tag - Delete one RF tag

USAGE:
   wnc delete rf-tag --name <name> [options]

DESCRIPTION:
   The name --name gives is read on the controller first, so a name it does not
   hold is a failure rather than a silent success. Pass --dry-run to report and
   change nothing.

OPTIONS:
   --name string  RF tag name, at most 32 characters
   --yes          act without the confirmation prompt
   --help, -h     show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc disable

```text
NAME:
   wnc disable - Disable an access point or one of its radios

USAGE:
   wnc disable [command [command options]]

COMMANDS:
   ap     Disable one access point
   radio  Disable one radio of one access point

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc disable ap

```text
NAME:
   wnc disable ap - Disable one access point

USAGE:
   wnc disable ap --ap-name <ap-name> [options]

DESCRIPTION:
   --ap-name is the ap_name column of wnc show ap. The controller resolves it
   first, so a name it holds no access point under is refused before the RPC.
   This sets the access point's admin state, not one radio's.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --ap-name string  access point name, as shown in the ap_name column of wnc show ap
   --yes             act without the confirmation prompt
   --help, -h        show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc disable radio

```text
NAME:
   wnc disable radio - Disable one radio of one access point

USAGE:
   wnc disable radio --ap-name <ap-name> --slot <n> [options]

DESCRIPTION:
   --ap-name is the ap_name column of wnc show ap, and --slot is the Slot column
   of wnc show overview. The band the RPC needs is read from the controller and
   follows the radio type, so a dual-band radio takes one number either way.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --ap-name string  access point name, as shown in the ap_name column of wnc show ap
   --yes             act without the confirmation prompt
   --slot int        radio slot, as shown in the Slot column of wnc show overview
   --help, -h        show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc enable

```text
NAME:
   wnc enable - Enable an access point or one of its radios

USAGE:
   wnc enable [command [command options]]

COMMANDS:
   ap     Enable one access point
   radio  Enable one radio of one access point

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc enable ap

```text
NAME:
   wnc enable ap - Enable one access point

USAGE:
   wnc enable ap --ap-name <ap-name> [options]

DESCRIPTION:
   --ap-name is the ap_name column of wnc show ap. The controller resolves it
   first, so a name it holds no access point under is refused before the RPC.
   This sets the access point's admin state, not one radio's.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --ap-name string  access point name, as shown in the ap_name column of wnc show ap
   --yes             act without the confirmation prompt
   --help, -h        show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc enable radio

```text
NAME:
   wnc enable radio - Enable one radio of one access point

USAGE:
   wnc enable radio --ap-name <ap-name> --slot <n> [options]

DESCRIPTION:
   --ap-name is the ap_name column of wnc show ap, and --slot is the Slot column
   of wnc show overview. The band the RPC needs is read from the controller and
   follows the radio type, so a dual-band radio takes one number either way.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --ap-name string  access point name, as shown in the ap_name column of wnc show ap
   --yes             act without the confirmation prompt
   --slot int        radio slot, as shown in the Slot column of wnc show overview
   --help, -h        show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc generate-token

```text
NAME:
   wnc generate-token - Print the Basic auth token for a controller account

USAGE:
   wnc generate-token [options]

OPTIONS:
   --username string, -u string  controller username [$WNC_USERNAME]
   --password string, -p string  controller password; prefer WNC_PASSWORD or piped stdin [$WNC_PASSWORD]
   --help, -h                    show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc reset

```text
NAME:
   wnc reset - Restart an access point or its controller session

USAGE:
   wnc reset [command [command options]]

COMMANDS:
   ap      Restart one access point
   capwap  Reset one access point's controller session

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc reset ap

```text
NAME:
   wnc reset ap - Restart one access point

USAGE:
   wnc reset ap --ap-name <ap-name> [options]

DESCRIPTION:
   --ap-name is the ap_name column of wnc show ap. The controller resolves it
   first, so a name it holds no access point under is refused before the RPC.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --ap-name string  access point name, as shown in the ap_name column of wnc show ap
   --yes             act without the confirmation prompt
   --help, -h        show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc reset capwap

```text
NAME:
   wnc reset capwap - Reset one access point's controller session

USAGE:
   wnc reset capwap --ap-name <ap-name> [options]

DESCRIPTION:
   --ap-name is the ap_name column of wnc show ap. The controller resolves it
   first, so a name it holds no access point under is refused before the RPC.
   The access point does not reboot: only its CAPWAP session is re-established.
   Pass --dry-run to name the target and change nothing.

OPTIONS:
   --ap-name string  access point name, as shown in the ap_name column of wnc show ap
   --yes             act without the confirmation prompt
   --help, -h        show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc save-config

```text
NAME:
   wnc save-config - Save the running configuration to the startup configuration

USAGE:
   wnc save-config [options]

DESCRIPTION:
   The startup configuration is the only destination: no file may be named, and
   every change on the controller is persisted rather than only what this CLI
   wrote. An access point's admin state is unaffected, being no part of the
   configuration. The save took at most 3.7 seconds on every release measured, so
   a --timeout a read survives can still refuse it. Pass --dry-run to name the
   controller and change nothing.

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --yes                                                              act without the confirmation prompt
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc set

```text
NAME:
   wnc set - Create or update a tag on a controller

USAGE:
   wnc set [command [command options]]

COMMANDS:
   policy-tag  Create or update one policy tag
   site-tag    Create or update one site tag
   rf-tag      Create or update one RF tag

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc set policy-tag

```text
NAME:
   wnc set policy-tag - Create or update one policy tag

USAGE:
   wnc set policy-tag --name <name> [options]

DESCRIPTION:
   A name --name gives that the controller does not hold is created and one it
   holds is updated, so the same command may be repeated. A field no flag names
   is left as it is rather than cleared. Pass --dry-run to report and change
   nothing.

OPTIONS:
   --name string            policy tag name, at most 32 characters
   --yes                    act without the confirmation prompt
   --description string     description for the policy tag
   --wlan string            WLAN profile to bind, required with --policy-profile
   --policy-profile string  policy profile the WLAN is bound to, required with --wlan
   --help, -h               show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc set site-tag

```text
NAME:
   wnc set site-tag - Create or update one site tag

USAGE:
   wnc set site-tag --name <name> [options]

DESCRIPTION:
   A name --name gives that the controller does not hold is created and one it
   holds is updated, so the same command may be repeated. A field no flag names
   is left as it is rather than cleared. Pass --dry-run to report and change
   nothing.

OPTIONS:
   --name string             site tag name, at most 32 characters
   --yes                     act without the confirmation prompt
   --description string      description for the site tag
   --ap-join-profile string  AP join profile to bind
   --flex-profile string     flex profile to bind, which clears --local-site
   --local-site              mark the site local
   --help, -h                show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc set rf-tag

```text
NAME:
   wnc set rf-tag - Create or update one RF tag

USAGE:
   wnc set rf-tag --name <name> [options]

DESCRIPTION:
   A name --name gives that the controller does not hold is created and one it
   holds is updated, so the same command may be repeated. A field no flag names
   is left as it is rather than cleared. Pass --dry-run to report and change
   nothing.

OPTIONS:
   --name string           RF tag name, at most 32 characters
   --yes                   act without the confirmation prompt
   --description string    description for the RF tag
   --profile-24ghz string  2.4 GHz RF profile to bind
   --profile-5ghz string   5 GHz RF profile to bind
   --profile-6ghz string   6 GHz RF profile to bind
   --help, -h              show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port] [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for the controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
```

## wnc show

```text
NAME:
   wnc show - Display controller state

USAGE:
   wnc show [command [command options]]

COMMANDS:
   overview, o    Per-radio RF summary across 2.4, 5 and 6 GHz
   ap, a          Associated access points
   ap-join, join  Join, discovery and DTLS outcome per access point, joined or not
   ap-tag, tag    Tag assignment and its resolved values, per access point
   client, c      Associated wireless clients
   wlan, w        Configured WLANs and their bound policy profiles
   policy-tag     Configured policy tags and the WLANs they bind
   site-tag       Configured site tags and the profiles they name
   rf-tag         Configured RF tags and their per-band RF profiles

OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
   --help, -h                                                         show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing
```

## wnc show overview

```text
NAME:
   wnc show overview - Per-radio RF summary across 2.4, 5 and 6 GHz

USAGE:
   wnc show overview [options]

DESCRIPTION:
   One row per access point radio, sorted by ap_name.
   A cell reading "-" is a value the controller did not send.
   Admin State is the radio's own state: an access-point-level disable leaves
   it Enabled with Oper State reading Down, so read wnc show ap for that state.

OPTIONS:
   --sort-by key, -b key      sort key (see --list-keys) (default: "ap_name")
   --list-keys                print the keys --sort-by and --columns accept, then exit
   --radio string, -r string  band filter (2.4|5|6)
   --help, -h                 show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show ap

```text
NAME:
   wnc show ap - Associated access points

USAGE:
   wnc show ap [options]

DESCRIPTION:
   One row per access point in capwap-data, sorted by ap_name.
   A cell reading "-" is a value the controller did not send.
   Admin State is the access point's own state, which an access-point-level
   disable changes and the Admin State column of wnc show overview does not.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "ap_name")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show ap-join

```text
NAME:
   wnc show ap-join - Join, discovery and DTLS outcome per access point, joined or not

USAGE:
   wnc show ap-join [options]

DESCRIPTION:
   One row per access point the controller remembers, joined or not, sorted
   by ap_name.
   A cell reading "-" is a value the controller did not send.
   It is the only view that reports an access point capwap-data has dropped.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "ap_name")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show ap-tag

```text
NAME:
   wnc show ap-tag - Tag assignment and its resolved values, per access point

USAGE:
   wnc show ap-tag [options]

DESCRIPTION:
   One row per access point, sorted by ap_name.
   A cell reading "-" is a value the controller did not send.
   The tag columns are the tags resolved onto the access point. The profile
   columns come from the site tag and agree only while Tag Source is Static.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "ap_name")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show client

```text
NAME:
   wnc show client - Associated wireless clients

USAGE:
   wnc show client [options]

DESCRIPTION:
   One row per associated client, sorted by ap_name.
   A cell reading "-" is a value the controller did not send.
   --radio, --ssid and --ap-name narrow the list. A client whose band the
   controller did not report is excluded by --radio, and the count is logged.

OPTIONS:
   --sort-by key, -b key      sort key (see --list-keys) (default: "ap_name")
   --list-keys                print the keys --sort-by and --columns accept, then exit
   --radio string, -r string  band filter (2.4|5|6)
   --ssid string, -s string   keep only clients on this SSID
   --ap-name string           keep only clients on this access point
   --help, -h                 show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show wlan

```text
NAME:
   wnc show wlan - Configured WLANs and their bound policy profiles

USAGE:
   wnc show wlan [options]

DESCRIPTION:
   One row per WLAN and each policy profile bound to it, sorted by wlan_id,
   so a WLAN bound under two tags appears twice and an unbound one appears
   once with its policy columns empty.
   A cell reading "-" is a value the controller did not send.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "wlan_id")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show policy-tag

```text
NAME:
   wnc show policy-tag - Configured policy tags and the WLANs they bind

USAGE:
   wnc show policy-tag [options]

DESCRIPTION:
   One row per WLAN binding the tag carries, sorted by policy_tag, so a tag
   binding three WLANs appears three times and one binding none appears once.
   A cell reading "-" is a value the controller did not send.
   WLAN Profile is the profile name the binding keys on, not always the SSID.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "policy_tag")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show site-tag

```text
NAME:
   wnc show site-tag - Configured site tags and the profiles they name

USAGE:
   wnc show site-tag [options]

DESCRIPTION:
   One row per site tag, sorted by site_tag.
   A cell reading "-" is a value the controller did not send.
   The read asks for the values in effect, so a leaf a tag left at its default
   is reported rather than arriving as an absence.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "site_tag")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```

## wnc show rf-tag

```text
NAME:
   wnc show rf-tag - Configured RF tags and their per-band RF profiles

USAGE:
   wnc show rf-tag [options]

DESCRIPTION:
   One row per RF tag, sorted by rf_tag.
   A cell reading "-" is a value the controller did not send.
   The read asks for the values in effect, because a plain read omits the
   built-in tag's per-band profile names.

OPTIONS:
   --sort-by key, -b key  sort key (see --list-keys) (default: "rf_tag")
   --list-keys            print the keys --sort-by and --columns accept, then exit
   --help, -h             show help

GLOBAL OPTIONS:
   --config string     path to the JSON configuration file [$WNC_CONFIG]
   --log-level string  log verbosity (error|warning|debug) (default: "warning")
   --dry-run           report what would happen and change nothing

INHERITED OPTIONS:
   --controller string, -c string [ --controller string, -c string ]  controller host[:port], repeatable [$WNC_CONTROLLER]
   --access-token string                                              Basic auth token for every controller [$WNC_ACCESS_TOKEN]
   --insecure, -k                                                     skip TLS certificate verification
   --format string, -o string, -f string                              output format (table|json) (default: "table")
   --pretty                                                           draw the table with borders and status glyphs
   --timeout duration, -t duration                                    request timeout (default: 1m0s)
   --sort-order string                                                sort direction (asc|desc) (default: "asc")
   --columns keys [ --columns keys ]                                  comma-separated keys to print, or all (see --list-keys)
```
