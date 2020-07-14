package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestCreateUpvote(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ufuel", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000))

	err := app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5, deposit)
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	upvote, found, err := app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[0])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")

	require.Equal(t, "25000000ufuel", upvote.VoteAmount.String())

	curatorBalance := app.BankKeeper.GetBalance(ctx, addrs[0], "ufuel")
	require.Equal(t, "1000000", curatorBalance.Amount.String())

}

func TestCreateUpvote_ExistingPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ufuel", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000))

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[1], addrs[1])
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5, deposit)
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	upvote, found, err := app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[0])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")

	require.Equal(t, "25000000ufuel", upvote.VoteAmount.String())

	creatorBalance := app.BankKeeper.GetBalance(ctx, addrs[1], "ufuel")
	require.Equal(t, "26000000", creatorBalance.Amount.String())

	curatorBalance := app.BankKeeper.GetBalance(ctx, addrs[0], "ufuel")
	require.Equal(t, "1000000", curatorBalance.Amount.String())

}
