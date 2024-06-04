#!/bin/sh

CHAINID=${CHAINID:-testing}
DENOM=${DENOM:-ustars}
BLOCK_GAS_LIMIT=${GAS_LIMIT:-75000000}

IAVL_CACHE_SIZE=${IAVL_CACHE_SIZE:-1562500}
QUERY_GAS_LIMIT=${QUERY_GAS_LIMIT:-5000000}
SIMULATION_GAS_LIMIT=${SIMULATION_GAS_LIMIT:-50000000}
MEMORY_CACHE_SIZE=${MEMORY_CACHE_SIZE:-1000}

# Build genesis file incl account for each address passed in
coins="10000000000000000$DENOM"
starsd init --chain-id $CHAINID $CHAINID
starsd keys add validator --keyring-backend="test"
starsd genesis add-genesis-account validator $coins --keyring-backend="test"

# create account for each passed in address
for addr in "$@"; do
  echo "creating genesis account: $addr"
  starsd genesis add-genesis-account $addr $coins --keyring-backend="test"
done

starsd genesis gentx validator 10000000000$DENOM --chain-id $CHAINID --keyring-backend="test"
starsd genesis collect-gentxs


# Set proper defaults and change ports
sed -i 's/"leveldb"/"goleveldb"/g' ~/.starsd/config/config.toml
sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.starsd/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.starsd/config/genesis.json
sed -i "s/\"max_gas\": \"-1\"/\"max_gas\": \"$BLOCK_GAS_LIMIT\"/" ~/.starsd/config/genesis.json
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.starsd/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.starsd/config/config.toml
sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.starsd/config/config.toml

sed -i "s/iavl-cache-size = 781250/iavl-cache-size = $IAVL_CACHE_SIZE/g" ~/.starsd/config/app.toml
sed -i "s/query_gas_limit = 50000000/query_gas_limit = $QUERY_GAS_LIMIT/g" ~/.starsd/config/app.toml
sed -i "s/simulation_gas_limit = 25000000/simulation_gas_limit = $SIMULATION_GAS_LIMIT/g" ~/.starsd/config/app.toml
sed -i "s/memory_cache_size = 512/memory_cache_size = $MEMORY_CACHE_SIZE/g" ~/.starsd/config/app.toml
sed -i "s/enable = false/enable = true/g" ~/.starsd/config/app.toml
sed -i "s/localhost:9090/0.0.0.0:9090/g" ~/.starsd/config/app.toml

# Start the stake
starsd start --pruning=nothing
