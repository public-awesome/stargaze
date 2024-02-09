set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/upgrade/v1.7/config.toml ~/.hermes/
hermes keys add --chain stargaze --key-name relayer --key-file $PWD/scripts/ci/hermes/stargaze.json
hermes keys add --chain gaia --key-name relayer --key-file $PWD/scripts/ci/hermes/gaia.json
hermes keys add --chain osmosis --key-name relayer --key-file $PWD/scripts/ci/hermes/osmosis.json
hermes keys add --chain icad --key-name relayer --key-file $PWD/scripts/ci/hermes/icad.json
hermes start

