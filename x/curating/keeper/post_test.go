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
	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ufuel", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", deposit, addrs[0], addrs[0])
	require.NoError(t, err)

	_, found := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.True(t, found, "post should be found")

	creatorBal := app.BankKeeper.GetBalance(ctx, addrs[0], "ufuel")
	require.Equal(t, "0", creatorBal.Amount.String())

	vps := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime())
	require.NotNil(t, vps)
}
