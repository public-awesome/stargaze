package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	appParams "github.com/public-awesome/stargaze/app/params"
)

func PrepareGenesisCmd(defaultNodeHome string, mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-genesis",
		Short: "Prepare a genesis file with initial setup",
		Long: `Prepare a genesis file with initial setup.
Examples include:
	- Setting module initial params
	- Setting denom metadata
Example:
	starsd prepare-genesis mainnet stargaze-1
	- Check input genesis:
		file is at ~/.starsd/config/genesis.json
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			depCdc := clientCtx.JSONMarshaler
			cdc := depCdc.(codec.Marshaler)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			// read genesis file
			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			// get genesis params
			var genesisParams GenesisParams
			network := args[0]
			if network == "testnet" {
				genesisParams = TestnetGenesisParams()
			} else if network == "mainnet" {
				genesisParams = MainnetGenesisParams()
			} else {
				return fmt.Errorf("please choose 'mainnet' or 'testnet'")
			}

			// get genesis params
			chainID := args[1]

			// run Prepare Genesis
			appState, genDoc, err = PrepareGenesis(clientCtx, appState, genDoc, genesisParams, chainID)

			// validate genesis state
			if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, appState); err != nil {
				return fmt.Errorf("error validating genesis file: %s", err.Error())
			}

			// save genesis
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			err = genutil.ExportGenesisFile(genDoc, genFile)
			return err
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func PrepareGenesis(clientCtx client.Context, appState map[string]json.RawMessage, genDoc *tmtypes.GenesisDoc, genesisParams GenesisParams, chainID string) (map[string]json.RawMessage, *tmtypes.GenesisDoc, error) {
	depCdc := clientCtx.JSONMarshaler
	cdc := depCdc.(codec.Marshaler)

	// chain params genesis
	genDoc.GenesisTime = genesisParams.GenesisTime

	genDoc.ConsensusParams = genesisParams.ConsensusParams

	// ---
	// staking module genesis
	stakingGenState := stakingtypes.GetGenesisStateFromAppState(depCdc, appState)
	stakingGenState.Params = genesisParams.StakingParams
	stakingGenStateBz, err := cdc.MarshalJSON(stakingGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal staking genesis state: %w", err)
	}
	appState[stakingtypes.ModuleName] = stakingGenStateBz

	// mint module genesis
	// mintGenState := minttypes.DefaultGenesisState()
	// mintGenState.Params = genesisParams.MintParams
	// mintGenStateBz, err := cdc.MarshalJSON(mintGenState)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to marshal mint genesis state: %w", err)
	// }
	// appState[minttypes.ModuleName] = mintGenStateBz

	// distribution module genesis
	distributionGenState := distributiontypes.DefaultGenesisState()
	distributionGenState.Params = genesisParams.DistributionParams
	// TODO Set initial community pool
	// distributionGenState.FeePool.CommunityPool = sdk.NewDecCoins()
	distributionGenStateBz, err := cdc.MarshalJSON(distributionGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal distribution genesis state: %w", err)
	}
	appState[distributiontypes.ModuleName] = distributionGenStateBz

	// gov module genesis
	govGenState := govtypes.DefaultGenesisState()
	govGenState.DepositParams = genesisParams.GovParams.DepositParams
	govGenState.TallyParams = genesisParams.GovParams.TallyParams
	govGenState.VotingParams = genesisParams.GovParams.VotingParams
	govGenStateBz, err := cdc.MarshalJSON(govGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal gov genesis state: %w", err)
	}
	appState[govtypes.ModuleName] = govGenStateBz

	// crisis module genesis
	crisisGenState := crisistypes.DefaultGenesisState()
	crisisGenState.ConstantFee = genesisParams.CrisisConstantFee
	// TODO Set initial community pool
	// distributionGenState.FeePool.CommunityPool = sdk.NewDecCoins()
	crisisGenStateBz, err := cdc.MarshalJSON(crisisGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal crisis genesis state: %w", err)
	}
	appState[crisistypes.ModuleName] = crisisGenStateBz

	// slashing module genesis
	slashingGenState := slashingtypes.DefaultGenesisState()
	slashingGenState.Params = genesisParams.SlashingParams
	slashingGenStateBz, err := cdc.MarshalJSON(slashingGenState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal slashing genesis state: %w", err)
	}
	appState[slashingtypes.ModuleName] = slashingGenStateBz

	// incentives module genesis
	// incentivesGenState := incentivestypes.GetGenesisStateFromAppState(depCdc, appState)
	// incentivesGenState.Params = genesisParams.IncentivesGenesis.Params
	// incentivesGenState.LockableDurations = genesisParams.IncentivesGenesis.LockableDurations
	// incentivesGenState.Gauges = genesisParams.IncentivesGenesis.Gauges
	// incentivesGenStateBz, err := cdc.MarshalJSON(incentivesGenState)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to marshal incentives genesis state: %w", err)
	// }
	// appState[incentivestypes.ModuleName] = incentivesGenStateBz

	// epochs module genesis
	// epochsGenState := epochstypes.DefaultGenesis()
	// epochsGenState.Epochs = genesisParams.Epochs
	// epochsGenStateBz, err := cdc.MarshalJSON(epochsGenState)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to marshal epochs genesis state: %w", err)
	// }
	// appState[epochstypes.ModuleName] = epochsGenStateBz

	// claim module genesis
	// claimGenState := claimtypes.GetGenesisStateFromAppState(depCdc, appState)
	// claimGenState.Params = genesisParams.ClaimParams
	// claimGenStateBz, err := cdc.MarshalJSON(claimGenState)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to marshal claim genesis state: %w", err)
	// }
	// appState[claimtypes.ModuleName] = claimGenStateBz

	// poolincentives module genesis
	// poolincentivesGenState := &genesisParams.PoolIncentivesGenesis
	// poolincentivesGenStateBz, err := cdc.MarshalJSON(poolincentivesGenState)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to marshal poolincentives genesis state: %w", err)
	// }
	// appState[poolincentivestypes.ModuleName] = poolincentivesGenStateBz

	// return appState and genDoc
	return appState, genDoc, nil
}

type GenesisParams struct {
	AirdropSupply sdk.Int

	StrategicReserveAccounts []banktypes.Balance

	ConsensusParams *tmproto.ConsensusParams

	GenesisTime         time.Time
	NativeCoinMetadatas []banktypes.Metadata

	StakingParams stakingtypes.Params
	// MintParams         minttypes.Params
	DistributionParams distributiontypes.Params
	GovParams          govtypes.Params

	CrisisConstantFee sdk.Coin

	SlashingParams slashingtypes.Params
	// IncentivesGenesis incentivestypes.GenesisState

	// PoolIncentivesGenesis poolincentivestypes.GenesisState

	// Epochs []epochstypes.EpochInfo

	// ClaimParams claimtypes.Params
}

func MainnetGenesisParams() GenesisParams {
	genParams := GenesisParams{}

	genParams.AirdropSupply = sdk.NewIntWithDecimal(5, 13)                // 5*10^13 ustars, 5*10^7 (50 million) stars
	genParams.GenesisTime = time.Date(2021, 6, 18, 17, 0, 0, 0, time.UTC) // Jun 18, 2021 - 17:00 UTC

	genParams.NativeCoinMetadatas = []banktypes.Metadata{
		{
			Description: fmt.Sprintf("The native token of Stargaze"),
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    appParams.BaseCoinUnit,
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    appParams.HumanCoinUnit,
					Exponent: appParams.StarsExponent,
					Aliases:  nil,
				},
			},
			Base:    appParams.BaseCoinUnit,
			Display: appParams.HumanCoinUnit,
		},
	}

	genParams.StrategicReserveAccounts = []banktypes.Balance{
		{
			Address: "stars1s4ckh9405q0a3jhkwx9wkf9hsjh66nmuu53dwe",
			Coins:   sdk.NewCoins(sdk.NewCoin(genParams.NativeCoinMetadatas[0].Base, sdk.NewInt(47_874_500_000_000))), // 47.8745 million STARS
		},
	}

	genParams.StakingParams = stakingtypes.DefaultParams()
	genParams.StakingParams.UnbondingTime = time.Hour * 24 * 7 * 2 // 2 weeks
	genParams.StakingParams.MaxValidators = 100
	genParams.StakingParams.BondDenom = genParams.NativeCoinMetadatas[0].Base
	genParams.StakingParams.MinCommissionRate = sdk.MustNewDecFromStr("0.05")

	genParams.DistributionParams = distributiontypes.DefaultParams()
	genParams.DistributionParams.BaseProposerReward = sdk.MustNewDecFromStr("0.01")
	genParams.DistributionParams.BonusProposerReward = sdk.MustNewDecFromStr("0.04")
	genParams.DistributionParams.CommunityTax = sdk.MustNewDecFromStr("0.05")
	genParams.DistributionParams.WithdrawAddrEnabled = true

	genParams.GovParams = govtypes.DefaultParams()
	genParams.GovParams.DepositParams.MaxDepositPeriod = time.Hour * 24 * 14 // 2 weeks
	genParams.GovParams.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(2_500_000_000),
	))
	genParams.GovParams.TallyParams.Quorum = sdk.MustNewDecFromStr("0.2") // 20%
	genParams.GovParams.VotingParams.VotingPeriod = time.Hour * 24 * 3    // 3 days

	genParams.CrisisConstantFee = sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(500_000_000_000),
	)

	genParams.SlashingParams = slashingtypes.DefaultParams()
	genParams.SlashingParams.SignedBlocksWindow = int64(30000)                       // 30000 blocks (~41 hr at 5 second blocks)
	genParams.SlashingParams.MinSignedPerWindow = sdk.MustNewDecFromStr("0.05")      // 5% minimum liveness
	genParams.SlashingParams.DowntimeJailDuration = time.Minute                      // 1 minute jail period
	genParams.SlashingParams.SlashFractionDoubleSign = sdk.MustNewDecFromStr("0.05") // 5% double sign slashing
	genParams.SlashingParams.SlashFractionDowntime = sdk.ZeroDec()                   // 0% liveness slashing

	// genParams.ClaimParams = claimtypes.Params{
	// 	AirdropStartTime:   genParams.GenesisTime,
	// 	DurationUntilDecay: time.Hour * 24 * 60,  // 60 days = ~2 months
	// 	DurationOfDecay:    time.Hour * 24 * 120, // 120 days = ~4 months
	// 	ClaimDenom:         genParams.NativeCoinMetadatas[0].Base,
	// }

	genParams.ConsensusParams = tmtypes.DefaultConsensusParams()
	genParams.ConsensusParams.Block.MaxBytes = 5 * 1024 * 1024
	genParams.ConsensusParams.Block.MaxGas = 6_000_000
	genParams.ConsensusParams.Evidence.MaxAgeDuration = genParams.StakingParams.UnbondingTime
	genParams.ConsensusParams.Evidence.MaxAgeNumBlocks = int64(genParams.StakingParams.UnbondingTime.Seconds()) / 3
	genParams.ConsensusParams.Version.AppVersion = 1

	return genParams
}

func TestnetGenesisParams() GenesisParams {

	genParams := MainnetGenesisParams()

	genParams.GenesisTime = time.Now()

	genParams.StakingParams.UnbondingTime = time.Hour * 24 * 7 * 2 // 2 weeks

	genParams.GovParams.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		sdk.NewInt(1000000), // 1 STARS
	))
	genParams.GovParams.TallyParams.Quorum = sdk.MustNewDecFromStr("0.0000000001") // 0.00000001%
	genParams.GovParams.VotingParams.VotingPeriod = time.Second * 300              // 300 seconds

	// genParams.ClaimParams.AirdropStartTime = genParams.GenesisTime
	// genParams.ClaimParams.DurationUntilDecay = time.Hour * 48 // 2 days
	// genParams.ClaimParams.DurationOfDecay = time.Hour * 48    // 2 days

	return genParams
}
