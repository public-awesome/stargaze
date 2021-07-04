#!/bin/sh

# configure CLI
starsd config chain-id localnet-1
starsd config keyring-backend test

# see contracts code that have been uploaded
starsd q wasm list-code

# download cw20-bonding contract code
curl -O https://github.com/CosmWasm/cosmwasm-plus/releases/download/v0.6.2/cw20_bonding.wasm

RES=$(starsd tx wasm store cw20_bonding.wasm --from user1 -y)

CODE_ID=$(echo $RES | jq -r '.logs[0].events[0].attributes[-1].value')
echo "uploaded code id: $CODE_ID"
