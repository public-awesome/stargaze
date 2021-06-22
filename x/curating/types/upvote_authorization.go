package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var (
	_ authz.Authorization = &UpvoteAuthorization{}
)

func (ua UpvoteAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgUpvote{})
}
func (ua UpvoteAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	return authz.AcceptResponse{Accept: true,
		Delete:  false,
		Updated: &UpvoteAuthorization{},
	}, nil
}
func (ua UpvoteAuthorization) ValidateBasic() error {
	return nil
}
