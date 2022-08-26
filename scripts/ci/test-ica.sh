MNEMONIC=$(cat hermes/icad.json | jq -r .mnemonic)
echo "$MNEMONIC" | icad keys add ica --recover --keyring-backend test
WALLET_ADDRESS=$(icad keys show ica -a --keyring-backend test)
icad tx intertx register --from $WALLET_ADDRESS  -connection-id connection-0 --chain-id ica --keyring-backend test
icad query intertx interchainaccounts connection-0 $WALLET_ADDRESS
