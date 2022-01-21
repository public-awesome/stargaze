#!/bin/sh

TXFLAG="--gas-prices 0.01ustars --gas auto --gas-adjustment 1.3 -y -b block"

CREATOR=$(starsd keys show creator -a)
INVESTOR=$(starsd keys show investor -a)

# see contracts code that have been uploaded
starsd q wasm list-code

# download cw20-bonding contract code
curl -LO https://github.com/CosmWasm/cosmwasm-plus/releases/download/v0.9.0/cw20_bonding.wasm

# upload contract code
starsd tx wasm store cw20_bonding.wasm --from validator $TXFLAG

# instantiate contract
INIT='{
  "name": "sirbobo",
  "symbol": "BOBO",
  "decimals": 2,
  "reserve_denom": "ustars",
  "reserve_decimals": 8,
  "curve_type": { "linear": { "slope": "1", "scale": 1 } }
}'
starsd tx wasm instantiate 1 "$INIT" --from creator --label "social token" $TXFLAG

# get contract address
starsd q wasm list-contract-by-code 1 --output json
CONTRACT=$(starsd q wasm list-contract-by-code 1 --output json | jq -r '.contracts[-1]')

# query contract
starsd q wasm contract-state smart $CONTRACT '{"token_info":{}}'
starsd q wasm contract-state smart $CONTRACT '{"curve_info":{}}'
starsd q wasm contract-state smart $CONTRACT "{\"balance\":{\"address\":\"$INVESTOR\"}}"

# execute a buy order
BUY='{"buy":{}}'
starsd tx wasm execute $CONTRACT $BUY --from investor --amount=500000000ustars $TXFLAG

# check balances
starsd q bank balances $INVESTOR
starsd q wasm contract-state smart $CONTRACT "{\"balance\":{\"address\":\"$INVESTOR\"}}"

# execute a burn / sell order
SELL='{"burn":{"amount":"500"}}'
starsd tx wasm execute $CONTRACT $SELL --from investor $TXFLAG
starsd q wasm contract-state smart $CONTRACT "{\"balance\":{\"address\":\"$INVESTOR\"}}"
starsd q bank balances $INVESTOR

