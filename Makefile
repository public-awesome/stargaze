.PHONY: build proto check_go_version install
#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
POST_ID ?= 1
STAKE_DENOM ?= ustarx

export GO111MODULE = on

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

# process build tags
build_tags = netgo
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

ifeq (cleveldb,$(findstring cleveldb,$(STARGAZE_BUILD_OPTIONS)))
  build_tags += gcc
else ifeq (rocksdb,$(findstring rocksdb,$(STARGAZE_BUILD_OPTIONS)))
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
		  -X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(STARGAZE_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
else ifeq (rocksdb,$(findstring rocksdb,$(STARGAZE_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
ifeq (,$(findstring nostrip,$(STARGAZE_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(STARGAZE_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

check_go_version:
	@echo "Go version: $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION)"
ifneq ($(GO_MINOR_VERSION),22)
	@echo "ERROR: Go version 1.22 is required for this version of Stargaze"
	exit 1
endif

all: install

install: check_go_version
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/starsd

build: check_go_version
	go build $(BUILD_FLAGS) -o bin/starsd ./cmd/starsd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run --tests=false --timeout 5m0s
	@go mod verify


full-lint: lint
	@gosec -exclude-dir=cmd ./...

build-linux: 
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) -o bin/starsd github.com/public-awesome/stargaze/cmd/starsd

build-docker-arm:
	docker buildx build -t publicawesome/stargaze:local -t publicawesome/stargaze:local-dev --platform linux/arm64 --load . 

build-docker:
	docker buildx build -t publicawesome/stargaze:local -t publicawesome/stargaze:local-dev --platform linux/amd64 --load .

docker-test: build-linux
	docker build -f docker/Dockerfile.test -t rocketprotocol/stargaze-relayer-test:latest .


test:
	go test -v -race github.com/public-awesome/stargaze/v15/x/...

test-pfm:
	cd e2e && go test -v -race -run TestPacketForwardMiddleware .

test-chain-upgrade:
	cd e2e && go test -v -race -run TestChainUpgrade .

test-ica:
	cd e2e && go test -v -race -run TestInterchainAccounts .

test-chain-conformance:
	cd e2e && go test -v -race -run TestStargazeConformance .

test-slinky:
	@cd e2e/slinky && go test -v -race .

.PHONY: test test-e2e build-linux docker-test lint build install format test-slinky

format:
	gofumpt -l -w .
###############################################################################
###                                Protobuf                                 ###
###############################################################################


ci-sign: 
	drone sign public-awesome/stargaze --save

.PHONY: build-readiness-checker

build-readiness-checker:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/readiness-checker ./testutil/readiness-checker
	docker build -t publicawesome/stargaze-readiness-checker -f docker/Dockerfile.readiness .

BUF_IMAGE=bufbuild/buf@sha256:3cb1f8a4b48bd5ad8f09168f10f607ddc318af202f5c057d52a45216793d85e5 #v1.4.0
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(BUF_IMAGE)


###############################################################################
###                                Protobuf                                 ###
###############################################################################
protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@$(protoImage) buf format --write proto/

proto-lint:
	@$(protoImage) buf lint --error-format=json


drone-generate:
	drone starlark --format --target .drone.star.yml
