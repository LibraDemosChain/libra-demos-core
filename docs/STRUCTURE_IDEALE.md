# Structure idéale du dépôt libra-demos-core

## 1. Racine du projet
- README.md (vision, architecture, innovations, liens clés)
- ROADMAP.md (phases, jalons, objectifs)
- LICENSE, CODE_OF_CONDUCT.md, CONTRIBUTING.md, SECURITY.md, GOVERNANCE.md
- docs/ (documentation technique, guides, schémas)

## 2. Backend blockchain (Cosmos SDK)
- blockchain/
  - app/ (application principale Cosmos SDK)
  - cmd/ (commandes CLI)
  - proto/ (protos gRPC)
  - x/ (modules avancés : democracy, privacy, ZK-KYC, IA)
  - README.md

## 3. Smart contracts (CosmWasm)
- smart-contracts/
  - privacy/
  - dao/
  - kyc/
  - README.md

## 4. Frontend (Next.js, CosmJS)
- frontend/
  - public/
  - src/
    - modules/
      - ai/ (NEOMA, modération, analyse)
      - privacy/ (gestion identité, ZK-KYC)
      - governance/ (vote, propositions, DAO)
      - multichain/ (interopérabilité, intégration autres blockchains)
  - README.md

## 5. Tests
- tests/
  - integration/
  - unit/

## 6. Ressources & formation
- NecessaryKnowledge/

## 7. Déploiement & scripts
- scripts/
- Dockerfile, docker-compose.yml

---

## Modules à intégrer/prioriser
- **blockchain/x/democracy/** : gouvernance décentralisée
- **blockchain/x/privacy/** : privacy stack (ZK, Anoma, Namada)
- **blockchain/x/ai/** : NEOMA (anti-désinformation, analyse)
- **blockchain/x/kyc/** : ZK-KYC
- **frontend/src/modules/** : IA, privacy, governance, multichain
- **smart-contracts/** : CosmWasm privacy, DAO, KYC

## Conseils de centralisation/nettoyage
- Supprimer les duplications entre prototypes
- Garder une structure modulaire claire (un dossier par domaine)
- Documenter chaque module (README)
- Séparer bien backend, frontend, smart contracts
- Utiliser des scripts de build/test/deploy communs

---

> Cette structure garantit une base professionnelle, évolutive, sécurisée et facile à contribuer.
