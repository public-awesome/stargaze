#!/bin/sh
set -eux
# create users
rm -rf $HOME/.starsd
STARSD_FILE=./bin/starsd
$STARSD_FILE config chain-id localnet-1
$STARSD_FILE config keyring-backend test
$STARSD_FILE config output json
yes | $STARSD_FILE keys add validator
yes | $STARSD_FILE keys add creator
yes | $STARSD_FILE keys add investor
yes | $STARSD_FILE keys add funder --pubkey "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AtObiFVE4s+9+RX5SP8TN9r2mxpoaT4eGj9CJfK7VRzN\"}"
VALIDATOR=$($STARSD_FILE keys show validator -a)
CREATOR=$($STARSD_FILE keys show creator -a)
INVESTOR=$($STARSD_FILE keys show investor -a)
FUNDER=$($STARSD_FILE keys show funder -a)
DENOM=ustars
# setup chain
$STARSD_FILE init stargaze --chain-id localnet-1
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.starsd/config/genesis.json
# modify config for development
config="$HOME/.starsd/config/config.toml"
if [ "$(uname)" = "Linux" ]; then
  sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
else
  sed -i '' "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
fi
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.starsd/config/genesis.json
# modify genesis params for localnet ease of use
# x/gov params change
# reduce voting period to 2 minutes
contents="$(jq '.app_state.gov.voting_params.voting_period = "120s"' $HOME/.starsd/config/genesis.json)" && echo "${contents}" >  $HOME/.starsd/config/genesis.json
# reduce minimum deposit amount to 10stake
contents="$(jq '.app_state.gov.deposit_params.min_deposit[0].amount = "10"' $HOME/.starsd/config/genesis.json)" && echo "${contents}" >  $HOME/.starsd/config/genesis.json
# reduce deposit period to 20seconds 
contents="$(jq '.app_state.gov.deposit_params.max_deposit_period = "20s"' $HOME/.starsd/config/genesis.json)" && echo "${contents}" >  $HOME/.starsd/config/genesis.json

$STARSD_FILE genesis add-genesis-account $VALIDATOR 10000000000000000ustars
$STARSD_FILE genesis add-genesis-account $CREATOR 10000000000000000ustars
$STARSD_FILE genesis add-genesis-account $INVESTOR 10000000000000000ustars
$STARSD_FILE genesis add-genesis-account $FUNDER 10000000000000000ustars
$STARSD_FILE genesis gentx validator 10000000000ustars --chain-id localnet-1 --keyring-backend test
$STARSD_FILE genesis collect-gentxs
$STARSD_FILE genesis validate-genesis
$STARSD_FILE start
