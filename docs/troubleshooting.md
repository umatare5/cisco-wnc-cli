# Troubleshooting

Every failure this CLI prints names its own cause, and the sections below index them.

## Reading a Failure

This CLI words every failure itself rather than passing the controller's through.
A failed read ends in `cause=…`, and a refusal is one `wnc: …` line.

```text
error: WNC1: the controller answered 401 Unauthorized (cause=auth)
```

The controller leads the sentence.
Where the failure cost only some cells rather than the whole row set, `endpoint=` joins the cause inside the parentheses.

`--log-level debug` renders the same record as logfmt and adds the SDK's own line above it.
It never carries the token or any header.

```text
level=debug msg="HTTP request failed" error="Get \"https://192.168.0.1/restconf/data/...\": context deadline exceeded" url="https://192.168.0.1/restconf/data/..."
level=error msg="the controller did not answer in time" cause=timeout controller=WNC1 status=0
```

## Error Causes

Every failed request carries one of these in its `cause=` field, and the set is closed.

| Cause                       | Meaning                                             |
| :-------------------------- | :-------------------------------------------------- |
| [`auth`](#auth)             | The controller rejected the token with 401          |
| [`forbidden`](#forbidden)   | The account authenticated but may not read          |
| [`not-found`](#not-found)   | The controller has no such node                     |
| [`timeout`](#timeout)       | No answer within `--timeout`                        |
| [`tls`](#tls)               | The certificate did not verify                      |
| [`connection`](#connection) | The host did not resolve, or refused the connection |
| [`http`](#http)             | A status none of the causes above name              |
| [`canceled`](#canceled)     | The run was interrupted                             |
| [`internal`](#internal)     | A fault in this CLI rather than in the exchange     |

### auth

The controller rejected the token with 401. Regenerate it and check the account is still valid:

```bash
printf '%s' "$WNC_PASSWORD" | wnc generate-token -u admin
```

A token built by hand with `echo` rather than `printf` carries a trailing newline, and the controller rejects it.

### forbidden

The account authenticated but is not permitted to read. RESTCONF needs privilege level 15 on IOS-XE.

### not-found

The controller has no such node.
Where the read is optional this is normal, and this CLI reports an empty result.
Seeing it as a failure means the whole collection is missing.

That usually means RESTCONF is enabled and the wireless feature set is not.

### timeout

No answer within `--timeout`, which defaults to 60 seconds per request.
A controller under load answers a whole-container read slowly, so raise it before suspecting the network:

```bash
wnc show client --timeout 120s
```

A joined view makes several sequential reads, so its ceiling is that multiple of the timeout.
Controllers are read concurrently, so adding one does not extend it.

> [!NOTE]
> A failure at 30 seconds every time is the connect stage, which the SDK pins and `--timeout` does not move.
> That run has no route to the controller, which [Measurements](measurements.md#configuration) records.

### tls

The certificate did not verify.
A controller usually presents a self-signed or internally issued certificate.
Trust its issuer rather than passing `--insecure`, which accepts an intercepted certificate too:

```bash
SSL_CERT_FILE=/path/to/ca-bundle.pem wnc show overview
```

> [!IMPORTANT]
> From Go 1.27 on macOS and Windows, `SSL_CERT_FILE` replaces the platform verifier rather than adding to it.
> So the bundle must carry every root the process needs.
> `GODEBUG=x509sslcertoverrideplatform=0` restores the platform verifier.

### connection

The host did not resolve, or refused the connection. Check the controller is reachable on 443:

```bash
curl -kIsS "https://<host>/restconf/" -o /dev/null -w '%{http_code}\n'
```

### http

The controller answered with a status none of the causes above name.
Examples are a `400` from a rejected parameter and a `5xx` from a controller under load.
The sentence carries the code, and `--log-level debug` puts it in a `status` field:

```bash
curl -kisS -H "Authorization: Basic $WNC_ACCESS_TOKEN" \
        -H "Accept: application/yang-data+json" \
        "https://<host>/restconf/data/<the endpoint the log named>" | head -1
```

### canceled

The run was interrupted. Nothing is printed and the exit code is 130.

### internal

A fault in this CLI rather than in the exchange. Re-run with `--log-level debug` and open an issue with that output.

## FAQ

Each question below names a result the CLI produced, and the answer gives the cause behind it.

### Why is every value `-`?

`-` is the controller reporting nothing, which [Architecture](architecture.md#absence) separates from zero.
Where a whole column is `-`, look for an `endpoint=` line: a secondary read that failed leaves exactly its own columns empty.

### Why is the row set smaller than expected?

These commands hide entries on purpose:

- **A remote-LAN port** – `show overview` drops it, the controller listing it among the radios with no RF
- **A dangling policy-tag binding** – `show wlan` drops it and reports the count as a warning

`show client -r` and `show overview -r` report how many rows they dropped for a missing band.
The `--ssid` and `--ap-name` filters report no such count.

### Why is the exit code 3 when the table is full?

At least one read failed while at least one succeeded, which [Architecture](architecture.md#exit-codes) sets out.
The stderr lines say which.

### Why does an action say the controller holds no access point of that name?

The target is resolved before anything is sent, so nothing was written. Both a stale record and another controller produce this.

- **Another controller holds it** – name that one with `--controller`
- **The name is not the controller's** – take it from the `ap_name` column of `wnc show ap`

### Why does `enable radio` say the controller reports no radio address?

The controller holds that access point and sent no base radio address for it.
The slot is keyed on that address, so nothing could be read and nothing was sent.
Read `wnc show ap` for the record as the controller has it.

### Why does `enable radio` say the controller holds no radio in that slot?

Read the `Slot` column of `wnc show overview` for the slots that exist.
This exits 1 rather than 2, because the controller answered before the slot was refused.

### Why does `enable radio` say the slot is a remote-LAN port?

The port carries no band and no admin state, so there is nothing to set.
`wnc show overview` drops it, which is why the slot appears in no table.

### Why does `enable radio` say the RPC has no band number for that radio type?

The RPC takes a band number 1 to 4, and the type the controller reported maps to none of them.
A 2.4/6 GHz XOR radio fits both the dual-band 3 and the 6 GHz 4, and the UWB and invalid types name no band.
[Measurements](measurements.md#radio) records the spellings that take no number.

Where the controller reported no type at all, the message says so rather than naming one.

### Why does `enable radio` say the band is unknown or unreported?

The controller reported no serving band for that radio, or a spelling outside 2.4, 5 and 6 GHz.
The message names which.
Read `wnc show overview` for the band the controller does report.

### Why is the radio accepted on other slots only?

The RPC's `must` clause forbids that band-and-slot pair.
The pair follows the radio type rather than the band a dual-band radio is serving.
So read the `Slot` column of `wnc show overview` and not its `Band` column, which [Measurements](measurements.md#radio) records.

### Why does `deauth` say the controller rejected the operation?

The client delete RPC is absent before 17.15.6, which [Measurements](measurements.md#client) records.
A release without it answers the post with a `400`.
The target was resolved on that controller a moment earlier, so the release is the only cause left.

### Why does an action refuse a piped stdin?

There is no terminal to answer the prompt on, so nothing was sent.
Pass `--yes` to act without one, or `--dry-run` to check the target.

### Why does `set` say the tag exists and there is nothing to change?

The name is already on the controller and no binding flag was given.
Name a field – `wnc set rf-tag --help` lists the ones that kind takes.

### Why does `delete` say the controller holds no tag of that name?

Tag names are case sensitive, and each kind is a separate list.
A policy tag and an RF tag may share a name without being related.
`wnc show ap-tag` names the tags in effect.

### Why is a tag write refused with 400?

This CLI checks the name's own pattern first, so a 400 that gets through is usually a binding the release does not accept.
Read the node back to see what it holds:

```bash
curl -kisS -H "Authorization: Basic $WNC_ACCESS_TOKEN" \
        -H "Accept: application/yang-data+json" \
        "https://<host>/restconf/data/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=<name>"
```

A profile name that exists nowhere is **not** a cause.
The controller keeps a dangling reference, which [Measurements](measurements.md#tag) records.

### Why does `save-config` report no result?

The controller answered the save and carried no `result` string, which is its whole account of what it did.
Reporting a save that may not have happened is the one answer this command must not give.
Whether it was saved is readable on the controller itself.

### Why does a command refuse a positional argument?

Every value a command takes is named by a flag, so a bare word is a fault wherever it lands and nothing was sent.
Where the leaf has a target flag the message names it, and elsewhere `--help` lists the flag the value belonged to.

### Why is a tag name refused for a leading or trailing space?

The key leaf's own pattern refuses one, and a flag value reaches that check with its spaces intact.
Quote the intended name, or drop the padding – an inner space is legal and passes.
