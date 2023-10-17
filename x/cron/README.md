# Cron

## Abstract

This module enables smart contracts on chain to receive callbacks at the beginning and end of each block. Useful for scheduling actions which need to happen every block.

## Concepts

### Priviledged contracts

An existing contract instantiated on chain which has been elevated to the status of priviledged contract via on-chain governance proposal or via whitelisted admin addresses.
Privileged contracts receive some extra superpowers where they can receive callbacks from the [`BeginBlocker & EndBlocker`](https://docs.cosmos.network/main/building-modules/beginblock-endblock.html) into their [`sudo`](https://book.cosmwasm.com/basics/entry-points.html?highlight=sudo#entry-points) entrypoint. More details on how to use this in [here](#how-to-use-in-cw-contract)

## State

The module state is quite simple. It stores the bech32 address of all the Cosmwasm smart contracts that have been elevated via governance or admins. Once the contract priviledge is demoted via governance, it is removed from state.

Storage keys:

- PriviledgedContract: `0x01 | contractAddress -> []byte{1}`

## Params

- AdminAddress: 
  List of `sdk.AccAddress` for accounts which are whitelisted to promote and demote contracts     

## Begin Blockers

In the ABCI beginblock, the module fetches the list of all priviledged contracts and loops through them and calls the `sudo` entry point with the `BeginBlocker` msg. None of the `abci.RequestBeginBlock` values are passed in.

## End Blockers

In the ABCI endblock, the module fetches the list of all priviledged contracts and loops through them and calls the `sudo` entry point with the `EndBlocker` msg.

## Events

The module emits the following events:

| Source type | Source name               |
| ----------- | ------------------------- |
| Keeper      | `SetContractPriviledge`   |
| Keeper      | `UnsetContractPriviledge` |

## Client

### CLI - Query

#### **list-privileged**

```
starsd q cron list-privileged
```

List all privileged contract addresses in bech32 format

#### **params**

```
starsd q cron params
```

List the module params

### CLI - Gov

```
starsd tx gov submit-proposal proposal.json --from {user}
```



You will need the x/gov module address to set as authority for the proposal. You can fetch it by running:

```starsd q auth module-account gov```

This will get you the following response
```jsonc
account:
  '@type': /cosmos.auth.v1beta1.ModuleAccount
  base_account:
    account_number: "7"
    address: stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz // This is the address you need to use for authority value
    pub_key: null
    sequence: "0"
  name: gov
  permissions:
  - burner
```
The expected format of the proposal.json is below. 

#### Promote Contract

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContract",
      "authority": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "contract": "stars14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9srsl6sm" // the contract to promote
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Promote contract",
    "summary": "Contract will get begin and end blocker callbacks"
}
```

#### Demote Contract

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContract",
      "authority": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "contract": "stars14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9srsl6sm" // the contract to demote
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Demote contract",
    "summary": "Contract will lose begin and end blocker callbacks"
}
```

#### Update Params

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.cron.v1.MsgUpdateParams",
      "authority": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "params": { // note: the entire params field needs to be filled
        "admin_address": [
            "stars1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cyt4fdd" 
        ]
      }
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Update module params",
    "summary": "This will add or remove the module admins"
}
```

### CLI - Tx

> **Note**
> Only whitelisted admin addresses can execute the following txs

#### **promote-to-privilege-contract** || **promote-contract**

```
starsd tx cron promote-contract {contractAddr} --from {adminUser}
starsd tx cron promote-contract stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420 --from admin
```

Immediately adds the given contract address to the priviledged contract list

#### **demote-from-privilege-contract** || **demote-contract**

```
starsd tx cron demote-contract {contractAddr} --from {adminUser}
starsd tx cron demote-contract stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420 --from admin
```

Immediately removes the given contract address to the priviledged contract list

## How to use in CW contract

To use the `BeginBlocker` & `EndBlocker` callback from this module, you need to add the following msg type to your contract msgs.

```rust
// msg.rs

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum SudoMsg {
    BeginBlock {},
    EndBlock {},
}
```

```rust
// contract.rs

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn sudo(deps: DepsMut, env: Env, msg: SudoMsg) -> Result<Response, ContractError> {
    match msg {
        SudoMsg::BeginBlock {} => !unimplemented(),
        SudoMsg::EndBlock {} => !unimplemented(),
    }
}
```

Ensure you have a sudo entrypoint in your contract and add your code to be called by the `SudoMsg::BeginBlock` & `SudoMsg::EndBlock`. Both endpoints are hit for every privileged contract in the block lifecycle, so **ensure both are implemented**.

This endpoint is hit at the beginning and end of every block by the x/cron module. The `deps` and `env` values are set as expected. It is not possible to receive any input in the `BeginBlock` & `EndBlock` call and as such the contract must maintain any relevant state it needs to use.

## Attribution

Thanks to Confio and the Tgrade open source project for providing a base for x/cron where most of it's bits were extracted and adapted to only do begin_blocker & end_blocker callbacks.

Found more about Tgrade's twasm module:

[x/twasm](https://github.com/confio/tgrade/tree/main/x/twasm)
