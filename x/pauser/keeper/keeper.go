package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	corestoretypes "cosmossdk.io/core/store"
	"github.com/public-awesome/stargaze/v18/internal/collcompat"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
)

// Keeper provides module state operations.
type Keeper struct {
	cdc             codec.Codec
	storeService    corestoretypes.KVStoreService
	wasmKeeper      types.WasmKeeper
	Schema          collections.Schema
	Params          collections.Item[types.Params]
	PausedContracts collections.Map[[]byte, types.PausedContract]
	PausedCodeIDs   collections.Map[uint64, types.PausedCodeID]
	authority       string // this should be the x/gov module account
}

// NewKeeper creates a new Keeper instance.
// Note: wasmKeeper is NOT passed at construction to break circular dependency.
// Use SetWasmKeeper() after WasmKeeper is initialized.
func NewKeeper(
	cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
	authority string,
) Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	keeper := Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		Params: collections.NewItem(
			sb,
			types.ParamsKey,
			"params",
			collcompat.ProtoValue[types.Params](cdc),
		),
		PausedContracts: collections.NewMap(
			sb,
			types.PausedContractPrefix,
			"pausedContracts",
			collections.BytesKey,
			collcompat.ProtoValue[types.PausedContract](cdc),
		),
		PausedCodeIDs: collections.NewMap(
			sb,
			types.PausedCodeIDPrefix,
			"pausedCodeIDs",
			collections.Uint64Key,
			collcompat.ProtoValue[types.PausedCodeID](cdc),
		),
	}
	return keeper
}

// SetWasmKeeper sets the wasm keeper after initialization to break circular dependency.
func (k *Keeper) SetWasmKeeper(wk types.WasmKeeper) {
	k.wasmKeeper = wk
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// IsExecutionPaused checks if execution is paused for the given contract address.
// It checks both direct contract pause and code ID pause.
func (k Keeper) IsExecutionPaused(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
	// 1. Check if the contract is directly paused
	if k.IsContractPaused(ctx, contractAddr) {
		return true
	}

	// 2. Check if the contract's code ID is paused
	if k.wasmKeeper == nil {
		return false
	}
	contractInfo := k.wasmKeeper.GetContractInfo(ctx, contractAddr)
	if contractInfo == nil {
		return false
	}
	return k.IsCodeIDPaused(ctx, contractInfo.CodeID)
}
