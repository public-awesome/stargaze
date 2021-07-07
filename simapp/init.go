package simapp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// SetupWithStakeDenom initializes a new SimApp. A Nop logger is set in SimApp.
func SetupWithStakeDenom(isCheckTx bool, stakeDenom string) *SimApp {
	db := dbm.NewMemDB()
	config := MakeEncodingConfig()
	app := NewSimApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 5, config, EmptyAppOptions{})
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := NewDefaultGenesisState()
		genesisState, err := initGenesis(config.Marshaler, genesisState, stakeDenom, DefaultUnbondingPeriod)

		if err != nil {
			panic(err)
		}
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")

		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func initGenesis(
	cdc codec.JSONCodec,
	appState GenesisState,
	stakeDenom,
	unbondingPeriod string,
) (GenesisState, error) {
	// migrate staking state
	if appState[stakingtypes.ModuleName] != nil {
		var stakingGenState stakingtypes.GenesisState
		err := cdc.UnmarshalJSON(appState[stakingtypes.ModuleName], &stakingGenState)
		if err != nil {
			return nil, err
		}

		stakingGenState.Params.BondDenom = stakeDenom

		d, err := time.ParseDuration(unbondingPeriod)
		if err != nil {
			return nil, fmt.Errorf("failed to parse unbonding period %w", err)
		}
		stakingGenState.Params.UnbondingTime = d

		appState[stakingtypes.ModuleName] = cdc.MustMarshalJSON(&stakingGenState)
	}

	// migrate crisis state
	if appState[crisistypes.ModuleName] != nil {
		var crisisGenState crisistypes.GenesisState
		err := cdc.UnmarshalJSON(appState[crisistypes.ModuleName], &crisisGenState)
		if err != nil {
			return nil, err
		}
		crisisGenState.ConstantFee.Denom = stakeDenom
		appState[crisistypes.ModuleName] = cdc.MustMarshalJSON(&crisisGenState)
	}

	// migrate gov state
	if appState[govtypes.ModuleName] != nil {
		var govGenState govtypes.GenesisState
		err := cdc.UnmarshalJSON(appState[govtypes.ModuleName], &govGenState)
		if err != nil {
			return nil, err
		}
		minDeposit := sdk.NewInt64Coin(stakeDenom, 10_000_000)
		govGenState.DepositParams.MinDeposit = sdk.NewCoins(minDeposit)
		appState[govtypes.ModuleName] = cdc.MustMarshalJSON(&govGenState)
	}
	// migrate mint state
	if appState[minttypes.ModuleName] != nil {
		var mintGenState minttypes.GenesisState
		err := cdc.UnmarshalJSON(appState[minttypes.ModuleName], &mintGenState)
		if err != nil {
			return nil, err
		}
		mintGenState.Params.MintDenom = stakeDenom
		appState[minttypes.ModuleName] = cdc.MustMarshalJSON(&mintGenState)
	}

	// migrate curating state
	if appState[curatingtypes.ModuleName] != nil {
		var curatingGenState curatingtypes.GenesisState
		err := cdc.UnmarshalJSON(appState[curatingtypes.ModuleName], &curatingGenState)
		if err != nil {
			return nil, err
		}

		curatingGenState.Params.StakeDenom = stakeDenom
		curatingGenState.Params.InitialRewardPool = sdk.NewCoin(stakeDenom, curatingGenState.Params.InitialRewardPool.Amount)

		appState[curatingtypes.ModuleName] = cdc.MustMarshalJSON(&curatingGenState)
	}

	return appState, nil
}
