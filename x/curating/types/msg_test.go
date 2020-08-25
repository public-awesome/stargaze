package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/public-awesome/stakebird/testdata"
	"github.com/public-awesome/stakebird/x/curating/types"
)

func TestPost(t *testing.T) {
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
