package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) FundDAO(ctx sdk.Context) {
	funder, err := sdk.AccAddressFromBech32(k.GetParams(ctx).Funder)
	if err != nil {
		panic(err)
	}
	availableFunds := k.bankKeeper.SpendableCoins(ctx, funder)

	if !availableFunds.Empty() {
		err = k.distKeeper.FundCommunityPool(ctx, availableFunds, funder)
		if err != nil {
			panic(err)
		}
	}
}
