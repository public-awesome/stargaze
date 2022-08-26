set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
hermes keys add stargaze-test-1 -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia-test-1 -f $PWD/scripts/ci/hermes/gaia.json
hermes keys add osmosis-test-1 -f $PWD/scripts/ci/hermes/osmosis.json
hermes keys add icad-test-1 -f $PWD/scripts/ci/hermes/icad.json
hermes create channel stargaze-test-1 gaia-test-1 --port-a transfer --port-b transfer
hermes create channel stargaze-test-1 osmosis-test-1 --port-a transfer --port-b transfer
hermes create channel stargaze-test-1 icad-test-1 --port-a transfer --port-b transfer
