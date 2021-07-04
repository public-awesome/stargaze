#!/bin/sh

# configure CLI
starsd config chain-id localnet-1
starsd config keyring-backend test

# see contracts code that have been uploaded
starsd q wasm list-code

# download cw20-bonding contract code
curl -LO https://github.com/CosmWasm/cosmwasm-plus/releases/download/v0.6.2/cw20_bonding.wasm

# upload contract code
starsd tx wasm store cw20_bonding.wasm --from user1 --gas-prices 0.01ustarx --gas auto --gas-adjustment 1.3 -y -b block
