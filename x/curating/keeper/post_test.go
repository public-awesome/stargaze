package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := uint64(500)
	vendorID := uint64(100)
	stake := sdk.NewInt64Coin("ufue", 100000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(100000))

	app.CuratingKeeper.CreatePost(ctx, postID, vendorID, "hash string", stake, addrs[0])

	_, found := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.True(t, found, "Post should be found")
}
