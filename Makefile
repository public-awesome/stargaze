PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

# TODO: Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=StakeApp \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=staked \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=stakecli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

create-wallet:
	staked keys add validator --keyring-backend test

init:
	rm -rf ~/.staked/config
	staked init stakebird --chain-id localnet-1
	staked add-genesis-account $(shell staked keys show validator -a --keyring-backend test) 10000000000000000ustb,10000000000000000uatom
	staked gentx validator --chain-id localnet-1 --amount 10000000000ustb --keyring-backend test
	staked collect-gentxs 

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/staked

install-with-faucet: go.sum
	go install -mod=readonly $(BUILD_FLAGS) -tags faucet ./cmd/staked

start:
	staked start --grpc.address 0.0.0.0:9091

build:
	go build $(BUILD_FLAGS) -o bin/staked ./cmd/staked

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run --tests=false
	@go mod verify


build-linux: 
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o bin/staked github.com/public-awesome/stakebird/cmd/staked

build-docker: build-linux
	docker build -f docker/Dockerfile -t publicawesome/stakebird .

docker-test: build-linux
	docker build -f docker/Dockerfile.test -t rocketprotocol/stakebird-relayer-test:latest .


test:
	go test github.com/public-awesome/stakebird/x/...

.PHONY: test build-linux docker-test lint  build init install

###############################################################################
###                                Protobuf                                 ###
###############################################################################
proto-all: proto-gen proto-lint proto-check-breaking

proto-gen:
	@./contrib/protocgen.sh

proto-lint:
	@buf check lint --error-format=json

proto-check-breaking:
	@buf check breaking --against-input '.git#branch=master'

.PHONY: proto-all proto-gen proto-lint proto-check-breaking

ci-sign: 
	drone sign public-awesome/stakebird --save

post: 
	staked tx curating post 1 1 "test" --from validator --keyring-backend test --chain-id localnet-1
