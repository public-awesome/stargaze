GAIA_TAG="goz-phase-3"
DENOM=ustarx
CHAINID=stargaze
RLYKEY=stb10dmk2q0numq3v0s7vwsx20dm4hq040vslyu4hy
make install
starsd version --long



# Setup stargaze
starsd init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.starsd/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.starsd/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.starsd/config/app.toml
starsd keys --keyring-backend test add validator

starsd add-genesis-account $(starsd keys --keyring-backend test show validator -a) 100000000000$DENOM,100000000000ucredits
starsd add-genesis-account $RLYKEY 100000000000$DENOM,100000000000ucredits
starsd gentx --name validator --keyring-backend test --amount 900000000$DENOM
starsd collect-gentxs

starsd start --pruning nothing
