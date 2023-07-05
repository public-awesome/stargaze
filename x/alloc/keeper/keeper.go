package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v11/x/alloc/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		stakingKeeper types.StakingKeeper
		distrKeeper   types.DistrKeeper

		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, stakingKeeper types.StakingKeeper, distrKeeper types.DistrKeeper,
	ps paramtypes.Subspace,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountKeeper, bankKeeper: bankKeeper, stakingKeeper: stakingKeeper, distrKeeper: distrKeeper,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetModuleAccountBalance gets the airdrop coin balance of module account
func (k Keeper) GetModuleAccountAddress(_ sdk.Context) sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.ModuleName)
}

// GetModuleAccountBalance gets the airdrop coin balance of module account
func (k Keeper) GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.AccountI {
	return k.accountKeeper.GetModuleAccount(ctx, moduleName)
}

func (k Keeper) sendToFairburnPool(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.FairburnPoolName, amount)
	return err
}

// DistributeInflation distributes module-specific inflation
func (k Keeper) DistributeInflation(ctx sdk.Context) error {

	denom := k.stakingKeeper.BondDenom(ctx)
	// get allocation params to retrieve distribution proportions
	params := k.GetParams(ctx)

	supplementPoolAddress := k.accountKeeper.GetModuleAccount(ctx, types.SupplementPoolName).GetAddress()
	supplementPoolBalance := k.bankKeeper.GetBalance(ctx, supplementPoolAddress, denom)

	// the amount that needs to be supplemented from the supplement pool
	supplementAmount := params.SupplementAmount.AmountOf(denom)

	distributionEvent := sdk.NewEvent(
		types.EventTypeDistribution,
	)
	// transfer supplement amount to be distributed to stakers if
	// 1- Supplement from params is not 0
	// 2- There is enough balance in the pool
	if !supplementAmount.IsZero() && supplementPoolBalance.Amount.GT(supplementAmount) {
		err := k.bankKeeper.SendCoinsFromModuleToModule(ctx,
			types.SupplementPoolName,
			authtypes.FeeCollectorName,
			sdk.NewCoins(sdk.NewCoin(denom, supplementAmount)),
		)
		if err != nil {
			return err
		}
		distributionEvent = distributionEvent.AppendAttributes(sdk.NewAttribute(types.AttributeKeySupplementAmount, supplementAmount.String()))
	}

	// retrieve balance from fee pool which is filled by minting new coins and by collecting transaction fees
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, denom)
	distributionEvent = distributionEvent.AppendAttributes(sdk.NewAttribute(types.AttributeKeyFeePoolAmount, blockInflation.String()))
	proportions := params.DistributionProportions

	if proportions.NftIncentives.GT(sdk.ZeroDec()) {
		incentiveRewards := k.GetProportions(ctx, blockInflation, proportions.NftIncentives)

		// Distribute NFT incentives to the community pool until a future update
		err := k.DistributeWeightedRewards(ctx, blockInflationAddr, incentiveRewards, params.WeightedIncentivesRewardsReceivers)
		if err != nil {
			return err
		}
		distributionEvent = distributionEvent.AppendAttributes(sdk.NewAttribute(types.AttributeKeyIncentivesAmount, incentiveRewards.String()))
	}

	// fund community pool if the value is not nil and greater than zero
	if !proportions.CommunityPool.IsNil() && proportions.CommunityPool.GT(sdk.ZeroDec()) {
		communityPoolTax := k.GetProportions(ctx, blockInflation, proportions.CommunityPool)
		err := k.distrKeeper.FundCommunityPool(ctx, sdk.NewCoins(communityPoolTax), blockInflationAddr)
		if err != nil {
			return err
		}
		distributionEvent = distributionEvent.AppendAttributes(sdk.NewAttribute(types.AttributeKeyCommunityPoolAmount, communityPoolTax.String()))
	}

	devRewards := k.GetProportions(ctx, blockInflation, proportions.DeveloperRewards)
	distributionEvent = distributionEvent.AppendAttributes(sdk.NewAttribute(types.AttributeKeyDevRewardsAmount, devRewards.String()))
	err := k.DistributeWeightedRewards(ctx, blockInflationAddr, devRewards, params.WeightedDeveloperRewardsReceivers)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		distributionEvent,
	})

	// fairburn pool
	fairburnPoolAddress := k.accountKeeper.GetModuleAccount(ctx, types.FairburnPoolName).GetAddress()
	collectedFairburnFees := k.bankKeeper.GetBalance(ctx, fairburnPoolAddress, denom)
	if collectedFairburnFees.IsZero() {
		return nil
	}
	// transfer collected fees from fairburn to the fee collector for distribution
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx,
		types.FairburnPoolName,
		authtypes.FeeCollectorName,
		sdk.NewCoins(collectedFairburnFees),
	)
	return err
}

// GetProportions gets the balance of the `MintedDenom` from minted coins
// and returns coins according to the `AllocationRatio`
func (k Keeper) GetProportions(_ sdk.Context, mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, mintedCoin.Amount.ToDec().Mul(ratio).TruncateInt())
}

func (k Keeper) DistributeWeightedRewards(ctx sdk.Context, feeCollectorAddress sdk.AccAddress, totalAllocation sdk.Coin, accounts []types.WeightedAddress) error {
	if totalAllocation.IsZero() {
		return nil
	}
	for _, w := range accounts {
		weightedReward := sdk.NewCoins(k.GetProportions(ctx, totalAllocation, w.Weight))
		if w.Address != "" {
			rewardAddress, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				return err
			}
			err = k.bankKeeper.SendCoins(ctx, feeCollectorAddress, rewardAddress, weightedReward)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (k Keeper) FundCommunityPool(ctx sdk.Context) error {
	// If this account exists and has coins, fund the community pool.
	// The address hardcoded here is randomly generated with no keypair behind it. It will be empty and unused after the genesis file is applied.
	funder, err := sdk.AccAddressFromHex("7C4954EAE77FF15A4C67C5F821C5241008ED966F") // stars103y4f6h80lc45nr8chuzr3fyzqywm9n0gnr394
	if err != nil {
		panic(err)
	}
	balances := k.bankKeeper.GetAllBalances(ctx, funder)
	if balances.IsZero() {
		return nil
	}
	return k.distrKeeper.FundCommunityPool(ctx, balances, funder)
}
