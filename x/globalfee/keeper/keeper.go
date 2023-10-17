package keeper

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/public-awesome/stargaze/v13/x/globalfee/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Keeper provides module state operations.
type Keeper struct {
	cdc        codec.Codec
	paramStore paramTypes.Subspace
	storeKey   storetypes.StoreKey
	wasmKeeper types.WasmKeeper
	authority  string // this should be the x/gov module account
}

// NewKeeper creates a new Keeper instance.
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, ps paramTypes.Subspace, wk types.WasmKeeper, authority string) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramStore: ps,
		wasmKeeper: wk,
		authority:  authority,
	}
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
