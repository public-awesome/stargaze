#!/bin/sh
set -eux
# create users
rm -rf .testnode/
mkdir -p .testnode
STARSD_FILE=./bin/starsd
STARGAZE_HOME="$PWD/.testnode"
$STARSD_FILE config set client chain-id localnet-1 --home $STARGAZE_HOME
$STARSD_FILE config set client keyring-backend test --home $STARGAZE_HOME
$STARSD_FILE config set client output json --home $STARGAZE_HOME

yes | $STARSD_FILE keys add validator --home $STARGAZE_HOME
yes | $STARSD_FILE keys add creator --home $STARGAZE_HOME
yes | $STARSD_FILE keys add investor --home $STARGAZE_HOME
yes | $STARSD_FILE keys add funder --home $STARGAZE_HOME --pubkey "{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"AtObiFVE4s+9+RX5SP8TN9r2mxpoaT4eGj9CJfK7VRzN\"}"
VALIDATOR=$($STARSD_FILE keys show validator -a --home $STARGAZE_HOME)
CREATOR=$($STARSD_FILE keys show creator -a --home $STARGAZE_HOME)
INVESTOR=$($STARSD_FILE keys show investor -a --home $STARGAZE_HOME)
FUNDER=$($STARSD_FILE keys show funder -a --home $STARGAZE_HOME)
DENOM=ustars
# setup chain
$STARSD_FILE init stargaze --chain-id localnet-1 --home $STARGAZE_HOME

# modify config for development
config="$STARGAZE_HOME/config/config.toml"
if [ "$(uname)" = "Linux" ]; then
  sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
  sed -i "s/\"stake\"/\"$DENOM\"/g" $STARGAZE_HOME/config/genesis.json
else
  sed -i '' "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
  sed -i '' "s/\"stake\"/\"$DENOM\"/g" $STARGAZE_HOME/config/genesis.json
fi

# modify genesis params for localnet ease of use
# x/gov params change
# reduce voting period to 2 minutes
contents="$(jq '.app_state.gov.voting_params.voting_period = "120s"' $STARGAZE_HOME/config/genesis.json)" && echo "${contents}" >  $STARGAZE_HOME/config/genesis.json
# reduce minimum deposit amount to 10stake
contents="$(jq '.app_state.gov.deposit_params.min_deposit[0].amount = "10"' $STARGAZE_HOME/config/genesis.json)" && echo "${contents}" >  $STARGAZE_HOME/config/genesis.json
# reduce deposit period to 20seconds 
contents="$(jq '.app_state.gov.deposit_params.max_deposit_period = "20s"' $STARGAZE_HOME/config/genesis.json)" && echo "${contents}" >  $STARGAZE_HOME/config/genesis.json

$STARSD_FILE genesis add-genesis-account $VALIDATOR 10000000000000000ustars --home $STARGAZE_HOME
$STARSD_FILE genesis add-genesis-account $CREATOR 10000000000000000ustars --home $STARGAZE_HOME
$STARSD_FILE genesis add-genesis-account $INVESTOR 10000000000000000ustars --home $STARGAZE_HOME
$STARSD_FILE genesis add-genesis-account $FUNDER 10000000000000000ustars --home $STARGAZE_HOME
$STARSD_FILE genesis gentx validator 10000000000ustars --chain-id localnet-1 --keyring-backend test --home $STARGAZE_HOME 
$STARSD_FILE genesis collect-gentxs --home $STARGAZE_HOME
$STARSD_FILE genesis validate-genesis --home $STARGAZE_HOME
$STARSD_FILE start --home $STARGAZE_HOME
