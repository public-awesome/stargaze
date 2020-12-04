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
	paramstore    paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryMarshaler, key sdk.StoreKey, ak types.AccountKeeper, bk types.BankKeeper,
	ps paramtypes.Subspace,
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

// RewardCreatorFromProtocol sends creator rewards from the protocol reward pool
func (k Keeper) RewardCreatorFromProtocol(
	ctx sdk.Context, account sdk.AccAddress, matchPool sdk.Dec) error {

	k.Logger(ctx).Debug(fmt.Sprintf("match pool: %v", matchPool))

	creatorShare := k.GetParams(ctx).CreatorAllocation
	creatorMatch := creatorShare.Mul(matchPool).TruncateInt()
	k.Logger(ctx).Debug(fmt.Sprintf("creator match: %v", creatorMatch))

	err := k.sendProtocolReward(ctx,
		account, sdk.NewCoin(k.GetParams(ctx).StakeDenom, creatorMatch))
	if err != nil {
		return err
	}

	return nil
}

// SendVotingReward sends the reward from quadratic voting to the user
func (k Keeper) SendVotingReward(
	ctx sdk.Context, account sdk.AccAddress, curatorReward sdk.Int) error {

	rewardCoin := sdk.NewCoin(types.DefaultVoteDenom, curatorReward)
	k.Logger(ctx).Debug(fmt.Sprintf("curator reward: %v", rewardCoin))

	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.VotingPoolName, account, sdk.NewCoins(rewardCoin))
	if err != nil {
		return err
	}

	return nil
}

// SendMatchingReward sends curator rewards from the protocol reward pool
func (k Keeper) SendMatchingReward(
	ctx sdk.Context, account sdk.AccAddress, matchReward sdk.Dec) error {

	curatorShare := sdk.OneDec().Sub(k.GetParams(ctx).CreatorAllocation)
	curatorMatch := curatorShare.Mul(matchReward).TruncateInt()
	k.Logger(ctx).Debug(fmt.Sprintf("curator match: %v", curatorMatch))

	err := k.sendProtocolReward(ctx,
		account, sdk.NewCoin(k.GetParams(ctx).StakeDenom, curatorMatch))
	if err != nil {
		return err
	}

	return nil
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
