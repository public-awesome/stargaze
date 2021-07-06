# Stargaze

[![Build Status](https://ci.publicawesome.com/api/badges/public-awesome/stargaze/status.svg)](https://ci.publicawesome.com/public-awesome/stargaze)
[![Go Report Card](https://goreportcard.com/badge/github.com/public-awesome/stargaze)](https://goreportcard.com/report/github.com/public-awesome/stargaze)
[![LOC](https://tokei.rs/b1/github/public-awesome/stargaze)](https://github.com/public-awesome/stargaze)
[![codecov](https://codecov.io/gh/public-awesome/stargaze/branch/master/graph/badge.svg)](https://codecov.io/gh/public-awesome/stargaze)

Stargaze is a protocol for incentivized content creation and curation. It creates attestations of content from social networks (currently Twitter), and enables their curation via quadratic voting. It is built as a sovereign proof-of-stake blockchain with it's own governance, that can interoperate with other blockchains such as Ethereum and Bitcoin via [IBC](https://cosmos.network/ibc).

## Install

### Run a local, single-node chain

```sh
# install binaries
make install

# create keys and initialize chain
make reset

# run
make start
```

### Test smart contracts

See [./scripts/wasm/README.md](./scripts/wasm/README.md).


