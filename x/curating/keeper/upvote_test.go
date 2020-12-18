package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestCreateUpvote(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID := "500"
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	err := app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
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
}

func TestCreateUpvote_ExistingPost(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID := "501"
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[1], addrs[1])
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

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

	postID := "501"
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[1], addrs[1])
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour*24*3 + 1))
	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.Error(t, err)
	serr, ok := err.(*sdkerrors.Error)
	require.True(t, ok)
	require.Equal(t, types.ErrPostExpired, serr)
}

func TestCreateUpvote_ExistingUpvote(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID := "502"
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(77_000_000))

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[1], addrs[1])
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5)
	require.NoError(t, err)
}
