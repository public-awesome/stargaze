set -ex
set -o pipefail
DENOM=ustars
CHAINID=stargaze
RLYKEY=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
starsd version --long
apk add -U --no-cache jq tree curl wget
STARGAZE_HOME=/stargaze/starsd
curl -s -v http://stargaze:8090/kill || echo "done"
sleep 10
sed -i 's/enable = false/enable = true/g' $STARGAZE_HOME/config/app.toml
sed -i 's/localhost:9090/0.0.0.0:9090/g' $STARGAZE_HOME/config/app.toml
sed -i 's/localhost:1317/0.0.0.0:1317/g' $STARGAZE_HOME/config/app.toml
cat $STARGAZE_HOME/config/app.toml | grep -A 10  grpc
cat $STARGAZE_HOME/config/app.toml | grep -A 10  api
starsd start --pruning nothing --home $STARGAZE_HOME --grpc.address 0.0.0.0:9090 --rpc.laddr tcp://0.0.0.0:26657 --skip-preferred-settings
