package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestDelegation(t *testing.T) {
	_, app, ctx := createTestInput()

	delAddrs := AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10000))
	valAddrs := ConvertAddrsToValAddrs(delAddrs)

	vendorID := uint64(100)
	postID := uint64(200)
	app.StakeKeeper.Delegate(ctx, vendorID, postID, delAddrs[0], valAddrs[0], votingPeriod, amount)
}
