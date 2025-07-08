# Scripts LibraDemosChain

Ce dossier centralise tous les scripts utiles pour le build, le test, le déploiement, la gestion des modules avancés (privacy, IA, DAO, multi-chain).

## Scripts principaux
- `deploy_contracts.sh` : déploiement des smart contracts CosmWasm
- `init_dao.sh` : initialisation d’une DAO
- `setup_privacy.sh` : configuration privacy stack (ZK, Anoma, Namada)
- `testnet_setup.sh` : déploiement d’un testnet multi-chaînes
- `build_all.sh` : build complet backend + frontend + smart contracts

## Utilisation
```sh
chmod +x scripts/*.sh
./scripts/build_all.sh
```

## Ressources
- Voir chaque script pour les instructions détaillées
- [docs/INSTALLATION.md](../docs/INSTALLATION.md)
- [docs/STRUCTURE_IDEALE.md](../docs/STRUCTURE_IDEALE.md)
