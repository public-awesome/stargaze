set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/upgrade/config.toml ~/.hermes/
hermes keys add stargaze -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia -f $PWD/scripts/ci/hermes/gaia.json
hermes keys add osmosis -f $PWD/scripts/ci/hermes/osmosis.json
hermes tx raw ft-transfer stargaze gaia transfer channel-0 9999 -d stake -o 1000 -n 2
hermes tx raw ft-transfer gaia stargaze transfer channel-0 9999 -d ustars -o 1000 -n 2
sleep 10
hermes tx raw ft-transfer stargaze osmosis transfer channel-0 9999 -d uosmo -o 1000 -n 2
sleep 10
hermes tx raw ft-transfer osmosis stargaze transfer channel-1 9999 -d ustars -o 1000 -n 2

sleep 30
export GAIA_ADDRESS=cosmos1wt3khka7cmn5zd592x430ph4zmlhf5gfztgha6
export STARGAZE_ADDRESS=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
export OSMOSIS_ADDRESS=osmo1qk2rqkk28z8v3d7npupz33zqc6dae6n9a2x5v4
curl -s http://gaia:1317/bank/balances/$GAIA_ADDRESS | jq '.'
curl -s http://stargaze-upgraded:1317/bank/balances/$STARGAZE_ADDRESS | jq '.'
curl -s http://osmosis:1317/bank/balances/$OSMOSIS_ADDRESS | jq '.'
