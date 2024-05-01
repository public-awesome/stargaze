package cmd

import (
	"encoding/json"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	minttypes "github.com/public-awesome/stargaze/v14/x/mint/types"

	// appParams "github.com/public-awesome/stargaze/app/params"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	alloctypes "github.com/public-awesome/stargaze/v14/x/alloc/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	globalfeetypes "github.com/public-awesome/stargaze/v14/x/globalfee/types"
)

const (
	HumanCoinUnit       = "stars"
	BaseCoinUnit        = "ustars"
	StarsExponent       = 6
	Bech32PrefixAccAddr = "stars"
)

type GenesisParams struct {
	GenesisTime         time.Time
	NativeCoinMetadatas []banktypes.Metadata

	StakingParams      stakingtypes.Params
	DistributionParams distributiontypes.Params

	SlashingParams slashingtypes.Params

	AllocParams alloctypes.Params

	MintParams      minttypes.Params
	GlobalFeeParams globalfeetypes.Params
	WasmParams      wasmtypes.Params

	GovParams govtypes.Params
}

func PrepareGenesis(
	clientCtx client.Context,
	appState map[string]json.RawMessage,
	genesisParams GenesisParams,
) map[string]json.RawMessage {
	// IBC transfer module genesis
	ibcGenState := ibctransfertypes.DefaultGenesisState()
	ibcGenState.Params.SendEnabled = true
	ibcGenState.Params.ReceiveEnabled = true
	ibcGenStateBz := clientCtx.Codec.MustMarshalJSON(ibcGenState)
	appState[ibctransfertypes.ModuleName] = ibcGenStateBz

	// mint module genesis
	mintGenState := minttypes.DefaultGenesisState()
	mintGenState.Params = genesisParams.MintParams
	mintGenStateBz := clientCtx.Codec.MustMarshalJSON(mintGenState)
	appState[minttypes.ModuleName] = mintGenStateBz

	return appState
}

