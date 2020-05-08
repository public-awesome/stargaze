package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/rocket-protocol/stakebird/x/stake/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the x/stake store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler
	// stakingKeeper types.StakingKeeper
	stakingKeeper stakingkeeper.Keeper
	paramspace    types.ParamSubspace
}

// NewKeeper creates a x/stake keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, stakingKeeper stakingkeeper.Keeper,
	paramspace types.ParamSubspace) Keeper {

	keeper := Keeper{
		storeKey:      key,
		cdc:           cdc,
		stakingKeeper: stakingKeeper,
		paramspace:    paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

// func (k Keeper) set(ctx sdk.Context, key string, value /* TODO: fill out this type */ ) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
// 	store.Set([]byte(key), bz)
// }

// func (k Keeper) delete(ctx sdk.Context, key string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Delete([]byte(key))
// }
