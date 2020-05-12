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
rocketd start
```

# CLI

The Stake module can be accessed via CLI and REST API.

```sh
rocketcli tx stake post -h
rocketcli tx stake delegate -h
```
