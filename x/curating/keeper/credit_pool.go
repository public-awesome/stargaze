package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// InitializeCreditPool sets up the credit pool from genesis
func (k Keeper) InitializeCreditPool(ctx sdk.Context, funds sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.CreditPoolName, sdk.NewCoins(funds))
}

// GetCreditPoolBalance returns the curation credit pool balance
func (k Keeper) GetCreditPoolBalance(ctx sdk.Context) sdk.Coin {
	return k.bankKeeper.GetBalance(ctx, k.GetCreditPool(ctx).GetAddress(), k.GetParams(ctx).CreditDenom)
}

// GetCreditPool returns the credit pool account
func (k Keeper) GetCreditPool(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.CreditPoolName)
}

// DistributeCredits distributes credits to all accounts
func (k Keeper) DistributeCredits(ctx sdk.Context) error {
	blocksPerYear := k.mintKeeper.GetParams(ctx).BlocksPerYear
	blocksPerDay := int64(blocksPerYear / 365.0)
	// week := blocksPerDay % 52
	// if it has been 24 hours...
	if (ctx.BlockHeight() % blocksPerDay) == 0 {
		k.accountKeeper.IterateAccounts(ctx, func(a authtypes.AccountI) (stop bool) {
			return true
		})
	}

	return nil
}
