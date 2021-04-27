package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/public-awesome/stargaze/simapp"
	"github.com/public-awesome/stargaze/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestInflateRewards(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	blockInflationAcct := app.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
	blockInflationAddr := blockInflationAcct.GetAddress()
	blockInflation := app.BankKeeper.GetBalance(ctx, blockInflationAddr, types.DefaultStakeDenom)
	require.True(t, blockInflation.Amount.IsZero())

	fakeInflationCoin := sdk.NewInt64Coin("ustb", 1000000)
	err := simapp.FundAccount(app, ctx, blockInflationAddr, sdk.NewCoins(fakeInflationCoin))
	app.AccountKeeper.SetModuleAccount(ctx, blockInflationAcct)
	require.NoError(t, err)
	blockInflation = app.BankKeeper.GetBalance(ctx, blockInflationAddr, types.DefaultStakeDenom)
	require.Equal(t, fakeInflationCoin, blockInflation)

	err = app.CuratingKeeper.InflateRewardPool(ctx)
	require.NoError(t, err)

	rewardPoolAddr := app.AccountKeeper.GetModuleAccount(ctx, types.RewardPoolName).GetAddress()
	rewardPool := app.BankKeeper.GetBalance(ctx, rewardPoolAddr, types.DefaultStakeDenom)
	require.Equal(t, "21000000500000", rewardPool.Amount.String())
}

func TestInflateRewardsNonDefault(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	blockInflationAcct := app.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)
	blockInflationAddr := blockInflationAcct.GetAddress()
	blockInflation := app.BankKeeper.GetBalance(ctx, blockInflationAddr, fakedenom)
	require.True(t, blockInflation.Amount.IsZero())

	fakeInflationCoin := sdk.NewInt64Coin(fakedenom, 1000000)
	err := simapp.FundAccount(app, ctx, blockInflationAddr, sdk.NewCoins(fakeInflationCoin))
	app.AccountKeeper.SetModuleAccount(ctx, blockInflationAcct)
	require.NoError(t, err)
	blockInflation = app.BankKeeper.GetBalance(ctx, blockInflationAddr, fakedenom)
	require.Equal(t, fakeInflationCoin, blockInflation)

	err = app.CuratingKeeper.InflateRewardPool(ctx)
	require.NoError(t, err)

	rewardPoolAddr := app.AccountKeeper.GetModuleAccount(ctx, types.RewardPoolName).GetAddress()

	rewardPool := app.BankKeeper.GetBalance(ctx, rewardPoolAddr, fakedenom)
	require.Equal(t, "21000000500000", rewardPool.Amount.String())
}
