package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyDistributionProportions  = []byte("DistributionProportions")
	KeyDeveloperRewardsReceiver = []byte("DeveloperRewardsReceiver")
	KeyIncentiveRewardsReceiver  = []byte("IncentiveRewardsReceiver")
)

// ParamTable for module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	distrProportions DistributionProportions,
	weightedDevRewardsReceivers []WeightedAddress,
) Params {
	return Params{
		DistributionProportions:           distrProportions,
		WeightedDeveloperRewardsReceivers: weightedDevRewardsReceivers,
	}
}

// default module parameters
func DefaultParams() Params {
	return Params{
		DistributionProportions: DistributionProportions{
			NftIncentives:    sdk.NewDecWithPrec(20, 2), // 20%
			DeveloperRewards: sdk.NewDecWithPrec(15, 2), // 15%
			CommunityPool:    sdk.NewDecWithPrec(5, 2),  // 5%
		},
		WeightedDeveloperRewardsReceivers:  []WeightedAddress{},
		WeightedIncentivesRewardsReceivers: []WeightedAddress{},
	}
}

// validate params
func (p Params) Validate() error {
	if err := validateDistributionProportions(p.DistributionProportions); err != nil {
		return err
	}
	if err := validateWeightedRewardsReceivers(p.WeightedDeveloperRewardsReceivers); err != nil {
		return err
	}
	return validateWeightedRewardsReceivers(p.WeightedIncentivesRewardsReceivers)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDistributionProportions, &p.DistributionProportions, validateDistributionProportions),
		paramtypes.NewParamSetPair(
			KeyDeveloperRewardsReceiver, &p.WeightedDeveloperRewardsReceivers, validateWeightedRewardsReceivers),
		paramtypes.NewParamSetPair(
			KeyIcentiveRewardsReceiver, &p.WeightedIncentivesRewardsReceivers, validateWeightedRewardsReceivers),
	}
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.NftIncentives.IsNegative() {
		return errors.New("NFT incentives distribution ratio should not be negative")
	}

	if v.DeveloperRewards.IsNegative() {
		return errors.New("developer rewards distribution ratio should not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("community pool ratio should not be negative")
	}

	totalProportions := v.NftIncentives.Add(v.DeveloperRewards).Add(v.CommunityPool)

	if totalProportions.GT(sdk.OneDec()) {
		return errors.New("total distributions can not be higher than 100%")
	}

	return nil
}

func validateWeightedRewardsReceivers(i interface{}) error {
	v, ok := i.([]WeightedAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// fund community pool when rewards address is empty
	if len(v) == 0 {
		return nil
	}

	weightSum := sdk.NewDec(0)
	for i, w := range v {
		// we allow address to be "" to go to community pool
		if w.Address != "" {
			_, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				return fmt.Errorf("invalid address at %dth", i)
			}
		}
		if !w.Weight.IsPositive() {
			return fmt.Errorf("non-positive weight at %dth", i)
		}
		if w.Weight.GT(sdk.NewDec(1)) {
			return fmt.Errorf("more than 1 weight at %dth", i)
		}
		weightSum = weightSum.Add(w.Weight)
	}

	if !weightSum.Equal(sdk.NewDec(1)) {
		return fmt.Errorf("invalid weight sum: %s", weightSum.String())
	}

	return nil
}
