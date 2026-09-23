# Testing

This document records how the suite is arranged, what it asserts, and the identities every fixture and sample uses.

```bash
make test-unit            # go test -race with coverage
make test-unit-coverage   # plus an HTML report under ./coverage
make lint                 # config verify, golangci-lint and go mod tidy
```

## Conventions

Tests sit next to the code they cover, in the same package.
So a test asserts an unexported rule directly, with no wrapper added just to reach it.

Tables and `t.Parallel()` are the default.
There is no assertion library and no golden file.

## Fixture Identities

A fixture copies the shape of a real controller response, with every identity replaced.
Every value below is synthetic, and this section is the source for the samples on every `command.*.md` page as well as for the test files.

### Reserved ranges

Each kind below draws on a reserved range, so none of these values is invented:

| Kind         | Range                                     | Reserved by                    |
| :----------- | :---------------------------------------- | :----------------------------- |
| MAC address  | `00:00:5e:00:53:00` – `00:00:5e:00:53:ff` | [RFC 7042 §2.1.2][rfc7042-sec] |
| IPv4 address | `192.168.0.0/16`                          | [RFC 1918][rfc1918]            |
| IPv6 address | `2001:db8::/32`                           | [RFC 3849][rfc3849]            |
| Domain name  | `example.internal`                        | ICANN private-use TLD          |

The MAC block's first octet has the I/G bit clear, so no fixture address is multicast.
The block sits inside IANA's own OUI, and [RFC 7042][rfc7042] reserves that part of it for documentation.
So no address here can collide with a vendor assignment.

Its last octet carries the role:

- **`:01`–`:0f`** – an access point's radio base address, `TEST-APnn` pairing with `:nn`
- **`:11`–`:1f`** – an access point's Ethernet address, `TEST-APnn` pairing with `:1n`
- **`:a1`–`:af`** – a client station

### Defined values

The rest have no standard to draw on, so this CLI defines them:

| Kind                 | Value                          |
| :------------------- | :----------------------------- |
| Controller name      | `WNC1` – `WNC4`                |
| Controller host      | `192.168.0.1` – `192.168.0.4`  |
| Access point name    | `TEST-AP01` – `TEST-AP99`      |
| Access point serial  | `TST0000AP01` – `TST0000AP99`  |
| Access point address | `192.168.0.11` onward          |
| Client address       | `192.168.0.21` onward          |
| Access token         | `TestToken0123456789ABCDEF==`  |
| Password             | `test-token-123`               |
| Client username      | `test-user` onward             |
| SSID                 | `test-essid01` onward          |
| Profile              | `test-<kind>-profile01` onward |
| Tag                  | a `test-` prefix               |

Some of those values carry a rule the table has no room for:

- **Prefix** – `labo-` is allowed wherever `test-` is, in the same case, so `LABO-AP01` fits too
- **Profile kind** – `<kind>` is `wlan`, `policy`, `rf`, `ap` or `flex`, so the value names its own kind
- **Controller name** – every prompt, report and `Controller` column shows it, not the host
- **Serial** – no `test-` prefix fits `[A-Z]{3}[0-9]{4}[A-Z0-9]{4}`, and no real serial has week 00 of year 00

### Exceptions

These categories are deliberately outside the scheme.

- **A grammar case is not an identity** – `internal/config` keeps `a.example` and `[2001:db8::1]` for parse assertions
- **A dialed address must be unroutable** – [RFC 1918][rfc1918] space can be live on a developer's own LAN

So a test whose host may be reached takes `192.0.2.0/24` from [RFC 5737][rfc5737] instead.
Where the assertion is that nothing answers, it takes `240.0.0.1` from [RFC 5735][rfc5735].

> [!IMPORTANT]
> Never paste a captured MAC address, serial number, hostname, username, SSID or tag name into a fixture or a sample transcript.
> Nothing in this CLI redacts one, so a value pasted by hand reaches the tree unchanged.

## RESTCONF Layer

`internal/wnc` is tested against a TLS test server serving canned responses, routed on the last element of the request path.
The SDK pins its own dialer, so no transport can be injected and the server has to be a real listener.
One test asserts that before the fixture-driven ones rely on it.

Fixtures deliberately include credential leaves, to assert the hand-written struct drops them at decode.

## Fan-out

`internal/show` tests the fan-out with a fetch function that never reaches the network.
Client construction still happens for real, against the unroutable range above.
So the tests cover outcome classification, reporting order and the print-nothing rule, with no server running.

## Command Tree

`internal/cli` drives the real command tree end to end and asserts the exit-code contract.
That covers usage faults, an unknown command, the help paths, the settings rejections and how `generate-token` takes a password.

