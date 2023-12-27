set -ex
set -o pipefail
DENOM=ustars
CHAINID=stargaze
RLYKEY=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
starsd version --long
apk add -U --no-cache jq tree
STARGAZE_HOME=/stargaze/starsd

# Setup stargaze
starsd init --chain-id $CHAINID $CHAINID --home $STARGAZE_HOME
cat $STARGAZE_HOME/config/genesis.json | jq '.app_state["interchainaccounts"]["host_genesis_state"]["params"]["allow_messages"]=["*"]' > $STARGAZE_HOME/config/tmp.json && mv $STARGAZE_HOME/config/tmp.json $STARGAZE_HOME/config/genesis.json
starsd config keyring-backend test --home $STARGAZE_HOME
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' $STARGAZE_HOME/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" $STARGAZE_HOME/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' $STARGAZE_HOME/config/app.toml
sed -i 's/enable = false/enable = true/g' $STARGAZE_HOME/config/app.toml
sed -i 's/localhost:9090/0.0.0.0:9090/g' $STARGAZE_HOME/config/app.toml
sed -i 's/172800s/60s/g'  $STARGAZE_HOME/config/genesis.json
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "100ms"/g'  $STARGAZE_HOME/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "100ms"/g'  $STARGAZE_HOME/config/config.toml
starsd keys --keyring-backend test add validator --home $STARGAZE_HOME
starsd keys --keyring-backend test add funder --home $STARGAZE_HOME
starsd add-genesis-account $(starsd keys --keyring-backend test show validator -a --home $STARGAZE_HOME) 10000000000000$DENOM --home $STARGAZE_HOME
starsd add-genesis-account $RLYKEY 10000000000000$DENOM --home $STARGAZE_HOME
starsd add-genesis-account stars1y8tcah6r989vna00ag65xcqn6mpasjjdekwfhm 1000000000000$DENOM --home $STARGAZE_HOME
starsd add-genesis-account stars103y4f6h80lc45nr8chuzr3fyzqywm9n0gnr394 20000000000000$DENOM  --home $STARGAZE_HOME
starsd add-genesis-account $(starsd keys --keyring-backend test show funder -a --home $STARGAZE_HOME) 21000000000000$DENOM --home $STARGAZE_HOME
starsd gentx validator 900000000$DENOM --keyring-backend test --chain-id stargaze --home $STARGAZE_HOME
starsd collect-gentxs --home $STARGAZE_HOME
/stargaze/bin/upgrade-watcher starsd start --pruning nothing --home $STARGAZE_HOME --grpc.address 0.0.0.0:9090 --rpc.laddr tcp://0.0.0.0:26657