// params only
func DefaultGenesisParams() GenesisParams {
	genParams := GenesisParams{}

	genParams.GenesisTime = time.Now()

	genParams.NativeCoinMetadatas = []banktypes.Metadata{
		{
			Description: "The native token of Stargaze",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    BaseCoinUnit,
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    HumanCoinUnit,
					Exponent: StarsExponent,
					Aliases:  nil,
				},
			},
			Name:    "Stargaze STARS",
			Base:    BaseCoinUnit,
			Display: HumanCoinUnit,
			Symbol:  "STARS",
		},
	}

	// alloc
	genParams.AllocParams = alloctypes.DefaultParams()
	genParams.AllocParams.DistributionProportions = alloctypes.DistributionProportions{
		NftIncentives:    math.LegacyNewDecWithPrec(45, 2), // 45%
		DeveloperRewards: math.LegacyNewDecWithPrec(15, 2), // 15%
	}
	genParams.AllocParams.WeightedDeveloperRewardsReceivers = []alloctypes.WeightedAddress{}

	// mint
	genParams.MintParams = minttypes.DefaultParams()
	genParams.MintParams.MintDenom = BaseCoinUnit
	genParams.MintParams.StartTime = genParams.GenesisTime.AddDate(1, 0, 0)
	genParams.MintParams.InitialAnnualProvisions = math.LegacyNewDec(1_000_000_000_000_000)
	genParams.MintParams.ReductionFactor = math.LegacyNewDec(2).QuoInt64(3)
	genParams.MintParams.BlocksPerYear = uint64(5737588)

	genParams.StakingParams = stakingtypes.DefaultParams()
	genParams.StakingParams.UnbondingTime = time.Hour * 24 * 7 * 2 // 2 weeks
	genParams.StakingParams.MaxValidators = 10
	genParams.StakingParams.BondDenom = genParams.NativeCoinMetadatas[0].Base
	genParams.StakingParams.MinCommissionRate = math.LegacyNewDecWithPrec(5, 2)

	genParams.DistributionParams = distributiontypes.DefaultParams()
	genParams.DistributionParams.CommunityTax = math.LegacyMustNewDecFromStr("0")
	genParams.DistributionParams.WithdrawAddrEnabled = true

	genParams.GovParams = govtypes.DefaultParams()
	votingPeriod := time.Hour * 24 * 14
	genParams.GovParams.MaxDepositPeriod = &votingPeriod
	genParams.GovParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		math.NewInt(1_000_000_000),
	))
	genParams.GovParams.Quorum = math.LegacyMustNewDecFromStr("0.2").String()
	genParams.GovParams.VotingPeriod = &votingPeriod

	// genParams.CrisisConstantFee = sdk.NewCoin(
	// 	genParams.NativeCoinMetadatas[0].Base,
	// 	sdk.NewInt(100_000_000_000),
	// )

	genParams.SlashingParams = slashingtypes.DefaultParams()
	genParams.SlashingParams.SignedBlocksWindow = int64(25000)                              // ~41 hr at 6 second blocks
	genParams.SlashingParams.MinSignedPerWindow = math.LegacyMustNewDecFromStr("0.05")      // 5% minimum liveness
	genParams.SlashingParams.DowntimeJailDuration = time.Minute                             // 1 minute jail period
	genParams.SlashingParams.SlashFractionDoubleSign = math.LegacyMustNewDecFromStr("0.05") // 5% double sign slashing
	genParams.SlashingParams.SlashFractionDowntime = math.LegacyMustNewDecFromStr("0.0001") // 0.01% liveness slashing

	// genParams.ConsensusParams = tmtypes.DefaultConsensusParams()
	// genParams.ConsensusParams.Block.MaxBytes = 25 * 1024 * 1024 // 26,214,400 for cosmwasm
	// genParams.ConsensusParams.Block.MaxGas = 100_000_000
	// genParams.ConsensusParams.Evidence.MaxAgeDuration = genParams.StakingParams.UnbondingTime
	// genParams.ConsensusParams.Evidence.MaxAgeNumBlocks = int64(genParams.StakingParams.UnbondingTime.Seconds()) / 3
	// genParams.ConsensusParams.Version.AppVersion = 1

	genParams.WasmParams = wasmtypes.DefaultParams()

	genParams.GlobalFeeParams = globalfeetypes.DefaultParams()

	return genParams
}

// params only
func TestnetGenesisParams() GenesisParams {
	genParams := DefaultGenesisParams()

	genParams.GenesisTime = time.Now()

	// mint
	genParams.MintParams.StartTime = genParams.GenesisTime.Add(time.Minute * 5)

	genParams.GovParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(
		genParams.NativeCoinMetadatas[0].Base,
		math.NewInt(1_000_000),
	))

	votingPeriod := time.Minute * 15
	genParams.GovParams.Quorum = math.LegacyMustNewDecFromStr("0.1").String() // 10%
	genParams.GovParams.VotingPeriod = &votingPeriod                          // 15 min

	// alloc
	genParams.AllocParams = alloctypes.DefaultParams()
	genParams.AllocParams.DistributionProportions = alloctypes.DistributionProportions{
		NftIncentives:    math.LegacyNewDecWithPrec(30, 2), // 30%
		DeveloperRewards: math.LegacyNewDecWithPrec(30, 2), // 30%
	}
	genParams.AllocParams.WeightedDeveloperRewardsReceivers = []alloctypes.WeightedAddress{
		// faucet
		{
			Address: "stars1qpeu488858wm3uzqfz9e6m76s5jmjjtcuwr8e2",
			Weight:  math.LegacyNewDecWithPrec(80, 2),
		},
		{
			Address: "stars1fayut6xzyka29zvznsumlgy5pl4vkn4fkmaznc",
			Weight:  math.LegacyNewDecWithPrec(20, 2),
		},
	}
	genParams.WasmParams.CodeUploadAccess = wasmtypes.AllowEverybody
	genParams.WasmParams.InstantiateDefaultPermission = wasmtypes.AccessTypeEverybody

	return genParams
}
