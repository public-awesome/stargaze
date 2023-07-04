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
	blockInflationAddr := k.accountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress()
	blockInflation := k.bankKeeper.GetBalance(ctx, blockInflationAddr, k.stakingKeeper.BondDenom(ctx))
	blockInflationDec := sdk.NewDecFromInt(blockInflation.Amount)

	params := k.GetParams(ctx)
	proportions := params.DistributionProportions

	if proportions.NftIncentives.GT(sdk.ZeroDec()) {
		nftIncentiveAmount := blockInflationDec.Mul(proportions.NftIncentives)
		nftIncentiveCoin := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), nftIncentiveAmount.TruncateInt())
		// Distribute NFT incentives to the community pool until a future update
		err := k.DistributeWeightedRewards(ctx, blockInflationAddr, nftIncentiveCoin, params.WeightedIncentivesRewardsReceivers)
		if err != nil {
			return err
		}

		// iterate over list of icentive addresses and proportions
		k.Logger(ctx).Debug("funded community pool", "amount", nftIncentiveCoin.String(), "from", blockInflationAddr)
	}

	// fund community pool if the value is not nil and greater than zero
	if !proportions.CommunityPool.IsNil() && proportions.CommunityPool.GT(sdk.ZeroDec()) {
		communityPoolTax := k.GetProportions(ctx, blockInflation, proportions.CommunityPool)
		err := k.distrKeeper.FundCommunityPool(ctx, sdk.NewCoins(communityPoolTax), blockInflationAddr)
		if err != nil {
			return err
		}
	}

	devRewards := k.GetProportions(ctx, blockInflation, proportions.DeveloperRewards)
	err := k.DistributeWeightedRewards(ctx, blockInflationAddr, devRewards, params.WeightedDeveloperRewardsReceivers)
	if err != nil {
		return err
	}

	// fairburn pool
	fairburnPoolAddress := k.accountKeeper.GetModuleAccount(ctx, types.FairburnPoolName).GetAddress()
	collectedFairburnFees := k.bankKeeper.GetBalance(ctx, fairburnPoolAddress, k.stakingKeeper.BondDenom(ctx))
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
