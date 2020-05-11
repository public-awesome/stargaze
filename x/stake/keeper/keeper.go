package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/rocket-protocol/stakebird/x/stake/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the x/stake store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.Marshaler
	sk         types.StakingKeeper
	paramstore paramtypes.Subspace
}

// NewKeeper creates a x/stake keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, stakingKeeper types.StakingKeeper,
	ps paramtypes.Subspace) Keeper {

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		sk:         stakingKeeper,
		paramstore: ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}
