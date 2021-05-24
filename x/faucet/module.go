package faucet

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/public-awesome/stargaze/x/faucet/client/cli"
	"github.com/public-awesome/stargaze/x/faucet/internal/keeper"
	"github.com/public-awesome/stargaze/x/faucet/internal/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic implements the sdk.AppModuleBasic interface
type AppModuleBasic struct{}

// Name returns the ModuleName
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the upgrade types on the LegacyAmino codec
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

// RegisterRESTRoutes registers all REST query handlers
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, r *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the upgrade module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if profile == TESTNET {
		//nolint:errcheck
		types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	}
}

// GetQueryCmd returns the cli query commands for this module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	if profile == TESTNET {
		return cli.GetQueryCmd()
	}
	return nil
}

// GetTxCmd returns the transaction commands for this module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	if profile == TESTNET {
		return cli.NewTxCmd()
	}
	return nil
}

// RegisterInterfaces registers the module's interface types
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// AppModule implements the sdk.AppModule interface
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

// RegisterInvariants does nothing, there are no invariants to enforce
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the staking module.
func (am AppModule) Route() sdk.Route {
	if profile == TESTNET {
		return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
	}
	return sdk.Route{}
}

// QuerierRoute returns the route we respond to for abci queries
func (AppModule) QuerierRoute() string { return types.QuerierKey }

// LegacyQuerierHandler registers a query handler to respond to the module-specific queries
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// InitGenesis is ignored, no sense in serializing future upgrades
func (am AppModule) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

// DefaultGenesis is an empty object
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return []byte("{}")
}

// ValidateGenesis is always successful, as we ignore the value
func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, config client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

// ExportGenesis is always empty, as InitGenesis does nothing either
func (am AppModule) ExportGenesis(_ sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return []byte("{}")
}

// BeginBlock returns the begin blocker for the curating module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {

}

// EndBlock does nothing
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }
