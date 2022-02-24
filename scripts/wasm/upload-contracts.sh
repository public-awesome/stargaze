#!/bin/bash

## CONFIG
BINARY="${BINARY:-starsd}"
DENOM="${DENOM:-ustars}"
CHAIN_ID="${CHAIN_ID:-stargaze-devnet-1}"
NODE="${NODE:-https://rpc.devnet.publicawesome.dev:443}"

if [ "$1" = "" ]
then
  echo "Usage: $0 1 arg required, address or key name."
  exit
fi

TXFLAG="--gas-prices 0.01$DENOM --gas auto --gas-adjustment 1.3 -y -b block --chain-id $CHAIN_ID --node $NODE --output json"

CONTRACTS_REPO=https://github.com/public-awesome/stargaze-contracts
CONTRACTS_TAG=v0.2.0
MARKETPLACE_REPO=https://github.com/public-awesome/marketplace
MARKETPLACE_TAG=v0.1.12

if [[ -z "$GITHUB_OAUTH_TOKEN" ]]; then
    echo "Must set GITHUB_OAUTH_TOKEN in environment" 1>&2
    exit 1
fi

if ! command -v fetch &> /dev/null
then
    echo "fetch could not be found: https://github.com/gruntwork-io/fetch"
    exit 1
fi

# Fetch contract wasm binaries
fetch --repo=$MARKETPLACE_REPO --tag=$MARKETPLACE_TAG --release-asset="collection_factory.wasm" .
fetch --repo=$MARKETPLACE_REPO --tag=$MARKETPLACE_TAG --release-asset="sg_marketplace.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="sg721.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="minter.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="whitelist.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="royalty_group.wasm" .
fetch --repo=https://github.com/CosmWasm/cw-nfts --tag=v0.11.0 --release-asset="cw721_metadata_onchain.wasm" .
fetch --repo=https://github.com/CosmWasm/cw-plus --tag=v0.11.1 --release-asset="cw4_group.wasm" .

# Store code on chain
CW721_CODE=$($BINARY tx wasm store cw721_metadata_onchain.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
MARKETPLACE_CODE=$($BINARY tx wasm store sg_marketplace.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
FACTORY_CODE=$($BINARY tx wasm store collection_factory.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
SG721_CODE=$($BINARY tx wasm store sg721.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
MINTER_CODE=$($BINARY tx wasm store minter.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
WHITELIST_CODE=$($BINARY tx wasm store whitelist.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
ROYALTY_GROUP_CODE=$($BINARY tx wasm store royalty_group.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
CW4_GROUP_CODE=$($BINARY tx wasm store cw4_group.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')

# Clean up
rm collection_factory.wasm cw721_metadata_onchain.wasm sg_marketplace.wasm sg721.wasm minter.wasm cw4_group.wasm royalty_group.wasm whitelist.wasm

# Instantiate marketplace
$BINARY tx wasm instantiate $MARKETPLACE_CODE "{}" --from $1 --label "marketplace" $TXFLAG --no-admin
MARKET_CONTRACT=$($BINARY q wasm list-contract-by-code $MARKETPLACE_CODE --node $NODE --chain-id $CHAIN_ID --output json | jq -r '.contracts[-1]')

# Instantiate factory
echo $PASSWORD | $BINARY tx wasm instantiate $FACTORY_CODE "{}" --from $1 --label "factory" $TXFLAG --no-admin
FACTORY_CONTRACT=$($BINARY q wasm list-contract-by-code $FACTORY_CODE --node $NODE --chain-id $CHAIN_ID --output json | jq -r '.contracts[-1]')

# Print out Code IDs
printf "\n ------------------------ \n"
printf "Code IDs: \n\n"
echo "CW721_CODE=$CW721_CODE"
echo "MARKETPLACE_CODE=$MARKETPLACE_CODE"
echo "FACTORY_CODE=$FACTORY_CODE"
echo "SG721_CODE=$SG721_CODE"
echo "MINTER_CODE=$MINTER_CODE"
echo "WHITELIST_CODE=$WHITELIST_CODE"
echo "ROYALTY_GROUP_CODE=$ROYALTY_GROUP_CODE"
echo "CW4_GROUP_CODE=$CW4_GROUP_CODE"

# Print out Contract Addresses
printf "\n ------------------------ \n"
printf "Contracts: \n\n"
echo "MARKETPLACE_CONTRACT=$MARKET_CONTRACT"
echo "FACTORY_CONTRACT=$FACTORY_CONTRACT"
