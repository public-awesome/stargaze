# Parameters

This section describes the module parameters

| Key                  | Type             | Default value |  Description                                                         |
| -------------------- | ---------------- | ------------- | -------------------------------------------------------------------- |
| PrivilegedAddresses  | []string         |  []           | List of `sdk.AccAddress` for accounts which can modify authorizations  |  
| MinimumGasPrices | []sdk.DecCoin | [] | The minimum gas price(s) for all TX on the chain (except the authorized contract executions). The list must be sorted by denoms asc. No duplicate denoms or zero amount values allowed. |