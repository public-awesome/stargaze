set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
hermes keys add stargaze -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia -f $PWD/scripts/ci/hermes/gaia.json
hermes keys add osmosis -f $PWD/scripts/ci/hermes/osmosis.json
hermes keys add icad -f $PWD/scripts/ci/hermes/icad.json
hermes create channel stargaze connection-0 --chain-b gaia --port-a transfer --port-b transfer --new-client-connection
hermes create channel stargaze connection-1 --chain-b osmosis --port-a transfer --port-b transfer --new-client-connection
hermes create channel stargaze connection-2 --chain-b icad --port-a transfer --port-b transfer --new-client-connection
