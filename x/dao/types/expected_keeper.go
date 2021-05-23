package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
When a module wishes to interact with another module, it is good practice to define what it will use
as an interface so the module cannot use things that are not permitted.
*/

// DistKeeper defines the expected interface for the distribution module
type DistKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}
