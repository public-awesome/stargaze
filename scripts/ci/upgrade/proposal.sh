set -ex
set -o pipefail
DENOM=ustars
CHAINID=stargaze
RLYKEY=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
starsd version --long
apk add -U --no-cache jq tree
STARGAZE_HOME=/stargaze/starsd
starsd config keyring-backend test --home $STARGAZE_HOME


echo "fund community pool"
starsd tx distribution fund-community-pool 20000000000000ustars  \
--gas-prices 1ustars --gas auto --gas-adjustment 1.5 --from funder  \
--chain-id stargaze -b block --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test

sleep 5
starsd q distribution community-pool --node http://stargaze:26657

HEIGHT=$(starsd status --node http://stargaze:26657 --home $STARGAZE_HOME | jq .SyncInfo.latest_block_height -r)

echo "current height $HEIGHT"
HEIGHT=$(expr $HEIGHT + 20) 
echo "submit with height $HEIGHT"
starsd tx gov submit-proposal software-upgrade v11 --upgrade-height $HEIGHT  \
--deposit 1000000000ustars \
--description "v11 Upgrade" \
--title "v11 Upgrade" \
--gas-prices 1ustars --gas auto --gas-adjustment 1.5 --from validator  \
--chain-id stargaze -b block --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test

starsd q gov proposals --node http://stargaze:26657 --home $STARGAZE_HOME


starsd tx gov vote 1 "yes" --gas-prices 1ustars --gas auto --gas-adjustment 1.5 --from validator  \
--chain-id stargaze -b block --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test
sleep 60
starsd q gov proposals --node http://stargaze:26657 --home $STARGAZE_HOME
sleep 60
