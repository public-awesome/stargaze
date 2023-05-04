package ante_test

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

func TestAnteHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AnteHandlerTestSuite))
}

func (s *AnteHandlerTestSuite) TestFeeDecoratorAntehandler() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()
	priv1, _, addr1 := testdata.KeyTestPubAddr()
	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	contractAddr := s.SetupContracts(addr1.String(), "counter.wasm")
	counterIncrementMsg := []byte(`{"increment": {}}`)
	counterResetMsg := []byte(`{"reset": 0}`)

	testCases := []struct {
		testCase    string
		minGasPrice sdk.DecCoins // min gas price configured by the validator
		globalFees  sdk.DecCoins // minimum gas price configured by x/globalfee module param
		feeSent     sdk.Coins    // the amount of fee sent by the user in the tx
		msg         []sdk.Msg
		expectErr   bool
	}{
		{
			"fail: min_gas_price: empty, globalfee: 5stake, feeSent: 1stake, not authorized contract exec",
			[]sdk.DecCoin{},
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterResetMsg,
				},
			},
			true,
		},
		{
			"ok: min_gas_price: empty, globalfee: 5stake, feeSent: 7stake, not authorized contract exec",
			[]sdk.DecCoin{},
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 7)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterResetMsg,
				},
			},
			false,
		},
		{
			"ok: min_gas_price: 0stake, globalfee: 5stake, feeSent: 0stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterResetMsg,
				},
			},
			false,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 1stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterResetMsg,
				},
			},
			true,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 3stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 3)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterResetMsg,
				},
			},
			false,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized code id",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterIncrementMsg,
				},
			},
			false,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized contract address",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterIncrementMsg,
				},
			},
			false,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, multiple authorized contract calls",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterIncrementMsg,
				},
			},
			false,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, one authorized contract + unauthorized msgs",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractAddr,
					Msg:      counterResetMsg,
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testCase, func() {
			_, antehandler := s.SetupTestGlobalFeeStoreAndMinGasPrice(tc.minGasPrice, tc.globalFees)
			s.Require().NoError(s.txBuilder.SetMsgs(tc.msg...))
			s.txBuilder.SetFeeAmount(tc.feeSent)
			s.txBuilder.SetGasLimit(1)
			tx, err := s.CreateTestTx(privs, accNums, accSeqs, s.ctx.ChainID())
			s.Require().NoError(err)

			_, err = antehandler(s.ctx, tx, false)

			if !tc.expectErr {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}
}
