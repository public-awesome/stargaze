# Mint

This module mints new tokens on each block. It implements a reduction factor to redunce emissions over time.

## Governance Parameters

Here's a table for the parameters of the module:

| Key                       | Type         | Description                            | Default                   |
| ------------------------- | ------------ | -------------------------------------- | ------------------------- |
| `MintDenom`               | string       | Denomination of the minted tokens      | `ustars`                  |
| `StartTime`               | string       | When to start emissions                |                           |
| `InitialAnnualProvisions` | string (dec) | The starting annual genesis allocation | `"1_000_000_000_000_000"` |
| `ReductionFactor`         | string (dec) | Ratio to reduce emissions by each year | `0.67`                    |
| `BlocksPerYear`           | uint64       | Number of blocks per year              | `6311520`                 |
