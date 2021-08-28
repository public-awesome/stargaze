#!/bin/sh

TXFLAG="--gas-prices 0.01ustarx --gas auto --gas-adjustment 1.3 -y -b block"

CREATOR=$(starsd keys show creator -a)
INVESTOR=$(starsd keys show investor -a)

# see contracts code that have been uploaded
starsd q wasm list-code

curl -LO https://github.com/CosmWasm/cosmwasm-plus/releases/download/v0.6.2/cw721_base.wasm

# upload contract code
starsd tx wasm store cw721_base.wasm --from validator $TXFLAG

# instantiate contract
INIT="{
  \"name\": \"usernames\",
  \"symbol\": \"USER\",
  \"minter\": \"$CREATOR\"
}"
starsd tx wasm instantiate 1 "$INIT" --from creator --label "register username" $TXFLAG

# get contract address
starsd q wasm list-contract-by-code 1 --output json
CONTRACT=$(starsd q wasm list-contract-by-code 1 --output json | jq -r '.contracts[-1]')

# query contract
starsd q wasm contract-state smart $CONTRACT '{"num_tokens":{}}'
starsd q wasm contract-state smart $CONTRACT '{"contract_info":{}}'

# execute a mint operation
MINT="{\"mint\":{\"token_id\":\"1\",\"owner\":\"$INVESTOR\",\"name\":\"bobo\"}}"
starsd tx wasm execute $CONTRACT $MINT --from creator $TXFLAG

# query contract
starsd q wasm contract-state smart $CONTRACT '{"owner_of":{"token_id":"1"}}'
starsd q wasm contract-state smart $CONTRACT '{"num_tokens":{}}'
starsd q wasm contract-state smart $CONTRACT '{"nft_info":{"token_id":"1"}}'
starsd q wasm contract-state smart $CONTRACT "{\"tokens\":{\"owner\":\"$INVESTOR\"}}"

# execute an NFT transfer
TRANSFER="{\"transfer_nft\":{\"recipient\":\"$CREATOR\",\"token_id\":\"1\"}}"
starsd tx wasm execute $CONTRACT $TRANSFER --from investor $TXFLAG

# owner should change from investor to creator
starsd q wasm contract-state smart $CONTRACT '{"owner_of":{"token_id":"1"}}'
