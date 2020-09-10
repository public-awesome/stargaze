package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/modules/incubator/faucet/internal/types"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

const FaucetStoreKey = "DefaultFaucetStoreKey"

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	SupplyKeeper  types.SupplyKeeper
	StakingKeeper types.StakingKeeper
	amount        int64         // set default amount for each mint.
	Limit         time.Duration // rate limiting for mint, etc 24 * time.Hours
	storeKey      sdk.StoreKey  // Unexposed key to access store from sdk.Context
	cdc           *codec.Codec  // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the Faucet Keeper
func NewKeeper(
	supplyKeeper types.SupplyKeeper,
	stakingKeeper types.StakingKeeper,
	amount int64,
	rateLimit time.Duration,
	storeKey sdk.StoreKey,
	cdc *codec.Codec) Keeper {
	return Keeper{
		SupplyKeeper:  supplyKeeper,
		StakingKeeper: stakingKeeper,
		amount:        amount,
		Limit:         rateLimit,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// MintAndSend mint coins and send to minter.
func (k Keeper) MintAndSend(ctx sdk.Context, minter sdk.AccAddress, mintTime int64) error {

	mining := k.getMining(ctx, minter)

	// refuse mint in 24 hours
	if k.isPresent(ctx, minter) &&
		time.Unix(mining.LastTime, 0).Add(k.Limit).UTC().After(time.Unix(mintTime, 0)) {
		return types.ErrWithdrawTooOften
	}

	denom := k.StakingKeeper.BondDenom(ctx)
	newCoin := sdk.NewCoin(denom, sdk.NewInt(k.amount))
	mining.Total = mining.Total.Add(newCoin)
	mining.LastTime = mintTime
	k.setMining(ctx, minter, mining)

	k.Logger(ctx).Info("Mint coin: %s", newCoin)

	err := k.SupplyKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minter, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) getMining(ctx sdk.Context, minter sdk.AccAddress) types.Mining {
	store := ctx.KVStore(k.storeKey)
	if !k.isPresent(ctx, minter) {
		denom := k.StakingKeeper.BondDenom(ctx)
		return types.NewMining(minter, sdk.NewCoin(denom, sdk.NewInt(0)))
	}
	bz := store.Get(minter.Bytes())
	var mining types.Mining
	k.cdc.MustUnmarshalBinaryBare(bz, &mining)
	return mining
}

func (k Keeper) setMining(ctx sdk.Context, minter sdk.AccAddress, mining types.Mining) {
	if mining.Minter.Empty() {
		return
	}
	if !mining.Total.IsPositive() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(minter.Bytes(), k.cdc.MustMarshalBinaryBare(mining))
}

// IsPresent check if the name is present in the store or not
func (k Keeper) isPresent(ctx sdk.Context, minter sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(minter.Bytes())
}

func (k Keeper) GetFaucetKey(ctx sdk.Context) types.FaucetKey {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(FaucetStoreKey))
	var faucet types.FaucetKey
	k.cdc.MustUnmarshalBinaryBare(bz, &faucet)
	return faucet
}

func (k Keeper) SetFaucetKey(ctx sdk.Context, armor string) {
	store := ctx.KVStore(k.storeKey)
	faucet := types.NewFaucetKey(armor)
	store.Set([]byte(FaucetStoreKey), k.cdc.MustMarshalBinaryBare(faucet))
}

func (k Keeper) HasFaucetKey(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(FaucetStoreKey))
}
