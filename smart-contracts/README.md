# Smart Contracts CosmWasm

Ce dossier contient les contrats intelligents CosmWasm pour LibraDemosChain (DAO, privacy, KYC, etc.).

## Déploiement d’un smart contract CosmWasm

Voir [../docs/howto_deploy_contract.md](../docs/howto_deploy_contract.md)

## Exemples CLI

```sh
# Stocker un contrat WASM
ldcd tx wasm store contract.wasm --from validator

# Instancier le contrat
ldcd tx wasm instantiate 1 '{}' --label "governance-contract" --from validator

# Exécuter le contrat
ldcd tx wasm execute ldc1... '{"method":"vote","params":{}}' --from citizen1
```
