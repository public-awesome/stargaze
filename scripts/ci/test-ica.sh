set -ex
MNEMONIC=$(cat $PWD/scripts/ci/hermes/ica-test.json | jq -r .mnemonic)
echo "$MNEMONIC" | icad keys add ica-test --recover --keyring-backend test
ICA_WALLET_ADDRESS=$(icad keys show ica-test -a --keyring-backend test)
echo $ICA_WALLET_ADDRESS

export ICA_CHAIN_ID=icad
icad config keyring-backend test
icad config chain-id $ICA_CHAIN_ID
icad config node http://icad:26657
icad status 
icad query account $ICA_WALLET_ADDRESS --node http://icad:26657
icad q bank balances $ICA_WALLET_ADDRESS --node http://icad:26657
icad tx intertx register --from $ICA_WALLET_ADDRESS --connection-id connection-0 --chain-id $ICA_CHAIN_ID --keyring-backend test  --generate-only >  tx.json  
cat tx.json | jq
icad tx intertx register --from $ICA_WALLET_ADDRESS --connection-id connection-0 --chain-id $ICA_CHAIN_ID --keyring-backend test --node http://icad:26657 -b block --yes
sleep 15
icad query intertx interchainaccounts connection-0 $ICA_WALLET_ADDRESS --node http://icad:26657
export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $ICA_WALLET_ADDRESS  -o json  --node http://icad:26657 | jq -r '.interchain_account_address') && echo $ICA_ADDR


starsd config keyring-backend test
starsd config chain-id stargaze
starsd config node http://stargaze:26657
starsd status
echo "$MNEMONIC" | starsd keys add ica-test --recover --keyring-backend test
STARGAZE_WALLET_ADDRESS=$(starsd keys show ica-test -a --keyring-backend test)
echo $STARGAZE_WALLET_ADDRESS