Nothing in that package runs in parallel, and that is deliberate.
Since urfave reads the `WNC_*` variables at parse time, the suite clears them with `t.Setenv`, which forbids `t.Parallel`.
`make test-unit` clears them again at the process level so a developer's own shell cannot change what the assertions see.

## Invariants

Some checks are about shape rather than behavior, and they exist because the failure they catch is silent.

- **Every declaration, one order** – the sort-key list, the column list and the json tags must agree
- **Banned outright** – `omitempty` drops a reported zero, an empty string and a reported false
- **Allowed on a pointer only** – `omitzero` there means nil, which encodes "not reported"
- **Banned as well** – every value of the json `format` tag is rejected at run time, not at compile time
- **Every command carries the usage hook** – urfave consults the running command's own and no other

> [!NOTE]
> `json/v2` drops a tag a sibling field repeats.
> That would leave a column in the table but missing from the JSON, with nothing failing.
> That silence is why those declarations are asserted rather than reviewed.

## Live Verification

The suite has no integration-test target, so these checks run outside `make test-unit`.

Every run follows the same steps.

1. Pick a controller to run against.
2. Run the command as the section below shows for its kind.
3. Check the output against that command's own `command.*.md` page.
4. Confirm the result on the controller itself.

Every command page carries an `Expected result` block, and step 3 matches it in shape rather than in value.
Those samples take their identities from [Fixture Identities](#fixture-identities).
So the shape and the wording have to agree, with this site's own values in place of the sample's.

### Show commands

Run the command, check its output against [`command.show.md`](command.show.md), then compare the values with the controller's own output:

```bash
wnc show overview -c "<host>" --insecure
```

`show ap dot11 5ghz summary`, `show ap uptime` and `show ap tag summary` each cover one view.
`show wlan id <n>` and `show wireless client summary` do the same.
The tag views compare against `show wireless tag {rf,site,policy} summary` and `detailed <name>`.

Some headings differ from the device's on purpose, because these views follow the YANG leaf and the write flag.
[`measurements.md`](measurements.md#device-heading-map) maps them in full.

### Action commands

Run the command with `--dry-run` first, then repeat it for real and check its prompt and report against [`command.action.md`](command.action.md):

```bash
wnc --dry-run disable radio --ap-name "<ap-name>" --slot 1 -c "<host>"  # names the target, posts nothing
wnc           disable radio --ap-name "<ap-name>" --slot 1 -c "<host>"  # prompts, then posts
wnc show overview -c "<host>" --insecure                                # read the slot back
```

Action commands interrupt service, so automating them is not realistic.
`reset ap` alone takes its clients off the air for about four minutes.

A tree that acts cannot be compared with the device's own output, because running it changes the controller.
`--dry-run` prints no prompt and stops before the RPC, so it covers everything except the write.
A `show` command reads the result back afterwards.

Each write was last taken on a release that [`measurements.md`](measurements.md) records, with the arm it used and what the post moved.

`deauth` is the only write whose support differs by release.
The client delete RPC arrives in the release [`measurements.md`](measurements.md#client) records, so an earlier one refuses the post.
Both arms are pinned on fixtures as well, in `internal/wnc/deauth_test.go` and `internal/cli/deauth_test.go`.

> [!NOTE]
> A dry run does not replace repeating the write on a release where it matters.

### Other commands

`generate-token` contacts no controller, so reproduce its output instead:

```bash
printf '%s' "<password>" | wnc generate-token -u admin  # the token
printf '%s' "admin:<password>" | base64                 # the same string
```

`save-config` does reach a controller and writes, so verify it the way an action command is verified.
`--dry-run` names the controller it would save, and the run itself reports `running configuration saved`.

`save-config` names no target, so every change on that controller is persisted, including ones this CLI did not make.
Pick a controller where that is acceptable.

### Administrative state

An administrative state has no second source to compare against, as [`measurements.md`](measurements.md#not-measured) records.

Read the state back with `wnc show ap` and `wnc show overview` instead.
After an access-point-level disable the two disagree by design, which [`measurements.md`](measurements.md#access-point) records.

[rfc1918]: https://datatracker.ietf.org/doc/html/rfc1918
[rfc3849]: https://datatracker.ietf.org/doc/html/rfc3849
[rfc5735]: https://datatracker.ietf.org/doc/html/rfc5735
[rfc5737]: https://datatracker.ietf.org/doc/html/rfc5737
[rfc7042]: https://datatracker.ietf.org/doc/html/rfc7042
[rfc7042-sec]: https://datatracker.ietf.org/doc/html/rfc7042#section-2.1.2
