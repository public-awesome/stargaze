package mint

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/mint/keeper"
	"github.com/public-awesome/stargaze/v17/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(goCtx context.Context, k keeper.Keeper) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter, err := k.GetMinter(ctx)
	if err != nil {
		panic(err)
	}
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}

	// recalculate annual provision
	minter.AnnualProvisions = minter.NextAnnualProvisions(ctx.BlockTime(), params)
	err = k.SetMinter(ctx, minter)
	if err != nil {
		panic(err)
	}

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err = k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
