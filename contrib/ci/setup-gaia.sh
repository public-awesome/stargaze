GAIA_TAG="goz-phase-3"

git clone https://github.com/cosmos/gaia
cd gaia 
git fetch --tags origin 
git checkout $GAIA_TAG
make install
gaiad version --long
