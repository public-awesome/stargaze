package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/user/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestNewMsgVouch(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	comment := "lorem ipsum"
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	msg := types.NewMsgVouch(addresses[0], addresses[1], comment)
	require.Equal(t, msg.GetVoucher(), addresses[0].String())
	require.Equal(t, msg.GetVouched(), addresses[1].String())
	require.Equal(t, msg.GetComment(), comment)
}

func TestValidateBasicMsgVouch_EmptyVoucher(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	comment := "lorem ipsum"

	msg := types.NewMsgVouch(nil, addresses[1], comment)
	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgVouch_EmptyVouched(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	comment := "lorem ipsum"

	msg := types.NewMsgVouch(addresses[0], nil, comment)
	err := msg.ValidateBasic()
	require.NotNil(t, err)
}

func TestValidateBasicMsgVouch_EmptyComment(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	comment := ""

	msg := types.NewMsgVouch(addresses[0], addresses[1], comment)
	err := msg.ValidateBasic()
	require.Nil(t, err) // asserting that comment is optional
}
