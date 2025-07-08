# Module Multi-Chain

Ce module gère l’interopérabilité entre LibraDemosChain et d’autres blockchains (Cosmos, Ethereum, Solana, etc.).

## Exemples d’utilisation CLI

```sh
# Créer un client IBC
ldcd tx ibc client create --client-id "07-tendermint-0" --client-type "07-tendermint" --from validator

# Établir une connexion
ldcd tx ibc connection open-init --connection-id "connection-0" --client-id "07-tendermint-0" --counterparty-client-id "07-tendermint-1" --from validator

# Créer un canal de transfert
ldcd tx ibc channel open-init --port-id "transfer" --channel-id "channel-0" --connection-id "connection-0" --from validator
```

## Fonctionnalités principales
- Connexion à plusieurs réseaux blockchain
- Transferts inter-chaînes
- Intégration IBC+, bridges, relayers

## Installation & configuration
- Dépendances à installer (voir `package.json`)
- Configuration : `frontend/src/modules/multichain/config.js`

## Développement
- Code source : `frontend/src/modules/multichain/`
- Tests : `tests/multichain/`

## Ressources
- [docs/ARCHITECTURE.md](../../../docs/ARCHITECTURE.md)
- [docs/STRUCTURE_IDEALE.md](../../../docs/STRUCTURE_IDEALE.md)
