package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/faucet/internal/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	defaultAmount int64 // set default amount for each mint.
	denomConfig   map[string]types.DenomConfig
	limit         time.Duration // rate limiting for mint, etc 24 * time.Hours
	storeKey      sdk.StoreKey  // Unexposed key to access store from sdk.Context

	cdc codec.BinaryCodec
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	defaultAmount int64,
	denomConfig map[string]types.DenomConfig,
	limit time.Duration) Keeper {
	return Keeper{
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
		defaultAmount: defaultAmount,
		denomConfig:   denomConfig,
		limit:         limit,
		storeKey:      key,
		cdc:           cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Limit returns the configured limit
func (k Keeper) Limit() time.Duration {
	return k.limit
}

func (k Keeper) getDenomConfig(denom string) types.DenomConfig {
	c, ok := k.denomConfig[denom]
	if !ok {
		return types.DenomConfig{
			Amount:         k.defaultAmount,
			BurnBeforeMint: false,
		}
	}
	return c
}

// MintAndSend mint coins and send to minter.
func (k Keeper) MintAndSend(ctx sdk.Context, minter sdk.AccAddress, mintTime int64, denom string) error {

	mining := k.getMining(ctx, minter, denom)

	// refuse mint in 24 hours
	if k.isPresent(ctx, minter, denom) &&
		time.Unix(mining.LastTime, 0).Add(k.limit).UTC().After(time.Unix(mintTime, 0)) {
		return types.ErrWithdrawTooOften
	}

	denomConfig := k.getDenomConfig(denom)
	balance := k.bankKeeper.GetBalance(ctx, minter, denom)
	if denomConfig.BurnBeforeMint && balance.IsValid() && !balance.IsZero() {
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, minter, types.ModuleName, sdk.NewCoins(balance))
		if err != nil {
			return err
		}
		err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
		if err != nil {
			return err
		}
	}
	newCoin := sdk.NewCoin(denom, sdk.NewInt(denomConfig.Amount))
	mining.Total = mining.Total.Add(newCoin)
	mining.LastTime = mintTime
	err := k.setMining(ctx, minter, mining, denom)
	if err != nil {
		return err
	}

	k.Logger(ctx).Info("mint coin: %s", newCoin)

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minter, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) getMining(ctx sdk.Context, minter sdk.AccAddress, denom string) types.Mining {
	store := ctx.KVStore(k.storeKey)
	if !k.isPresent(ctx, minter, denom) {
		return types.NewMining(minter, sdk.NewCoin(denom, sdk.NewInt(0)))
	}
	bz := store.Get(minterKey(minter, denom))
	var mining types.Mining
	k.cdc.MustUnmarshal(bz, &mining)
	return mining
}

func (k Keeper) setMining(ctx sdk.Context, minter sdk.AccAddress, mining types.Mining, denom string) error {
	if !mining.Total.IsPositive() {
		return types.ErrInvalidCoinAmount
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(minterKey(minter, denom), k.cdc.MustMarshal(&mining))
	return nil
}

func minterKey(minter sdk.AccAddress, denom string) []byte {
	mBytes := minter.Bytes()
	return append(mBytes, []byte(denom)...)
}

// IsPresent check if the name is present in the store or not
func (k Keeper) isPresent(ctx sdk.Context, minter sdk.AccAddress, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(minterKey(minter, denom))
}

// GetFaucetKey retrieves the faucet key from the store
func (k Keeper) GetFaucetKey(ctx sdk.Context) types.FaucetKey {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FaucetStoreKey)
	var faucet types.FaucetKey
	k.cdc.MustUnmarshal(bz, &faucet)
	return faucet
}

// SetFaucetKey sets the faucet key
func (k Keeper) SetFaucetKey(ctx sdk.Context, armor string) {
	store := ctx.KVStore(k.storeKey)
	faucet := types.NewFaucetKey(armor)
	store.Set(types.FaucetStoreKey, k.cdc.MustMarshal(&faucet))
}

// HasFaucetKey checks if faucet key is already stored.
func (k Keeper) HasFaucetKey(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.FaucetStoreKey)
}
