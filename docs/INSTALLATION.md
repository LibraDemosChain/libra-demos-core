# Guide d’installation LibraDemosChain

## Prérequis

- Go (>=1.20)
- Node.js (>=18)
- npm
- Docker (optionnel)
- Rust (pour CosmWasm et ZK)

## Installation backend (Cosmos SDK, multi-blockchain, privacy)

```sh
cd blockchain
# Installer les dépendances Go
go mod tidy
# (Optionnel) Installer les modules privacy/ZK/Anoma/Namada
# Voir blockchain/x/privacy/README.md
# Lancer le nœud local (mainnet/testnet)
go run main.go
```

- Exemple de configuration : voir `blockchain/data/config/genesis.json` et `blockchain/data/config/app.toml` pour personnaliser le réseau, les modules privacy, et la gouvernance.
- Pour l’interopérabilité multi-chaînes, voir `blockchain/x/multichain/README.md`.

## Installation smart contracts (CosmWasm, privacy, DAO, KYC)

```sh
cd smart-contracts
# Compiler les contrats CosmWasm
cargo build --release
# Déployer sur la blockchain (voir scripts/deploy_contracts.sh)
```

- Exemples de contrats privacy/DAO/KYC dans `smart-contracts/`.
- Scripts de déploiement avancés : `scripts/deploy_contracts.sh`, `scripts/init_dao.sh`.

## Installation frontend (Next.js, IA NEOMA, multi-chain)

```sh
cd frontend
npm install
# (Optionnel) Configurer l’IA NEOMA et la privacy (voir frontend/src/modules/ai/README.md)
npm run dev
```

- Pour l’intégration multi-blockchain et privacy, voir `frontend/src/modules/multichain/README.md` et `frontend/src/modules/privacy/README.md`.
- Pour l’activation de l’IA NEOMA, voir `frontend/src/modules/ai/README.md`.

## Déploiement testnet/mainnet

- Voir [docs/DEPLOYMENT.md](DEPLOYMENT.md) pour la configuration multi-blockchain, privacy, IA, scripts de déploiement, et exemples de fichiers genesis/app.toml.
- Exemples de configuration disponibles dans `blockchain/data/config/` et `scripts/`.
- Pour déployer un réseau complet :

```sh
# Lancer tous les services (backend, frontend, smart contracts)
docker-compose up --build
```

## Ressources complémentaires

- [docs/STRUCTURE_IDEALE.md](STRUCTURE_IDEALE.md)
- [docs/SECURITY.md](SECURITY.md)
- [docs/GOVERNANCE.md](GOVERNANCE.md)
- [blockchain/x/privacy/README.md](../blockchain/x/privacy/README.md)
- [frontend/src/modules/ai/README.md](../frontend/src/modules/ai/README.md)
- [blockchain/x/multichain/README.md](../blockchain/x/multichain/README.md)
- [frontend/src/modules/multichain/README.md](../frontend/src/modules/multichain/README.md)
- [smart-contracts/README.md](../smart-contracts/README.md)
- [scripts/README.md](../scripts/README.md)

## Rapports & analyses avancées

- [docs/ARCHITECTURE.md](ARCHITECTURE.md)
- [docs/WHITEPAPER.md](WHITEPAPER.md)
- [docs/ROADMAP.md](../ROADMAP.md)
- [docs/SECURITY_AUDIT_CHECKLIST.md](../SECURITY_AUDIT_CHECKLIST.md)
- [docs/STRUCTURE_IDEALE.md](STRUCTURE_IDEALE.md)

> Pour toute question ou contribution, voir la documentation de chaque module ou ouvrir une issue sur GitHub.
