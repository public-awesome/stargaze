PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

# TODO: Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NewApp \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=rocketd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=rocketcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

create-wallet:
	bin/rocketcli keys add validator --keyring-backend test

init:
	rm -rf ~/.rocketd
	bin/rocketd init rocket
	bin/rocketd add-genesis-account $(shell bin/rocketcli keys show validator -a --keyring-backend test) 10000000000ufuel --keyring-backend test
	bin/rocketd gentx --name=validator --amount 10000000000ufuel --keyring-backend test
	bin/rocketd collect-gentxs 

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/rocketd
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/rocketcli

build:
		go build -o bin/rocketd ./cmd/rocketd
		go build -o bin/rocketcli ./cmd/rocketcli

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
	GOARCH=amd64 GOOS=linux go build -o bin/rocketd github.com/rocket-protocol/stakebird/cmd/rocketd
	GOARCH=amd64 GOOS=linux  go build -o bin/rocketcli github.com/rocket-protocol/stakebird/cmd/rocketcli

docker-test: build-linux
	docker build -f docker/Dockerfile.test -t rocketprotocol/rocketzone-relayer-test:latest .

###############################################################################
###                                Protobuf                                 ###
###############################################################################
proto-all: proto-gen proto-lint proto-check-breaking

proto-gen:
	@./scripts/protocgen.sh

proto-lint:
	@buf check lint --error-format=json

proto-check-breaking:
	@buf check breaking --against-input '.git#branch=master'

.PHONY: proto-all proto-gen proto-lint proto-check-breaking
