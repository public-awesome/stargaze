package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestVouch(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voucher := addresses[0]
	vouched := addresses[1]

	err := app.UserKeeper.CreateVouch(ctx, voucher, vouched, "")
	require.NoError(t, err)

	vouchByVouched, found, err := app.UserKeeper.GetVouchByVouched(ctx, vouched)
	require.NoError(t, err)
	require.True(t, found, "vouch should be found")

	vouchesByVoucher := app.UserKeeper.GetVouchesByVoucher(ctx, voucher)
	require.NoError(t, err)
	require.Contains(t, vouchesByVoucher, vouchByVouched, "vouch should be found")
}

func TestIsVouched(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	addresses := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	voucher := addresses[0]
	vouched := addresses[1]

	is := app.UserKeeper.IsVouched(ctx, vouched)
	require.False(t, is)

	err := app.UserKeeper.CreateVouch(ctx, voucher, addresses[1], "")
	require.NoError(t, err)

	is = app.UserKeeper.IsVouched(ctx, vouched)
	require.True(t, is)
}

func TestCanVouch(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	addresses := testdata.AddTestAddrsIncremental(app, ctx, 4, sdk.NewInt(1000000))
	voucher := addresses[0]

	can := app.UserKeeper.CanVouch(ctx, voucher)
	require.True(t, can)

	err := app.UserKeeper.CreateVouch(ctx, voucher, addresses[1], "")
	require.NoError(t, err)
	err = app.UserKeeper.CreateVouch(ctx, voucher, addresses[2], "")
	require.NoError(t, err)
	err = app.UserKeeper.CreateVouch(ctx, voucher, addresses[3], "")
	require.NoError(t, err)

	can = app.UserKeeper.CanVouch(ctx, voucher)
	require.False(t, can)
}
