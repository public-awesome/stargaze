package ante_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stargazeapp "github.com/public-awesome/stargaze/v11/app"
	"github.com/public-awesome/stargaze/v11/testutil/simapp"
	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestMinDepositDecorator(t *testing.T) {
	priv1 := secp256k1.GenPrivKey()
	pub1 := priv1.PubKey()
	addr1 := sdk.AccAddress(pub1.Address())

	priv2 := secp256k1.GenPrivKey()
	pub2 := priv2.PubKey()
	addr2 := sdk.AccAddress(pub2.Address())

	genTokens := sdk.TokensFromConsensusPower(5000, sdk.DefaultPowerReduction)
	bondTokens := sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction)
	genCoin := sdk.NewCoin(sdk.DefaultBondDenom, genTokens)
	stars := sdk.NewCoin("ustars", sdk.NewInt(5_000_000_000))
	bondCoin := sdk.NewCoin(sdk.DefaultBondDenom, bondTokens)

	acc1 := &authtypes.BaseAccount{Address: addr1.String()}
	acc2 := &authtypes.BaseAccount{Address: addr2.String()}
	accs := authtypes.GenesisAccounts{acc1, acc2}
	balances := []banktypes.Balance{
		{
			Address: addr1.String(),
			Coins:   sdk.Coins{genCoin, stars},
		},
		{
			Address: addr2.String(),
			Coins:   sdk.Coins{genCoin, stars},
		},
	}

	app := simapp.SetupWithGenesisAccounts(t, t.TempDir(), accs, balances...)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	params := govtypes.DefaultDepositParams()
	params.MinDeposit = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(5_000_000_000)))
	app.GovKeeper.SetDepositParams(ctx, params)
	app.Commit()
	content := govtypes.ContentFromProposalType("Prop Title", "Description", govtypes.ProposalTypeText)

	createProposalMsg, err := govtypes.NewMsgSubmitProposal(content, sdk.NewCoins(bondCoin), addr1)

	require.NoError(t, err)
	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	encoding := stargazeapp.MakeEncodingConfig()
	txGen := encoding.TxConfig
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createProposalMsg}, "", []uint64{0}, []uint64{0}, true, false, false, priv1)
	require.EqualError(t, err, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("min deposit cannot be lower than %d %s", 1_000_000_000, sdk.DefaultBondDenom)).Error())

	createProposalMsg, err = govtypes.NewMsgSubmitProposal(content, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1_000_000_000)), addr1)
	require.NoError(t, err)

	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createProposalMsg}, "", []uint64{0}, []uint64{0}, false, true, true, priv1)
	require.NoError(t, err)

	app = simapp.SetupWithGenesisAccounts(t, t.TempDir(), accs, balances...)
	ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	params = govtypes.DefaultDepositParams()
	params.MinDeposit = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20_000_000_000)))
	app.GovKeeper.SetDepositParams(ctx, params)
	app.Commit()

	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createProposalMsg}, "", []uint64{0}, []uint64{0}, true, false, false, priv1)
	require.EqualError(t, err, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("min deposit cannot be lower than %d %s", 4000000000, sdk.DefaultBondDenom)).Error())

	createProposalMsg, err = govtypes.NewMsgSubmitProposal(content, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 4_000_000_000)), addr2)
	require.NoError(t, err)
	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createProposalMsg}, "", []uint64{1}, []uint64{0}, false, true, true, priv2)
	require.NoError(t, err)
}
