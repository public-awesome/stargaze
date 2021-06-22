
#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= false
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
BUILDDIR ?= $(CURDIR)/build
FAUCET_ENABLED ?= false
DOCKER := $(shell which docker)
POST_ID ?= 1
STAKE_DENOM ?= ustarx

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

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=stargaze \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=starsd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

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
	./bin/starsd keys add validator --keyring-backend test
	./bin/starsd keys add user1 --keyring-backend test
	./bin/starsd keys add dao --keyring-backend test --pubkey starspub1addwnpepqfqzrt4mzg8e6w0ltk557ysys3dl6vf40lfkqrlpmz0m83xs5fgjy42tt8y

reset: clean create-wallet init
clean:
	rm -rf $(HOME)/.starsd/config
	rm -rf $(HOME)/.starsd/data
	rm -rf $(HOME)/.starsd/keyring-test

init:
	./bin/starsd init stargaze --stake-denom $(STAKE_DENOM) --chain-id localnet-1
	./bin/starsd add-genesis-account $(shell ./bin/starsd keys show validator -a --keyring-backend test) 10000000000000000$(STAKE_DENOM),10000000000000000ucredits
	./bin/starsd add-genesis-account $(shell ./bin/starsd keys show user1 -a --keyring-backend test) 10000000000000$(STAKE_DENOM),10000000000000ucredits,10000000000000uatom
	./bin/starsd add-genesis-account stars1czlu4tvr3dg3ksuf8zak87eafztr2u004zyh5a 300000000000000$(STAKE_DENOM)
	./bin/starsd gentx validator 10000000000$(STAKE_DENOM) --chain-id localnet-1  --keyring-backend test
	./bin/starsd collect-gentxs 

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/starsd

start:
	./bin/starsd start --grpc.address 0.0.0.0:9091

build:
	go build $(BUILD_FLAGS) -o bin/starsd ./cmd/starsd

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
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o bin/starsd github.com/public-awesome/stargaze/cmd/starsd

build-docker:
	docker build -t publicawesome/stargaze .

docker-test: build-linux
	docker build -f docker/Dockerfile.test -t rocketprotocol/stargaze-relayer-test:latest .


test:
	go test github.com/public-awesome/stargaze/x/...

fake-post:
	./bin/starsd tx curating post  1 $(POST_ID) "post body"  --from validator --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -b block -y

fake-upvote:
	./bin/starsd tx curating upvote 1 $(POST_ID) 10  --from validator --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -b block -y

fake-upvote-user:
	./bin/starsd tx curating upvote 1 $(POST_ID) 10  --from user1 --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -b block -y

fake-stake:
	./bin/starsd tx stake stake 1 $(POST_ID) 100 $(shell ./bin/starsd keys show validator --keyring-backend test --bech val --output json | jq -r '.address') --from validator --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -b block -y

fake-unstake:
	./bin/starsd tx stake unstake 1 $(POST_ID) 10  --from validator --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -b block -y

fake-create-pool1:
	./bin/starsd tx liquidity create-pool 1 100000000$(STAKE_DENOM),100000000ucredits --from validator --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -b block -y

fake-create-pool2:
	./bin/starsd tx liquidity create-pool 1 100000000$(STAKE_DENOM),100000000uatom --from user1 --keyring-backend test --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') -y

fake-swap:
	./bin/starsd tx liquidity swap 2 1 1000ustarx uatom 1.15 0.003 --from validator --chain-id $(shell ./bin/starsd status | jq -r '.NodeInfo.network') --keyring-backend test -y

.PHONY: test build-linux docker-test lint  build init install

###############################################################################
###                                Protobuf                                 ###
###############################################################################
proto-all: proto-gen proto-lint proto-check-breaking



proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace tendermintdev/docker-build-proto \
	find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;


proto-lint:
	@buf check lint --error-format=json

proto-check-breaking:
	@buf check breaking --against-input '.git#branch=master'

.PHONY: proto-all proto-gen proto-lint proto-check-breaking

ci-sign: 
	drone sign public-awesome/stargaze --save

post: 
	starsd tx curating post 1 1 "test" --from validator --keyring-backend test --chain-id localnet-1

containerProtoVer=v0.2
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=cosmos-sdk-proto-gen-$(containerProtoVer)
containerProtoGenSwagger=cosmos-sdk-proto-gen-swagger-$(containerProtoVer)
containerProtoFmt=cosmos-sdk-proto-fmt-$(containerProtoVer)


proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./contrib/protocgen.sh; fi