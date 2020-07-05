package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/funding/types"
)

// Keeper of the funding store
type Keeper struct {
	storeKey           sdk.StoreKey
	cdc                codec.Marshaler
	BankKeeper         types.BankKeeper
	ChannelKeeper      types.ChannelKeeper
	DistributionKeeper types.DistributionKeeper
	paramspace         types.ParamSubspace
}

// NewKeeper creates a funding keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper, distKeeper types.DistributionKeeper, paramspace types.ParamSubspace) Keeper {

	keeper := Keeper{
		storeKey:           key,
		cdc:                cdc,
		BankKeeper:         bankKeeper,
		ChannelKeeper:      channelKeeper,
		DistributionKeeper: distKeeper,
		paramspace:         paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
