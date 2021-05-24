package keeper

import "github.com/public-awesome/stargaze/x/curating/types"

// MustMarshalPost attempts to encode a Post object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalPost(post types.Post) []byte {
	return k.cdc.MustMarshal(&post)
}

// MustUnmarshalPost attempts to decode a Post object and return it. It panics on error.
func (k Keeper) MustUnmarshalPost(data []byte, post *types.Post) {
	k.cdc.MustUnmarshal(data, post)
}

// MustMarshalUpvote attempts to encode an Upvote object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalUpvote(upvote types.Upvote) []byte {
	return k.cdc.MustMarshal(&upvote)
}
