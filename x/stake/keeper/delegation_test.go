package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDelegation(t *testing.T) {
	_, app, ctx := createTestInput()

	delAddrs := AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10000))
	valAddrs := ConvertAddrsToValAddrs(delAddrs)

	vendorID := uint64(100)
	postID := uint64(200)
	votingPeriod := time.Hour * 24 * 7
	amount := sdk.NewInt64Coin("ufuel", 10000)
	err := app.StakeKeeper.Delegate(ctx, vendorID, postID, delAddrs[0], valAddrs[0], votingPeriod, amount)

	require.NoError(t, err)
}
