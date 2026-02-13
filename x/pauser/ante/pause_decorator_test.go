package ante_test

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stargazeapp "github.com/public-awesome/stargaze/v17/app"
	"github.com/public-awesome/stargaze/v17/testutil/simapp"
	pauserante "github.com/public-awesome/stargaze/v17/x/pauser/ante"
	pausertypes "github.com/public-awesome/stargaze/v17/x/pauser/types"
	"github.com/stretchr/testify/suite"
)

type PauseDecoratorTestSuite struct {
	suite.Suite

	app       *stargazeapp.App
	msgServer wasmtypes.MsgServer
	ctx       sdk.Context
	clientCtx client.Context
	txBuilder client.TxBuilder
}

func TestPauseDecoratorTestSuite(t *testing.T) {
	suite.Run(t, new(PauseDecoratorTestSuite))
}

type storeCache struct {
	sync.Mutex
	contracts map[string][]byte
}

var contractsCache = storeCache{contracts: make(map[string][]byte)}

type counterInstantiateMsg struct {
	Count uint64 `json:"count"`
}

func (s *PauseDecoratorTestSuite) SetupTest() {
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
			MaxGas: 225_000_000,
		},
	})

	encodingConfig := stargazeapp.MakeEncodingConfig()

	s.app = app
	s.ctx = ctx
	s.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
}

func (s *PauseDecoratorTestSuite) SetupWasmMsgServer() {
	wasmParams := s.app.Keepers.WasmKeeper.GetParams(s.ctx)
	wasmParams.CodeUploadAccess = wasmtypes.AllowEverybody
	err := s.app.Keepers.WasmKeeper.SetParams(s.ctx, wasmParams)
	s.Require().NoError(err)
	s.msgServer = wasmkeeper.NewMsgServerImpl(&s.app.Keepers.WasmKeeper)
}

func (s *PauseDecoratorTestSuite) DeployContract(senderAddr string) (string, uint64) {
	b, err := getContractBytes("../../globalfee/ante/counter.wasm")
	s.Require().NoError(err)

	storeRes, err := s.msgServer.StoreCode(s.ctx, &wasmtypes.MsgStoreCode{
		Sender:       senderAddr,
		WASMByteCode: b,
	})
	s.Require().NoError(err)

	initMsg := counterInstantiateMsg{Count: 0}
	initMsgRaw, err := json.Marshal(&initMsg)
	s.Require().NoError(err)

	instantiateRes, err := s.msgServer.InstantiateContract(s.ctx, &wasmtypes.MsgInstantiateContract{
		Sender: senderAddr,
		Admin:  senderAddr,
		CodeID: storeRes.CodeID,
		Label:  "Counter Contract",
		Msg:    initMsgRaw,
		Funds:  sdk.NewCoins(),
	})
	s.Require().NoError(err)

	return instantiateRes.Address, storeRes.CodeID
}

func (s *PauseDecoratorTestSuite) CreateTestTx(
	ctx sdk.Context, privs []cryptotypes.PrivKey,
	accNums, accSeqs []uint64,
	chainID string, signMode signing.SignMode,
) (xauthsigning.Tx, error) {
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

func (s *PauseDecoratorTestSuite) TestPausedContractRejected() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, _ := s.DeployContract(addr1.String())

	// Pause the contract
	err := s.app.Keepers.PauserKeeper.SetPausedContract(s.ctx, pausertypes.PausedContract{
		ContractAddress: contractAddr,
		PausedBy:        addr1.String(),
	})
	s.Require().NoError(err)

	// Build a MsgExecuteContract tx
	executeMsg := []byte(`{"increment": {}}`)
	s.txBuilder.SetMsgs(&wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractAddr,
		Msg:      executeMsg,
	})
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().Error(err)
	s.Require().ErrorIs(err, pausertypes.ErrContractPaused)
}

func (s *PauseDecoratorTestSuite) TestUnpausedContractAllowed() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, _ := s.DeployContract(addr1.String())

	// Contract is NOT paused - tx should pass through
	executeMsg := []byte(`{"increment": {}}`)
	s.txBuilder.SetMsgs(&wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractAddr,
		Msg:      executeMsg,
	})
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().NoError(err)
}

