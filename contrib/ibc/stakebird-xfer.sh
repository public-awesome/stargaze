#!/bin/bash

if [[ ! -x "$(which jq)" ]]; then
  echo "please install jq: https://stedolan.github.io/jq/"
  exit 1
fi

RELAYER_TAG="goz-phase-3"

# start gaia + stakebird
./contrib/ibc/gaia-stakebird.sh
sleep 20

set -e

echo "Downloading and installing relayer tag $RELAYER_TAG..."
(cd /tmp; \
curl -OL https://github.com/iqlusioninc/relayer/archive/$RELAYER_TAG.zip &> /dev/null; \
unzip -o $RELAYER_TAG.zip &> /dev/null; \
cd relayer-$RELAYER_TAG; \
make install &> /dev/null)

# rm -rf ~/.relayer/
rly cfg init

# NOTE: you may want to look at the config between these steps to see
# what is added in each step. The config is located at ~/.relayer/config/config.yaml
cat ~/.relayer/config/config.yaml

rly cfg add-dir contrib/ibc/relayer/configs/stakebird-xfer/

# NOTE: you may want to look at the config between these steps
cat ~/.relayer/config/config.yaml

rly keys restore ibc0 testkey "$(jq -r '.secret' data/ibc0/n0/gaiacli/key_seed.json)"
rly keys restore ibc1 testkey "$(jq -r '.secret' data/ibc1/n0/stakecli/key_seed.json)"

rly lite init ibc0 -f
rly lite init ibc1 -f

# At this point the relayer --home directory is ready for normal operations between
# ibc0 and ibc1. Looking at the folder structure of the relayer at this point is helpful
tree ~/.relayer

# Now you can connect the two chains with one command:
# rly -d tx link demo

# Check the token balances on both chains
# rly q bal ibc0
# rly q bal ibc1

# Then send some tokens between the chains
# rly tx xfer ibc0 ibc1 100000stake true $(rly keys show ibc1 testkey)

# See that the transfer has completed
# rly q bal ibc0
# rly q bal ibc1
