package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/pkg/errors"
	"github.com/rocket-protocol/mothership/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/types"
)

const (
	flagStakeDenom = "stake-denom"
)

func initGenesis(cdc codec.JSONMarshaler, genDoc *types.GenesisDoc, stakeDenom string) (genutil.AppMap, error) {
	var appState genutil.AppMap
	if err := cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
		return appState, errors.Wrap(err, "failed to JSON unmarshal initial genesis state")
	}

	// migrate staking state
	if appState[staking.ModuleName] != nil {
		var stakingGenState staking.GenesisState
		err := cdc.UnmarshalJSON(appState[staking.ModuleName], &stakingGenState)
		if err != nil {
			return appState, err
		}
		stakingGenState.Params.BondDenom = stakeDenom
		appState[staking.ModuleName] = cdc.MustMarshalJSON(stakingGenState)
	}
	// migrate crisis state
	if appState[crisis.ModuleName] != nil {
		var crisisGenState crisis.GenesisState
		err := cdc.UnmarshalJSON(appState[crisis.ModuleName], &crisisGenState)
		if err != nil {
			return appState, err
		}
		crisisGenState.ConstantFee.Denom = stakeDenom
		appState[crisis.ModuleName] = cdc.MustMarshalJSON(crisisGenState)
	}

	// migrate gov state
	if appState[gov.ModuleName] != nil {
		var govGenState gov.GenesisState
		err := cdc.UnmarshalJSON(appState[gov.ModuleName], &govGenState)
		if err != nil {
			return appState, err
		}
		minDeposit := sdk.NewInt64Coin(stakeDenom, 10_000_000)
		govGenState.DepositParams.MinDeposit = sdk.NewCoins(minDeposit)
		appState[gov.ModuleName] = cdc.MustMarshalJSON(govGenState)
	}
	// migrate mint state
	if appState[mint.ModuleName] != nil {
		var mintGenState mint.GenesisState
		err := cdc.UnmarshalJSON(appState[mint.ModuleName], &mintGenState)
		if err != nil {
			return appState, err
		}
		mintGenState.Params.MintDenom = stakeDenom
		appState[mint.ModuleName] = cdc.MustMarshalJSON(mintGenState)
	}
	return appState, nil
}

// InitCmd wraps the genutil.InitCmd to inject specific settings for rocket chain
func InitCmd(ctx *server.Context, cdc codec.JSONMarshaler, mbm module.BasicManager,
	defaultNodeHome string) *cobra.Command {

	init := genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome)
	init.PostRunE = func(cmd *cobra.Command, args []string) error {
		config := ctx.Config
		config.SetRoot(viper.GetString(cli.HomeFlag))

		genFile := config.GenesisFile()
		genDoc := &types.GenesisDoc{}

		if _, err := os.Stat(genFile); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			genDoc, err = types.GenesisDocFromFile(genFile)
			if err != nil {
				return errors.Wrap(err, "Failed to read genesis doc from file")
			}
		}
		stakeDenom := viper.GetString(flagStakeDenom)
		appState, err := initGenesis(cdc, genDoc, stakeDenom)
		if err != nil {
			return err
		}
		genDoc.AppState, err = cdc.MarshalJSON(appState)
		if err != nil {
			return errors.Wrap(err, "failed to JSON marshal migrated genesis state")
		}
		if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
			return errors.Wrap(err, "Failed to export gensis file")
		}
		return nil

	}

	init.Flags().String(flagStakeDenom, app.DefaultStakeDenom, "app's stake denom")
	return init
}
