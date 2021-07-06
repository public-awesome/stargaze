#!/bin/sh

# create users
rm -rf $HOME/.starsd
starsd config chain-id localnet-1
starsd config keyring-backend test
starsd config output json
yes | starsd keys add validator
yes | starsd keys add creator
yes | starsd keys add investor
VALIDATOR=$(starsd keys show validator -a)
CREATOR=$(starsd keys show creator -a)
INVESTOR=$(starsd keys show investor -a)

# setup chain
starsd init stargaze --stake-denom ustarx --chain-id localnet-1
starsd add-genesis-account $VALIDATOR 10000000000000000ustarx
starsd add-genesis-account $CREATOR 10000000000000000ustarx
starsd add-genesis-account $INVESTOR 10000000000000000ustarx
starsd gentx validator 10000000000ustarx --chain-id localnet-1 --keyring-backend test
starsd collect-gentxs
starsd validate-genesis
starsd start
