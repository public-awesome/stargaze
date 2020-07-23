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
	stakecli config keyring-backend test
	stakecli config chain-id localnet-1
	stakecli keys add validator

init:
	staked init stakebird --chain-id localnet-1
	staked add-genesis-account $(shell bin/stakecli keys show validator -a) 10000000000000000ufuel --keyring-backend test
	staked gentx --name=validator --amount 10000000000ufuel --keyring-backend test
	staked collect-gentxs 

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/staked
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/stakecli

start:
	staked start --pruning=nothing --log_level "curating:info,funding:info,main:info,state:info,*:error"

build:
	go build $(BUILD_FLAGS) -o bin/staked ./cmd/staked
	go build $(BUILD_FLAGS) -o bin/stakecli ./cmd/stakecli

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify


build-linux: 
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o bin/staked github.com/public-awesome/stakebird/cmd/staked
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o bin/stakecli github.com/public-awesome/stakebird/cmd/stakecli

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
