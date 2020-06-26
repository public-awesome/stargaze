package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
)

/*
When a module wishes to interact with another module, it is good practice to define what it will use
as an interface so the module cannot use things that are not permitted.
*/

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	IterateAccounts(ctx sdk.Context, process func(authexported.Account) (stop bool))
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account // only used for simulation

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authexported.ModuleAccountI

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, authexported.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SetBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) error
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
