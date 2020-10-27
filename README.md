# Stakebird DAO

[![Build Status](https://ci.publicawesome.com/api/badges/public-awesome/stakebird/status.svg)](https://ci.publicawesome.com/public-awesome/stakebird)
[![Go Report Card](https://goreportcard.com/badge/github.com/public-awesome/stakebird)](https://goreportcard.com/report/github.com/public-awesome/stakebird)
[![](https://tokei.rs/b1/github/public-awesome/stakebird)](https://github.com/public-awesome/stakebird)
[![codecov](https://codecov.io/gh/public-awesome/stakebird/branch/master/graph/badge.svg)](https://codecov.io/gh/public-awesome/stakebird)

Stakebird is a [content curation DAO](https://ethresear.ch/t/prediction-markets-for-content-curation-daos/1312).

Testnet coming soon.

## Install

Stakebird is built as a sovereign proof-of-stake blockchain that aims to interoperate with Cosmos Hub, Ethereum, and Bitcoin. 

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

### Run a local testnet with a connection to Gaia (Cosmos Hub)

Stakebird requires [Gaia](https://github.com/cosmos/gaia) (Cosmos Hub), and an IBC [relayer](https://github.com/iqlusioninc/relayer) to facilitate connections.

To setup a local testnet running Gaia and Stakebird, run:
```
./contrib/ibc/gaia-stakebird.sh
```

To setup the relayer and do a token transfer between chains, run:
```
./contrib/ibc/stakebird-xfer.sh
```

## CLI
The curating module can be accessed via CLI and REST API.

```sh
staked tx curating post -h
staked tx curating upvote -h
```
