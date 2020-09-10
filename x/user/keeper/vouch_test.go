package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voucher := addresses[0]
	vouched := addresses[1]

	err := app.UserKeeper.CreateVouch(ctx, voucher, vouched, "")
	require.NoError(t, err)

	vouchByVoucher, found, err := app.UserKeeper.GetVouchByVoucher(ctx, voucher)
	require.NoError(t, err)
	require.True(t, found, "vouch should be found")

	vouchByVouched, found, err := app.UserKeeper.GetVouchByVouched(ctx, vouched)
	require.NoError(t, err)
	require.True(t, found, "vouch should be found")

	require.Equal(t, vouchByVoucher, vouchByVouched)
}