func (s *PauseDecoratorTestSuite) TestPausedCodeIDRejected() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, codeID := s.DeployContract(addr1.String())

	// Pause via code ID
	err := s.app.Keepers.PauserKeeper.SetPausedCodeID(s.ctx, pausertypes.PausedCodeID{
		CodeID:   codeID,
		PausedBy: addr1.String(),
	})
	s.Require().NoError(err)

	// Build a MsgExecuteContract tx
	executeMsg := []byte(`{"increment": {}}`)
	s.txBuilder.SetMsgs(&wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractAddr,
		Msg:      executeMsg,
	})
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().Error(err)
	s.Require().ErrorIs(err, pausertypes.ErrContractPaused)
}

func (s *PauseDecoratorTestSuite) TestNonExecuteMsgAllowed() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, _ := s.DeployContract(addr1.String())

	// Pause the contract
	err := s.app.Keepers.PauserKeeper.SetPausedContract(s.ctx, pausertypes.PausedContract{
		ContractAddress: contractAddr,
		PausedBy:        addr1.String(),
	})
	s.Require().NoError(err)

	// Send a MsgInstantiateContract (not MsgExecuteContract) - should pass through
	initMsg, _ := json.Marshal(&counterInstantiateMsg{Count: 0})
	s.txBuilder.SetMsgs(&wasmtypes.MsgInstantiateContract{
		Sender: addr1.String(),
		Admin:  addr1.String(),
		CodeID: 1,
		Label:  "Another Contract",
		Msg:    initMsg,
		Funds:  sdk.NewCoins(),
	})
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().NoError(err)
}

func (s *PauseDecoratorTestSuite) TestAuthzExecPausedContractRejected() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, _ := s.DeployContract(addr1.String())

	// Pause the contract
	err := s.app.Keepers.PauserKeeper.SetPausedContract(s.ctx, pausertypes.PausedContract{
		ContractAddress: contractAddr,
		PausedBy:        addr1.String(),
	})
	s.Require().NoError(err)

	// Wrap MsgExecuteContract inside authz.MsgExec
	innerMsg := &wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractAddr,
		Msg:      []byte(`{"increment": {}}`),
	}
	innerMsgAny, err := codectypes.NewAnyWithValue(innerMsg)
	s.Require().NoError(err)

	authzExec := &authz.MsgExec{
		Grantee: addr1.String(),
		Msgs:    []*codectypes.Any{innerMsgAny},
	}

	s.txBuilder.SetMsgs(authzExec)
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	// authz-wrapped MsgExecuteContract to a paused contract must be rejected
	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().Error(err)
	s.Require().ErrorIs(err, pausertypes.ErrContractPaused)
}

func (s *PauseDecoratorTestSuite) TestAuthzExecUnpausedContractAllowed() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, _ := s.DeployContract(addr1.String())

	// Contract is NOT paused
	innerMsg := &wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractAddr,
		Msg:      []byte(`{"increment": {}}`),
	}
	innerMsgAny, err := codectypes.NewAnyWithValue(innerMsg)
	s.Require().NoError(err)

	authzExec := &authz.MsgExec{
		Grantee: addr1.String(),
		Msgs:    []*codectypes.Any{innerMsgAny},
	}

	s.txBuilder.SetMsgs(authzExec)
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().NoError(err)
}

func (s *PauseDecoratorTestSuite) TestAuthzExecNestedTooDeepRejected() {
	s.SetupTest()
	s.SetupWasmMsgServer()
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	priv1, _, addr1 := testdata.KeyTestPubAddr()

	contractAddr, _ := s.DeployContract(addr1.String())

	// Build a MsgExecuteContract at the innermost level
	innerMsg := &wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractAddr,
		Msg:      []byte(`{"increment": {}}`),
	}
	innerMsgAny, err := codectypes.NewAnyWithValue(innerMsg)
	s.Require().NoError(err)

	// Nest 3 levels of authz.MsgExec (exceeds maxNestedMsgDepth of 2)
	current := &authz.MsgExec{
		Grantee: addr1.String(),
		Msgs:    []*codectypes.Any{innerMsgAny},
	}
	for i := 0; i < 2; i++ {
		wrapped, err := codectypes.NewAnyWithValue(current)
		s.Require().NoError(err)
		current = &authz.MsgExec{
			Grantee: addr1.String(),
			Msgs:    []*codectypes.Any{wrapped},
		}
	}

	s.txBuilder.SetMsgs(current)
	s.txBuilder.SetGasLimit(200_000)

	testTx, err := s.CreateTestTx(s.ctx, []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}, "", signing.SignMode_SIGN_MODE_DIRECT)
	s.Require().NoError(err)

	pauseDecorator := pauserante.NewPauseDecorator(s.app.Keepers.PauserKeeper)
	anteHandler := sdk.ChainAnteDecorators(pauseDecorator)

	_, err = anteHandler(s.ctx, testTx, false)
	s.Require().Error(err)
	s.Require().ErrorIs(err, pausertypes.ErrNestedMsgTooDeep)
}

