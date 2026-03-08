package app

import (
	"time"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/bytes"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	minttypes "github.com/public-awesome/stargaze/v17/x/mint/types"
)

// InitStargazeAppForTestnet is broken down into two sections:
// Required Changes: Changes that, if not made, will cause the testnet to halt or panic
// Optional Changes: Changes to customize the testnet to one's liking
func InitStargazeAppForTestnet(app *App, newValAddr bytes.HexBytes, newValPubKey crypto.PubKey, newOperatorAddress, upgradeToTrigger string) *App {
	//
	// Required Changes:
	//

	ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})

	pubkey := &ed25519.PubKey{Key: newValPubKey.Bytes()}
	pubkeyAny, err := codectypes.NewAnyWithValue(pubkey)
	if err != nil {
		tmos.Exit(err.Error())
	}

	// STAKING
	//

	// Create Validator struct for our new validator.
	_, bz, err := bech32.DecodeAndConvert(newOperatorAddress)
	if err != nil {
		tmos.Exit(err.Error())
	}
	bech32Addr, err := bech32.ConvertAndEncode(Bech32PrefixValAddr, bz)
	if err != nil {
		tmos.Exit(err.Error())
	}
	newVal := stakingtypes.Validator{
		OperatorAddress: bech32Addr,
		ConsensusPubkey: pubkeyAny,
		Jailed:          false,
		Status:          stakingtypes.Bonded,
		Tokens:          math.NewInt(900000000000000),
		DelegatorShares: math.LegacyMustNewDecFromStr("10000000"),
		Description: stakingtypes.Description{
			Moniker: "Testnet Validator",
		},
		Commission: stakingtypes.Commission{
			CommissionRates: stakingtypes.CommissionRates{
				Rate:          math.LegacyMustNewDecFromStr("0.05"),
				MaxRate:       math.LegacyMustNewDecFromStr("0.1"),
				MaxChangeRate: math.LegacyMustNewDecFromStr("0.05"),
			},
		},
		MinSelfDelegation: math.OneInt(),
	}

	// Remove all validators from power store
	stakingKey := app.GetKey(stakingtypes.ModuleName)
	stakingStore := ctx.KVStore(stakingKey)
	iterator, err := app.Keepers.StakingKeeper.ValidatorsPowerStoreIterator(ctx)
	if err != nil {
		tmos.Exit(err.Error())
	}
	for ; iterator.Valid(); iterator.Next() {
		stakingStore.Delete(iterator.Key())
	}
	iterator.Close()

	// Remove all validators from last validators store
	iterator, err = app.Keepers.StakingKeeper.LastValidatorsIterator(ctx)
	if err != nil {
		tmos.Exit(err.Error())
	}
	for ; iterator.Valid(); iterator.Next() {
		stakingStore.Delete(iterator.Key())
	}
	iterator.Close()

	// Add our validator to power and last validators store
	if err = app.Keepers.StakingKeeper.SetValidator(ctx, newVal); err != nil {
		tmos.Exit(err.Error())
	}
	if err = app.Keepers.StakingKeeper.SetValidatorByConsAddr(ctx, newVal); err != nil {
		tmos.Exit(err.Error())
	}
	if err = app.Keepers.StakingKeeper.SetValidatorByPowerIndex(ctx, newVal); err != nil {
		tmos.Exit(err.Error())
	}

	valAddr, err := sdk.ValAddressFromBech32(newVal.GetOperator())
	if err != nil {
		tmos.Exit(err.Error())
	}
	if err = app.Keepers.StakingKeeper.SetLastValidatorPower(ctx, valAddr, 0); err != nil {
		tmos.Exit(err.Error())
	}
	if err := app.Keepers.StakingKeeper.Hooks().AfterValidatorCreated(ctx, valAddr); err != nil {
		panic(err)
	}

	// DISTRIBUTION
	//

	// Initialize records for this validator across all distribution stores
	if err = app.Keepers.DistrKeeper.SetValidatorHistoricalRewards(ctx, valAddr, 0, distrtypes.NewValidatorHistoricalRewards(sdk.DecCoins{}, 1)); err != nil {
		tmos.Exit(err.Error())
	}
	if err = app.Keepers.DistrKeeper.SetValidatorCurrentRewards(ctx, valAddr, distrtypes.NewValidatorCurrentRewards(sdk.DecCoins{}, 1)); err != nil {
		tmos.Exit(err.Error())
	}
	if err = app.Keepers.DistrKeeper.SetValidatorAccumulatedCommission(ctx, valAddr, distrtypes.InitialValidatorAccumulatedCommission()); err != nil {
		tmos.Exit(err.Error())
	}
	if err = app.Keepers.DistrKeeper.SetValidatorOutstandingRewards(ctx, valAddr, distrtypes.ValidatorOutstandingRewards{Rewards: sdk.DecCoins{}}); err != nil {
		tmos.Exit(err.Error())
	}

	// SLASHING
	//

	// Set validator signing info for our new validator.
	newConsAddr := sdk.ConsAddress(newValAddr.Bytes())
	newValidatorSigningInfo := slashingtypes.ValidatorSigningInfo{
		Address:     newConsAddr.String(),
		StartHeight: app.LastBlockHeight() - 1,
		Tombstoned:  false,
	}
	if err = app.Keepers.SlashingKeeper.SetValidatorSigningInfo(ctx, newConsAddr, newValidatorSigningInfo); err != nil {
		tmos.Exit(err.Error())
	}

	//
	// Optional Changes:
	//

	// GOV — shorten voting period for testnet
	govParams, err := app.Keepers.GovKeeper.Params.Get(ctx)
	if err != nil {
		tmos.Exit(err.Error())
	}
	newExpeditedVotingPeriod := time.Minute
	newVotingPeriod := time.Minute * 2
	govParams.ExpeditedVotingPeriod = &newExpeditedVotingPeriod
	govParams.VotingPeriod = &newVotingPeriod
	govParams.MinDeposit = sdk.NewCoins(sdk.NewInt64Coin("ustars", 100000000))
	govParams.ExpeditedMinDeposit = sdk.NewCoins(sdk.NewInt64Coin("ustars", 150000000))

	if err = app.Keepers.GovKeeper.Params.Set(ctx, govParams); err != nil {
		tmos.Exit(err.Error())
	}

	// BANK — fund test accounts
	defaultCoins := sdk.NewCoins(
		sdk.NewInt64Coin("ustars", 1000000000000),                                                            // 1M STARS
		sdk.NewInt64Coin("ibc/9DF365E2C0EF4EA02FA771F638BB9C0C830EFCD354629BDC017F79B348B4E989", 1000000000), // ATOM
		sdk.NewInt64Coin("ibc/14D1406D84227FDF4B055EA5CB2298095BBCA3F3BC3EF583AE6DF36F0FB179C8", 1000000000), // TIA
		sdk.NewInt64Coin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", 1000000000), // OSMO
		sdk.NewInt64Coin("ibc/4A1C18CA7F50544760CF306189B810CE4C1CB156C7FC870143D401FE7280E591", 1000000000), // USDC
	)

	testAccounts := []sdk.AccAddress{
		sdk.MustAccAddressFromBech32("stars1s8qx0zvz8yd6e4x0mqmqf7fr9vvfn622wtp3g3"),
		sdk.MustAccAddressFromBech32("stars1g4762t55kn0hrxyrg0myyt68pll79cfykl7lvp"),
		sdk.MustAccAddressFromBech32("stars1kqzatlz85yfwvxexgq7qrhee03pgtmplazlcsc"),
	}

	for _, account := range testAccounts {
		if err := app.Keepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, defaultCoins); err != nil {
			tmos.Exit(err.Error())
		}
		if err := app.Keepers.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, account, defaultCoins); err != nil {
			tmos.Exit(err.Error())
		}
	}

	// UPGRADE
	if upgradeToTrigger != "" {
		upgradePlan := upgradetypes.Plan{
			Name:   upgradeToTrigger,
			Height: app.LastBlockHeight() + 10,
		}
		if err = app.Keepers.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan); err != nil {
			panic(err)
		}
	}

	return app
}
