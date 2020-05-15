# Stakebird Chain

Blockchain that powers Stakebird, as well as the zone for [Game Of Zones](https://cosmos.network/goz).

# Run

```sh
# install binaries
make install
# create keys
make create-wallet

# init
make init

# run
staked start
```

# CLI

The Stake module can be accessed via CLI and REST API.

```sh
stakecli tx stake post -h
stakecli tx stake delegate -h
```
