# Stake Module Specification

## Abstract

This module allows staking on posts to earn yield.

## Concepts

This module wraps the Cosmos staking module to provide delegation functionality. A post stake action wraps around a delegation, and a post unstake action wraps around an undelegation.
## State

### Parameters

n/a

### Base Types

```go
type Deleation struct {
	Validator      sdk.ValAddress
	Amount         sdk.Coin
}
```

### Stores

- A mapping from `0x01|vendorID|postID|delegator` to `Delegation`.

## Messages

### MsgStake

```go
type MsgStake struct {
	VendorID      uint32 // 1 = twitter, 2 = reddit, etc.
	PostID        string
	Delegator     sdk.AccAddress
	Validator     sdk.ValAddress
	Amount        sdk.Coin
}

type MsgUnstake struct {
	VendorID      uint32 // 1 = twitter, 2 = reddit, etc.
	PostID        string
	Delegator     sdk.AccAddress
	Amount        sdk.Coin
}
```

Delegations are persisted with the key `0x01|vendorID|postID|delegator`.


## Events

### Handlers

#### MsgStake

| Type     | Attribute Key  | Attribute Value    |
| -------- | -------------- | ------------------ |
| stake    | vendor_id      | {vendorID}         |
| stake    | post_id        | {postID}           |
| stake    | delegator      | {delegatorAddress} |
| stake    | validator      | {validatorAddress} |
| stake    | amount         | {amount}           |
| message  | module         | stake              |
| message  | action         | stake              |
| message  | sender         | {delegatorAddress} |

#### MsgUnstake

| Type     | Attribute Key | Attribute Value    |
| -------- | ------------- | ------------------ |
| stake    | vendor_id      | {vendorID}         |
| stake    | post_id        | {postID}           |
| stake    | delegator      | {delegatorAddress} |
| stake    | amount         | {amount}           |
| message  | module         | stake              |
| message  | action         | stake              |
| message  | sender         | {delegatorAddress} |
