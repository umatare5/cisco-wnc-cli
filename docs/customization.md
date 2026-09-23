# Customization

This CLI reads a setting from a flag first, then from the environment, then from the configuration file.

## Flags

[`help.md`](help.md) transcribes `--help` for every command except `completion`.
The flags a command takes follow from what that command does:

| Set                     | Flags                                                       |
| :---------------------- | :---------------------------------------------------------- |
| Every command           | `--config`, `--log-level`, `--dry-run`                      |
| Contacting a controller | `--controller`, `--access-token`, `--insecure`, `--timeout` |
| Reading, so `show` only | `--format`, `--pretty`, `--columns`                         |
| Sorting, so `show` only | `--sort-by`, `--sort-keys`, `--sort-order`                  |
| Acting, so asking first | `--yes`                                                     |

Without a command, `--dry-run` validates the configuration and contacts nothing.
With one, a `show` command names the controllers it would read and contacts none, while [Architecture](architecture.md#acting-on-a-controller) sets out where an action stops.

`--columns` selects the keys the table and the JSON carry, in the order given, rather than the view's default set.
That set leaves out client and access point addresses, serial numbers, coordinates, usernames and LLDP neighbors, because each identifies a device, a person or a site.
It also leaves out a fixed set of other columns on `show ap`, `show ap-join`, `show client` and `show wlan`, so the default table stays narrow.

The `Expected result` of each view in [Show commands](command.show.md) shows the set that remains, and `--sort-keys` lists every key whether the default set prints it or not.
`--sort-by` still takes a key the selection leaves out, because sorting runs before it.

`--insecure` drops certificate verification, so it accepts an interception as readily as a private CA.
Trust the issuer instead, which [Troubleshooting](troubleshooting.md#tls) sets out.

> [!WARNING]
> `--ap-name` reaches the controller in the request URL and `deauth --username` in the body, and neither is bounded locally.
> A secret mistyped into one has already left the host.

## Environment Variables

`WNC_CONFIG` reaches every command, and `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN` reach the commands that contact a controller:

| Variable           | Description                                             |
| :----------------- | :------------------------------------------------------ |
| `WNC_CONTROLLER`   | Controller `host[:port]`, comma separated for several   |
| `WNC_ACCESS_TOKEN` | Basic auth token applied to every controller            |
| `WNC_CONFIG`       | Configuration file path, replacing the default location |

These are read by `wnc generate-token` alone, which contacts no controller:

| Variable       | Description         |
| :------------- | :------------------ |
| `WNC_USERNAME` | Controller username |
| `WNC_PASSWORD` | Controller password |

## Configuration File

A file keeps the controller list out of every invocation.
[`SECURITY.md`](../SECURITY.md#exposure) ranks it first among the places the token can sit.
[`examples/config.json`](../examples/config.json) is a working file to copy.

- **The path** – `--config`, then `$WNC_CONFIG`, then `$XDG_CONFIG_HOME/wnc/config.json`, then `~/.config/wnc/config.json`
- **One token covers the file** – every controller it lists is read with that credential, and so is a host named on `--controller`
- **The read is strict** – an unknown, duplicated or case-differing key, a comment and a trailing comma are each rejected

Check a hand-edited file without contacting anything:

```bash
wnc --config ./config.json --dry-run
```
