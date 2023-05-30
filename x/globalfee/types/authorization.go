package types

import (
	"regexp"
	"strings"

	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// validateMethods validates the method names for the code and contract authorizations
// valid values are alphanumeric names and '*' which is wildcard for all methods of given codeid/contract
// e.g "mint", "test_burn", "*", "list2"
func validateMethods(caMethods []string) error {
	if len(caMethods) == 0 {
		return sdkErrors.Wrap(ErrInvalidMethods, "empty methods provided")
	}

	pattern := `^(\*|[\w]+)$` // only allow method names or "*"
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	for _, method := range caMethods {
		if len(strings.TrimSpace(method)) == 0 {
			return sdkErrors.Wrap(ErrInvalidMethods, "method name is empty")
		}

		if !matcher.MatchString(method) {
			return sdkErrors.Wrap(ErrInvalidMethods, "invalid method name")
		}
	}

	return nil
}
