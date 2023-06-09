package globalfee

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/public-awesome/stargaze/v11/x/globalfee/client/cli"
	"github.com/public-awesome/stargaze/v11/x/globalfee/keeper"
	"github.com/public-awesome/stargaze/v11/x/globalfee/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.AppModule      = AppModule{}
)

// AppModuleBasic defines the basic application module for this module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the module's name.
func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the module's types on the given LegacyAmino codec.
func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(amino)
}

// RegisterInterfaces registers the module's interface types.
func (a AppModuleBasic) RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns default genesis state as raw bytes for the module.
func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the module.
func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var state types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &state); err != nil {
		return fmt.Errorf("failed to unmarshal x/%s genesis state: %w", types.ModuleName, err)
	}

	return state.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, serveMux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), serveMux, types.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

// GetTxCmd returns the root tx command for the module.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns no root query command for the module.
func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// AppModule implements an application module for this module.
type AppModule struct {
	AppModuleBasic

	_      codec.Codec
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// RegisterInvariants registers the module invariants.
func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the module.
// Deprecated.
func (AppModule) Route() sdk.Route { return sdk.Route{} }

// QuerierRoute returns the module's querier route name.
func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// LegacyQuerierHandler returns the staking module sdk.Querier.
func (a AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers the module services.
func (a AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(a.keeper))
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(a.keeper))
}

// InitGenesis performs genesis initialization for the module. It returns no validator updates.
func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(bz, &genesisState)

	InitGenesis(ctx, a.keeper, genesisState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the module.
func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	state := ExportGenesis(ctx, a.keeper)
	return cdc.MustMarshalJSON(state)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (a AppModule) ConsensusVersion() uint64 {
	return 1
}

// BeginBlock returns the begin blocker for the module.
func (a AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the module. It returns no validator updates.
func (a AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the module.
func (a AppModule) GenerateGenesisState(_ *module.SimulationState) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (a AppModule) ProposalContents(_ module.SimulationState) []simTypes.WeightedProposalContent {
	return []simTypes.WeightedProposalContent{}
}

// RandomizedParams creates randomized param changes for the simulator.
func (a AppModule) RandomizedParams(_ *rand.Rand) []simTypes.ParamChange {
	return []simTypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder for the module's types.
func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the module operations with their respective weights.
func (a AppModule) WeightedOperations(_ module.SimulationState) []simTypes.WeightedOperation {
	return []simTypes.WeightedOperation{}
}
