package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/public-awesome/stargaze/simapp"
	"github.com/public-awesome/stargaze/x/stake/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestNewMsgStake(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postID := "100"
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addresses)

	msg := types.NewMsgStake(vendorID, postID, addresses[0], valAddrs[0], sdk.NewInt(1000000))
	require.Equal(t, msg.GetVendorID(), vendorID)
	require.Equal(t, msg.GetPostID(), postID)
	require.Equal(t, msg.GetDelegator(), addresses[0].String())
}
