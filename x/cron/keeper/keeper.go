package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v15/internal/collcompat"
	"github.com/public-awesome/stargaze/v15/x/cron/types"
)

type (
	Keeper struct {
		cdc                 codec.BinaryCodec
		storeService        corestoretypes.KVStoreService
		wasmKeeper          types.WasmKeeper
		Schema              collections.Schema
		Params              collections.Item[types.Params]
		PrivilegedContracts collections.Map[[]byte, []byte]
		authority           string // this should be the x/gov module account
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
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
		PrivilegedContracts: collections.NewMap(
			sb,
			types.PrivilegedContractsPrefix,
			"privilegedContracts",
			collections.BytesKey,
			collections.BytesValue,
		),
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	keeper.Schema = schema
	return keeper
}

// GetAuthority returns the x/wasm module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ModuleLogger(ctx)
}

func ModuleLogger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
