package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/x/stake/types"
)

// Keeper of the x/curating store
type Keeper struct {
	storeKey       sdk.StoreKey
	cdc            codec.BinaryCodec
	stakingKeeper  types.StakingKeeper
	curatingKeeper types.CurationKeeper
	bankKeeper     types.BankKeeper
	paramstore     paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	ck types.CurationKeeper,
	sk types.StakingKeeper,
	bk types.BankKeeper,
	ps paramtypes.Subspace) Keeper {

	keeper := Keeper{
		storeKey:       key,
		cdc:            cdc,
		stakingKeeper:  sk,
		curatingKeeper: ck,
		bankKeeper:     bk,
		paramstore:     ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
