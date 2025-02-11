package ante_test

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
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
	contractWithCodeAuth := s.SetupContractWithCodeAuth(addr1.String(), "counter.wasm", []string{"increment"})
	contractWithAddrAuth := s.SetupContractWithContractAuth(addr1.String(), "counter.wasm", []string{"increment"})
	contractWithAddrAuthAll := s.SetupContractWithContractAuth(addr1.String(), "counter.wasm", []string{"*"})
	counterIncrementMsg := []byte(`{"increment": {}}`)
	counterResetMsg := []byte(`{"reset": 0}`)

	strPtr := func(s string) *string {
		return &s
	}
	testCases := []struct {
		testCase           string
		minGasPrice        sdk.DecCoins // min gas price configured by the validator
		globalFees         sdk.DecCoins // minimum gas price configured by x/globalfee module param
		feeSent            sdk.Coins    // the amount of fee sent by the user in the tx
		msg                []sdk.Msg
		expectErr          bool
		gasWanted          int64
		expectedConainsErr *string
	}{
		{
			"fail: min_gas_price: empty, globalfee: 5stake, feeSent: 1stake, not authorized contract exec",
			[]sdk.DecCoin{},
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			1,
			nil,
		},
		{
			"ok: min_gas_price: empty, globalfee: 5stake, feeSent: 7stake, not authorized contract exec",
			[]sdk.DecCoin{},
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 7)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"fail: min_gas_price: 0stake, globalfee: 5stake, feeSent: 0stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 0stake, globalfee: 5stake, feeSent: 5stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 1stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			1,
			nil,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 3stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 3)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 5stake, not authorized contract exec",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized code id",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized contract address but not auth msg",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithAddrAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized contract address with auth msg",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithAddrAuth,
					Msg:      counterIncrementMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized contract address with auth all (*)",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithAddrAuthAll,
					Msg:      counterIncrementMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized contract address with auth all (*)",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithAddrAuthAll,
					Msg:      counterResetMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, multiple authorized contract calls",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithAddrAuth,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithAddrAuthAll,
					Msg:      counterResetMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, one authorized contract + unauthorized msgs",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			1,
			nil,
		},
		{
			"ok: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, one authorized contract + unauthorized msgs",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			false,
			1,
			nil,
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, authorized code id but overallocated gas",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
			},
			true,
			50_000_000,
			strPtr("overallocated gas value"),
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, mixed authz and overallocated gas ",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			50_000_000,
			strPtr("overallocated gas value"),
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 500STAKE, mixed authz and overallocated gas  ",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 500_000_000)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
			},
			true,
			50_000_000,
			strPtr("overallocated gas value"),
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake, mixed authz and overallocated gas inv first",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)),
			[]sdk.Msg{

				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
			},
			true,
			50_000_000,
			strPtr("overallocated gas value"),
		},
		{
			"fail: min_gas_price: 2stake, globalfee: 5stake, feeSent: 500STAKE, mixed authz and overallocated gas inv first message",
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 2)),
			sdk.NewDecCoins(sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 5)),
			sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 500_000_000)),
			[]sdk.Msg{
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterResetMsg,
				},
				&wasmtypes.MsgExecuteContract{
					Sender:   addr1.String(),
					Contract: contractWithCodeAuth,
					Msg:      counterIncrementMsg,
				},
			},
			true,
			50_000_000,
			strPtr("overallocated gas value"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testCase, func() {
			_, antehandler := s.SetupTestGlobalFeeStoreAndMinGasPrice(tc.minGasPrice, tc.globalFees)
			s.Require().NoError(s.txBuilder.SetMsgs(tc.msg...))
			s.txBuilder.SetFeeAmount(tc.feeSent)
			s.txBuilder.SetGasLimit(uint64(tc.gasWanted))
			tx, err := s.CreateTestTx(s.ctx, privs, accNums, accSeqs, s.ctx.ChainID(), signing.SignMode_SIGN_MODE_DIRECT)
			s.Require().NoError(err)

			_, err = antehandler(s.ctx, tx, false)

			if !tc.expectErr {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
				if tc.expectedConainsErr != nil {
					s.Require().Contains(err.Error(), *tc.expectedConainsErr)
				}
			}
		})
	}
}
