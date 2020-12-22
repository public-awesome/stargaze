package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

// Keeper of the x/curating store
type Keeper struct {
	storeKey       sdk.StoreKey
	cdc            codec.BinaryMarshaler
	accountKeeper  types.AccountKeeper
	bankKeeper     types.BankKeeper
	stakingKeeper  types.StakingKeeper
	curatingKeeper types.CurationKeeper
	paramstore     paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, ak types.AccountKeeper,
	ck types.CurationKeeper, bk types.BankKeeper, ps paramtypes.Subspace) Keeper {

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	keeper := Keeper{
		storeKey:       key,
		cdc:            cdc,
		accountKeeper:  ak,
		bankKeeper:     bk,
		curatingKeeper: ck,
		paramstore:     ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
