module github.com/public-awesome/stargaze

go 1.15

require (
	github.com/CosmWasm/wasmd v0.15.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coinbase/rosetta-sdk-go v0.5.9 // indirect
	github.com/cosmos/cosmos-sdk v0.42.6
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect
	github.com/hdevalence/ed25519consensus v0.0.0-20210204194344-59a8610d2b87 // indirect
	github.com/improbable-eng/grpc-web v0.14.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/cosmos-rosetta-gateway v0.3.0-rc2.0.20210304154332-87d6ca4410df // indirect
	github.com/tendermint/liquidity v0.1.0-rc0.0.20210206010652-2ca7f61d0f15
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
