set -ex
cp ./scripts/ci/hermes/config.toml ~/.hermes/
hermes tx raw ft-transfer stargaze gaia transfer channel-0 9999 -d stake -o 1000 -n 2
