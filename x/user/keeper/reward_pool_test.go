package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/public-awesome/stakebird/x/user"
	"github.com/public-awesome/stakebird/x/user/types"
	"github.com/stretchr/testify/require"
)

func TestInflateRewards(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	blockInflationAcct := app.AccountKeeper.GetModuleAccount(ctx, auth.FeeCollectorName)
	blockInflationAddr := blockInflationAcct.GetAddress()
	blockInflation := app.BankKeeper.GetBalance(ctx, blockInflationAddr, types.DefaultStakeDenom)
	require.True(t, blockInflation.Amount.IsZero())

	fakeInflationCoin := sdk.NewInt64Coin("ustb", 1000000)
	err := app.BankKeeper.SetBalance(ctx, blockInflationAddr, fakeInflationCoin)
	app.AccountKeeper.SetModuleAccount(ctx, blockInflationAcct)
	require.NoError(t, err)
	blockInflation = app.BankKeeper.GetBalance(ctx, blockInflationAddr, types.DefaultStakeDenom)
	require.Equal(t, fakeInflationCoin, blockInflation)

	err = app.userKeeper.InflateRewardPool(ctx)
	require.NoError(t, err)

	rewardPoolAddr := app.AccountKeeper.GetModuleAccount(ctx, user.RewardPoolName).GetAddress()
	rewardPool := app.BankKeeper.GetBalance(ctx, rewardPoolAddr, types.DefaultStakeDenom)
	require.Equal(t, "500000", rewardPool.Amount.String())
}
