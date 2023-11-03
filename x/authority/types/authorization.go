package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetMsgAuthorization(msgTypeURL string, authorizations []*Authorization) (*Authorization, bool) {
	for _, auth := range authorizations {
		if auth.MsgTypeUrl == msgTypeURL {
			return auth, true
		}
	}
	return &Authorization{}, false
}

func (a Authorization) IsAuthorized(proposer string) bool {
	for _, addr := range a.GetAddresses() {
		if addr == proposer {
			return true
		}
	}
	return false
}

func (a Authorization) Validate() error {
	if len(a.GetAddresses()) == 0 {
		return fmt.Errorf("addresses cannot be empty")
	}
	for _, addr := range a.GetAddresses() {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	if len(a.MsgTypeUrl) == 0 {
		return fmt.Errorf("msg type url cannot be empty")
	}
	return nil
}
