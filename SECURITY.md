# Security Policy

This policy covers `wnc` and the container image published for it.

## Supported Versions

Only the most recent tagged release carries fixes, because no older tag gets a patch branch.
Reproduce a finding against that release before reporting it.

## Reporting a Vulnerability

Report privately through [GitHub Security Advisories](https://github.com/umatare5/cisco-wnc-cli/security/advisories/new), never through an issue or a pull request.

The response is best effort, and no reply time is promised.
The advisory goes out once the fix ships, carries a CVE request, and credits the reporter unless they ask otherwise.

## What to Include

**Redact these first.** Everyone invited to an advisory thread can read it, so none of them belongs in a report.

- The access token, from a log line, a command's output or a configuration file
- A device serial number, from an access point or from the controller
- A client-derived hostname, username or IPv6 address

Then include the following.

- **Affected versions** – the release or image tag reproduced against, and the controller's IOS-XE version.
- **Reproduction steps** – the command and its flags, and whether `--dry-run` reproduces it.
- **Output** – the table or the JSON, with every value above removed.
- **Impact** – state the exploit scenario, and what it reaches.
- **Suggested fix** – propose a remediation where you have one; this one is optional.
- **Disclosure status** – say whether it is shared elsewhere, and give your plan for sharing it.

## Exposure

The token is the base64 of `user:password`, carrying the password instead of a scoped token.
Keep it out of the places another person can read, in order of preference.

1. **A config file at mode `0600`.** It avoids the shell history and process list. A looser mode warns.
2. **`$WNC_ACCESS_TOKEN`.** Visible in the process environment.
3. **`--access-token`.** Visible in the process list, so use it interactively only.

`wnc generate-token` reads `--password`, `$WNC_PASSWORD`, then piped stdin. Leave the first two unset.

- **Posture** – each exposure is documented, not accidental, so keep the token on a controlled path.
- **Transport** – a controller is reached over HTTPS alone, and `--insecure` drops the certificate check.

> [!IMPORTANT]
> A leaked token is a leaked password. Rotating it requires changing the controller account password.

## In Scope

- The access token reaching a log line, a help text, or the output of any command except `generate-token`
- The access token reaching the process table through anything except `--access-token`
- Certificate verification weakened by anything except `--insecure` or the configuration file's `insecure` key
- A request reaching a controller, access point, radio or tag the command did not name
- The published container image, because a defect in the image `ghcr.io` serves may not exist in the source

## Out of Scope

The following fall outside this policy.

- A controller-side or IOS-XE defect, which belongs to Cisco PSIRT
- A dependency advisory with no code path reachable from `./cmd`, unless you show the reachable path
- An operator's own configuration, which [`docs/customization.md`](docs/customization.md) covers
