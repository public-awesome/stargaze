# Stakebird DAO

Stakebird is a [content curation DAO](https://ethresear.ch/t/prediction-markets-for-content-curation-daos/1312).

Testnet coming soon.

## Run

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

## CLI
The curating module can be accessed via CLI and REST API.

```sh
stakecli tx curating post -h
stakecli tx curating upvote -h
```
