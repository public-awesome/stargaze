#!/usr/bin/env bash

set -eo pipefail

SWAGGER_DIR=./swagger-proto
mkdir -p "$SWAGGER_DIR"
THIRD_PARTY_DIR="$SWAGGER_DIR/third_party"
mkdir -p "$THIRD_PARTY_DIR"
buf export buf.build/cosmos/cosmos-sdk:v0.47.0 -o "$THIRD_PARTY_DIR"

mkdir -p "$SWAGGER_DIR/proto"
cp -r ./proto/osmosis "$SWAGGER_DIR/proto"
cp -r ./proto/publicawesome "$SWAGGER_DIR/proto"
printf "version: v1\ndirectories:\n  - proto\n  - third_party" > "$SWAGGER_DIR/buf.work.yaml"
printf "version: v1\nname: buf.build/public-awesome/stargaze\n" > "$SWAGGER_DIR/proto/buf.yaml"
cp ./proto/buf.gen.swagger.yaml "$SWAGGER_DIR/proto/buf.gen.swagger.yaml"
mkdir -p ./tmp-swagger-gen
cd "$SWAGGER_DIR"


# create swagger files on an individual basis  w/ `buf build` and `buf generate` (needed for `swagger-combine`)
proto_dirs=$(find ./proto ./third_party -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ -n "$query_file" ]]; then
    buf generate --template proto/buf.gen.swagger.yaml "$query_file"
  fi
done

cd ..
swagger-combine ./docs/config.json -o ./docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true
# clean swagger files
rm -rf ./tmp-swagger-gen
rm -rf "$SWAGGER_DIR"