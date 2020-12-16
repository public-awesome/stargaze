package curating_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/curating"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/stretchr/testify/assert"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestHandlePost(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs = simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))

	handler := curating.NewHandler(app.CuratingKeeper)
	msgPost := types.NewMsgPost(1, "123", addrs[0], nil, "testbody")
	_, err := handler(ctx, msgPost)
	assert.NoError(t, err)
}

func TestHandleUpvote(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs = simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))

	handler := curating.NewHandler(app.CuratingKeeper)
	msgUpvote := types.NewMsgUpvote(1, "123", addrs[0], nil, 1)
	_, err := handler(ctx, msgUpvote)
	assert.NoError(t, err)
}

// func TestHandleStake(t *testing.T) {
// 	app := simapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// 	addrs = simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))
// 	addrVals := simapp.ConvertAddrsToValAddrs(addrs)

// 	initPower := int64(1000000)
// 	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

// 	// // create validator
// 	// PKs := simapp.CreateTestPubKeys(500)
// 	// initBond := tstaking.CreateValidatorWithValPower(addrVals[0], PKs[0], initPower, true)

// 	// // must end-block
// 	// updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
// 	// require.NoError(t, err)
// 	// require.Equal(t, 1, len(updates))

// 	handler := curating.NewHandler(app.CuratingKeeper)
// 	msgStake := types.NewMsgStake(1, "123", addrs[0], addrVals[0], sdk.NewInt(1_000_000))
// 	_, err := handler(ctx, msgStake)

// 	assert.NoError(t, err)
// }
