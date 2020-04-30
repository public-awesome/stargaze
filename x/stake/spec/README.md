# x/stake module specification

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

### DelegationQueue (FIFO)

* DelegationQueue: 0x01 | format(expire_time) | vendor_id | post_id | stake_id -> Delegation

`stake_id` is an auto-incrementing `uint32`.

```go
type Delegation struct {
    DelegatorAddr sdk.AccAddress
    ValidatorAddr sdk.ValAddress
    Amount        sdk.Dec
}
```

## Messages

### MsgDelegate

* Call `Delegate()` in the staking module
* Add a new `Delegation` to the end of `DelegationQueue`

## End-Block

* Check `expire_time` in `DelegationQueue` against the current block time
* Process all delegations that have an `expire_time` that exceeds the current block time
    - Undelegate
    - Distribute rewards

## Events

### EndBlocker

| Type              | Attribute Key         | Attribute Value           |
| ----------------- | --------------------- | ------------------------- |
| complete_stake    | amount                | {Amount}                  |
| complete_stake    | delegator             | {delegatorAddress}        |

## Parameters

* `VotingPeriod` = 3 days
