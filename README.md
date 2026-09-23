<div align="center">

  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/umatare5/cisco-wnc-cli/main/docs/assets/logo_dark.png" width="115px" />
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/umatare5/cisco-wnc-cli/main/docs/assets/logo.png" width="115px" />
    <img alt="wnc" src="https://raw.githubusercontent.com/umatare5/cisco-wnc-cli/main/docs/assets/logo.png" width="115px" />
  </picture>

  <h1>cisco-wnc-cli</h1>

  <p>A command-line interface for Cisco Catalyst 9800 Wireless Network Controllers.</p>

  <p>
    <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/umatare5/cisco-wnc-cli?label=Latest%20version" />
    <a href="https://github.com/umatare5/cisco-wnc-cli/actions/workflows/go-test-build.yml"><img alt="Test and Build" src="https://github.com/umatare5/cisco-wnc-cli/actions/workflows/go-test-build.yml/badge.svg?branch=main" /></a>
    <a href="https://github.com/umatare5/cisco-wnc-cli/actions/workflows/go-vulncheck.yml"><img alt="govulncheck" src="https://github.com/umatare5/cisco-wnc-cli/actions/workflows/go-vulncheck.yml/badge.svg?branch=main" /></a><br>
    <img alt="Test Coverage" src="https://raw.githubusercontent.com/umatare5/cisco-wnc-cli/main/docs/assets/coverage.svg" />
    <a href="https://www.bestpractices.dev/projects/10820"><img alt="OpenSSF Best Practices" src="https://www.bestpractices.dev/projects/10820/badge" /></a>
    <a href="./LICENSE"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" /></a>
    <a href="https://developer.cisco.com/codeexchange/github/repo/umatare5/wnc"><img alt="Published" src="https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg" /></a>
  </p>

</div>

## Overview

