#!/bin/sh

CHAINID=${CHAINID:-testing}
DENOM=${DENOM:-ustars}

# TODO: Allow for multiple genesis account addrs to be passed in

# Build genesis file incl account for passed address
coins="10000000000000000$DENOM"
starsd init --chain-id $CHAINID $CHAINID
starsd keys add validator --keyring-backend="test"
starsd add-genesis-account validator $coins --keyring-backend="test"
starsd add-genesis-account $1 $coins --keyring-backend="test"
starsd gentx validator 10000000000$DENOM --chain-id $CHAINID --keyring-backend="test"
starsd collect-gentxs


# Set proper defaults and change ports
sed -i 's/"leveldb"/"goleveldb"/g' ~/.starsd/config/config.toml
sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.starsd/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.starsd/config/genesis.json
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.starsd/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.starsd/config/config.toml
sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.starsd/config/config.toml

# Start the stake
starsd start --pruning=nothing
