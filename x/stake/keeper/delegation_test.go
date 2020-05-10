package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
)

func TestDelegation(t *testing.T) {
	_, app, ctx := createTestInput()

	addrDels := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10000))
	// valAddrs := simapp.ConvertAddrsToValAddrs(addrDels)

	vendorID := uint64(100)
	postID := uint64(200)
	app.StakeKeeper.Delegate(ctx, vendorID, postID, delAddr, valAddr, votingPeriod, amount)
}
