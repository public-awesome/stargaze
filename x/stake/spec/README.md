# Stake module specification

## Abstract

The stake module performs time-constrained validator delegations.

## Concepts

This module wraps the Cosmos `x/staking` module to perform time-bound delegations. It is designed for end users to perform staking and secure the network, without having to manually undelegate. 

## Dependencies

* x/bank
* x/staking
    - `UnbondingPeriod` = 3 days
    - `MaxEntries` = 7+

## State

### Post

* Posts: 0x01 | vendor_id | post_id -> Post

```go
type Post struct {
	ID              uint64 
	VendorID        uint32 
	Body            string 
	VotingPeriod    time.Duration 
    VotingStartTime time.Time
}
```

A post can optionally include a delegation.

### DelegationQueue (FIFO)

* DelegationQueue: 0x02 | format(expire_time) | vendor_id | post_id | stake_id -> Delegation

`stake_id` is an auto-incrementing `uint32`.

```go
type Delegation struct {
    DelegatorAddr sdk.AccAddress
    ValidatorAddr sdk.ValAddress
    Amount        sdk.Dec
}
```

## Messages

### MsgPost

* Persist the post
* Add `Delegation` if the post message includes one

```go
type MsgPost struct {
	ID              uint64 
	VendorID        uint32 
	Body            string 
	VotingPeriod    time.Duration 
    VotingStartTime time.Time
    Delegation      *staking.Delegation
}
```

### MsgDelegate

* If post doesn't exist, create the `Post` and start the voting period.
* Call `Delegate()` in the staking module
* Add a new `Delegation` to the end of `DelegationQueue`

```go
type MsgDelegate struct {
	VendorID      uint32
	PostID        uint64
	DelegatorAddr sdk.AccAddress
	ValidatorAddr sdk.ValAddress
	Amount        sdk.Coin
}
```

## End-Block

* Check `expire_time` in `DelegationQueue` against the current block time
* Process all delegations that have an `expire_time` that exceeds the current block time
    - Undelegate
    - Distribute rewards

## Events

### EndBlocker

| Type                      | Attribute Key         | Attribute Value           |
| ------------------------- | --------------------- | ------------------------- |
| voting_period_start       | post                  | {Post}                    |
| voting_period_end         | post                  | {Post}                    |

## Parameters

