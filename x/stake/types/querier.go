package types

// query endpoints
const (
	QueryStake = "stake"
)

func NewQueryStakeRequest(vendorID uint32, postID string) *QueryStakeRequest {
	return &QueryStakeRequest{VendorId: vendorID, PostId: postID}
}
