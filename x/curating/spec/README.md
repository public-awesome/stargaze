# Curating Module Specification

## Abstract

This module provides functionality to curate social media content and earn rewards in cryptocurrency via market-based mechanisms.

## Concepts

First, a post is created with a `post` transaction, taking a hash of the content, and a deposit. The deposit is returned after a curation window.

Posts are curated with `upvote` transactions and `moderate` transactions.

Upvotes are based on quadratic voting, and take number of votes as an argument. They also takes a deposit that is returned after the curation window. Deposits are meant to keep the system honest. For example, governance may choose to slash the deposit of upvotes that are engaging in collusion or a Sybil attack.

Moderate transactions are for signaling that a post contains spam or illegal content. Moderated posts immediately go to governance for adjudication. Curation is paused during this time. For posts deemed as inappropriate, curators get slashed and the moderator gets a reward. For posts not deemed as inappropriate, curation resumes while the moderator gets slashed.

Only data for posts currently being curated are stored on-chain. Everything else is deleted after the curation window completes and rewards are paid out.

## State

### Parameters

Each of these parameters can be voted on by governance, but will be assigned sane defaults for genesis.

```go
type Params struct {
	CurationWindow       	 time.Duration
	PostDeposit              sdk.Coin
	UpvoteDepost  	         sdk.Coin
	ModerationDeposit        sdk.Coin
	VoteAmount               sdk.Coin // amount for 1 vote
	MaxNumVotes              uint32 // upper-limit on QV voting
	MaxVendors               uint32 // upper-limit on vendor_ids
	RewardPoolAllocation 	 sdk.Dec // from inflation
}
```

### Base Types

```go
type Post struct {
	VendorID        uint32
	PostID          string
	Creator         sdk.AccAddress
	RewardAccount	sdk.AccAddress
	Body            string
	Deposit         sdk.Coin
	CurationEndTime time.Time
}
```

```go
type Upvote struct {
	VendorID        uint32
	PostID          string
	Curator         sdk.AccAddress
	VoteAmount      sdk.Coin
	Deposit         sdk.Coin
	CurationEndTime time.Time
}
```

### Stores

_Stores are KVStores in the multi-store. The key to find the store is the first
parameter in the list_.

We will use one KVStore `curating` to store two mappings:

- A mapping from `0x01|vendorID|postID` to `Post`.
- A mapping from `0x02|vendorID|postID|curator` to `Upvote`.

### Post Curation Queue

All queues objects are sorted by timestamp. The time used within any queue is
first rounded to the nearest nanosecond then sorted. The sortable time format
used is a slight modification of the RFC3339Nano and uses the the format string
`"2006-01-02T15:04:05.000000000"`. Notably this format:

- right pads all zeros
- drops the time zone info (uses UTC)

In all cases, the stored timestamp represents the maturation time of the queue
element.

For the purpose of tracking the end of post curation windows the post curation queue is kept.

- PostCurationQueue: `0x41 | format(CurationEndTime) -> []VPPair`

```go
type VPPair struct {
    VendorID uint32
    PostID   string
}
```

During each `EndBlock`, all the posts that have reached the end of their curation window are processed. The upvotes are tallied, and rewards are paid out to via a mix of quadratic voting and quadratic finance matching.

## Messages

### MsgPost

A social media post to be curated is registered on-chain with `MsgPost`. Note that the actual data of the post is not stored, just a hash of the data, and a vendor and post identifier.

Post curation begins as soon as the post is registered on-chain.

```go
type MsgPost struct {
	VendorID      uint32 // 1 = twitter, 2 = reddit, etc.
	PostID        string
	Creator       sdk.AccAddress
	RewardAccount sdk.AccAddress // optional
	Body          string // hex string, optional
	Deposit       sdk.Coin
}
```

Posts are persisted with the key `0x01|vendorID|postID`.

An optional reward account enables rewards to go to a different account than the creator. 

Validate the post, calculate the curation end time, and insert the vendor/post_id pair into the post curation queue.

### MsgUpvote

