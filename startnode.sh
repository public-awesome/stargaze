#!/bin/sh

# create users
rm -rf $HOME/.starsd
starsd config chain-id localnet-1
starsd config keyring-backend test
starsd config output json
yes | starsd keys add validator
yes | starsd keys add creator
yes | starsd keys add investor
yes | starsd keys add funder --pubkey starspub1addwnpepqwmnprxqj8at8rgnejj5y7kay5xt7u0r74eqnj4dwvkkcwtyf9nxsve82v3
VALIDATOR=$(starsd keys show validator -a)
CREATOR=$(starsd keys show creator -a)
INVESTOR=$(starsd keys show investor -a)
FUNDER=$(starsd keys show funder -a)

# setup chain
starsd init stargaze --stake-denom ustarx --chain-id localnet-1

# modify config for development
config="$HOME/.starsd/config/config.toml"
if [ "$(uname)" = "Linux" ]; then
  sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
else
  sed -i '' "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
fi

# start
starsd add-genesis-account $VALIDATOR 10000000000000000ustarx
starsd add-genesis-account $CREATOR 10000000000000000ustarx
starsd add-genesis-account $INVESTOR 10000000000000000ustarx
starsd add-genesis-account $FUNDER 10000000000000000ustarx
starsd gentx validator 10000000000ustarx --chain-id localnet-1 --keyring-backend test
starsd collect-gentxs
starsd validate-genesis
starsd start
