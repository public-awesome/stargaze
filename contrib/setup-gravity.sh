#!/bin/sh

echo "Starting chain.."
killall starsd
make install
make reset
sleep 5
starsd start --log_level debug --log_format json 2> log &

sleep 5

echo "Setting up Gravity Bridge.."
[ ! -d "$HOME/.gbt" ] && gbt init

echo "Register orchestrator address.."
gbt -a stars keys register-orchestrator-address --validator-phrase "$(sed '6!d' val-phrase)" --fees=125000ustarx

sleep 5

val_key=$(starsd keys show validator --keyring-backend test -a)
echo "validator operator key: $val_key"

orch_key=$(gbt -a stars keys show 2>&1 | head -n 1 | awk '{print $7}')
echo "orchestrator delegate key: $orch_key"

eth_key=$(gbt -a stars keys show 2>&1 | sed -n '2p' | awk '{print $7}')
echo "Ethereum key: $eth_key, PLEASE FUND THIS ACCOUNT"

echo "Funding orchestrator account.."
starsd tx bank send validator $orch_key 25000000000ucredits --chain-id=localnet-1 --keyring-backend=test -y
starsd q bank balances $orch_key

#echo "Funding validator Ethereum account.."
#curl -vv -XPOST http://testnet2.althea.net/get_eth/$eth_key

# killall -HUP geth
# wget -O https://github.com/althea-net/althea-chain/raw/main/docs/althea/configs/geth-light-config.toml
# geth --syncmode "light" --goerli --http --config geth-light-config.toml


