package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
)

// SetPausedCodeID stores a paused code ID.
func (k Keeper) SetPausedCodeID(ctx sdk.Context, pc types.PausedCodeID) error {
	return k.PausedCodeIDs.Set(ctx, pc.CodeID, pc)
}

// GetPausedCodeID returns a paused code ID.
func (k Keeper) GetPausedCodeID(ctx sdk.Context, codeID uint64) (types.PausedCodeID, error) {
	return k.PausedCodeIDs.Get(ctx, codeID)
}

// DeletePausedCodeID removes a paused code ID.
func (k Keeper) DeletePausedCodeID(ctx sdk.Context, codeID uint64) error {
	return k.PausedCodeIDs.Remove(ctx, codeID)
}

// IsCodeIDPaused checks if a code ID is paused.
func (k Keeper) IsCodeIDPaused(ctx sdk.Context, codeID uint64) bool {
	has, err := k.PausedCodeIDs.Has(ctx, codeID)
	if err != nil {
		return false
	}
	return has
}

// IteratePausedCodeIDs iterates over all paused code IDs.
func (k Keeper) IteratePausedCodeIDs(ctx sdk.Context, cb func(types.PausedCodeID) bool) {
	iterator, err := k.PausedCodeIDs.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		pc, err := iterator.Value()
		if err != nil {
			panic(err)
		}
		if cb(pc) {
			return
		}
	}
}
