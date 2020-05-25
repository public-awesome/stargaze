package types

// Query endpoints supported by the bondcurve querier
const (
	// TODO: Describe query parameters, update Params with your query
	QueryParams = "params"
)

/*
Below you will be able how to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

*/
