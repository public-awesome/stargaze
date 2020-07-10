# Stakebird DAO

Stakebird is a [content curation DAO](https://ethresear.ch/t/prediction-markets-for-content-curation-daos/1312).

Testnet coming soon.

## Install

Stakebird is a sovereign chain that connects to other chains via [IBC](https://cosmos.network/ibc). At minimum, it requires [Gaia](https://github.com/cosmos/gaia) (Cosmos Hub), and a [relayer](https://github.com/iqlusioninc/relayer) to facilitate connections.

### Run a local, single-node chain

```sh
# install binaries
make install

# create keys
make create-wallet

# initialize chain
make init

# run
staked start
```

### Run a local testnet with IBC

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
stakecli tx curating post -h
stakecli tx curating upvote -h
```
