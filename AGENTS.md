# Repository Instructions

> [!IMPORTANT]
> Read [`README.md`](README.md) for project overview.

## Tech Stack

The list below covers the toolchain, the modules the binary links, and the release builder.

- Go 1.27+ (see [`go.mod`](go.mod))
- [`umatare5/cisco-ios-xe-wireless-go`](https://github.com/umatare5/cisco-ios-xe-wireless-go) – sole RESTCONF SDK for Cisco C9800 WNC
- [`urfave/cli/v3`](https://github.com/urfave/cli) – command tree, flags and application lifecycle
- [`sirupsen/logrus`](https://github.com/sirupsen/logrus) – process logger and the `slog` bridge the SDK takes
- [`olekukonko/tablewriter`](https://github.com/olekukonko/tablewriter) – borderless table writer
- [`goreleaser`](https://goreleaser.com/) – cross-platform release builds, configured by [`.goreleaser.yml`](.goreleaser.yml)

## Repository Structure

Read from [`cmd/main.go`](cmd/main.go) – each package is named for what it owns.

- [`cmd/`](cmd) – entry point, exiting on what `internal/cli` returns
- [`internal/cli/`](internal/cli) – urfave command tree, flag definitions, exit codes, and the `version` string ldflags sets
- [`internal/config/`](internal/config) – flag, environment and file resolution in that precedence, and the grammars it parses
- [`internal/log/`](internal/log) – logrus setup and the `*slog.Logger` the SDK takes
- [`internal/wnc/`](internal/wnc) – sole importer of the SDK, with one `fetch_*.go` per read and one file per action beside it
- [`internal/show/`](internal/show) – per-command row building, the concurrent controller fan-out, and the enum display tables
- [`internal/render/`](internal/render) – `Column[T]`, shared by the table and the JSON writer so a column cannot exist in one only
- [`docs/`](docs) – reference pages behind the README, indexed by [`docs/README.md`](docs/README.md)
- [`examples/`](examples) – a working configuration file whose `note` fields carry what JSON has no comment syntax for
- [`scripts/`](scripts) – helper scripts the pre-commit hooks run

## Setup and Commands

Run `make pre-commit-install` first.

- Read [`Makefile`](Makefile) for every make target and its requirements.
- Read [`CONTRIBUTING.md`](CONTRIBUTING.md) for the contribution rules.

## Code Style

Follow [Effective Go](https://go.dev/doc/effective_go) conventions and the software development principles DRY/YAGNI/SRP.

- Keep code simple and readable, avoiding clever tricks that obscure intent.
- Keep every change minimal, in code, tests, comments and documentation.
- Write simple comments that explain the reasoning behind the code, not just what it does.
- A claim in [`docs/`](docs) cites the Go file and line, or the controller reading it came from.

## Testing

Follow [`CONTRIBUTING.md`](CONTRIBUTING.md).

- Run `make lint` and `make test-unit` before creating a commit.
- Take every fixture and sample identity from [`docs/testing.md`](docs/testing.md#fixture-identities), never a value read off a device.

## Commits and PRs

Follow [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `chore(deps):`, etc.).

- Sign off every commit with `Signed-off-by:` (DCO).
- Open PRs against `main`. Create Draft PR as default.
