#!/bin/sh

# create users
rm -rf $HOME/.starsd
starsd config chain-id localnet-1
starsd config keyring-backend test
starsd config output json
yes | starsd keys add validator
yes | starsd keys add creator
yes | starsd keys add investor
yes | starsd keys add funder --pubkey "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AtObiFVE4s+9+RX5SP8TN9r2mxpoaT4eGj9CJfK7VRzN\"}"
VALIDATOR=$(starsd keys show validator -a)
CREATOR=$(starsd keys show creator -a)
INVESTOR=$(starsd keys show investor -a)
FUNDER=$(starsd keys show funder -a)

# setup chain
starsd init stargaze --chain-id localnet-1
# modify config for development
config="$HOME/.starsd/config/config.toml"
if [ "$(uname)" = "Linux" ]; then
  sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
  sed -i "s/enable = false/enable = true/g" $config
else
  sed -i '' "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
  sed -i '' "s/enable = false/enable = true/g" $config
fi

# modify genesis params for localnet ease of use
# x/gov params change
# reduce voting period to 2 minutes
contents="$(jq '.app_state.gov.voting_params.voting_period = "120s"' $HOME/.starsd/config/genesis.json)" && echo "${contents}" >  $HOME/.starsd/config/genesis.json
# reduce minimum deposit amount to 10stake
contents="$(jq '.app_state.gov.deposit_params.min_deposit[0].amount = "10"' $HOME/.starsd/config/genesis.json)" && echo "${contents}" >  $HOME/.starsd/config/genesis.json
# reduce deposit period to 20seconds 
contents="$(jq '.app_state.gov.deposit_params.max_deposit_period = "20s"' $HOME/.starsd/config/genesis.json)" && echo "${contents}" >  $HOME/.starsd/config/genesis.json

starsd add-genesis-account $VALIDATOR 10000000000000000stake
starsd add-genesis-account $CREATOR 10000000000000000stake
starsd add-genesis-account $INVESTOR 10000000000000000stake
starsd add-genesis-account $FUNDER 10000000000000000stake
starsd gentx validator 10000000000stake --chain-id localnet-1 --keyring-backend test
starsd collect-gentxs
starsd validate-genesis
starsd start
