package bondcurve_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rocket-protocol/stakebird/testdata"
	"github.com/rocket-protocol/stakebird/x/bondcurve"
	"github.com/rocket-protocol/stakebird/x/bondcurve/types"
	"github.com/stretchr/testify/require"
)

func TestHandleMsgBuy(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10000))
	sender := addrs[0]

	ibcCoin := sdk.NewCoin("transfer/ibczeroxfer/stake", sdk.NewInt(10000))
	msgBuy := types.NewMsgBuy(ibcCoin, sender)

	handler := bondcurve.NewHandler(app.BondCurveKeeper)
	res, err := handler(ctx, msgBuy)
	require.NoError(t, err)
	require.NotNil(t, res)

	// moduleAccount := app.AccountKeeper.GetModuleAccount(ctx, bondcurve.ModuleName)
	// addr := moduleAccount.GetAddress()
	// coin := app.BankKeeper.GetBalance(ctx, addr, "transfer/ibczeroxfer/stake")
	// fmt.Println(coin)
	// coin = app.BankKeeper.GetBalance(ctx, addr, "stake")
	// fmt.Println(coin)
	// coin = app.BankKeeper.GetBalance(ctx, addr, "ufuel")
	// fmt.Println(coin)

	// coin = app.BankKeeper.GetBalance(ctx, sender, "transfer/ibczeroxfer/stake")
	// fmt.Println(coin)
	// coin = app.BankKeeper.GetBalance(ctx, sender, "stake")
	// fmt.Println(coin)
	// coin = app.BankKeeper.GetBalance(ctx, sender, "ufuel")
	// fmt.Println(coin)

	fuelCoin := sdk.NewCoin("ufuel", sdk.NewInt(10000))
	msgSell := types.NewMsgSell(fuelCoin, sender)
	res, err = handler(ctx, msgSell)
	require.NoError(t, err)
	require.NotNil(t, res)
}
