package keeper_test

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
	"testing"

	"github.com/public-awesome/stakebird/x/curating/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var postID = "500"

func TestPost(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[0], addrs[0])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[0], addrs[0])
	require.Equal(t, types.ErrDuplicatePost, err)
}

func TestPost_EmptyCreator(t *testing.T) {
	app := simapp.Setup(false)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", nil, addrs[1])
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

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[0], nil)
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

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[0], addrs[1])
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

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[0], addrs[1])
	require.NoError(t, err)

	_, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	postIDBz, err := postIDBytes(postID)
	require.NoError(t, err)
	err = app.CuratingKeeper.DeletePost(ctx, vendorID, postIDBz)
	require.NoError(t, err)

	_, found, err = app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.False(t, found, "post should NOT be found")
}

func TestInsertCurationQueue(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postIDBz, err := postIDBytes(postID)
	require.NoError(t, err)

	curationWindow := app.CuratingKeeper.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	app.CuratingKeeper.InsertCurationQueue(ctx, vendorID, postIDBz, curationEndTime)

	timeSlice := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, curationEndTime)
	require.Len(t, timeSlice, 1)

	vpPair := types.VPPair{VendorID: vendorID, PostID: postIDBz}
	require.Equal(t, vpPair, timeSlice[0])
}

func TestCurationQueueTimeSlice(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postIDBz, err := hash(postID)
	require.NoError(t, err)
	vpPair := types.VPPair{VendorID: vendorID, PostID: postIDBz}

	curationWindow := app.CuratingKeeper.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	app.CuratingKeeper.SetCurationQueueTimeSlice(ctx, curationEndTime, []types.VPPair{vpPair})

	timeSlice := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, curationEndTime)
	require.Len(t, timeSlice, 1)
	require.Equal(t, vpPair, timeSlice[0])
}

func hash(body string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// postIDBytes returns the byte representation of a postID
func postIDBytes(postID string) ([]byte, error) {
	postIDInt64, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		return nil, err
	}

	postIDBz := make([]byte, 8)
	binary.BigEndian.PutUint64(postIDBz, uint64(postIDInt64))

	return postIDBz, nil
}
