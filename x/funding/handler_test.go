package funding_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/public-awesome/stakebird/x/funding"
	"github.com/public-awesome/stakebird/x/funding/types"
	"github.com/stretchr/testify/require"
)

func TestHandleMsgBuy(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(10000))
	sender := addrs[0]

	ibcCoin := sdk.NewCoin("transfer/ibczeroxfer/stake", sdk.NewInt(10000))
	msgBuy := types.NewMsgBuy(ibcCoin, sender)

	handler := funding.NewHandler(app.FundingKeeper)
	res, err := handler(ctx, msgBuy)
	require.NoError(t, err)
	require.NotNil(t, res)

	communityPool := app.FundingKeeper.DistributionKeeper.GetFeePool(ctx).CommunityPool
	require.Equal(t, "10000.000000000000000000", communityPool.AmountOf("transfer/ibczeroxfer/stake").String())
}
