# Client

This section describes interactions with the module by the user

## CLI

### Query

The `query` commands allows a user to query the module state.

Use the `-h` / `--help` flag to get a help description of a command.

```bash
starsd q globalfee -h
```

#### params

Get the module parameters.

Usage:

```bash
starsd q globalfee params [flags]
```

Example output:

```yaml
privileged_address:
  - stars1hjlswfagegg7ymu9p2vn84wqzcaf83w2qj89tq
  - stars1k7nky435jc55th95uhgk94zvchxdezqg3vp07e
  - stars185yv4ka7ymn3f6meq7k3g42nxrgfeykjkua59c
```

#### auth_by_code_id

Gets the authorizations for the given code id

Usage:

```bash
starsd q globalfee auth_by_code_id [code-id] [flags]
```

Example output:

```yaml
code_id: 2
methods:
  - mint
  - unlist
```

#### auth_by_contract_address

Gets the authorizations for the given contract address

Usage:

```bash
starsd q globalfee auth_by_contract_address [contract-address] [flags]
```

Example output:

```yaml
contract_address: stars1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfsvgv0pl
methods:
- *
```

#### auth_all

Gets all the authorizations which have been set up

Usage:

```bash
starsd q globalfee auth_all [flags]
```

Example output:

```yaml
code_authorizations:
- code_id: 2
  methods:
  - mint
  - unlist
contract_authorizations:
- contract_address: stars1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfsvgv0pl
  methods:
  - *
```

### Gov

```starsd tx gov submit-proposal proposal.json --from {user}```

You will need the x/gov module address to set as authority for the proposal. You can fetch it by running: `starsd q auth module-account gov`

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
      "@type": "/publicawesome.stargaze.globalfee.v1.MsgUpdateParams",
      "authority": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "params": { // note: the entire params field needs to be filled
        "privileged_addresses": [
            "stars1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cyt4fdd" 
        ],
        "minimum_gas_prices": [
          "1ustar"
        ]
      }
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Update module params",
    "summary": "This will update the module admins as well as the minimum gas price"
}
```

#### Set Code Authorization

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.globalfee.v1.MsgSetCodeAuthorization",
      "sender": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "code_authorization": {
        "code_id": 1,
        "methods": [
          "mint"
        ]
      }
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Set Code Authorization",
    "summary": "Set Code Authorization for all contracts of given code id"
}
```

#### Remove Code Authorization

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.globalfee.v1.MsgRemoveCodeAuthorization",
      "sender": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "code_id": 1
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Remove Code Authorization",
    "summary": "Remove Code Authorization for all contracts of given code id"
}
```

#### Set Contract Authorization

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.globalfee.v1.MsgSetContractAuthorization",
      "sender": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "contract_authorization": {
        "contract_address": "stars14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9srsl6sm",
        "methods": [
          "*"
        ]
      }
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Set Contract Authorization",
    "summary": "Set contract Authorization for contract of given address"
}
```

#### Remove Contract Authorization

```jsonc
{
    "messages": [
     {
      "@type": "/publicawesome.stargaze.globalfee.v1.MsgRemoveContractAuthorization",
      "sender": "stars10d07y265gmmuvt4z0w9aw880jnsr700jw7ycaz", // x/gov address
      "contract_address": "stars14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9srsl6sm"
     }
    ],
    "metadata": "metadata",
    "deposit": "1000stake",
    "title": "Remove Contract Authorization",
    "summary": "Remove contract Authorization for contract of given address"
}
```


### Transactions

The `tx` commands allows a user to interact with the module.

Use the `-h` / `--help` flag to get a help description of a command.

```bash
starsd tx globalfee -h
```

#### set-code-authorization

Creates or updates the gas free operation authorization for the given code ID and for the provided methods.
The methods should be comma separated values.

> **Note**
> Only whitelisted address can perform this operation.

Usage:

```bash
starsd tx globalfee set-code-authorization [code-id] [methods] [flags]
```

Example:

```bash
starsd tx globalfee set-code-authorization 3 "mint,unlist"  \
  --from myAccountKey \
  --fees 1500ustars
```

#### remove-code-authorization

Removes the gas free operation authorization for the given code ID.

> **Note**
> Only whitelisted address can perform this operation.

Usage:

```bash
starsd tx globalfee remove-code-authorization [code-id]  [flags]
```

Example:

```bash
starsd tx globalfee remove-code-authorization 3  \
  --from myAccountKey \
  --fees 1500ustars
```

#### set-contract-authorization

Creates or updates the gas free operation authorization for the given contract address and for the provided methods.
The methods should be comma separated values.

> **Note**
> Only whitelisted address can perform this operation.

Usage:

```bash
starsd tx globalfee set-contract-authorization [contract-address] [methods] [flags]
```

Example:

```bash
starsd tx globalfee set-contract-authorization stars1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfsvgv0pl "*"  \
  --from myAccountKey \
  --fees 1500ustars
```

#### remove-contract-authorization

Removes the gas free operation authorization for the given contract address.

> **Note**
> Only whitelisted address can perform this operation.

Usage:

```bash
starsd tx globalfee remove-contract-authorization [contract-address] [flags]
```

Example:

```bash
starsd tx globalfee remove-contract-authorization stars1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfsvgv0pl \
  --from myAccountKey \
  --fees 1500ustars
```
