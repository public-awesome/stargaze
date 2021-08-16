set -ex
DENOM=stake
CHAINID=gaia
RLYKEY=cosmos1wt3khka7cmn5zd592x430ph4zmlhf5gfztgha6
gaiad version --long

# Setup gaia
gaiad init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.gaia/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.gaia/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.gaia/config/app.toml
sed -i 's/enable = false/enable = true/g' ~/.gaia/config/app.toml
gaiad keys --keyring-backend test add validator

gaiad add-genesis-account $(gaiad keys --keyring-backend test show validator -a) 100000000000$DENOM
gaiad add-genesis-account $RLYKEY 100000000000$DENOM
gaiad gentx validator 900000000$DENOM --keyring-backend test --chain-id gaia
gaiad collect-gentxs

gaiad start --pruning nothing
