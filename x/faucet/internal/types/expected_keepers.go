package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper is required for mining coin
type BankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(
		ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins,
	) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

// StakingKeeper is required for getting Denom
type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
}
