# Client

THis section describes interactions with the module by the user

## CLI

### Query

The `query` commands allows a user to query the module state.

Use the `-h` / `--help` flag to get a help description of a command.

```bash
starsd q globalfee -h
```

> You can add the `-o json` for the JSON output format.

#### params

Get the module parameters.

Usage:

```bash
starsd q globalfee params [flags]
```

Example output:

```yaml
```

### Transations

The `tx` commands allows a user to interact with the module.

Use the `-h` / `--help` flag to get a help description of a command.

```bash
starsd tx globalfee -h
```

#### blah

blah

Usage:

```bash
starsd tx globalfee blah [flags]
```

Command specific flags:

* `--` - blah;

Example:

```bash
starsd tx globalfee blah  \
  --from myAccountKey \
  --fees 1500ustars
```