Every transaction on Stakebird includes a deposit which could be slashed by governance for bad behavior. This helps keep the network honest.

The amount of vote credits (`VoteAmount`) is quadratically associated to `VoteNum`, and is used for the upvote amount.

| `VoteNum` | `VoteAmount`  |
|-----------|---------------|
|   2       | 4             |
|   3       | 9             |
|   4       | 16            |


```go
type MsgUpvote struct {
	VendorID uint32
	PostID   string
	Curator  sdk.AccAddress
	VoteNum  uint32 // 1 = 1, 2 = 4, 3 = 9 (quadratic)
	Deposit  sdk.Coin
}
```

Upvotes are persisted with the key `0x02|vendorID|postID|curator`.

If the corresponding post doesn't exist yet for an upvote, then also create the post with the curator being the `RewardAccount` for the creator. This way, the curator earns creator rewards as well.

### MsgModerate

`MsgModerate` is to signify spam or illegal content. Moderation delegates to the Cosmos governance module, creating a new text proposal for each moderation message.

Text proposal:
```
title: "Content moderation proposal {vendor_id} - {post_id}"
description: hash of the post
type: "Text"
deposit: `MsgModerate.Deposit`
```

Governance will have to match the hash of the post with the actual post, and moderate the content.

[TODO] Figure out how to handle the result of moderation (i.e: re-start curation or delete/flag post). See https://github.com/public-awesome/stakebird/issues/51.

```go
type MsgModerate struct {
	VendorID  uint32
	PostID    string
	Moderator sdk.AccAddress
	Deposit   sdk.Coin
}
```

## End-Block

### Ending Post Curations

Iterate `PostCurationQueue` and process all upvotes that have completed their curation window with the following procedure:

- find all `VPPairs` that have ended curation
- find all associated `Upvote` objects by iterating the `0x02|vendorID|postID|curator` store
- calculate the `VotingPool` according to quadratic voting
- calculate the `MatchedPool` with quadratic finance matching
- distribute rewards to the creator of the post
- distribute rewards to each curator

## Events

### EndBlocker

| Type                  | Attribute Key         | Attribute Value           |
| --------------------- | --------------------- | ------------------------- |
| complete_curation     | vendor_id             | {vendorID}                |
| complete_curation     | post_id               | {postID}                  |
| complete_curation     | vote_count            | {totalVoteCount}          |
| complete_curation     | vote_amount           | {totalVoteAmount}         |
| complete_curation     | voting_pool           | {totalVotingPool}         |
| complete_curation     | matching_pool         | {totalMatchingPool}       |
| creator_reward        | creator               | {creatorAddress}          |
| creator_reward        | reward                | {rewardAmount}            |
| curator_reward        | curator               | {curatorAddress}          |
| curator_reward        | reward                | {rewardAmount}            |

### Handlers

#### MsgPost

| Type     | Attribute Key  | Attribute Value    |
| -------- | -------------- | ------------------ |
| post     | vendor_id      | {vendorID}         |
| post     | post_id        | {postID}           |
| post     | creator        | {creatorAddress}   |
| post     | reward_account | {rewardAddress}    |
| message  | module         | curating           |
| message  | action         | post               |
| message  | sender         | {creatorAddress}   |

#### MsgUpvote

| Type     | Attribute Key | Attribute Value    |
| -------- | ------------- | ------------------ |
| upvote   | vendor_id     | {vendorID}         |
| upvote   | post_id       | {postID}           |
| upvote   | curator       | {curatorAddress}   |
| message  | module        | curating           |
| message  | action        | upvote             |
| message  | sender        | {curatorAddress}   |

#### MsgModerate

| Type     | Attribute Key | Attribute Value    |
| -------- | ------------- | ------------------ |
| moderate | vendor_id     | {vendorID}         |
| moderate | post_id       | {postID}           |
| moderate | moderator     | {moderatorAddress} |
| message  | module        | curating           |
| message  | action        | moderate           |
| message  | sender        | {moderatorAddress} |