// DeployPauserTestContract stores and instantiates the pauser test contract.
// It returns (contractAddr, codeID).
func (s *PauseDecoratorTestSuite) DeployPauserTestContract(senderAddr string) (string, uint64) {
	b, err := getContractBytes("pauser_test_contract.wasm")
	s.Require().NoError(err)

	storeRes, err := s.msgServer.StoreCode(s.ctx, &wasmtypes.MsgStoreCode{
		Sender:       senderAddr,
		WASMByteCode: b,
	})
	s.Require().NoError(err)

	initMsgRaw := []byte(`{"count": 0}`)
	instantiateRes, err := s.msgServer.InstantiateContract(s.ctx, &wasmtypes.MsgInstantiateContract{
		Sender: senderAddr,
		Admin:  senderAddr,
		CodeID: storeRes.CodeID,
		Label:  "Pauser Test Contract",
		Msg:    initMsgRaw,
		Funds:  sdk.NewCoins(),
	})
	s.Require().NoError(err)

	return instantiateRes.Address, storeRes.CodeID
}

// InstantiatePauserTestContract instantiates another instance of an already-stored code ID.
func (s *PauseDecoratorTestSuite) InstantiatePauserTestContract(senderAddr string, codeID uint64) string {
	initMsgRaw := []byte(`{"count": 0}`)
	instantiateRes, err := s.msgServer.InstantiateContract(s.ctx, &wasmtypes.MsgInstantiateContract{
		Sender: senderAddr,
		Admin:  senderAddr,
		CodeID: codeID,
		Label:  "Pauser Test Contract 2",
		Msg:    initMsgRaw,
		Funds:  sdk.NewCoins(),
	})
	s.Require().NoError(err)
	return instantiateRes.Address
}

func (s *PauseDecoratorTestSuite) TestContractToContractPausedRejected() {
	s.SetupTest()
	s.SetupWasmMsgServer()

	_, _, addr1 := testdata.KeyTestPubAddr()

	// Deploy two instances of the pauser test contract
	contractA, codeID := s.DeployPauserTestContract(addr1.String())
	contractB := s.InstantiatePauserTestContract(addr1.String(), codeID)

	// Pause contract B
	err := s.app.Keepers.PauserKeeper.SetPausedContract(s.ctx, pausertypes.PausedContract{
		ContractAddress: contractB,
		PausedBy:        addr1.String(),
	})
	s.Require().NoError(err)

	// Contract A calls CallIncrement targeting paused contract B
	callMsg := fmt.Sprintf(`{"call_increment": {"contract": "%s"}}`, contractB)
	_, err = s.msgServer.ExecuteContract(s.ctx, &wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractA,
		Msg:      []byte(callMsg),
		Funds:    sdk.NewCoins(),
	})

	// The wasm handler decorator should block the contract-to-contract call
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "contract is paused")
}

func (s *PauseDecoratorTestSuite) TestContractToContractUnpausedAllowed() {
	s.SetupTest()
	s.SetupWasmMsgServer()

	_, _, addr1 := testdata.KeyTestPubAddr()

	// Deploy two instances of the pauser test contract
	contractA, codeID := s.DeployPauserTestContract(addr1.String())
	contractB := s.InstantiatePauserTestContract(addr1.String(), codeID)

	// Contract B is NOT paused - call should succeed
	callMsg := fmt.Sprintf(`{"call_increment": {"contract": "%s"}}`, contractB)
	_, err := s.msgServer.ExecuteContract(s.ctx, &wasmtypes.MsgExecuteContract{
		Sender:   addr1.String(),
		Contract: contractA,
		Msg:      []byte(callMsg),
		Funds:    sdk.NewCoins(),
	})
	s.Require().NoError(err)

	// Verify contract B's count was incremented
	queryRes, err := s.app.Keepers.WasmKeeper.QuerySmart(s.ctx, sdk.MustAccAddressFromBech32(contractB), []byte(`{"get_count": {}}`))
	s.Require().NoError(err)
	s.Require().Contains(string(queryRes), `"count":1`)
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
	bz, err := os.ReadFile(contract)
	if err != nil {
		return nil, err
	}
	contractsCache.contracts[contract] = bz
	return bz, nil
}
