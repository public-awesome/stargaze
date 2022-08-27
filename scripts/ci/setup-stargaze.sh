set -ex
DENOM=ustars
CHAINID=stargaze
RLYKEY=stars12g0xe2ld0k5ws3h7lmxc39d4rpl3fyxp5qys69
LEDGER_ENABLED=false make install
starsd version --long



# Setup stargaze
starsd init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.starsd/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.starsd/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.starsd/config/app.toml
sed -i 's/enable = false/enable = true/g' ~/.starsd/config/app.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.starsd/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.starsd/config/config.toml
sed -i -e 's/\"allow_messages\":.*/\"allow_messages\": [\"\/cosmos.bank.v1beta1.MsgSend\", \"\/cosmos.staking.v1beta1.MsgDelegate\"]/g' ~/.starsd/config/genesis.json
starsd keys --keyring-backend test add validator

starsd add-genesis-account $(starsd keys --keyring-backend test show validator -a) 1000000000000$DENOM
starsd add-genesis-account $RLYKEY 1000000000000$DENOM
starsd add-genesis-account stars1y8tcah6r989vna00ag65xcqn6mpasjjdekwfhm 1000000000000$DENOM
starsd gentx validator 900000000$DENOM --keyring-backend test --chain-id $CHAINID
starsd collect-gentxs

starsd start --pruning nothing
