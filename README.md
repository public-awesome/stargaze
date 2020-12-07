# Stakebird

[![Build Status](https://ci.publicawesome.com/api/badges/public-awesome/stakebird/status.svg)](https://ci.publicawesome.com/public-awesome/stakebird)
[![Go Report Card](https://goreportcard.com/badge/github.com/public-awesome/stakebird)](https://goreportcard.com/report/github.com/public-awesome/stakebird)
[![LOC](https://tokei.rs/b1/github/public-awesome/stakebird)](https://github.com/public-awesome/stakebird)
[![codecov](https://codecov.io/gh/public-awesome/stakebird/branch/master/graph/badge.svg)](https://codecov.io/gh/public-awesome/stakebird)

Stakebird is a protocol for incentivized content creation and curation. It creates attestations of content from social networks (currently Twitter), and enables their curation via quadratic voting. It is built as a sovereign proof-of-stake blockchain with it's own governance, that can interoperate with other blockchains such as Ethereum and Bitcoin via [IBC](https://cosmos.network/ibc).

## Install

### Run a local, single-node chain

```sh
# install binaries
make

# create keys
make create-wallet

# initialize chain
make init

# run
staked start
```

## CLI

The curating module can be accessed via CLI and REST API.

```sh
staked tx curating post -h
staked tx curating upvote -h
```
