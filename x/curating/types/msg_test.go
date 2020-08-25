package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/public-awesome/stakebird/testdata"
	"github.com/public-awesome/stakebird/x/curating/types"
)

func TestNewMsgPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	body := "lorem ipsum"

	msg := types.NewMsgPost(vendorID, postID, addresses[0], addresses[0], body)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetBody(), body)
	require.Equal(t, msg.GetCreator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
}

func TestValidateBasicMsgPost_EmptyBody(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	body := ""

	msg := types.NewMsgPost(vendorID, postID, addresses[0], addresses[0], body)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetCreator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgPost_EmptyCreator(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	body := "lorem ipsum"

	msg := types.NewMsgPost(vendorID, postID, nil, addresses[0], body)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetBody(), body)

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgPost_EmptyPostID(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := ""
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	body := "lorem ipsum"

	msg := types.NewMsgPost(vendorID, postID, addresses[0], addresses[0], body)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetCreator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetBody(), body)

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgPost_InvalidVendorID(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(0)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	body := "lorem ipsum"

	msg := types.NewMsgPost(vendorID, postID, addresses[0], addresses[0], body)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetCreator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetBody(), body)

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestNewMsgUpvote(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voteNum := int32(2)

	msg := types.NewMsgUpvote(vendorID, postID, addresses[0], addresses[0], voteNum)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetCurator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetVoteNum(), voteNum)
}

func TestValidateBasicMsgUpvote_EmptyCurator(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voteNum := int32(2)

	msg := types.NewMsgUpvote(vendorID, postID, nil, addresses[0], voteNum)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetVoteNum(), voteNum)

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgUpvote_EmptyPostID(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := ""
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voteNum := int32(2)

	msg := types.NewMsgUpvote(vendorID, postID, addresses[0], addresses[0], voteNum)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetCurator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetVoteNum(), voteNum)

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgUpvote_InvalidVendorID(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(0)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voteNum := int32(2)

	msg := types.NewMsgUpvote(vendorID, postID, addresses[0], addresses[0], voteNum)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetCurator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])
	require.Equal(t, msg.GetVoteNum(), voteNum)

	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgUpvote_InvalidVoteNum(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	vendorID := uint32(1)
	postID := "100"
	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	// zero votes
	voteNum := int32(0)

	msg := types.NewMsgUpvote(vendorID, postID, addresses[0], addresses[0], voteNum)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetCurator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])

	err := msg.ValidateBasic()
	require.NotNil(t, err)

	// negative votes
	voteNum = int32(-1)

	msg = types.NewMsgUpvote(vendorID, postID, addresses[0], addresses[0], voteNum)
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetCurator(), addresses[0])
	require.Equal(t, msg.GetRewardAccount(), addresses[0])

	err = msg.ValidateBasic()
	require.NotNil(t, err)
}
