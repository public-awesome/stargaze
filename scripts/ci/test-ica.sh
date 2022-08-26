MNEMONIC=$(cat $PWD/scripts/ci/hermes/icad.json | jq -r .mnemonic)
echo "$MNEMONIC" | icad keys add ica --recover --keyring-backend test
ICA_WALLET_ADDRESS=$(icad keys show ica -a --keyring-backend test)
echo $ICA_WALLET_ADDRESS

icad config keyring-backend test
icad tx intertx register --from $ICA_WALLET_ADDRESS  --connection-id connection-0 --chain-id ica --keyring-backend test --node http://icad:26657 -b block --yes
icad query intertx interchainaccounts connection-0 $ICA_WALLET_ADDRESS --node http://icad:26657 --keyring-backend test
export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $ICA_WALLET_ADDRESS  -o json  --node http://icad:26657 | jq -r '.interchain_account_address') && echo $ICA_ADDR


starsd config keyring-backend test
echo "$MNEMONIC" | starsd keys add stargaze --recover --keyring-backend test
STARGAZE_WALLET_ADDRESS=$(starsd keys show stargaze -a --keyring-backend test)
echo $STARGAZE_WALLET_ADDRESS
