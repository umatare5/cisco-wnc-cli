# Contributing

Thank you for considering a contribution.

## Development

Install [`gotestsum`](https://github.com/gotestyourself/gotestsum), [`golangci-lint`](https://golangci-lint.run/docs/welcome/install/local/), [`pre-commit`](https://pre-commit.com/#install) and [`gitleaks`](https://github.com/gitleaks/gitleaks#installing), then run `make pre-commit-install`.

- **The module and the artifacts differ** – imports carry `cisco-wnc-cli`, and every artifact is `wnc`.
- **Hook order** – the branch guard, `golangci-lint`, `actionlint`, `gitleaks`, then `markdownlint-cli2`.
- **The guard carries `fail_fast`** – a commit on `main` stops there, so work on a branch.
- **Only `gitleaks` comes from `PATH`** – pre-commit builds the rest at the versions it pins.
- **The markdown hook runs `--fix`** – it rewrites files, so reach it with `make pre-commit-test`.
- **One install arms every worktree** – the hook path is shared, so it passes `--allow-missing-config`.
- **`make build` skips a rebuild** – the file target does nothing while the binary exists.

> [!WARNING]
> `make clean` removes `./tmp` whole, and a worktree lives there too.
> Check what is there before running it.

## Commands

`make help` prints this list together with the tools each target needs.

| Command                     | Description                                              |
| :-------------------------- | :------------------------------------------------------- |
| `make help`                 | Display available targets and requirements               |
| `make build`                | Build the binary into `./tmp/wnc`                        |
| `make lint`                 | Verify the lint config, run golangci-lint, tidy `go.mod` |
| `make test-unit`            | Run unit tests with coverage using gotestsum             |
| `make test-unit-coverage`   | Generate the HTML coverage report                        |
| `make snapshot`             | Build a GoReleaser snapshot                              |
| `make clean`                | Remove the build and coverage artifacts                  |
| `make image`                | Build the Docker image                                   |
| `make pre-commit-install`   | Install the pre-commit hooks                             |
| `make pre-commit-test`      | Run every hook across the tree                           |
| `make pre-commit-uninstall` | Remove the pre-commit hooks                              |

## Build

`make image` builds the container image for the host architecture.

- **Context** – the binary cross-compiles into `./tmp/image/linux/<arch>`, and the build runs from `./tmp/image`.
- **Layout** – the `Dockerfile` expects the GoReleaser one, `linux/<arch>/wnc` beside `LICENSE` and `NOTICE`.
- **Tag** – the local image is `$USER/wnc`, and GoReleaser pushes released ones to `ghcr.io/umatare5/wnc`.

## Testing

Ship a test with the change it covers, and run the suite before every commit.

1. Add the test beside the code it covers, in the same package.
2. Take every fixture and sample identity from [`docs/testing.md`](docs/testing.md#fixture-identities), never from a device.
3. Run `make test-unit` – every package under `gotestsum`, with `-race` and a coverage profile.
4. Run `make test-unit-coverage` for the HTML report under `./coverage`.

Note the following as well.

- **The floor is enforced in CI** – the [coverage workflow](.github/workflows/go-test-coverage.yml) sets `coverage_threshold`, and fails below it.
- **The environment is cleared** – `make test-unit` unsets `WNC_*`, so a new one is added there and in `cli_test.go`.
- **The suite is documented** – [`docs/testing.md`](docs/testing.md) covers the column invariants and the TLS harness that stands in for a controller.

## Code Style

`golangci-lint` enforces what [`.golangci.yml`](.golangci.yml) configures, and `make lint` verifies that config before running it.

## Documentation

Every fact has one page that owns it, and the other pages link to it rather than restating it.

- **Headings are pinned** – [`.markdownlint-cli2.jsonc`](.markdownlint-cli2.jsonc) fixes the heading order wherever it sets `MD043`.
- **Contracts travel** – a heading change ships with its contract in the same pull request.
- **The transcript is verbatim** – [`docs/help.md`](docs/help.md) carries the binary's own `--help`, so a flag or usage change updates it.
- **Its `VERSION:` line reads `dev`** – `make build` stamps the version into the binary, so the transcript comes from `go build`.
- **[`NOTICE`](NOTICE) tracks the module set** – a change to what the binary links updates it.
- **Links are checked in CI only** – that run reaches third-party hosts, and `lychee .` reproduces a failure locally.

## Release

A release is prepared in one pull request, and merging it publishes everything.

1. Rename `## [Unreleased]` in [`CHANGELOG.md`](CHANGELOG.md) to `## [vX.Y.Z]`.
2. List the pull requests it carries, and add that version's link at the foot.
3. Update the version in the [`VERSION`](VERSION) file.
4. Refresh the coverage badge.

```bash
make test-unit
octocov badge coverage --config .octocov.yml > docs/assets/coverage.svg
```

A push to `main` touching `VERSION` runs the [release workflow](https://github.com/umatare5/cisco-wnc-cli/actions/workflows/go-release.yml), which tags the commit and publishes the release in the same run.

- **There is no manual trigger** – the push runs it, and the weekly snapshot build tags nothing.
- **Nothing automates the badge** – the coverage workflow enforces the floor but writes none.
- **The release links 404 until the merge** – [`lychee.toml`](lychee.toml) excludes the release-tag pattern.

## Pull Requests

Open a pull request against `main`, as a draft by default.

1. Fork the repository and create a feature branch.
2. Commit with a [Conventional Commits](https://www.conventionalcommits.org/) subject and a `Signed-off-by:` trailer.
3. Add tests and update the documentation beside the code.
4. Run `make lint` and `make test-unit`, then rebase against `main`.
5. Open the pull request.

Nothing in a commit identifies a real device or carries a credential.

- **`gitleaks` reads shapes** – a token a document or a fixture needs is registered in [`.gitleaks.toml`](.gitleaks.toml) by both path and text.
- **Each entry sets `condition = "AND"`** – registering by path alone would exempt every secret in that file.
- **A token that opens a real account stays out of the repository** – no allowlist entry makes it safe.
