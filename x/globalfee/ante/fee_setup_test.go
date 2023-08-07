package ante_test

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stargazeapp "github.com/public-awesome/stargaze/v11/app"
	"github.com/public-awesome/stargaze/v11/testutil/simapp"
	"github.com/public-awesome/stargaze/v11/x/globalfee/ante"
	"github.com/public-awesome/stargaze/v11/x/globalfee/types"
	"github.com/stretchr/testify/suite"
)

type AnteHandlerTestSuite struct {
	suite.Suite

	app       *stargazeapp.App
	msgServer wasmtypes.MsgServer
	ctx       sdk.Context
	clientCtx client.Context
	txBuilder client.TxBuilder
}

type storeCache struct {
	sync.Mutex
	contracts map[string][]byte
}

var contractsCache = storeCache{contracts: make(map[string][]byte)}

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
		Time:    time.Now(),
	})

	encodingConfig := stargazeapp.MakeEncodingConfig()

	s.app = app
	s.ctx = ctx
	s.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
}

func (s *AnteHandlerTestSuite) SetupTestGlobalFeeStoreAndMinGasPrice(minGasPrice []sdk.DecCoin, globalFees sdk.DecCoins) (ante.FeeDecorator, sdk.AnteHandler) {
	s.app.GlobalFeeKeeper.SetParams(s.ctx, types.Params{MinimumGasPrices: globalFees})
	s.ctx = s.ctx.WithMinGasPrices(minGasPrice).WithIsCheckTx(true)

	// build fee decorator
	feeDecorator := ante.NewFeeDecorator(s.app.AppCodec(), s.app.GlobalFeeKeeper, s.app.StakingKeeper)

	// chain fee decorator to antehandler
	antehandler := sdk.ChainAnteDecorators(feeDecorator)

	return feeDecorator, antehandler
}

func (s *AnteHandlerTestSuite) SetupWasmMsgServer() {
	wasmParams := s.app.WasmKeeper.GetParams(s.ctx)
	wasmParams.CodeUploadAccess = wasmtypes.AllowEverybody
	s.app.WasmKeeper.SetParams(s.ctx, wasmParams)
	s.msgServer = wasmkeeper.NewMsgServerImpl(wasmkeeper.NewDefaultPermissionKeeper(s.app.WasmKeeper))
}

func (s *AnteHandlerTestSuite) SetupContractWithCodeAuth(senderAddr string, contractBinary string, authMethods []string) string {
	codeId, err := storeContract(s.ctx, s.msgServer, senderAddr, contractBinary)
	s.Require().NoError(err)

	instantiageMsg := CounterInsantiateMsg{Count: 0}
	instantiateMsgRaw, err := json.Marshal(&instantiageMsg)
	s.Require().NoError(err)

	initMsg := wasmtypes.MsgInstantiateContract{Sender: senderAddr, Admin: senderAddr, CodeID: codeId, Label: "Counter Contract", Msg: instantiateMsgRaw, Funds: sdk.NewCoins()}
	instantiateRes, err := s.msgServer.InstantiateContract(sdk.WrapSDKContext(s.ctx), &initMsg)
	s.Require().NoError(err)

	err = s.app.GlobalFeeKeeper.SetCodeAuthorization(s.ctx, types.CodeAuthorization{
		CodeID:  codeId,
		Methods: authMethods,
	})
	s.Require().NoError(err)

	return instantiateRes.Address
}

func (s *AnteHandlerTestSuite) SetupContractWithContractAuth(senderAddr string, contractBinary string, authMethods []string) string {
	codeId, err := storeContract(s.ctx, s.msgServer, senderAddr, contractBinary)
	s.Require().NoError(err)

	instantiageMsg := CounterInsantiateMsg{Count: 0}
	instantiateMsgRaw, err := json.Marshal(&instantiageMsg)
	s.Require().NoError(err)

	initMsg := wasmtypes.MsgInstantiateContract{Sender: senderAddr, Admin: senderAddr, CodeID: codeId, Label: "Counter Contract", Msg: instantiateMsgRaw, Funds: sdk.NewCoins()}
	instantiateRes, err := s.msgServer.InstantiateContract(sdk.WrapSDKContext(s.ctx), &initMsg)
	s.Require().NoError(err)

	err = s.app.GlobalFeeKeeper.SetContractAuthorization(s.ctx, types.ContractAuthorization{
		ContractAddress: instantiateRes.Address,
		Methods:         authMethods,
	})
	s.Require().NoError(err)

	return instantiateRes.Address
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

func storeContract(ctx sdk.Context, msgServer wasmtypes.MsgServer, creator string, contract string) (uint64, error) {
	b, err := getContractBytes(contract)
	if err != nil {
		return 0, err
	}
	res, err := msgServer.StoreCode(sdk.WrapSDKContext(ctx), &wasmtypes.MsgStoreCode{
		Sender:       creator,
		WASMByteCode: b,
	})
	if err != nil {
		return 0, err
	}
	return res.CodeID, nil
}

func getTestAccount() (privateKey secp256k1.PrivKey, publicKey crypto.PubKey, accountAddress sdk.AccAddress) {
	privateKey = secp256k1.GenPrivKey()
	publicKey = privateKey.PubKey()
	accountAddress = sdk.AccAddress(publicKey.Address())
	return
}

func getContractBytes(contract string) ([]byte, error) {
	contractsCache.Lock()
	bz, found := contractsCache.contracts[contract]
	contractsCache.Unlock()
	if found {
		return bz, nil
	}
	contractsCache.Lock()
	defer contractsCache.Unlock()
	var err error
	bz, err = os.ReadFile(contract)
	if err != nil {
		return nil, err
	}
	contractsCache.contracts[contract] = bz
	return bz, nil
}

type CounterInsantiateMsg struct {
	Count uint64 `json:"count"`
}
