set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
hermes keys add stargaze -f ./scripts/ci/hermes/stargaze.json
hermes keys add gaia -f ./scripts/ci/hermes/gaia.json
hermes create channel stargaze gaia --port-a transfer --port-b transfer
