package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(100)
	deposit := sdk.NewInt64Coin("ufuel", 100000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(100000))

	app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[0])

	_, found := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.True(t, found, "Post should be found")

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}
