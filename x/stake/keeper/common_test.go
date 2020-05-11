package keeper_test

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/davecgh/go-spew/spew"
	"github.com/rocket-protocol/stakebird/x/stake/testdata"
	"github.com/rocket-protocol/stakebird/x/stake/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

var (
	PKs = simapp.CreateTestPubKeys(5)

	valConsPk1 = PKs[0]
	valConsPk2 = PKs[1]
	valConsPk3 = PKs[2]

	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())

	distrAcc = auth.NewEmptyModuleAccount(types.ModuleName)
)

// type GenesisState map[string]json.RawMessage

// func NewDefaultGenesisState() GenesisState {
// 	cdc := std.MakeCodec(testdata.ModuleBasics)
// 	return testdata.ModuleBasics.DefaultGenesis(cdc)
// }

func createTestInput() (*codec.Codec, *testdata.SimApp, sdk.Context) {
	db := dbm.NewMemDB()
	logger := log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout))

	opts := []func(*baseapp.BaseApp){baseapp.SetPruning(store.PruneNothing)}
	app := testdata.NewSimApp(logger, db, nil, true, 0, map[int64]bool{}, simapp.DefaultNodeHome, opts...)

	genesisState := testdata.ModuleBasics.DefaultGenesis(app.Codec())
	// genesisState := NewDefaultGenesisState()
	// spew.Dump(genesisState)
	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: time.Now()}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := app.NewContext(false, header)

	spew.Dump("params", app.StakeKeeper.GetParams(ctx))

	return codec.New(), app, ctx
}

type GenerateAccountStrategy func(int) []sdk.AccAddress

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrsIncremental(app *testdata.SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	return addTestAddrs(app, ctx, accNum, accAmt, createIncrementalAccounts)
}

// ConvertAddrsToValAddrs converts the provided addresses to ValAddress.
func ConvertAddrsToValAddrs(addrs []sdk.AccAddress) []sdk.ValAddress {
	valAddrs := make([]sdk.ValAddress, len(addrs))

	for i, addr := range addrs {
		valAddrs[i] = sdk.ValAddress(addr)
	}

	return valAddrs
}

func addTestAddrs(app *testdata.SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int, strategy GenerateAccountStrategy) []sdk.AccAddress {
	testAddrs := strategy(accNum)

	initCoins := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), accAmt))
	setTotalSupply(app, ctx, accAmt, accNum)

	// fill all the addresses with some coins, set the loose pool tokens simultaneously
	for _, addr := range testAddrs {
		saveAccount(app, ctx, addr, initCoins)
	}

	return testAddrs
}

// setTotalSupply provides the total supply based on accAmt * totalAccounts.
func setTotalSupply(app *testdata.SimApp, ctx sdk.Context, accAmt sdk.Int, totalAccounts int) {
	totalSupply := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), accAmt.MulRaw(int64(totalAccounts))))
	prevSupply := app.BankKeeper.GetSupply(ctx)
	app.BankKeeper.SetSupply(ctx, bank.NewSupply(prevSupply.GetTotal().Add(totalSupply...)))
}

// saveAccount saves the provided account into the simapp with balance based on initCoins.
func saveAccount(app *testdata.SimApp, ctx sdk.Context, addr sdk.AccAddress, initCoins sdk.Coins) {
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	app.AccountKeeper.SetAccount(ctx, acc)
	_, err := app.BankKeeper.AddCoins(ctx, addr, initCoins)
	if err != nil {
		panic(err)
	}
}

// createIncrementalAccounts is a strategy used by addTestAddrs() in order to generated addresses in ascending order.
func createIncrementalAccounts(accNum int) []sdk.AccAddress {
	var addresses []sdk.AccAddress
	var buffer bytes.Buffer

	// start at 100 so we can make up to 999 test addresses with valid test addresses
	for i := 100; i < (accNum + 100); i++ {
		numString := strconv.Itoa(i)
		buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") //base address string

		buffer.WriteString(numString) //adding on final two digits to make addresses unique
		res, _ := sdk.AccAddressFromHex(buffer.String())
		bech := res.String()
		addr, _ := FakeAddr(buffer.String(), bech)

		addresses = append(addresses, addr)
		buffer.Reset()
	}

	return addresses
}

func FakeAddr(addr string, bech string) (sdk.AccAddress, error) {
	res, err := sdk.AccAddressFromHex(addr)
	if err != nil {
		return nil, err
	}
	bechexpected := res.String()
	if bech != bechexpected {
		return nil, fmt.Errorf("bech encoding doesn't match reference")
	}

	bechres, err := sdk.AccAddressFromBech32(bech)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(bechres, res) {
		return nil, err
	}

	return res, nil
}
