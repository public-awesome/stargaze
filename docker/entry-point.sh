#!/bin/sh

CHAINID=${CHAINID:-testing}
DENOM=${DENOM:-ustars}
BLOCK_GAS_LIMIT=${GAS_LIMIT:-75000000}

IAVL_CACHE_SIZE=${IAVL_CACHE_SIZE:-1562500}
QUERY_GAS_LIMIT=${QUERY_GAS_LIMIT:-5000000}
SIMULATION_GAS_LIMIT=${SIMULATION_GAS_LIMIT:-50000000}
MEMORY_CACHE_SIZE=${MEMORY_CACHE_SIZE:-1000}

NATS_URL=${NATS_URL}

# Build genesis file incl account for each address passed in
coins="10000000000000000$DENOM"
starsd init --chain-id $CHAINID $CHAINID
starsd keys add validator --keyring-backend="test"
starsd add-genesis-account validator $coins --keyring-backend="test"

# create account for each passed in address
for addr in "$@"; do
  echo "creating genesis account: $addr"
  starsd add-genesis-account $addr $coins --keyring-backend="test"
done

starsd gentx validator 10000000000$DENOM --chain-id $CHAINID --keyring-backend="test"
starsd collect-gentxs


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

# if natsurl is set, then configure nats
if [ -z "$NATS_URL" ]; then
  echo "NATS_URL not set, not configuring nats"
else
  sed -i 's/streamers = \[\]/streamers = ["nats"]/g' ~/.starsd/config/app.toml
  sed -i 's/\[streamers.file\]/[streamers.nats]/g' ~/.starsd/config/app.toml
  sed -i 's/keys = \["\*", \]/keys = ["*"]/g' ~/.starsd/config/app.toml
  sed -i "s/write_dir = \"\"/url = [ \"$NATS_URL\" ]/g" ~/.starsd/config/app.toml
  sed -i 's/prefix = ""//' ~/.starsd/config/app.toml
fi

# Start the stake
starsd start --pruning=nothing