This CLI manages wireless LANs across multiple [Catalyst 9800 Wireless Network Controllers](https://www.cisco.com/site/us/en/products/networking/wireless/wireless-lan-controllers/catalyst-9800-series/index.html).

- 📤 **Shell Friendly**: A borderless table `awk` and `cut` read, and a JSON array keyed by the sort names
- 🌐 **Multi Controller**: Reads concurrently and labels per WLC, so offline hosts drop rows and not the run
- 🔭 **Joined Views**: `show overview`, `show ap`, `show client` and `show wlan` join what the device splits
- 🎨 **Colorful Output**: `--pretty` borders and glyphs the table, and `--format json` feeds machines and AIs

## Supported Environment

Cisco Catalyst 9800 Wireless Network Controller running on:

- **Cisco IOS-XE 17.12.5 or later** – Last verified on 17.12.8, with no `wnc deauth` before 17.15.6
- **Cisco IOS-XE 17.15.6 or later** – Last verified on 17.15.6
- **Cisco IOS-XE 17.18.4a or later** – Last verified on 17.18.4a

> [!IMPORTANT]
> This CLI requires these minimum versions due to RESTCONF defects in earlier releases.
> It fails on **17.15.4b** and **17.18.1**.
> See [cisco-ios-xe-wireless-go #28](https://github.com/umatare5/cisco-ios-xe-wireless-go/issues/28) and [cisco-ios-xe-wireless-go #29](https://github.com/umatare5/cisco-ios-xe-wireless-go/issues/29) for details.

## Installation

This CLI supports container images and OS-specific binaries.

```bash
docker pull ghcr.io/umatare5/wnc
```

Or, download the binaries from [Releases](https://github.com/umatare5/cisco-wnc-cli/releases).
`(linux|darwin)_(amd64|arm64)` and `windows_amd64` are supported.

## Quick Start

This CLI reaches a controller over RESTCONF alone, so enable RESTCONF and HTTPS on the Catalyst 9800 first.

See the [Programmability Configuration Guide, Cisco IOS XE 17.15.x – RESTCONF](https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/prog/configuration/1715/b_1715_programmability_cg/restconf_protocol.html#id_125840) for enabling them.

### 1. Generate a Basic Auth token

```bash
read -rs WNC_PASSWORD # < your-password
printf '%s' "$WNC_PASSWORD" | wnc generate-token -u admin
# Output: YWRtaW46eW91ci1wYXNzd29yZA== (admin:your-password)
```

### 2. Set the environment variables

```bash
export WNC_CONTROLLER="wnc1.example.internal"
export WNC_ACCESS_TOKEN="YWRtaW46eW91ci1wYXNzd29yZA=="
```

### 3. Print a wireless overview

```bash
wnc show overview
```

> [!TIP]
> `wnc completion <shell>` writes the script to stdout, and `--help` names each shell and the line it needs.

## CLI Reference

This CLI groups commands by purpose. Each links to a section showing its output.

### Show commands

These commands read a controller and print a table or JSON. See [Show commands](docs/command.show.md) for details.

| Command                                      | Description                                                 |
| :------------------------------------------- | :---------------------------------------------------------- |
| [`wnc show overview`][wnc-show-overview]     | One row per radio, with RF summary across 2.4, 5 and 6 GHz  |
| [`wnc show ap`][wnc-show-ap]                 | One row per associated access point                         |
| [`wnc show ap-join`][wnc-show-ap-join]       | One row per access point's join, discovery and DTLS outcome |
| [`wnc show ap-tag`][wnc-show-ap-tag]         | One row per access point, with assigned and resolved tags   |
| [`wnc show client`][wnc-show-client]         | One row per associated wireless client                      |
| [`wnc show wlan`][wnc-show-wlan]             | One row per WLAN and its bound policy profile               |
| [`wnc show policy-tag`][wnc-show-policy-tag] | One row per policy tag and the WLANs it binds               |
| [`wnc show site-tag`][wnc-show-site-tag]     | One row per site tag and the profiles it names              |
| [`wnc show rf-tag`][wnc-show-rf-tag]         | One row per RF tag and its per-band RF profiles             |

### Action commands

These commands act on a controller in [order](docs/architecture.md#acting-on-a-controller). See [Action commands](docs/command.action.md) for details.

| Command                                                                      | Description                                     |
| :--------------------------------------------------------------------------- | :---------------------------------------------- |
| [`wnc reset ap`][wnc-reset-ap]                                               | Restart one access point                        |
| [`wnc reset capwap`][wnc-reset-capwap]                                       | Reset one access point's controller session     |
| [`wnc (enable\|disable) (ap\|radio)`][wnc-enable-wnc-disable]                | Enable or disable one access point or one radio |
| [`wnc set (policy\|site\|rf)-tag`][wnc-set-policy-tag-site-tag-rf-tag]       | Create or update one tag on a controller        |
| [`wnc delete (policy\|site\|rf)-tag`][wnc-delete-policy-tag-site-tag-rf-tag] | Delete one tag from a controller                |
| [`wnc deauth`][wnc-deauth]                                                   | Deauthenticate one client by MAC or username    |

### Other commands

These commands stand outside both groups for specific reasons. See [Other commands](docs/command.other.md) for details.

| Command                                    | Description                                                 |
| :----------------------------------------- | :---------------------------------------------------------- |
| [`wnc generate-token`][wnc-generate-token] | Generate the Basic auth token for a controller account      |
| [`wnc save-config`][wnc-save-config]       | Save the running configuration to the startup configuration |

### Help

`--help` lists commands and flags. See [Help](docs/help.md) for all transcripts except `completion`.

## Customization

This CLI reads its settings from flags, environment variables and a config file. See [Customization](docs/customization.md) for the details.

## Troubleshooting

See [Troubleshooting](docs/troubleshooting.md) for every error message this CLI prints, and what each one means.

## Documentation

Both pages below are written for a contributor rather than an operator.

- **[Architecture](docs/architecture.md)** – the output contract, the absence rule, the exit codes, and the write order
- **[Measurements](docs/measurements.md)** – every reading taken on a live controller, and the gaps

## Contributing

See [`CONTRIBUTING.md`](CONTRIBUTING.md) for the development setup, the test conventions and the release steps.

## License

MIT. The binary statically links MIT and BSD 3-Clause dependencies, so their notices are reproduced in [`NOTICE`](NOTICE).
[`LICENSE`](LICENSE) ships beside it in every release archive and container image.

[wnc-show-overview]: docs/command.show.md#wnc-show-overview
[wnc-show-ap]: docs/command.show.md#wnc-show-ap
[wnc-show-ap-join]: docs/command.show.md#wnc-show-ap-join
[wnc-show-ap-tag]: docs/command.show.md#wnc-show-ap-tag
[wnc-show-client]: docs/command.show.md#wnc-show-client
[wnc-show-wlan]: docs/command.show.md#wnc-show-wlan
[wnc-show-policy-tag]: docs/command.show.md#wnc-show-policy-tag
[wnc-show-site-tag]: docs/command.show.md#wnc-show-site-tag
[wnc-show-rf-tag]: docs/command.show.md#wnc-show-rf-tag
[wnc-reset-ap]: docs/command.action.md#wnc-reset-ap
[wnc-reset-capwap]: docs/command.action.md#wnc-reset-capwap
[wnc-enable-wnc-disable]: docs/command.action.md#wnc-enable-wnc-disable
[wnc-set-policy-tag-site-tag-rf-tag]: docs/command.action.md#wnc-set-policy-tag-site-tag-rf-tag
[wnc-delete-policy-tag-site-tag-rf-tag]: docs/command.action.md#wnc-delete-policy-tag-site-tag-rf-tag
[wnc-deauth]: docs/command.action.md#wnc-deauth
[wnc-generate-token]: docs/command.other.md#wnc-generate-token
[wnc-save-config]: docs/command.other.md#wnc-save-config
