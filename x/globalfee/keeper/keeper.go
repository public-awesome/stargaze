package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/log"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v15/internal/collcompat"

	corestoretypes "cosmossdk.io/core/store"
	"github.com/public-awesome/stargaze/v15/x/globalfee/types"
)

// Keeper provides module state operations.
type Keeper struct {
	cdc                    codec.Codec
	storeService           corestoretypes.KVStoreService
	wasmKeeper             types.WasmKeeper
	Schema                 collections.Schema
	Params                 collections.Item[types.Params]
	CodeAuthorizations     collections.Map[uint64, types.CodeAuthorization]
	ContractAuthorizations collections.Map[[]byte, types.ContractAuthorization]
	authority              string // this should be the x/gov module account
}

// NewKeeper creates a new Keeper instance.
func NewKeeper(
	cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
	wk types.WasmKeeper,
	authority string,
) Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	keeper := Keeper{
		cdc:          cdc,
		storeService: storeService,
		wasmKeeper:   wk,
		authority:    authority,
		Params: collections.NewItem(
			sb,
			types.ParamsKey,
			"params",
			collcompat.ProtoValue[types.Params](cdc),
		),
		CodeAuthorizations: collections.NewMap(
			sb,
			types.CodeAuthorizationPrefix,
			"codeAuthorizations",
			collections.Uint64Key,
			collcompat.ProtoValue[types.CodeAuthorization](cdc),
		),
		ContractAuthorizations: collections.NewMap(
			sb,
			types.ContractAuthorizationPrefix,
			"contractAuthorizations",
			collections.BytesKey,
			collcompat.ProtoValue[types.ContractAuthorization](cdc),
		),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) GetContractInfo(ctx sdk.Context, contractAddr sdk.AccAddress) *wasmtypes.ContractInfo {
	return k.wasmKeeper.GetContractInfo(ctx, contractAddr)
}

// GetAuthority returns the x/wasm module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
