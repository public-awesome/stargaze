package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

func TestDelegation(t *testing.T) {
	_, app, ctx := createTestInput()

	// create fake addresses
	delAddrs := AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10000))
	valAddrs := ConvertAddrsToValAddrs(delAddrs)

	// create validator with 50% commission
	commission := staking.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	msg := staking.NewMsgCreateValidator(
		valAddrs[0], valConsPk1, sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)), staking.Description{}, commission, sdk.OneInt(),
	)
	sh := staking.NewHandler(app.StakingKeeper)
	res, err := sh(ctx, msg)
	require.NoError(t, err)
	require.NotNil(t, res)
	// end block to bond validator
	staking.EndBlocker(ctx, app.StakingKeeper)
	// next block
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	// historical count should be 2 (once for validator init, once for delegation init)
	require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))

	// perform end-user delegation
	vendorID := uint64(1)
	postID := uint64(1)
	votingPeriod := time.Hour * 24 * 7
	amount := sdk.NewInt64Coin("ufuel", 10000)
	err = app.StakeKeeper.Delegate(ctx, vendorID, postID, delAddrs[0], valAddrs[0], votingPeriod, amount)
	require.NoError(t, err)

	// check if delegation is stored in staking store
	delegations := app.StakingKeeper.GetAllDelegations(ctx)
	require.Len(t, delegations, 1)

	// check if delegation is in voting delegation queue
	// endTime := ctx.BlockTime().Add(votingPeriod * 5 * time.Hour) // after block time
	endTime := ctx.BlockTime().Add(votingPeriod * -5 * time.Hour) // before block time

	app.StakeKeeper.IterateVotingDelegationQueue(ctx, endTime, func(vendorID, postID uint64, delegation stakingtypes.Delegation) bool {
		require.True(t, delegation.GetShares().Equal(amount.Amount.ToDec()))
		return true
	})

	// TODO: test end blocker
}

func TestABCI(t *testing.T) {
	// // check if delegation is in voting queue
	// endTime := ctx.BlockTime()
	// app.StakeKeeper.IterateVotingDelegationQueue(ctx, endTime, func(vendorID, postID uint64, delegation stakingtypes.Delegation) bool {
	// 	spew.Dump(delegation)
	// 	require.True(t, delegation.GetShares().Equal(amount.Amount.ToDec()))
	// 	return true
	// })
	// require.Fail(t, "failureMessage strings")
}
