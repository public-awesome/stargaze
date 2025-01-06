set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/upgrade/v1.7/config.toml ~/.hermes/
hermes keys add --chain stargaze --key-name relayer --key-file $PWD/scripts/ci/hermes/stargaze.json
hermes keys add --chain gaia --key-name relayer --key-file $PWD/scripts/ci/hermes/gaia.json
hermes keys add --chain osmosis --key-name relayer --key-file $PWD/scripts/ci/hermes/osmosis.json
hermes keys add --chain icad --key-name relayer --key-file $PWD/scripts/ci/hermes/icad.json
sleep 10
hermes tx ft-transfer --dst-chain gaia --src-chain stargaze --src-port transfer --src-channel channel-0 --amount 1000 --denom ustars --timeout-seconds 30
sleep 10
hermes tx ft-transfer --dst-chain stargaze --src-chain gaia --src-port transfer --src-channel channel-0 --amount 1000 --denom stake --timeout-seconds 30
sleep 10
hermes tx ft-transfer --dst-chain osmosis --src-chain stargaze --src-port transfer --src-channel channel-1 --amount 1000 --denom ustars --timeout-seconds 30
sleep 10
hermes tx ft-transfer --dst-chain stargaze --src-chain osmosis --src-port transfer --src-channel channel-0 --amount 1000 --denom uosmo --timeout-seconds 30
# hermes tx raw ft-transfer stargaze gaia transfer channel-0 9999 -d stake -o 1000 -n 2
# sleep 10
# hermes tx raw ft-transfer gaia stargaze transfer channel-0 9999 -d ustars -o 1000 -n 2
# sleep 10
# hermes tx raw ft-transfer stargaze osmosis transfer channel-0 9999 -d uosmo -o 1000 -n 2
# sleep 10
# hermes tx raw ft-transfer osmosis stargaze transfer channel-1 9999 -d ustars -o 1000 -n 2

sleep 30
export GAIA_ADDRESS=cosmos1wt3khka7cmn5zd592x430ph4zmlhf5gfztgha6
export STARGAZE_ADDRESS=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
export OSMOSIS_ADDRESS=osmo1qk2rqkk28z8v3d7npupz33zqc6dae6n9a2x5v4
curl -s http://gaia:1317/bank/balances/$GAIA_ADDRESS | jq '.'
curl -s http://stargaze-upgraded:1317/cosmos/bank/v1beta1/balances/$STARGAZE_ADDRESS | jq '.'
curl -s http://osmosis:1317/bank/balances/$OSMOSIS_ADDRESS | jq '.'
