set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/v1.7/config.toml ~/.hermes/
hermes keys add --chain stargaze --key-name relayer --key-file $PWD/scripts/ci/hermes/stargaze.json
hermes keys add --chain gaia --key-name relayer --key-file $PWD/scripts/ci/hermes/gaia.json
hermes keys add --chain osmosis --key-name relayer --key-file $PWD/scripts/ci/hermes/osmosis.json
hermes keys add --chain icad --key-name relayer --key-file $PWD/scripts/ci/hermes/icad.json
hermes create connection --a-chain stargaze --b-chain gaia
hermes create connection --a-chain stargaze --b-chain osmosis
hermes create connection --a-chain stargaze --b-chain icad
hermes create channel --port-a transfer --port-b transfer --a-chain stargaze connection-0
hermes create channel --port-a transfer --port-b transfer --a-chain stargaze connection-1

