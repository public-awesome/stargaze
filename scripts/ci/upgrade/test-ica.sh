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
icad tx intertx register --from $ICA_WALLET_ADDRESS --connection-id connection-0 --chain-id $ICA_CHAIN_ID --keyring-backend test --node http://icad:26657 -b block --yes
sleep 30
icad query intertx interchainaccounts connection-0 $ICA_WALLET_ADDRESS --node http://icad:26657
export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $ICA_WALLET_ADDRESS  -o json  --node http://icad:26657 | jq -r '.interchain_account_address') && echo $ICA_ADDR


starsd config keyring-backend test
starsd config chain-id stargaze
starsd config node http://stargaze-upgraded:26657
starsd status
echo "$MNEMONIC" | starsd keys add ica-test --recover --keyring-backend test
STARGAZE_WALLET_ADDRESS=$(starsd keys show ica-test -a --keyring-backend test)
echo $STARGAZE_WALLET_ADDRESS
starsd q bank balances $ICA_ADDR 
starsd tx bank send ica-test $ICA_ADDR 100000000ustars --chain-id stargaze -y --from ica-test
starsd q bank balances $ICA_ADDR 

VALIDATOR=$(starsd q  staking validators --limit 1 -o json | jq '.validators[0].operator_address' -r)
echo "Delegate to validator $VALIDATOR"

TX_MSG=$(cat <<EOF
{
    "@type":"/cosmos.bank.v1beta1.MsgSend",
    "from_address":"$ICA_ADDR",
    "to_address":"stars1ly5qeh4xjept0udwny9edwzgw95qmvekms3na8",
    "amount": [
        {
            "denom": "ustars",
            "amount": "1000"
        }
    ]
}
EOF
)
echo "$TX_MSG" > send.json
starsd q bank balances stars1ly5qeh4xjept0udwny9edwzgw95qmvekms3na8
# Submit a bank send tx using the interchain account via ibc
icad tx intertx submit send.json --connection-id connection-0 --from $ICA_WALLET_ADDRESS --chain-id icad -y -b block
sleep 20
starsd q bank balances stars1ly5qeh4xjept0udwny9edwzgw95qmvekms3na8



DELEGATE_MSG=$(cat <<EOF
{
    "@type":"/cosmos.staking.v1beta1.MsgDelegate",
    "delegator_address":"$ICA_ADDR",
    "validator_address":"$VALIDATOR",
    "amount": {
        "denom": "ustars",
        "amount": "1000"
    }
}
EOF
)
echo $DELEGATE_MSG > delegate.json
# Submit a staking delegation tx using the interchain account via ibc
icad tx intertx submit delegate.json --connection-id connection-0 --from $ICA_WALLET_ADDRESS --chain-id icad -y -b block --timeout-height 5000
sleep 30
starsd query staking delegations $ICA_ADDR
