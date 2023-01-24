package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v8/testutil/keeper"
	"github.com/public-awesome/stargaze/v8/testutil/sample"
)

func Test_SetPrivileged(t *testing.T) {
	k, ctx := keeper.CronKeeper(t)

	acc1 := sample.AccAddress() // contract doesnt exist
	err := k.SetPrivileged(ctx, acc1)
	if err == nil {
		t.Errorf("expected %s to not exist, and fail to set privilege", acc1)
	}

	acc2 := sdk.MustAccAddressFromBech32("cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du")
	_ = k.SetPrivileged(ctx, acc2)
	if !k.IsPrivileged(ctx, acc2) {
		t.Errorf("expected %s to be privileged", acc1)
	}
}

func Test_UnsetPrivileged(t *testing.T) {
	k, ctx := keeper.CronKeeper(t)

	acc1 := sample.AccAddress() // contract doesnt exist
	err := k.UnsetPrivileged(ctx, acc1)
	if err == nil {
		t.Errorf("expected %s to not exist, and fail to set privilege", acc1)
	}

	acc2 := sdk.MustAccAddressFromBech32("cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du")
	err = k.UnsetPrivileged(ctx, acc2)
	if err == nil {
		t.Errorf("expected %s to not be privileged, and fail to unset privilege", acc1)
	}

	_ = k.SetPrivileged(ctx, acc2)
	_ = k.UnsetPrivileged(ctx, acc2)
	if k.IsPrivileged(ctx, acc1) {
		t.Errorf("expected %s to not be privileged", acc1)
	}
}

func Test_IteratePrivileged(t *testing.T) {
	k, ctx := keeper.CronKeeper(t)
	acc1 := sdk.MustAccAddressFromBech32("cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du")
	acc2 := sdk.MustAccAddressFromBech32("cosmos1hfml4tzwlc3mvynsg6vtgywyx00wfkhrtpkx6t")
	acc3 := sdk.MustAccAddressFromBech32("cosmos144sh8vyv5nqfylmg4mlydnpe3l4w780jsrmf4k")
	_ = k.SetPrivileged(ctx, acc1)
	_ = k.SetPrivileged(ctx, acc2)
	_ = k.SetPrivileged(ctx, acc3)

	_ = k.UnsetPrivileged(ctx, acc2)
	expectedContractCount := 2

	count := 0
	k.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		count += 1
		return false
	})
	if count != 2 {
		t.Errorf("expected %d, got %d", expectedContractCount, count)
	}
}
