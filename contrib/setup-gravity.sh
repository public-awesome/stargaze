#!/bin/sh
set -eux

BIN=starsd

CHAINID="gravity-test"

# setup chain
coins="10000000000ustarx,100000000000ufoo"
$BIN init --chain-id $CHAINID $CHAINID

ARGS="--keyring-backend test"

# generate validator, orchestrator, and eth keys for validator
$BIN keys add val $ARGS 2>> ~/.$BIN/val-phrase
$BIN keys add orch $ARGS 2>> ~/.$BIN/orch-phrase
$BIN eth_keys add >> ~/.$BIN/val-eth-key

# set genesis accounts
VAL_KEY=$($BIN keys show val -a $ARGS)
ORCH_KEY=$($BIN keys show orch -a $ARGS)
$BIN add-genesis-account $VAL_KEY $coins
$BIN add-genesis-account $ORCH_KEY $coins

# generate genesis txns
ETH_KEY=$(grep address ~/.$BIN/val-eth-key | sed -n 1p | sed 's/.*://')
$BIN gentx val 500000000ustarx $ETH_KEY $ORCH_KEY --chain-id=$CHAINID $ARGS

$BIN collect-gentxs
echo "Collected gentx"

# # Set proper defaults and change ports
# sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.$BIN/config/config.toml
# sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.$BIN/config/config.toml
# sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.$BIN/config/config.toml
# sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.$BIN/config/config.toml

# # Start the chain
# $BIN start --pruning=nothing