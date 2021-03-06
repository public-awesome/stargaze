package keeper_test

import (
	"encoding/json"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/simapp"
	"github.com/public-awesome/stargaze/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestCreatePost(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	ctx = ctx.WithBlockTime(time.Now())

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	_, err = app.CuratingKeeper.CreatePost(ctx, vendorID, &postID, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	post, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")
	require.Equal(t, "500", post.PostID.String())
	obj, err := json.MarshalIndent(post.PostID, "", " ")
	if err != nil {
		panic(err)
	}
	require.Equal(t, "\"500\"", string(obj))

	curatingQueue := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime().Add(10*time.Minute))
	require.Equal(t, curatingQueue[0].PostID.String(), "500")
	require.Equal(t, curatingQueue[0].VendorID, uint32(1))

	// add another post
	postID, err = types.PostIDFromString("501")
	_, err = app.CuratingKeeper.CreatePost(ctx, vendorID, &postID, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	// fast forward block time
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour*24*3 + 1))

	// test expired post iterator
	app.CuratingKeeper.IterateExpiredPosts(ctx, func(post types.Post) (stop bool) {
		require.NotNil(t, post)
		return false
	})
}

func TestCreateVendor0Post(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("1")
	require.NoError(t, err)
	vendorID := uint32(0)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	ctx = ctx.WithBlockTime(time.Now())

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	_, err = app.CuratingKeeper.CreatePost(ctx, vendorID, nil, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	post, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")
	require.Equal(t, "1", post.PostID.String())
	obj, err := json.MarshalIndent(post.PostID, "", " ")
	if err != nil {
		panic(err)
	}
	require.Equal(t, "\"1\"", string(obj))

	curatingQueue := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime().Add(10*time.Minute))
	require.Equal(t, curatingQueue[0].PostID.String(), "1")
	require.Equal(t, curatingQueue[0].VendorID, uint32(0))

	// add another post
	postID, err = types.PostIDFromString("2")
	_, err = app.CuratingKeeper.CreatePost(ctx, vendorID, nil, bodyHash, body, addrs[0], addrs[0], "", nil, "", nil)
	require.NoError(t, err)

	// fast forward block time
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour*24*3 + 1))

	// test expired post iterator
	app.CuratingKeeper.IterateExpiredPosts(ctx, func(post types.Post) (stop bool) {
		require.NotNil(t, post)
		return false
	})
}
