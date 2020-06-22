package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/public-awesome/stakebird/x/stake"
	"github.com/public-awesome/stakebird/x/stake/types"
	"github.com/stretchr/testify/require"
)

func TestInflateRewards(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	blockInflationAcct := app.AccountKeeper.GetModuleAccount(ctx, auth.FeeCollectorName)
	blockInflationAddr := blockInflationAcct.GetAddress()
	blockInflation := app.BankKeeper.GetBalance(ctx, blockInflationAddr, types.StakeDenom)
	require.True(t, blockInflation.Amount.IsZero())

	fakeInflationCoin := sdk.NewInt64Coin("ufuel", 1000000)
	err := app.BankKeeper.SetBalance(ctx, blockInflationAddr, fakeInflationCoin)
	app.AccountKeeper.SetModuleAccount(ctx, blockInflationAcct)
	require.NoError(t, err)
	blockInflation = app.BankKeeper.GetBalance(ctx, blockInflationAddr, types.StakeDenom)
	require.Equal(t, fakeInflationCoin, blockInflation)

	app.StakeKeeper.InflateRewardPool(ctx)

	rewardPoolAddr := app.AccountKeeper.GetModuleAccount(ctx, stake.RewardPoolName).GetAddress()
	rewardPool := app.BankKeeper.GetBalance(ctx, rewardPoolAddr, types.StakeDenom)
	require.Equal(t, "500000", rewardPool.Amount.String())
}
