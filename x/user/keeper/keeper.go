package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/x/user/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the x/user store
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	paramstore    paramtypes.Subspace
}

// NewKeeper creates a x/user keeper
func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper, ps paramtypes.Subspace) Keeper {

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	keeper := Keeper{
		storeKey:      key,
		cdc:           cdc,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		paramstore:    ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}
