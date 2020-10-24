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
