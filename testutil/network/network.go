package network

// COME BACK

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	tmdb "github.com/cometbft/cometbft-db"
// 	tmrand "github.com/cometbft/cometbft/libs/rand"
// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/crypto/hd"
// 	"github.com/cosmos/cosmos-sdk/crypto/keyring"
// 	servertypes "github.com/cosmos/cosmos-sdk/server/types"
// 	storetypes "cosmossdk.io/store/types"
// 	"github.com/cosmos/cosmos-sdk/testutil/network"
// 	simapp "github.com/cosmos/cosmos-sdk/testutil/sims"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

// 	"github.com/public-awesome/stargaze/v12/app"
// )

// type (
// 	Network = network.Network
// 	Config  = network.Config
// )

// // New creates instance with fully configured cosmos network.
// // Accepts optional config, that will be used in place of the DefaultConfig() if provided.
// func New(t *testing.T, configs ...network.Config) *network.Network {
// 	t.Helper()
// 	if len(configs) > 1 {
// 		panic("at most one config should be provided")
// 	}
// 	var cfg network.Config
// 	if len(configs) == 0 {
// 		cfg = DefaultConfig()
// 	} else {
// 		cfg = configs[0]
// 	}
// 	net := network.New(t, cfg)
// 	t.Cleanup(net.Cleanup)
// 	return net
// }

// // DefaultConfig will initialize config for the network with custom application,
// // genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
// func DefaultConfig() network.Config {
// 	encoding := app.MakeEncodingConfig()
// 	return network.Config{
// 		Codec:             encoding.Codec,
// 		TxConfig:          encoding.TxConfig,
// 		LegacyAmino:       encoding.Amino,
// 		InterfaceRegistry: encoding.InterfaceRegistry,
// 		AccountRetriever:  authtypes.AccountRetriever{},
// 		AppConstructor: func(val network.ValidatorI) servertypes.Application {
// 			return app.NewStargazeApp(
// 				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
// 				encoding,
// 				simapp.EmptyAppOptions{},
// 				app.EmptyWasmOpts,
// 				app.GetEnabledProposals(),
// 				baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
// 				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
// 			)
// 		},
// 		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Codec),
// 		TimeoutCommit:   2 * time.Second,
// 		ChainID:         "chain-" + tmrand.NewRand().Str(6),
// 		NumValidators:   1,
// 		BondDenom:       sdk.DefaultBondDenom,
// 		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
// 		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
// 		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
// 		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
// 		PruningStrategy: storetypes.PruningOptionNothing,
// 		CleanupDir:      true,
// 		SigningAlgo:     string(hd.Secp256k1Type),
// 		KeyringOptions:  []keyring.Option{},
// 	}
// }
