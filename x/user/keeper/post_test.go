package keeper_test

import (
	"crypto/sha256"
	"testing"

	"github.com/public-awesome/stakebird/x/user/types"
	userTypes "github.com/public-awesome/stakebird/x/user/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

var postID = "500"

func TestPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ustb", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.userKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[0])
	require.NoError(t, err)

	_, found, err := app.userKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "0", creatorBal.Amount.String())

	vps := app.userKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestPost_EmptyCreator(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ustb", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.userKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, nil, addrs[1])
	require.NoError(t, err)

	_, found, err := app.userKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "1000000", creatorBal.Amount.String())

	rewardAccountBal := app.BankKeeper.GetBalance(ctx, addrs[1], "ustb")
	require.Equal(t, "1000000", rewardAccountBal.Amount.String())

	vps := app.userKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestPost_EmptyRewardAccount(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ustb", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.userKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], nil)
	require.NoError(t, err)

	_, found, err := app.userKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "0", creatorBal.Amount.String())

	vps := app.userKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestPost_WithRewardAccount(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ustb", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.userKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[1])
	require.NoError(t, err)

	_, found, err := app.userKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ustb")
	require.Equal(t, "0", creatorBal.Amount.String())

	rewardAccountBal := app.BankKeeper.GetBalance(ctx, addrs[1], "ustb")
	require.Equal(t, "1000000", rewardAccountBal.Amount.String())

	vps := app.userKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}

func TestDeletePost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ustb", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.userKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[1])
	require.NoError(t, err)

	_, found, err := app.userKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")

	postIDHash, err := hash(postID)
	require.NoError(t, err)
	err = app.userKeeper.DeletePost(ctx, vendorID, postIDHash)
	require.NoError(t, err)

	_, found, err = app.userKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.False(t, found, "post should NOT be found")
}

func TestInsertCurationQueue(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postIDHash, err := hash(postID)
	require.NoError(t, err)

	curationWindow := app.userKeeper.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	app.userKeeper.InsertCurationQueue(ctx, vendorID, postIDHash, curationEndTime)

	timeSlice := app.userKeeper.GetCurationQueueTimeSlice(ctx, curationEndTime)
	require.Len(t, timeSlice, 1)

	vpPair := userTypes.VPPair{VendorID: vendorID, PostIDHash: postIDHash}
	require.Equal(t, vpPair, timeSlice[0])
}

func TestCurationQueueTimeSlice(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postIDHash, err := hash(postID)
	require.NoError(t, err)
	vpPair := userTypes.VPPair{VendorID: vendorID, PostIDHash: postIDHash}

	curationWindow := app.userKeeper.GetParams(ctx).CurationWindow
	curationEndTime := ctx.BlockTime().Add(curationWindow)
	app.userKeeper.SetCurationQueueTimeSlice(ctx, curationEndTime, []types.VPPair{vpPair})

	timeSlice := app.userKeeper.GetCurationQueueTimeSlice(ctx, curationEndTime)
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
