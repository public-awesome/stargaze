package ante_test

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stargazeapp "github.com/public-awesome/stargaze/v9/app"
	"github.com/public-awesome/stargaze/v9/testutil/simapp"
	"github.com/public-awesome/stargaze/v9/x/globalfee/ante"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type AnteHandlerTestSuite struct {
	suite.Suite

	app       *stargazeapp.App
	ctx       sdk.Context
	clientCtx client.Context
	txBuilder client.TxBuilder
}

func (s *AnteHandlerTestSuite) SetupTest() {
	_, _, acc1_addr := getTestAccount()
	_, _, acc2_addr := getTestAccount()
	genAccounts := authtypes.GenesisAccounts{
		&authtypes.BaseAccount{Address: acc1_addr.String()},
		&authtypes.BaseAccount{Address: acc2_addr.String()},
	}
	genBalances := []banktypes.Balance{
		{
			Address: acc1_addr.String(),
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5_000_000_000)),
		},
		{
			Address: acc2_addr.String(),
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5_000_000_000)),
		},
	}
	app := simapp.SetupWithGenesisAccounts(s.T(), s.T().TempDir(), genAccounts, genBalances...)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{
		ChainID: "ante-test-1",
		Height:  1,
	})

	encodingConfig := cosmoscmd.MakeEncodingConfig(stargazeapp.ModuleBasics)

	s.app = app
	s.ctx = ctx
	s.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
}

func (s *AnteHandlerTestSuite) SetupTestGlobalFeeStoreAndMinGasPrice(minGasPrice []sdk.DecCoin, globalFees sdk.DecCoins) (ante.FeeDecorator, sdk.AnteHandler) {
	subspace := s.app.GetSubspace(types.ModuleName)
	subspace.SetParamSet(s.ctx, &types.Params{MinimumGasPrices: globalFees})
	s.ctx = s.ctx.WithMinGasPrices(minGasPrice).WithIsCheckTx(true)

	// build fee decorator
	feeDecorator := ante.NewFeeDecorator(s.app.AppCodec(), nil, s.app.StakingKeeper)

	// chain fee decorator to antehandler
	antehandler := sdk.ChainAnteDecorators(feeDecorator)

	return feeDecorator, antehandler
}

func (s *AnteHandlerTestSuite) CreateTestTx(privs []cryptotypes.PrivKey, accNums []uint64, accSeqs []uint64, chainID string) (xauthsigning.Tx, error) {
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  s.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}

	if err := s.txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}

	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(
			s.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
			signerData,
			s.txBuilder,
			priv,
			s.clientCtx.TxConfig,
			accSeqs[i],
		)
		if err != nil {
			return nil, err
		}

		sigsV2 = append(sigsV2, sigV2)
	}

	if err := s.txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}

	return s.txBuilder.GetTx(), nil
}

func getTestAccount() (privateKey secp256k1.PrivKey, publicKey crypto.PubKey, accountAddress sdk.AccAddress) {
	privateKey = secp256k1.GenPrivKey()
	publicKey = privateKey.PubKey()
	accountAddress = sdk.AccAddress(publicKey.Address())
	return
}
