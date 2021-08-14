set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
cat $PWD/scripts/ci/hermes/stargaze.json
hermes keys add stargaze -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia -f $PWD/scripts/ci/hermes/gaia.json
sleep 300
hermes create channel stargaze gaia --port-a transfer --port-b transfer
