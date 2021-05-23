package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/curating/types"
)

// InitGenesis initializes the module state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	amount := sdk.NewCoins(sdk.NewInt64Coin("ustarx", 5_000_000_000))
	sender, err := sdk.AccAddressFromBech32("stars1czlu4tvr3dg3ksuf8zak87eafztr2u004zyh5a")
	if err != nil {
		panic(err)
	}

	k.distKeeper.FundCommunityPool(ctx, amount, sender)
}
