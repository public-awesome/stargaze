package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stargaze/simapp"
	"github.com/public-awesome/stargaze/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestCreateUpvote(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	upvote, found, err := app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[0])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")

	require.Equal(t, "25000000ucredits", upvote.VoteAmount.String())

	curatorBalance := app.BankKeeper.GetBalance(ctx, addrs[0], "ucredits")
	require.Equal(t, "2000000", curatorBalance.Amount.String())

	// test iterate
	app.CuratingKeeper.IterateUpvotes(ctx, vendorID, postID, func(u types.Upvote) (stop bool) {
		require.Equal(t, "500", u.PostID.String())
		return false
	})
}

func TestCreateUpvote_ExistingPost(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("501")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[1], addrs[1])
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.NoError(t, err)

	post, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")
	require.Equal(t, uint64(5), post.TotalVotes)
	require.Equal(t, uint64(1), post.TotalVoters)
	require.Equal(t, sdk.NewInt64Coin("ucredits", 25_000_000), post.TotalAmount)

	upvote, found, err := app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[0])
	require.NoError(t, err)
	require.True(t, found, "upvote should be found")

	require.Equal(t, "25000000ucredits", upvote.VoteAmount.String())

	creatorBalance := app.BankKeeper.GetBalance(ctx, addrs[1], fakedenom)
	require.Equal(t, "27000000", creatorBalance.Amount.String())

	curatorBalance := app.BankKeeper.GetBalance(ctx, addrs[0], "ucredits")
	require.Equal(t, "2000000", curatorBalance.Amount.String())
}
func TestCreateUpvote_ExpiredPost(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("501")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[1], addrs[1])
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour*24*3 + 1))
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.Error(t, err)
	serr, ok := err.(*sdkerrors.Error)
	require.True(t, ok)
	require.Equal(t, types.ErrPostExpired, serr)
}

func TestMultipleUpvotes(t *testing.T) {
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

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[1], addrs[1])
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
}

func TestCreateUpvote_ExistingUpvote(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("502")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000))

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[1], addrs[1])
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.Error(t, types.ErrAlreadyVoted, err)
}
