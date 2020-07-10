package types

// Query endpoints supported by the curating querier
const (
	QueryParams  = "parameters"
	QueryPost    = "post"
	QueryPosts   = "posts"
	QueryUpvotes = "upvotes"
)

/*
How to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

*/
