# Cron

## Abstract

This module enables smart contracts on chain to receive callbacks at the end of each block. Useful for scheduling actions which need to happen every block.


## Concepts

### Priviledged contracts
An existing contract instantiated on chain which has been elevated to the status of priviledged contract via on-chain governance proposal. 
Privileged contracts receive some extra superpowers where they can receive callbacks from the [`EndBlocker`](https://docs.cosmos.network/main/building-modules/beginblock-endblock.html) into their [`sudo`](https://book.cosmwasm.com/basics/entry-points.html?highlight=sudo#entry-points) entrypoint. More details on how to use this in [here](#how-to-use-in-cw-contract)

## State

The module state is quite simple. It stores an array of bech32 address of all the Cosmwasm smart contracts that have been elevated via governance. Once the contract priviledge is demoted via governance, it is removed from state.

## End Blockers

In the ABCI endblock, the module fetches the list of all priviledged contracts and loops through them and calls the `sudo` entry point with the `EndBlocker` msg.

## Events

The module emits the following events:

| Source type | Source name                |
| ----------- | -------------------------- |
| Keeper      | `SetContractPriviledge`    |
| Keeper      | `UnsetContractPriviledge`  |

## Client

### CLI - Query

#### **list-privileged**

```
starsd q cron list-privileged
```

List all privileged contract addresses in bech32 format

### CLI - Tx

#### **promote-to-privilege-contract**

```
starsd tx cron promote-to-privilege-contract {contractAddr} --title {proposalTitle} --deposit {depositAmount}
starsd tx cron promote-to-privilege-contract stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420 --title "Promote Contract Proposal" --deposit 1000ustars
```

Creates a governance proposal which on passing will add the given contract address to the priviledged contract list

#### **demote-from-privilege-contract**

```
starsd tx cron demote-from-privilege-contract {contractAddr} --title {proposalTitle} --deposit {depositAmount}
starsd tx cron demote-from-privilege-contract  stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420 --title "Demote Contract Proposal" --deposit 1000ustars
```

Creates a governance proposal which on passing will remove the given contract address from the priviledged contract list

## How to use in CW contract

To use the `EndBlocker` callback from this module, you need to add the following msg type to your contract msgs. 

`msg.rs`
```rust
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum SudoMsg {
    EndBlock {},
}
```
`contract.rs`
```rust
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn sudo(deps: DepsMut, env: Env, msg: SudoMsg) -> Result<Response, ContractError> {
    match msg {
        SudoMsg::EndBlock {} => !unimplemented(),
    }
}
```
Ensure you have a specific sudo entrypoint in your contract and add your code to be called by the `SudoMsg::EndBlock`.
This endpoint is hit at the end of every block by the x/cron module. The `deps` and `env` values are set as expected. It is not possible to receive any input in the `EndBlock` call and as such the contract must maintain any relevant state it needs to use. 