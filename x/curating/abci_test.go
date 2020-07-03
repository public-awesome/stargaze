package curating_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/public-awesome/stakebird/x/curating"
	"github.com/stretchr/testify/require"
)

func TestEndBlockerExpiringPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ufuel", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[0])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ufuel")
	require.Equal(t, "0", creatorBal.Amount.String())

	// curator1
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[1], addrs[1], 5, deposit)
	require.NoError(t, err)
	_, found, err = app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[1])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")
	curatorBal := app.BankKeeper.GetBalance(ctx, addrs[1], "ufuel")
	require.Equal(t, "0", curatorBal.Amount.String())

	// curator2
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[2], addrs[2], 5, deposit)
	require.NoError(t, err)
	_, found, err = app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[2])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")
	curatorBal = app.BankKeeper.GetBalance(ctx, addrs[2], "ufuel")
	require.Equal(t, "0", curatorBal.Amount.String())

	// fast-forward blocktime to simulate end of curation window
	h := ctx.BlockHeader()
	h.Time = ctx.BlockHeader().Time.Add(app.CuratingKeeper.GetParams(ctx).CurationWindow)
	ctx = ctx.WithBlockHeader(h)

	curating.EndBlocker(ctx, app.CuratingKeeper)

	creatorBal = app.BankKeeper.GetBalance(ctx, addrs[0], "ufuel")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	curatorBal = app.BankKeeper.GetBalance(ctx, addrs[1], "ufuel")
	require.Equal(t, "1000000", curatorBal.Amount.String())

	curatorBal = app.BankKeeper.GetBalance(ctx, addrs[2], "ufuel")
	require.Equal(t, "1000000", curatorBal.Amount.String())
}
