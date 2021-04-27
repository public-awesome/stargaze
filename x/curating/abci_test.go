package curating_test

import (
	"testing"
	"time"

	"github.com/public-awesome/stargaze/x/curating"
	"github.com/public-awesome/stargaze/x/curating/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/public-awesome/stargaze/simapp"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var addrs = []sdk.AccAddress{}

func setup(t *testing.T) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	vendorID := uint32(1)

	addrs = simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10_000_000))

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	_, err = app.CuratingKeeper.CreatePost(
		ctx,
		vendorID,
		&postID,
		bodyHash,
		body,
		addrs[0],
		addrs[0],
		"chain id",
		nil,
		"metadata",
		nil,
	)
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetAllBalances(ctx, addrs[0])
	require.Equal(t, "10000000", creatorBal.AmountOf("ucredits").String())
	require.Equal(t, "10000000", creatorBal.AmountOf("ustb").String())

	// curator1
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[1], addrs[1], 1)
	require.NoError(t, err)
	_, found, err = app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[1])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")
	curator1Bal := app.BankKeeper.GetBalance(ctx, addrs[1], "ucredits")
	require.Equal(t, "9000000", curator1Bal.Amount.String(),
		"10 (initial bal) - 1 (upvote)")

	// curator2
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[2], addrs[2], 3)
	require.NoError(t, err)
	_, found, err = app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[2])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")
	curator2Bal := app.BankKeeper.GetBalance(ctx, addrs[2], "ucredits")
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
// creator  = 10 credits, 10 stars
// curator1 = 10 credits, 10 stars, upvote 1 credits
// curator2 = 10 credits, 10 stars, upvote 9 credits
//
// qvf
// voting_pool  = 10 credits
// root_sum     = 4
// match_pool   = 4^2 - 10 = 6
//
// match_reward_per_vote = match_pool / 4 = 1.5 stars
// curator 1 match reward = 1.5 stars
// curator 2 match reward = 4.5 stars
func TestEndBlockerExpiringPost(t *testing.T) {
	app, ctx := setup(t)

	// add funds to reward pool
	funds := sdk.NewInt64Coin("ustb", 10_000_000_000)
	err := app.BankKeeper.MintCoins(ctx, types.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	curating.EndBlocker(ctx, app.CuratingKeeper)

	creatorBal := app.BankKeeper.GetAllBalances(ctx, addrs[0])
	require.Equal(t, "10000000", creatorBal.AmountOf("ucredits").String(),
		"10 (bal)")

	require.Equal(t, "10000000", creatorBal.AmountOf("ustb").String(),
		"10 (bal)")

	curator1Bal := app.BankKeeper.GetAllBalances(ctx, addrs[1])
	require.Equal(t, "9000000", curator1Bal.AmountOf("ucredits").String(),
		"9 (bal)")
	require.Equal(t, "11500000", curator1Bal.AmountOf("ustb").String(),
		"9 (bal) + 1 (deposit) + 1.5 (match reward)")

	curator2Bal := app.BankKeeper.GetAllBalances(ctx, addrs[2])
	require.Equal(t, "1000000", curator2Bal.AmountOf("ucredits").String(),
		"1 (bal)")
	require.Equal(t, "14500000", curator2Bal.AmountOf("ustb").String(),
		"9 (bal) + 1 (deposit) + 4.5 (match reward)")
}

func TestEndBlockerExpiringPostWithSmolRewardPool(t *testing.T) {
	app, ctx := setup(t)

	// burn existing funds from the reward pool to reset it
	funds := app.CuratingKeeper.GetRewardPoolBalance(ctx)
	err := app.BankKeeper.BurnCoins(ctx, types.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	// add funds to reward pool
	funds = sdk.NewInt64Coin("ustb", 1_000_000)
	err = app.BankKeeper.MintCoins(ctx, types.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	curating.EndBlocker(ctx, app.CuratingKeeper)

	// creator match reward = 0.05 * match_reward = 3 STB
	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "10000000", creatorBal.Amount.String(),
		"10 (initial)")

	curator1Bal := app.BankKeeper.GetAllBalances(ctx, addrs[1])
	require.Equal(t, "9000000", curator1Bal.AmountOf("ucredits").String(),
		"9 (bal)")
	require.Equal(t, "10000250", curator1Bal.AmountOf("ustb").String(),
		"9 (bal) + 1 (deposit) + 250u (match reward)")

	curator2Bal := app.BankKeeper.GetAllBalances(ctx, addrs[2])
	require.Equal(t, "1000000", curator2Bal.AmountOf("ucredits").String(),
		"1 (bal)")
	require.Equal(t, "10000750", curator2Bal.AmountOf("ustb").String(),
		"9 (bal) + 1 (deposit) + 750u (match reward)")
}

func TestEndBlocker_RemoveFromExpiredQueue(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs = simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10_000_000))

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	postID, err := types.PostIDFromString("777")
	require.NoError(t, err)
	_, err = app.CuratingKeeper.CreatePost(ctx, uint32(1), &postID, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	postID, err = types.PostIDFromString("888")
	require.NoError(t, err)
	_, err = app.CuratingKeeper.CreatePost(ctx, uint32(1), &postID, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	// force 2 different keys in the iterator underlying store
	b := ctx.BlockHeader()
	b.Time = ctx.BlockHeader().Time.Add(time.Second)

	postID, err = types.PostIDFromString("999")
	require.NoError(t, err)
	_, err = app.CuratingKeeper.CreatePost(ctx.WithBlockHeader(b), uint32(1), &postID, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	// fast-forward blocktime to simulate end of curation window
	h := ctx.BlockHeader()
	h.Time = ctx.BlockHeader().Time.Add(
		app.CuratingKeeper.GetParams(ctx).CurationWindow + time.Minute)
	ctx = ctx.WithBlockHeader(h)

	posts := make([]types.Post, 0)
	curatingEndTimes := make(map[time.Time]bool)
	app.CuratingKeeper.IterateExpiredPosts(ctx, func(p types.Post) bool {
		curatingEndTimes[p.GetCuratingEndTime()] = true
		posts = append(posts, p)
		return false
	})

	require.Len(t, curatingEndTimes, 2, "it should have 2 different end times")
	require.Len(t, posts, 3, "it should have 3 posts")

	curating.EndBlocker(ctx, app.CuratingKeeper)
	posts = make([]types.Post, 0)
	app.CuratingKeeper.IterateExpiredPosts(ctx, func(p types.Post) bool {
		posts = append(posts, p)
		return false
	})

	require.Len(t, posts, 0, "posts should have been removed from queue")
}

func TestEndblocker_BurnCoinsFromVotingPool(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("501")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 5, sdk.NewInt(27_000_000), fakedenom)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	_, err = app.CuratingKeeper.CreatePost(ctx, vendorID, &postID, bodyHash, body, addrs[1], addrs[1], "", nil, "", nil)
	require.NoError(t, err)

	// amt = 1
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 1)
	require.NoError(t, err)

	// amt = 4
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[1], addrs[1], 2)
	require.NoError(t, err)

	// amt = 9
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[2], addrs[2], 3)
	require.NoError(t, err)

	// amt = 16
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[3], addrs[3], 4)
	require.NoError(t, err)

	post, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")
	require.Equal(t, uint64(10), post.TotalVotes)
	require.Equal(t, uint64(4), post.TotalVoters)
	// 1 + 4 + 9 + 16
	require.Equal(t, sdk.NewInt64Coin("ucredits", 30_000_000), post.TotalAmount)
	supply := app.BankKeeper.GetSupply(ctx, "ucredits")

	// 5 initial accounts funded with 27CREDITS
	require.Equal(t, sdk.NewInt(27_000_000*5).String(), supply.Amount.String())

	// pool should have 30CREDITS
	require.Equal(t,
		sdk.NewInt64Coin("ucredits", 30_000_000),
		app.BankKeeper.GetBalance(ctx, authtypes.NewModuleAddress(types.VotingPoolName), "ucredits"),
	)

	// fast-forward blocktime to simulate end of curation window
	h := ctx.BlockHeader()
	h.Time = ctx.BlockHeader().Time.Add(
		app.CuratingKeeper.GetParams(ctx).CurationWindow)
	ctx = ctx.WithBlockHeader(h)

	// run endblocker
	curating.EndBlocker(ctx, app.CuratingKeeper)

	poolBalance := app.BankKeeper.GetBalance(ctx, authtypes.NewModuleAddress(types.VotingPoolName), "ucredits")
	require.Equal(t,
		"0ucredits",
		poolBalance.String(),
	)

	// supply should decrease
	supply = app.BankKeeper.GetSupply(ctx, "ucredits")
	// initial funded amount - 30CREDITS
	require.Equal(t, sdk.NewInt(27_000_000*5-30_000_000).String(), supply.Amount.String())

}
