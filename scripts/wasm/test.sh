#!/bin/sh

TXFLAG="--gas-prices 0.01ustarx --gas auto --gas-adjustment 1.3 -y -b block"

# configure CLI
starsd config chain-id localnet-1
starsd config keyring-backend test

# see contracts code that have been uploaded
starsd q wasm list-code

# download cw20-bonding contract code
curl -LO https://github.com/CosmWasm/cosmwasm-plus/releases/download/v0.6.2/cw20_bonding.wasm

# upload contract code
starsd tx wasm store cw20_bonding.wasm --from user1 $TXFLAG

# instantiate contract
INIT='{
  "name": "sirbobo",
  "symbol": "BOBO",
  "decimals": 8,
  "reserve_denom": "ustarx",
  "reserve_decimals": 6,
  "curve_type": { "square_root": { "slope": "1", "scale": 1 } }
}'
starsd tx wasm instantiate 1 "$INIT" --from user1 --label "social token" $TXFLAG

# query contract
starsd q wasm list-contract-by-code 1 --output json
CONTRACT=$(starsd q wasm list-contract-by-code 1 --output json | jq -r '.contracts[-1]')
starsd q wasm contract $CONTRACT
starsd q wasm contract-state all $CONTRACT

# execute contract
