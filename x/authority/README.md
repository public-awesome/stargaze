# Authority

## Abstract

This module enables governance functionality to be delegated to addresses and not have to rely on the sdk x/gov module.
E.g A wallet or multisig or a smart contract can be given the authority to execute a governance operation without going through on-chain governance proposal.

## Concepts

With SDK 47, modules take in a string address of their "authority". This authority has the power to execute "authority-gated" operations. E.g Updating the module parameters. By default all the modules are configured with x/gov module as their authority.

Our custom x/authority module allows replacement of the x/gov module from the authority role. The module also allows configuration of authority not just at a per module basis but at a per Msg basis.

e.g x/wasmd has many authority gated operations. Using x/authority module instead of x/gov allows us to have different admins for each of the operations.

`UpdateParams` authorization could be given to the dev team multisig

`SudoContract` authorization could be given to a smart contract DAO

> **Note**
>
> The way the authority mechanism works, if a module has x/authority as its authority, then all its gov msgs should be explicitly authorized at the x/authority module params. It a particular msg does not have authorization, it cannot be executed at all.


## Params

- Authorizations

  This is a list of all the msg type URLs and the addresses that can execute them

## Client

### CLI - Query

#### **params**

```sh
starsd q authority params --output json | jq
```

```json
starsd q authority params 
{
  "params": {
    "authorizations": [
      {
        "msgTypeUrl": "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
        "addresses": [
          "stars1x394ype3x8nt9wz0j78m8c8kcezpslrcjmmrc0"
        ]
      },
      {
        "msgTypeUrl": "/cosmwasm.wasm.v1.MsgUpdateParams",
        "addresses": [
          "stars1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cyt4fdd"
        ]
      }
    ]
  }
}
```

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

#### Update Params

```jsonc
{
  "messages": [
    {
      "@type": "/publicawesome.stargaze.authority.v1.MsgUpdateParams",
      "authority": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "params": {
        "authorizations": [
          {
            "msgTypeUrl": "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
            "addresses": [
              "stars1x394ype3x8nt9wz0j78m8c8kcezpslrcjmmrc0"
            ]
          }
        ]
      }
    }
  ],
  "metadata": "metadata",
  "deposit": "1000stake",
  "title": "Update Authority Params",
  "summary": ""
}
```

### CLI - Tx

#### **submit-proposal**

```
starsd tx authority submit-proposal [path/to/proposal.json]
```

The format of proposal.json is same as the one used for x/gov module to keep the parity in usage.

Ensure you are setting the "authority" or "sender" value to x/authority address and not x/gov address