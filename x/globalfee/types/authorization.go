package types

import (
	"regexp"
	"strings"

	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func validateMethods(caMethods []string) error {
	if len(caMethods) <= 0 {
		return sdkErrors.Wrap(ErrInvalidMethods, "empty methods provided")
	}

	for _, method := range caMethods {
		if len(strings.TrimSpace(method)) == 0 {
			return sdkErrors.Wrap(ErrInvalidMethods, "method name is empty")
		}

		pattern := `^(\*|[\w]+)$`
		matchFound, err := regexp.MatchString(pattern, method)
		if err != nil {
			return sdkErrors.Wrap(err, "")
		}
		if !matchFound {
			return sdkErrors.Wrap(ErrInvalidMethods, "invalid method name")
		}
	}

	return nil
}
