#!/bin/bash

## CONFIG
BINARY="${BINARY:-starsd}"
DENOM="${DENOM:-ustars}"
CHAIN_ID="${CHAIN_ID:-double-double-1}"
NODE="${NODE:-https://rpc.double-double-1.stargaze-apis.com:443}"

if [ "$1" = "" ]
then
  echo "Usage: $0 1 arg required, address or key name."
  exit
fi

TXFLAG="--gas-prices 0.01$DENOM --gas auto --gas-adjustment 1.3 -y -b block --chain-id $CHAIN_ID --node $NODE --output json"

CONTRACTS_REPO=https://github.com/public-awesome/stargaze-contracts
CONTRACTS_TAG=v0.12.4-alpha
MARKETPLACE_REPO=https://github.com/public-awesome/marketplace
MARKETPLACE_TAG=v0.5.1

if ! command -v fetch &> /dev/null
then
    echo "fetch could not be found: https://github.com/gruntwork-io/fetch"
    exit 1
fi

# Fetch contract wasm binaries
fetch --repo=$MARKETPLACE_REPO --tag=$MARKETPLACE_TAG --release-asset="sg_marketplace.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="sg721.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="minter.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="whitelist.wasm" .
fetch --repo=$CONTRACTS_REPO --tag=$CONTRACTS_TAG --release-asset="claim.wasm" .
# Store code on chain
MARKETPLACE_CODE=$($BINARY tx wasm store sg_marketplace.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
SG721_CODE=$($BINARY tx wasm store sg721.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
MINTER_CODE=$($BINARY tx wasm store minter.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
WHITELIST_CODE=$($BINARY tx wasm store whitelist.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
CLAIM_CODE=$($BINARY tx wasm store claim.wasm --from $1 $TXFLAG | jq -r '.logs[0].events[-1].attributes[0].value')
# Clean up
rm sg_marketplace.wasm sg721.wasm minter.wasm whitelist.wasm claim.wasm

# Instantiate marketplace
$BINARY tx wasm instantiate $MARKETPLACE_CODE '{"trading_fee_percent": 2, "ask_expiry": [86400,15552000], "bid_expiry": [86400,15552000], "operators": ["'$($BINARY keys show -a $1)'"]}' --from $1 --label "marketplace" $TXFLAG --no-admin
MARKET_CONTRACT=$($BINARY q wasm list-contract-by-code $MARKETPLACE_CODE --node $NODE --chain-id $CHAIN_ID --output json | jq -r '.contracts[-1]')

# Instantiate claim contract
$BINARY tx wasm instantiate $CLAIM_CODE "{\"marketplace_addr\":\"$MARKET_CONTRACT\"}" --from $1 --label "claim" $TXFLAG --no-admin
CLAIM_CONTRACT=$($BINARY q wasm list-contract-by-code $CLAIM_CODE --node $NODE --chain-id $CHAIN_ID --output json | jq -r '.contracts[-1]')

# Print out Code IDs
printf "\n ------------------------ \n"
printf "Code IDs: \n\n"
echo "MARKETPLACE_CODE=$MARKETPLACE_CODE"
echo "SG721_CODE=$SG721_CODE"
echo "MINTER_CODE=$MINTER_CODE"
echo "WHITELIST_CODE=$WHITELIST_CODE"
echo "CLAIM_CODE=$CLAIM_CODE"

# Print out Contract Addresses
printf "\n ------------------------ \n"
printf "Contracts: \n\n"
echo "MARKETPLACE_CONTRACT=$MARKET_CONTRACT"
echo "CLAIM_CONTRACT=$CLAIM_CONTRACT"
