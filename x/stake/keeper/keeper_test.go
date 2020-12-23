package keeper_test

import (
	"testing"

	"github.com/bwmarrin/snowflake"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	PKs = simapp.CreateTestPubKeys(500)
)

func TestPerformStakeAndUnstake(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	valConsPk1 := PKs[0]

	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(1000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)

	// create validator with 50% commission
	tstaking.Commission = types.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, 100, true)
	addrDels := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(3))

	vendorID := uint32(1)
	postID := "500"
	postIDBz, err := postIDBytes(postID)
	require.NoError(t, err)

	delAddr := addrDels[1]
	valAddr := valAddrs[0]
	amount := sdk.NewInt(2)

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", delAddr, delAddr)
	require.NoError(t, err)

	staking.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	err = app.StakeKeeper.PerformStake(ctx, vendorID, postIDBz, delAddr, valAddr, amount)
	require.NoError(t, err)
	s, found, err := app.StakeKeeper.GetStake(ctx, vendorID, postIDBz, delAddr)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, sdk.NewInt(2), s.Amount)

	// withdraw half of the stake
	err = app.StakeKeeper.PerformUnstake(ctx, vendorID, postIDBz, delAddr, valAddr, amount.QuoRaw(2))
	require.NoError(t, err)
	s, found, err = app.StakeKeeper.GetStake(ctx, vendorID, postIDBz, delAddr)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, sdk.NewInt(1), s.Amount)

	// withdraw the other half, stake should be gone
	err = app.StakeKeeper.PerformUnstake(ctx, vendorID, postIDBz, delAddr, valAddr, amount.QuoRaw(2))
	require.NoError(t, err)
	s, found, err = app.StakeKeeper.GetStake(ctx, vendorID, postIDBz, delAddr)
	require.NoError(t, err)
	require.False(t, found)
}

// postIDBytes returns the byte representation of a postID int64
func postIDBytes(postID string) ([]byte, error) {
	pID, err := snowflake.ParseString(postID)
	if err != nil {
		return nil, err
	}

	temp := pID.IntBytes()

	return temp[:], nil
}
