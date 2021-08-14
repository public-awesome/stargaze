set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
cat $PWD/scripts/ci/hermes/stargaze.json
hermes keys add stargaze -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia -f $PWD/scripts/ci/hermes/gaia.json
sleep 360
curl -s http://gaia:26657/status
curl -s http://stargaze:26657/status
hermes create channel stargaze gaia --port-a transfer --port-b transfer
