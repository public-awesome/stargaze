set -ex
set -o pipefail
DENOM=ustars
CHAINID=stargaze
RLYKEY=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
starsd version --long
apk add -U --no-cache jq tree
STARGAZE_HOME=/stargaze/starsd
starsd config keyring-backend test --home $STARGAZE_HOME

HEIGHT=$(starsd status --node http://stargaze:26657 --home $STARGAZE_HOME | jq .SyncInfo.latest_block_height -r)

echo "current height $HEIGHT"
HEIGHT=$(expr $HEIGHT + 20) 
echo "submit with height $HEIGHT"
starsd tx gov submit-proposal software-upgrade v6 --upgrade-height $HEIGHT  \
--deposit 10000000ustars \
--description "V6 Upgrade" \
--title "V6 Upgrade" \
--gas-prices 0.025ustars --gas auto --gas-adjustment 1.5 --from validator  \
--chain-id stargaze -b block --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test

starsd q gov proposals --node http://stargaze:26657 --home $STARGAZE_HOME


starsd tx gov vote 1 "yes" --gas-prices 0.025ustars --gas auto --gas-adjustment 1.5 --from validator  \
--chain-id stargaze -b block --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test
sleep 60
starsd q gov proposals --node http://stargaze:26657 --home $STARGAZE_HOME
sleep 60
