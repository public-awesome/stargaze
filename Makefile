
#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= false
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
BUILDDIR ?= $(CURDIR)/build
FAUCET_ENABLED ?= false

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(FAUCET_ENABLED),true)
 build_tags += faucet
endif

ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(GAIA_BUILD_OPTIONS)))
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=stakebird \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=staked \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/tendermint/tendermint/version.TMCoreSemVer=0.34.0-rc6

ifeq (cleveldb,$(findstring cleveldb,$(GAIA_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (,$(findstring nostrip,$(GAIA_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(GAIA_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif


all: install

create-wallet:
	./bin/staked keys add validator --keyring-backend test

reset: clean init
clean:
	rm -rf ~/.staked/config
	rm -rf ~/.staked/data

init:
	./bin/staked init stakebird --stake-denom uegg --chain-id localnet-1
	./bin/staked add-genesis-account $(shell ./bin/staked keys show validator -a --keyring-backend test) 10000000000000000uegg,10000000000000000ucredits
	./bin/staked gentx validator --chain-id localnet-1 --amount 10000000000uegg --keyring-backend test
	./bin/staked collect-gentxs 

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/staked

start:
	./bin/staked start --grpc.address 0.0.0.0:9091

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
	@golangci-lint run --tests=false --skip-dirs="simapp"
	@go mod verify


build-linux: 
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o bin/staked github.com/public-awesome/stakebird/cmd/staked

build-docker: build-linux
	docker build -f docker/Dockerfile -t publicawesome/stakebird .

docker-test: build-linux
	docker build -f docker/Dockerfile.test -t rocketprotocol/stakebird-relayer-test:latest .


test:
	go test github.com/public-awesome/stakebird/x/...

fake-post:
	./bin/staked tx curating post  1 $(POST_ID) "post body"  --from validator --keyring-backend test --chain-id $(shell ./bin/staked status | jq '.node_info.network') -b block -y

fake-upvote:
	./bin/staked tx curating upvote 1 $(POST_ID) 1  --from validator --keyring-backend test --chain-id $(shell ./bin/staked status | jq '.node_info.network') -b block -y

fake-stake:
	./bin/staked tx stake stake 1 $(POST_ID) 100 $(VAL) --from validator --keyring-backend test --chain-id $(shell ./bin/staked status | jq '.node_info.network') -b block -y

fake-unstake:
	./bin/staked tx stake unstake 1 $(POST_ID) 10  --from validator --keyring-backend test --chain-id $(shell ./bin/staked status | jq '.node_info.network') -b block -y

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
