<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [Unreleased]

## [v3.0.0](https://github.com/public-awesome/stargaze/releases/tag/v3.0.0) - 2022-02-22

- [#531](https://github.com/public-awesome/stargaze/pull/537) Add CosmWasm support to wasdm v0.23.0 and custom messages
- [#538](https://github.com/public-awesome/stargaze/pull/535) Upgrade to Cosmos SDK `V0.45.1`
- [#538](https://github.com/public-awesome/stargaze/pull/535) Upgrade spm to allow setting iavl-cache-size
- [#537](https://github.com/public-awesome/stargaze/pull/537) Upgrade to ibc-go `v2.0.3`

## [v2.0.0](https://github.com/public-awesome/stargaze/releases/tag/v2.0.0) - 2022-01-21

- [#528](https://github.com/public-awesome/stargaze/pull/528) Add rocksdb as compile option
- [#524](https://github.com/public-awesome/stargaze/pull/524) Upgrade to Cosmos SDK `V0.45.0`
- [#523](https://github.com/public-awesome/stargaze/pull/523) Upgrade Module Version to v2
- [#522](https://github.com/public-awesome/stargaze/pull/522) Fix Amino for Claim Tx
- [#495](https://github.com/public-awesome/stargaze/issues/495) Adds denom metadata as part of the migration
- [#502](https://github.com/public-awesome/stargaze/issues/502) Upgrades to ibc-go v2
- [#517](https://github.com/public-awesome/stargaze/pull/517) Fix IBC gov proposals routing

## [v1.1.2](https://github.com/public-awesome/stargaze/releases/tag/v1.1.2) - 2022-01-05

- [#519](https://github.com/public-awesome/stargaze/pull/519) Fix missing GRPC Routes for claim and alloc modules

## [v1.1.1](https://github.com/public-awesome/stargaze/releases/tag/v1.1.1) - 2021-12-30

- [#511](https://github.com/public-awesome/stargaze/pull/511) Bump Cosmos SDK to `v0.44.5` and ibc-go to `v1.2.5`

## [v1.1.0](https://github.com/public-awesome/stargaze/releases/tag/v1.1.0) - 2021-11-29

- Add tendermint `rollback` command to help operators restore to previous height in case of apphash errors
- Bump Cosmos SDK to `v0.44.4`
- Bump ibc-go to `v1.2.3`
- Fix CLI output to stdout/stderr preventing integration with tooling like `jq`

## [v1.0.0](https://github.com/public-awesome/stargaze/releases/tag/v1.0.0) - 2021-10-29

Initial Release
