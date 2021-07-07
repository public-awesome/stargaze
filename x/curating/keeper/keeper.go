package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/x/curating/types"
)

// Keeper of the x/curating store
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	paramstore    paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, ak types.AccountKeeper, bk types.BankKeeper,
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

func (k Keeper) validatePostBodyLength(ctx sdk.Context, body string) error {
	if uint32(len(body)) > k.GetParams(ctx).MaxPostBodyLength {
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest, fmt.Sprintf("post body contains too many characters: %s", body))
	}

	return nil
}

// SendMatchingReward sends curator rewards from the protocol reward pool
func (k Keeper) SendMatchingReward(
	ctx sdk.Context, account sdk.AccAddress, matchReward sdk.Dec) (sdk.Coin, error) {

	reward := sdk.NewCoin(k.GetParams(ctx).StakeDenom, matchReward.TruncateInt())
	err := k.sendProtocolReward(ctx,
		account, reward)
	if err != nil {
		return sdk.Coin{}, err
	}

	return reward, nil
}

// BurnFromVotingPool burns an amount of CREDITS from the voting pool
func (k Keeper) BurnFromVotingPool(ctx sdk.Context, amount sdk.Int) error {
	votingCoin := sdk.NewCoin(types.DefaultVoteDenom, amount)
	err := k.bankKeeper.BurnCoins(ctx, types.VotingPoolName, sdk.NewCoins(votingCoin))
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
