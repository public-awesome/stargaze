package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v8/testutil/keeper"
	"github.com/public-awesome/stargaze/v8/testutil/sample"
)

func Test_SetPrivileged(t *testing.T) {
	k, ctx := keeper.CronKeeper(t)

	acc1 := sample.AccAddress()
	k.SetPrivileged(ctx, acc1)

	k.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		if addr.String() != acc1.String() {
			t.Errorf("expected %s, got %s", acc1, addr)
		}
		return false
	})
}

func Test_UnsetPrivileged(t *testing.T) {
	k, ctx := keeper.CronKeeper(t)

	acc1 := sample.AccAddress()
	k.SetPrivileged(ctx, acc1)
	k.UnsetPrivileged(ctx, acc1)

	k.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		if addr.String() == acc1.String() {
			t.Errorf("expected %s to not be there, got %s", acc1, addr)
		}
		return false
	})
}

func Test_IteratePrivileged(t *testing.T) {
	k, ctx := keeper.CronKeeper(t)
	acc1 := sample.AccAddress()
	k.SetPrivileged(ctx, acc1)
	acc2 := sample.AccAddress()
	k.SetPrivileged(ctx, acc2)
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
