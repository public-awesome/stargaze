set -ex
DENOM=uosmo
CHAINID=osmosis
RLYKEY=osmo1qk2rqkk28z8v3d7npupz33zqc6dae6n9a2x5v4
osmosisd version --long

# Setup Osmosis
osmosisd init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.osmosisd/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.osmosisd/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.osmosisd/config/app.toml
sed -i 's/enable = false/enable = true/g' ~/.osmosisd/config/app.toml
osmosisd keys --keyring-backend test add validator

osmosisd add-genesis-account $(osmosisd keys --keyring-backend test show validator -a) 100000000000$DENOM
osmosisd add-genesis-account $RLYKEY 100000000000$DENOM
osmosisd prepare-genesis mainnet $CHAINID
osmosisd gentx validator 900000000$DENOM --keyring-backend test --chain-id $CHAINID
osmosisd collect-gentxs

osmosisd start --pruning nothing
