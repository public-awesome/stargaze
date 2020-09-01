package curating_test

import (
	"testing"

	"github.com/public-awesome/stakebird/x/curating"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

var addrs = []sdk.AccAddress{}

func setup(t *testing.T) (*testdata.SimApp, sdk.Context) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(1)

	deposit := sdk.NewInt64Coin("ustb", 1_000_000)
	addrs = testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10_000_000))

	err := app.CuratingKeeper.CreatePost(
		ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[0])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "9000000", creatorBal.Amount.String())

	// curator1
	err = app.CuratingKeeper.CreateUpvote(
		ctx, vendorID, postID, addrs[1], addrs[1], 1, deposit)
	require.NoError(t, err)
	_, found, err = app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[1])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")
	curator1Bal := app.BankKeeper.GetBalance(ctx, addrs[1], "uatom")
	require.Equal(t, "9000000", curator1Bal.Amount.String(),
		"10 (initial bal) - 1 (upvote)")

	// curator2
	err = app.CuratingKeeper.CreateUpvote(
		ctx, vendorID, postID, addrs[2], addrs[2], 3, deposit)
	require.NoError(t, err)
	_, found, err = app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[2])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")
	curator2Bal := app.BankKeeper.GetBalance(ctx, addrs[2], "uatom")
	require.Equal(t, "1000000", curator2Bal.Amount.String(),
		"10 (initial bal) - 9 (upvote)")

	// fast-forward blocktime to simulate end of curation window
	h := ctx.BlockHeader()
	h.Time = ctx.BlockHeader().Time.Add(
		app.CuratingKeeper.GetParams(ctx).CurationWindow)
	ctx = ctx.WithBlockHeader(h)

	return app, ctx
}

// initial state
// creator  = 10 STB
// curator1 = 10 STB, upvote 1 STB
// curator2 = 10 STB, upvote 9 STB
//
// qvf
// voting_pool  = 10 STB
// root_sum     = 4
// match_pool   = 4^2 - 10 = 6
// voter_reward = 5 STB
// match_reward = match_pool / 2 = 3
func TestEndBlockerExpiringPost(t *testing.T) {
	app, ctx := setup(t)

	// add funds to reward pool
	funds := sdk.NewInt64Coin("ustb", 10_000_000_000)
	err := app.BankKeeper.MintCoins(ctx, curating.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	curating.EndBlocker(ctx, app.CuratingKeeper)

	// creator match reward = 0.5 * match_reward = 3 STB
	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "13000000", creatorBal.Amount.String(),
		"10 (initial) + 3 (creator match reward)")

	// curator1Bal := app.BankKeeper.GetBalance(ctx, addrs[1], "ustb")
	// require.Equal(t, "15500000", curator1Bal.Amount.String(),
	// "8 (bal) + 1 (deposit) + 5 (voting reward) + 1.5 (match reward)")

	// curator2Bal := app.BankKeeper.GetBalance(ctx, addrs[2], "ustb")
	// require.Equal(t, "7500000", curator2Bal.Amount.String(),
	// "0 (bal) + 1 (deposit) + 5 (voter reward) + 1.5 (match reward)")
}

func TestEndBlockerExpiringPostWithSmolRewardPool(t *testing.T) {
	app, ctx := setup(t)

	// add funds to reward pool
	funds := sdk.NewInt64Coin("ustb", 1_000_000)
	err := app.BankKeeper.MintCoins(ctx, curating.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	curating.EndBlocker(ctx, app.CuratingKeeper)

	// creator match reward = 0.5 * match_reward = 3 STB
	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "10000500", creatorBal.Amount.String(),
		"10 (initial) + 3 (creator match reward)")

	// curator1Bal := app.BankKeeper.GetBalance(ctx, addrs[1], "ustb")
	// require.Equal(t, "14000249", curator1Bal.Amount.String(),
	// "8 (bal) + 1 (deposit) + 5 (voting reward) + 1.5 (match reward)")

	// curator2Bal := app.BankKeeper.GetBalance(ctx, addrs[2], "ustb")
	// require.Equal(t, "6000249", curator2Bal.Amount.String(),
	// "0 (bal) + 1 (deposit) + 5 (voter reward) + 1.5 (match reward)")
}
