package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given annual provisions values
func NewMinter(annualProvisions sdk.Dec) Minter {
	return Minter{
		AnnualProvisions: annualProvisions,
	}
}

// InitialMinter returns an initial Minter object
func InitialMinter() Minter {
	return NewMinter(
		sdk.NewDec(0),
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
func DefaultInitialMinter() Minter {
	return InitialMinter()
}

// validate minter
func ValidateMinter(minter Minter) error {
	return nil
}

// NextAnnualProvisions returns the next annual provisions
func (m Minter) NextAnnualProvisions(blockTime time.Time, params Params) sdk.Dec {
	if params.StartTime.After(blockTime) {
		return sdk.ZeroDec()
	}

	return params.InitialAnnualProvisions.
		Mul(params.ReductionFactor.Power(currentYear(blockTime, params.StartTime)))
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params) sdk.Coin {
	provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))
	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}

func currentYear(blockTime time.Time, startTime time.Time) uint64 {
	delta := blockTime.Sub(startTime)
	year := sdk.NewInt(int64(delta)).QuoRaw(int64(365 * 24 * time.Hour))

	return year.Uint64()
}
