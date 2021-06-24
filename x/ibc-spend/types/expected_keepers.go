package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
)

// AccountKeeper defines the account contract that must be fulfilled when
// creating a x/bank keeper.
type AccountKeeper interface {
	NewAccount(sdk.Context, types.AccountI) types.AccountI
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI

	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetAllAccounts(ctx sdk.Context) []types.AccountI
	SetAccount(ctx sdk.Context, acc types.AccountI)

	IterateAccounts(ctx sdk.Context, process func(types.AccountI) bool)

	ValidatePermissions(macc types.ModuleAccountI) error

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (types.ModuleAccountI, []string)
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc types.ModuleAccountI)
}

type TransferKeeper interface {
	SendTransfer(
		ctx sdk.Context,
		sourcePort,
		sourceChannel string,
		token sdk.Coin,
		sender sdk.AccAddress,
		receiver string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
	) error
}

type DistributionKeeper interface {
	DistributeFromFeePool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error
}
