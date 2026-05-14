// Package sdkutil holds general-purpose Cosmos SDK helpers.
//
// The ApplyFuncIfNoError / IsOutOfGasError / PrintPanicRecoveryError helpers
// in this file are copied from Osmosis (osmoutils/cache_ctx.go,
// https://github.com/osmosis-labs/osmosis), licensed under Apache 2.0. They
// let callers run a function inside a cache context so that state changes are
// dropped if the function returns an error or panics, instead of halting the
// chain.
package sdkutil

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"

	"cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ApplyFuncIfNoError runs f inside a cache context. If f returns an error or
// panics, the state changes inside the cache context are discarded and the
// error is logged. If f succeeds the cache context is written.
//
// Avoid iterators inside f because of cache-store semantics. Out-of-gas panics
// are re-panicked to preserve normal tx execution flow; this is safe for
// BeginBlock / EndBlock code which is not gas-metered.
func ApplyFuncIfNoError(ctx sdk.Context, f func(ctx sdk.Context) error) (err error) {
	return applyFunc(ctx, f, ctx.Logger().Error)
}

// ApplyFuncIfNoErrorLogToDebug is the same as ApplyFuncIfNoError, but sends logs to debug instead of error if there is an error.
func ApplyFuncIfNoErrorLogToDebug(ctx sdk.Context, f func(ctx sdk.Context) error) (err error) {
	return applyFunc(ctx, f, ctx.Logger().Debug)
}

func applyFunc(ctx sdk.Context, f func(ctx sdk.Context) error, logFunc func(string, ...any)) (err error) {
	defer func() {
		if recoveryError := recover(); recoveryError != nil {
			if isErr, _ := IsOutOfGasError(recoveryError); isErr {
				panic(recoveryError)
			} else {
				PrintPanicRecoveryError(ctx, recoveryError)
				err = errors.New("panic occurred during execution")
			}
		}
	}()
	cacheCtx, write := ctx.CacheContext()
	err = f(cacheCtx)
	if err != nil {
		logFunc(err.Error())
	} else {
		write()
	}
	return err
}

// IsOutOfGasError returns true when the recovered panic value is an SDK
// out-of-gas / gas-overflow sentinel. The SDK encodes these as struct values
// that do not implement the error interface, so callers can't use errors.As.
func IsOutOfGasError(err any) (bool, string) {
	switch e := err.(type) {
	case types.ErrorOutOfGas:
		return true, e.Descriptor
	case types.ErrorGasOverflow:
		return true, e.Descriptor
	default:
		return false, ""
	}
}

// PrintPanicRecoveryError error logs the recoveryError, along with the stacktrace, if it can be parsed.
// If not emits them to stdout.
func PrintPanicRecoveryError(ctx sdk.Context, recoveryError any) {
	errStackTrace := string(debug.Stack())
	switch e := recoveryError.(type) {
	case types.ErrorOutOfGas:
		ctx.Logger().Debug("out of gas error inside panic recovery block: " + e.Descriptor)
		return
	case string:
		ctx.Logger().Error("Recovering from (string) panic: " + e)
	case runtime.Error:
		ctx.Logger().Error("recovered (runtime.Error) panic: " + e.Error())
	case error:
		ctx.Logger().Error("recovered (error) panic: " + e.Error())
	default:
		ctx.Logger().Error("recovered (default) panic. Could not capture logs in ctx, see stdout")
		fmt.Println("Recovering from panic ", recoveryError)
		debug.PrintStack()
		return
	}
	ctx.Logger().Error("stack trace: " + errStackTrace)
}
