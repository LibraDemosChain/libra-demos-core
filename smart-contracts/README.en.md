# CosmWasm Smart Contracts

[🇫🇷 Version française](./README.md)

This folder contains the CosmWasm smart contracts for LibraDemosChain (DAO, privacy, KYC, etc.).

## Deploying a CosmWasm Smart Contract

See [../docs/howto_deploy_contract.md](../docs/howto_deploy_contract.md)

## CLI Examples

```sh
# Store a WASM contract
ldcd tx wasm store contract.wasm --from validator

# Instantiate the contract
ldcd tx wasm instantiate 1 '{}' --label "governance-contract" --from validator

# Execute the contract
ldcd tx wasm execute ldc1... '{"method":"vote","params":{}}' --from citizen1
```
