package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// Keeper of the x/curating store
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryMarshaler
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	paramstore    paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, ak types.AccountKeeper,
	bk types.BankKeeper, stakingKeeper types.StakingKeeper, ps paramtypes.Subspace,
) Keeper {

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	// ensure reward pool module account is set
	if addr := ak.GetModuleAddress(types.RewardPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.RewardPoolName))
	}

	keeper := Keeper{
		storeKey:      key,
		cdc:           cdc,
		accountKeeper: ak,
		bankKeeper:    bk,
		stakingKeeper: stakingKeeper,
		paramstore:    ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) validateVendorID(ctx sdk.Context, vendorID uint32) error {
	if vendorID > k.GetParams(ctx).MaxVendors {
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid vendor_id %d", vendorID))
	}

	return nil
}

// RewardCreatorFromVotingPool sends creator rewards from the voting pool
func (k Keeper) RewardCreatorFromVotingPool(
	ctx sdk.Context, account sdk.AccAddress, votingPool sdk.Int) (sdk.Coin, error) {

	k.Logger(ctx).Debug(fmt.Sprintf("voting pool: %v", votingPool))

	creatorShare := k.GetParams(ctx).CreatorVotingRewardAllocation
	creatorAlloc := creatorShare.MulInt(votingPool).TruncateInt()
	k.Logger(ctx).Debug(fmt.Sprintf("creator allocation: %v", creatorAlloc))

	reward, err := k.SendVotingReward(ctx, account, creatorAlloc)
	if err != nil {
		return sdk.Coin{}, err
	}

	return reward, nil
}

// RewardCreatorFromProtocol sends creator rewards from the protocol reward pool
func (k Keeper) RewardCreatorFromProtocol(
	ctx sdk.Context, account sdk.AccAddress, matchPool sdk.Dec) (sdk.Coin, error) {

	k.Logger(ctx).Debug(fmt.Sprintf("match pool: %v", matchPool))

	creatorShare := k.GetParams(ctx).CreatorProtocolRewardAllocation
	creatorMatch := creatorShare.Mul(matchPool).TruncateInt()
	k.Logger(ctx).Debug(fmt.Sprintf("creator match: %v", creatorMatch))

	reward := sdk.NewCoin(k.GetParams(ctx).StakeDenom, creatorMatch)
	err := k.sendProtocolReward(ctx,
		account, reward)
	if err != nil {
		return sdk.Coin{}, err
	}

	return reward, nil
}

// SendVotingReward sends the reward from quadratic voting to the user
func (k Keeper) SendVotingReward(
	ctx sdk.Context, account sdk.AccAddress, reward sdk.Int) (sdk.Coin, error) {

	rewardCoin := sdk.NewCoin(types.DefaultVoteDenom, reward)
	k.Logger(ctx).Debug(fmt.Sprintf("reward coin: %v", rewardCoin))

	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.VotingPoolName, account, sdk.NewCoins(rewardCoin))
	if err != nil {
		return sdk.Coin{}, err
	}

	return rewardCoin, nil
}

// SendMatchingReward sends curator rewards from the protocol reward pool
func (k Keeper) SendMatchingReward(
	ctx sdk.Context, account sdk.AccAddress, matchReward sdk.Dec) (sdk.Coin, error) {

	curatorShare := sdk.OneDec().Sub(k.GetParams(ctx).CreatorProtocolRewardAllocation)
	curatorMatch := curatorShare.Mul(matchReward).TruncateInt()
	k.Logger(ctx).Debug(fmt.Sprintf("curator match: %v", curatorMatch))

	reward := sdk.NewCoin(k.GetParams(ctx).StakeDenom, curatorMatch)
	err := k.sendProtocolReward(ctx,
		account, reward)
	if err != nil {
		return sdk.Coin{}, err
	}

	return reward, nil
}

// sendProtocolReward sends the quadratic finance matching reward to the user
func (k Keeper) sendProtocolReward(ctx sdk.Context, account sdk.AccAddress, amt sdk.Coin) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.RewardPoolName, account, sdk.NewCoins(amt))
	if err != nil {
		return sdkerrors.Wrapf(err, "spending from reward pool")
	}

	return nil
}
