#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen
cd proto

proto_dirs=$(find ./publicawesome -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ..

# Fetching the stargaze version to tag in the swagger doc
stargazeTemplate="{stargaze-version}"
stargazeVersion=$(echo $(git describe --tags) | sed 's/^v//')
sed "s/$stargazeTemplate/$stargazeVersion/g" ./docs/config.json > ./tmp-swagger-gen/config.json

# Fetching the cosmos-sdk version to use the appropriate swagger file
sdkTemplate="{sdk-version}"
sdkVersion=$(go list -m -f '{{ .Version }}' github.com/cosmos/cosmos-sdk)
sed -i "s/$sdkTemplate/$sdkVersion/g" ./tmp-swagger-gen/config.json

# Fetching the wasmd version to use the appropriate swagger file
wasmdTemplate="{wasmd-version}"
wasmdVersion=$(go list -m -f '{{ .Version }}' github.com/CosmWasm/wasmd)
sed -i "s/$wasmdTemplate/$wasmdVersion/g" ./tmp-swagger-gen/config.json

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./tmp-swagger-gen/config.json -o ./docs/static/swagger.yaml -f yaml --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen