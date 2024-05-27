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
--chain-id stargaze -b sync --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test

sleep 5
starsd q distribution community-pool --node http://stargaze:26657

# HEIGHT=$(starsd status -o json --node http://stargaze:26657 --home $STARGAZE_HOME | jq .sync_info.latest_block_height -r)
HEIGHT=$(starsd status --node http://stargaze:26657 --home $STARGAZE_HOME | jq .SyncInfo.latest_block_height -r)

echo "current height $HEIGHT"
HEIGHT=$(expr $HEIGHT + 700) 
echo "submit with height $HEIGHT"
cat <<EOT >> proposal.json
{
  "messages": [
    {
      "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
      "authority": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz",
      "plan": {
        "name": "v14",
        "height": "$HEIGHT",
        "info": ""
      }
    }
  ],

  "deposit": "1000000000ustars",
  "title": "Upgrade",
  "summary": "Upgrade"
}
EOT
cat proposal.json
starsd tx gov submit-proposal proposal.json  \
--gas-prices 1ustars --gas auto --gas-adjustment 1.5 --from validator  \
--chain-id stargaze -b sync --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test
sleep 10
starsd q gov proposals --node http://stargaze:26657 --home $STARGAZE_HOME
starsd tx gov vote 1 "yes" --gas-prices 1ustars --gas auto --gas-adjustment 1.5 --from validator  \
--chain-id stargaze -b sync --yes --node http://stargaze:26657 --home $STARGAZE_HOME --keyring-backend test
sleep 30
starsd q gov proposals --node http://stargaze:26657 --home $STARGAZE_HOME
sleep 30
