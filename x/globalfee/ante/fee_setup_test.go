package ante_test

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stargazeapp "github.com/public-awesome/stargaze/v15/app"
	"github.com/public-awesome/stargaze/v15/testutil/simapp"
	"github.com/public-awesome/stargaze/v15/x/globalfee/ante"
	"github.com/public-awesome/stargaze/v15/x/globalfee/types"
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
	_, _, acc1Addr := getTestAccount()
	_, _, acc2Addr := getTestAccount()
	genAccounts := authtypes.GenesisAccounts{
		&authtypes.BaseAccount{Address: acc1Addr.String()},
		&authtypes.BaseAccount{Address: acc2Addr.String()},
	}
	genBalances := []banktypes.Balance{
		{
			Address: acc1Addr.String(),
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5_000_000_000)),
		},
		{
			Address: acc2Addr.String(),
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5_000_000_000)),
		},
	}
	app := simapp.SetupWithGenesisAccounts(s.T(), s.T().TempDir(), genAccounts, genBalances...)
	h := cmtproto.Header{Height: app.LastBlockHeight() + 1}
	ctx := sdk.NewContext(app.CommitMultiStore(), h, false, app.Logger()).WithBlockTime(time.Now()).WithConsensusParams(cmtproto.ConsensusParams{
		Block: &cmtproto.BlockParams{
			MaxGas: 225_000_000, // 225M
		},
	})

	encodingConfig := stargazeapp.MakeEncodingConfig()

	s.app = app
	s.ctx = ctx
	s.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
}

func (s *AnteHandlerTestSuite) SetupTestGlobalFeeStoreAndMinGasPrice(minGasPrice []sdk.DecCoin, globalFees sdk.DecCoins) (ante.FeeDecorator, sdk.AnteHandler) {
	err := s.app.Keepers.GlobalFeeKeeper.SetParams(s.ctx, types.Params{MinimumGasPrices: globalFees})
	s.Require().NoError(err)

	s.ctx = s.ctx.WithMinGasPrices(minGasPrice).WithIsCheckTx(true).WithConsensusParams(cmtproto.ConsensusParams{
		Block: &cmtproto.BlockParams{
			MaxGas: 225_000_000, // 225M
		},
	})

	// build fee decorator
	feeDecorator := ante.NewFeeDecorator(s.app.AppCodec(), s.app.Keepers.GlobalFeeKeeper, s.app.Keepers.StakingKeeper)

	// chain fee decorator to antehandler
	antehandler := sdk.ChainAnteDecorators(feeDecorator)

	return feeDecorator, antehandler
}

func (s *AnteHandlerTestSuite) SetupWasmMsgServer() {
	wasmParams := s.app.Keepers.WasmKeeper.GetParams(s.ctx)
	wasmParams.CodeUploadAccess = wasmtypes.AllowEverybody
	err := s.app.Keepers.WasmKeeper.SetParams(s.ctx, wasmParams)
	s.Require().NoError(err)
	s.msgServer = wasmkeeper.NewMsgServerImpl(&s.app.Keepers.WasmKeeper)
}

func (s *AnteHandlerTestSuite) SetupContractWithCodeAuth(senderAddr string, contractBinary string, authMethods []string) string {
	codeID, err := storeContract(s.ctx, s.msgServer, senderAddr, contractBinary)
	s.Require().NoError(err)

	instantiageMsg := CounterInsantiateMsg{Count: 0}
	instantiateMsgRaw, err := json.Marshal(&instantiageMsg)
	s.Require().NoError(err)

	initMsg := wasmtypes.MsgInstantiateContract{Sender: senderAddr, Admin: senderAddr, CodeID: codeID, Label: "Counter Contract", Msg: instantiateMsgRaw, Funds: sdk.NewCoins()}
	instantiateRes, err := s.msgServer.InstantiateContract(s.ctx, &initMsg)
	s.Require().NoError(err)

	err = s.app.Keepers.GlobalFeeKeeper.SetCodeAuthorization(s.ctx, types.CodeAuthorization{
		CodeID:  codeID,
		Methods: authMethods,
	})
	s.Require().NoError(err)

	return instantiateRes.Address
}

func (s *AnteHandlerTestSuite) SetupContractWithContractAuth(senderAddr string, contractBinary string, authMethods []string) string {
	codeID, err := storeContract(s.ctx, s.msgServer, senderAddr, contractBinary)
	s.Require().NoError(err)

	instantiageMsg := CounterInsantiateMsg{Count: 0}
	instantiateMsgRaw, err := json.Marshal(&instantiageMsg)
	s.Require().NoError(err)

	initMsg := wasmtypes.MsgInstantiateContract{Sender: senderAddr, Admin: senderAddr, CodeID: codeID, Label: "Counter Contract", Msg: instantiateMsgRaw, Funds: sdk.NewCoins()}
	instantiateRes, err := s.msgServer.InstantiateContract(s.ctx, &initMsg)
	s.Require().NoError(err)

	err = s.app.Keepers.GlobalFeeKeeper.SetContractAuthorization(s.ctx, types.ContractAuthorization{
		ContractAddress: instantiateRes.Address,
		Methods:         authMethods,
	})
	s.Require().NoError(err)

	return instantiateRes.Address
}

func (s *AnteHandlerTestSuite) CreateTestTx(
	ctx sdk.Context, privs []cryptotypes.PrivKey,
	accNums, accSeqs []uint64,
	chainID string, signMode signing.SignMode,
) (xauthsigning.Tx, error) {
	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sigsV2 := make([]signing.SignatureV2, 0, len(privs))
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  signMode,
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err := s.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			Address:       sdk.AccAddress(priv.PubKey().Address()).String(),
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
			PubKey:        priv.PubKey(),
		}
		sigV2, err := tx.SignWithPrivKey(
			ctx, signMode, signerData,
			s.txBuilder, priv, s.clientCtx.TxConfig, accSeqs[i])
		if err != nil {
			return nil, err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = s.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	return s.txBuilder.GetTx(), nil
}

func storeContract(ctx sdk.Context, msgServer wasmtypes.MsgServer, creator string, contract string) (uint64, error) {
	b, err := getContractBytes(contract)
	if err != nil {
		return 0, err
	}
	res, err := msgServer.StoreCode(ctx, &wasmtypes.MsgStoreCode{
		Sender:       creator,
		WASMByteCode: b,
	})
	if err != nil {
		return 0, err
	}
	return res.CodeID, nil
}

func getTestAccount() (privateKey secp256k1.PrivKey, publicKey crypto.PubKey, accountAddress sdk.AccAddress) { //nolint:golint,unparam
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
