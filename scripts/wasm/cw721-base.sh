#!/bin/sh

TXFLAG="--gas-prices 0.01ustars --gas auto --gas-adjustment 1.3 -y -b block"

CREATOR=$(starsd keys show creator -a)
INVESTOR=$(starsd keys show investor -a)

# see contracts code that have been uploaded
starsd q wasm list-code

curl -LO https://github.com/CosmWasm/cosmwasm-plus/releases/download/v0.9.0/cw721_base.wasm

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
MINT1="{\"mint\":{\"token_id\":\"1\",\"owner\":\"$INVESTOR\",\"name\":\"bobo1\",\"description\":\"Ape556\",\"image\":\"https://lh3.googleusercontent.com/Xv5i9YaUJO73claDwpJ-cqkx6-xhrrZnF9QAD_Qn4aPJHSxklQp4anlJuV-GABs4ZkB1lIsmkodUT3V8ER0wivNrwBswUHZzrpVOaQ=h1328\"}}"
MINT2="{\"mint\":{\"token_id\":\"2\",\"owner\":\"$INVESTOR\",\"name\":\"bobo2\",\"description\":\"Ape557\",\"image\":\"https://images.squarespace-cdn.com/content/v1/5c12933f365f02733c923e4e/1623457826739-RFS8YBP06I1W5WW2CSCG/tyler-hobbs-fidenza-612.png?format=750w\"}}"
MINT3="{\"mint\":{\"token_id\":\"3\",\"owner\":\"$INVESTOR\",\"name\":\"bobo3\",\"description\":\"Ape558\",\"image\":\"https://miro.medium.com/max/1400/1*Fg9-AmnE5X_929CqE1xp9w.png\"}}"
MINT4="{\"mint\":{\"token_id\":\"4\",\"owner\":\"$INVESTOR\",\"name\":\"bobo4\",\"description\":\"Ape559\",\"image\":\"https://lh3.googleusercontent.com/Zau-70Ga57u021g4xxx9UqHyiwwpxuFI-W1q0BWetxhmhm8_rTERCPsCfQled_nxBDIN40U7x1hDX3CvVkMeLe4Pxg=w600\"}}"
MINT5="{\"mint\":{\"token_id\":\"5\",\"owner\":\"$INVESTOR\",\"name\":\"bobo5\",\"description\":\"Ape554\",\"image\":\"https://lh3.googleusercontent.com/sjXMZ1J39uU7UlpXjm8TmGclwNbC6YiITFZXAmj0brUE9Yw_cLKeyFPTMRur2etRGfifYdIFdDpN7qZ9Vn9I3GTfUTP8DX0S05mv8w=w600\"}}"
starsd tx wasm execute $CONTRACT $MINT1 --from creator $TXFLAG
starsd tx wasm execute $CONTRACT $MINT2 --from creator $TXFLAG
starsd tx wasm execute $CONTRACT $MINT3 --from creator $TXFLAG
starsd tx wasm execute $CONTRACT $MINT4 --from creator $TXFLAG
starsd tx wasm execute $CONTRACT $MINT5 --from creator $TXFLAG

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
