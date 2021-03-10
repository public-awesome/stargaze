package keeper_test

import (
	"testing"

	"github.com/public-awesome/stargaze/x/curating/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/simapp"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestPost(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[0], addrs[0])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[0], addrs[0])
	require.Equal(t, types.ErrDuplicatePost, err)
}

func TestPost_EmptyCreator(t *testing.T) {
	app := simapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	vendorID := uint32(1)

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, nil, addrs[1])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	rewardAccountBal := app.BankKeeper.GetBalance(ctx, addrs[1], "ustb")
	require.Equal(t, "1000000", rewardAccountBal.Amount.String())

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestPost_EmptyRewardAccount(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[0], nil)
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestPost_WithRewardAccount(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[0], addrs[1])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	rewardAccountBal := app.BankKeeper.GetBalance(ctx, addrs[1], "ustb")
	require.Equal(t, "1000000", rewardAccountBal.Amount.String())

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestDeletePost(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	body := "body string"
	bodyHash, err := types.BodyHashFromString(body)
	require.NoError(t, err)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, bodyHash, body, addrs[0], addrs[1])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	err = app.CuratingKeeper.DeletePost(ctx, vendorID, postID)
	require.NoError(t, err)

	_, found, err = app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.False(t, found, "post should NOT be found")
}

func TestInsertCurationQueue(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	curationWindow := app.CuratingKeeper.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	app.CuratingKeeper.InsertCurationQueue(ctx, vendorID, postID, curationEndTime)

	timeSlice := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, curationEndTime)
	require.Len(t, timeSlice, 1)

	vpPair := types.VPPair{VendorID: vendorID, PostID: postID}
	require.Equal(t, vpPair, timeSlice[0])
}

func TestCurationQueueTimeSlice(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)
	vpPair := types.VPPair{VendorID: vendorID, PostID: postID}

	curationWindow := app.CuratingKeeper.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	app.CuratingKeeper.SetCurationQueueTimeSlice(ctx, curationEndTime, []types.VPPair{vpPair})

	timeSlice := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, curationEndTime)
	require.Len(t, timeSlice, 1)
	require.Equal(t, vpPair, timeSlice[0])
}
