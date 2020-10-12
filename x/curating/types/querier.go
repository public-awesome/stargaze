package types

// query endpoints
const (
	QueryPost   = "post"
	QueryUpvote = "upvote"
)

func NewQueryPostRequest(vendorID uint32, postID string) *QueryPostRequest {
	return &QueryPostRequest{VendorId: vendorID, PostId: postID}
}
