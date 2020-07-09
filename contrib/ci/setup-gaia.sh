GAIA_TAG="goz-phase-3"
DENOM=stake
CHAINID=gaia

# install gaia
git clone https://github.com/cosmos/gaia
cd gaia 
git fetch --tags origin 
git checkout $GAIA_TAG
make install
gaiad version --long



# Setup gaia
gaiad init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.gaiad/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.gaiad/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.gaiad/config/app.toml
gaiacli keys --keyring-backend test add validator

gaiad add-genesis-account $(gaiacli keys --keyring-backend test show validator -a) 100000000000$DENOM
gaiad gentx --name validator --keyring-backend test --amount 900000000$DENOM
gaiad collect-gentxs
