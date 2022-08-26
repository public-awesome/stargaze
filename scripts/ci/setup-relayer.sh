set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
hermes keys add stargaze -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia -f $PWD/scripts/ci/hermes/gaia.json
hermes keys add osmosis -f $PWD/scripts/ci/hermes/osmosis.json
hermes keys add icad -f $PWD/scripts/ci/hermes/icad.json
hermes create connection stargaze gaia
hermes create connection stargaze osmosis
hermes create connection icad stargaze 
hermes create channel --port-a transfer --port-b transfer stargaze connection-0
hermes create channel --port-a transfer --port-b transfer stargaze connection-1
hermes create channel --port-a transfer --port-b transfer stargaze connection-2
hermes create channel --port-a icacontroller --port-b icahost icad connection-0
