package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	stargaze "github.com/public-awesome/stargaze/app"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/types"
)

const (
	flagStakeDenom      = "stake-denom"
	flagUnbondingPeriod = "unbonding-period"
)

// InitCmd wraps the genutil.InitCmd to inject specific settings for stargaze chain
func InitCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {

	init := genutilcli.InitCmd(mbm, defaultNodeHome)

	init.PostRunE = func(cmd *cobra.Command, args []string) error {
		clientCtx := client.GetClientContextFromCmd(cmd)
		cdc := clientCtx.JSONCodec

		serverCtx := server.GetServerContextFromCmd(cmd)
		config := serverCtx.Config

		config.SetRoot(clientCtx.HomeDir)

		genFile := config.GenesisFile()
		genDoc := &types.GenesisDoc{}

		if _, err := os.Stat(genFile); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			genDoc, err = types.GenesisDocFromFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to read genesis doc from file %w", err)
			}
		}
		stakeDenom := viper.GetString(flagStakeDenom)
		unbondingPeriod := viper.GetString(flagUnbondingPeriod)
		appState, err := initGenesis(cdc, genDoc, stakeDenom, unbondingPeriod)
		if err != nil {
			return err
		}
		genDoc.AppState = appState
		if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
			return fmt.Errorf("failed to export gensis file %w", err)
		}
		return nil
	}

	init.Flags().String(flagStakeDenom, stargaze.DefaultStakeDenom, "app's stake denom")
	init.Flags().String(flagUnbondingPeriod, stargaze.DefaultUnbondingPeriod, "app's unbonding period")
	return init
}

func initGenesis(
	cdc codec.JSONCodec,
	genDoc *types.GenesisDoc,
	stakeDenom,
	unbondingPeriod string,
) (json.RawMessage, error) {
	appState := make(map[string]json.RawMessage)
	if err := tmjson.Unmarshal(genDoc.AppState, &appState); err != nil {
		return nil, fmt.Errorf("failed to JSON unmarshal initial genesis state %w", err)
	}
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

	// migrate liquidity state
	// if appState[liquiditytypes.ModuleName] != nil {
	// 	var liquidityGenState liquiditytypes.GenesisState
	// 	err := cdc.UnmarshalJSON(appState[liquiditytypes.ModuleName], &liquidityGenState)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	liquidityGenState.Params.LiquidityPoolCreationFee = sdk.NewCoins(sdk.NewCoin(stakeDenom, sdk.NewInt(100_000_000)))
	// 	appState[liquiditytypes.ModuleName] = cdc.MustMarshalJSON(&liquidityGenState)
	// }

	return tmjson.Marshal(appState)
}
