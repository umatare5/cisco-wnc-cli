# Changelog

Each section covers one release and lists the pull requests it carries.

## [Unreleased]

## [v0.4.1]

- [#109](https://github.com/umatare5/cisco-wnc-cli/pull/109) – Write American English in the comments and the build config
- [#110](https://github.com/umatare5/cisco-wnc-cli/pull/110) – Render the coverage badge in CI instead of committing it
- [#111](https://github.com/umatare5/cisco-wnc-cli/pull/111) – Center pretty glyphs and tell settings turned off from faults

## [v0.4.0]

This release renames 17 keys and 5 headings and changes the policy_status values (#107), so check a script that reads the JSON or names a key before upgrading.

- [#106](https://github.com/umatare5/cisco-wnc-cli/pull/106) – Rebuild on every make build and keep worktrees on make clean
- [#107](https://github.com/umatare5/cisco-wnc-cli/pull/107) – Rename the show keys and headings to the controller's vocabulary

## [v0.3.0]

This release shrinks the default column set (#100), renames --sort-keys to --list-keys (#102) and renames six keys (#102, #103), so check a script that reads the JSON or names a key before upgrading.

- [#97](https://github.com/umatare5/cisco-wnc-cli/pull/97) – Bump umatare5/common to v0.21.1 to fix the CodeQL workflow
- [#98](https://github.com/umatare5/cisco-wnc-cli/pull/98) – Restructure the documentation set and rewrite the prose it carries
- [#99](https://github.com/umatare5/cisco-wnc-cli/pull/99) – Accept --dry-run after the subcommand and honor it on show
- [#100](https://github.com/umatare5/cisco-wnc-cli/pull/100) – Add --columns to show and hide identifiers by default
- [#101](https://github.com/umatare5/cisco-wnc-cli/pull/101) – Add access point coordinates, height and floor to wnc show ap
- [#102](https://github.com/umatare5/cisco-wnc-cli/pull/102) – Rename --sort-keys to --list-keys and drop unit suffixes from keys
- [#103](https://github.com/umatare5/cisco-wnc-cli/pull/103) – Rename tx_power_dbm to txpower and the ChUtil heading to Utilization

## [v0.2.1]

- [#85](https://github.com/umatare5/cisco-wnc-cli/pull/85) – Update dependency golangci/golangci-lint to v2.13.2
- [#86](https://github.com/umatare5/cisco-wnc-cli/pull/86) – Extend the shared Renovate profile and pin the Alpine tag
- [#88](https://github.com/umatare5/cisco-wnc-cli/pull/88) – Run the build and tests weekly
- [#89](https://github.com/umatare5/cisco-wnc-cli/pull/89) – Report failed scheduled runs and build a weekly release snapshot
- [#90](https://github.com/umatare5/cisco-wnc-cli/pull/90) – Unify the license spelling and repin the README heading contract
- [#91](https://github.com/umatare5/cisco-wnc-cli/pull/91) – Update all patch dependencies
- [#92](https://github.com/umatare5/cisco-wnc-cli/pull/92) – Update module golang.org/x/sys to v0.44.0 for a security advisory
- [#93](https://github.com/umatare5/cisco-wnc-cli/pull/93) – Update dependency goreleaser/goreleaser to v2.18.1
- [#94](https://github.com/umatare5/cisco-wnc-cli/pull/94) – Update umatare5/common action to v0.21.0
- [#95](https://github.com/umatare5/cisco-wnc-cli/pull/95) – Follow the repository rename to umatare5/cisco-wnc-cli

## [v0.2.0]

- [#80](https://github.com/umatare5/cisco-wnc-cli/pull/80) – Rebuild the CLI and modernize the development base
- [#81](https://github.com/umatare5/cisco-wnc-cli/pull/81) – Restore the link check triggers and fix the image COPY path
- [#82](https://github.com/umatare5/cisco-wnc-cli/pull/82) – Pin the documentation skeletons and ship the license notices
- [#83](https://github.com/umatare5/cisco-wnc-cli/pull/83) – Restructure the documentation set and pin its heading skeletons

[v0.4.1]: https://github.com/umatare5/cisco-wnc-cli/releases/tag/v0.4.1
[v0.4.0]: https://github.com/umatare5/cisco-wnc-cli/releases/tag/v0.4.0
[v0.3.0]: https://github.com/umatare5/cisco-wnc-cli/releases/tag/v0.3.0
[v0.2.1]: https://github.com/umatare5/cisco-wnc-cli/releases/tag/v0.2.1
[v0.2.0]: https://github.com/umatare5/cisco-wnc-cli/releases/tag/v0.2.0
