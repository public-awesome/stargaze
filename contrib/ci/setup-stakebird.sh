GAIA_TAG="goz-phase-3"
DENOM=ustb
CHAINID=stakebird
RLYKEY=stake10dmk2q0numq3v0s7vwsx20dm4hq040vsawpj35
make install
staked version --long



# Setup stakebird
staked init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.staked/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.staked/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.staked/config/app.toml
stakecli keys --keyring-backend test add validator

staked add-genesis-account $(stakecli keys --keyring-backend test show validator -a) 100000000000$DENOM
staked add-genesis-account $RLYKEY 100000000000$DENOM
staked gentx --name validator --keyring-backend test --amount 900000000$DENOM
staked collect-gentxs

staked start --pruning nothing
