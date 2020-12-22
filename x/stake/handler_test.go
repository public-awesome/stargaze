package stake_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/stake"
	"github.com/public-awesome/stakebird/x/stake/types"
	"github.com/stretchr/testify/assert"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestHandleMsgStake(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))

	handler := stake.NewHandler(app.StakeKeeper)
	msg := types.NewMsgStake(1, "123", addrs[0], sdk.NewInt(1_000_000))
	_, err := handler(ctx, msg)
	assert.NoError(t, err)
}

func TestHandleMsgUnstake(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))

	handler := stake.NewHandler(app.StakeKeeper)
	msg := types.NewMsgUnstake(1, "123", addrs[0], sdk.NewInt(1_000_000))
	_, err := handler(ctx, msg)
	assert.NoError(t, err)
}
