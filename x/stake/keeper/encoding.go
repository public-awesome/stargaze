package keeper

import "github.com/public-awesome/stargaze/x/stake/types"

// MustMarshalStake attempts to encode a Stake object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalStake(stake types.Stake) []byte {
	return k.cdc.MustMarshal(&stake)
}

// MustUnmarshalStake attempts to decode a Stake object and return it. It panics on error.
func (k Keeper) MustUnmarshalStake(data []byte, stake *types.Stake) {
	k.cdc.MustUnmarshal(data, stake)
}
