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

#### set-code-authorization-proposal

Creates a gov proposal to create or update the zero gas fee operation authorization for the given code ID and for the provided methods. The methods should be comma separated values.

Any stargaze address can perform this operation.

Usage:

```bash
starsd tx globalfee set-code-authorization-proposal [code-id] [methods]  [flags]
```

Example:

```bash
starsd tx globalfee set-code-authorization-proposal 3 "mint,unlist"  \
  --title "Adding new code authorization" \
  --deposit 1000ustars \
  --from myAccountKey \
  --fees 1500ustars
```

#### remove-code-authorization-proposal

Creates a gov proposal to remove the zero gas fee operation authorization for the given code ID.

Any stargaze address can perform this operation.

Usage:

```bash
starsd tx globalfee remove-code-authorization-proposal [code-id] [flags]
```

Example:

```bash
starsd tx globalfee remove-code-authorization-proposal 3 \
  --title "Removing the code authorization" \
  --deposit 1000ustars \
  --from myAccountKey \
  --fees 1500ustars
```

#### set-contract-authorization-proposal

Creates a gov proposal to create or update the gasless operation authorization for the given contract adress and for the provided methods. The methods should be comma separated values.

Any stargaze address can perform this operation.

Usage:

```bash
starsd tx globalfee set-contract-authorization-proposal [contract-address] [methods] [flags]
```

Example:

```bash
starsd tx globalfee set-contract-authorization-proposal stars1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfsvgv0pl "*"  \
  --title "Adding new contract authorization" \
  --deposit 1000ustars \
  --from myAccountKey \
  --fees 1500ustars
```

#### remove-contract-authorization-proposal

Creates a gov proposal to remove the zero gas fee operation authorization for the given contract adress.

Any stargaze address can perform this operation.

Usage:

```bash
starsd tx globalfee remove-contract-authorization-proposal [contract-address] [flags]
```

Example:

```bash
starsd tx globalfee remove-contract-authorization-proposal stars1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfsvgv0pl  \
  --title "Removing the contract authorization" \
  --deposit 1000ustars \
  --from myAccountKey \
  --fees 1500ustars
```
