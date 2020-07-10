set -ex
RELAYER_TAG="goz-phase-3"
GAIA_CHAINID=gaia
STAKEBIRD_CHAINID=stakebird

git clone https://github.com/iqlusioninc/relayer
cd relayer
make install
cd

rly version 

sleep 60
RLYKEY=integration-test
DIRECTORY=`dirname $0`
MNEMONIC=$(head -n 1 $DIRECTORY/mnemonic.txt)
rly keys restore $GAIA_CHAINID $RLYKEY "$MNEMONIC"
rly keys restore $STAKEBIRD_CHAINID $RLYKEY "$MNEMONIC"
rly cfg add-dir $DIRECTORY/chains/
cat ~/.relayer/config/config.yaml
rly lite init $GAIA_CHAINID -f
rly lite init $STAKEBIRD_CHAINID -f
rly pth gen $GAIA_CHAINID transfer $STAKEBIRD_CHAINID transfer integration-test -f
rly tx link integration-test

rly q bal $GAIA_CHAINID
rly q bal $STAKEBIRD_CHAINID


rly tx xfer $GAIA_CHAINID $STAKEBIRD_CHAINID 100stake true $(rly keys show $STAKEBIRD_CHAINID $RLYKEY)


rly q bal $GAIA_CHAINID
rly q bal $STAKEBIRD_CHAINID
