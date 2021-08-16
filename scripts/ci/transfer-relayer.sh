set -ex
mkdir -p ~/.hermes/
cp ./scripts/ci/hermes/config.toml ~/.hermes/
hermes keys add stargaze -f $PWD/scripts/ci/hermes/stargaze.json
hermes keys add gaia -f $PWD/scripts/ci/hermes/gaia.json
hermes tx raw ft-transfer stargaze gaia transfer channel-0 9999 -d stake -o 1000 -n 2
sleep 30
