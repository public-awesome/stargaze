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
starsd keys --keyring-backend test add validator

starsd add-genesis-account $(starsd keys --keyring-backend test show validator -a) 100000000000$DENOM
starsd add-genesis-account $RLYKEY 100000000000$DENOM
starsd gentx validator 900000000$DENOM --keyring-backend test --chain-id stargaze
starsd collect-gentxs

starsd start --pruning nothing
