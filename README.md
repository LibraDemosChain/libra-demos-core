[🇫🇷 Lire en français](LISEZ_MOI.md)

# libra-demos-core
LibraDemosChain core project – decentralized democratic governance platform.

---

## 🚀 Vision & Objectifs

LibraDemosChain vise à révolutionner la démocratie mondiale en créant un système de gouvernance décentralisé, transparent et accessible à tous. Inspiré par Bitcoin, le projet va plus loin : chaque citoyen peut participer directement, sans intermédiaire, grâce à la blockchain, l’IA et la confidentialité avancée.

- **Multi-blockchain** : Architecture modulaire, intégration de blockchains orientées privacy (Anoma, Namada, ZK, etc.) et haute performance (Solana, Polygon, Ethereum…)
- **IA pour la gouvernance (NEOMA)** : Anti-désinformation, analyse de propositions, modération intelligente
- **ZK-KYC & Privacy Stack** : Preuves à divulgation nulle de connaissance, gestion de l’identité et des votes anonymes
- **Déploiement par phases** : National (France) → Continental (Europe) → Mondial

---

## 🛠 Guides pratiques

- [Déploiement d’un smart contract CosmWasm](docs/howto_deploy_contract.md)
- [Initialisation du nœud & configuration genesis](docs/INSTALLATION.md)
- [Connexion IBC & interopérabilité](blockchain/x/multichain/README.md)
- [Utilisation des modules avancés (privacy, IA, DAO, KYC)](docs/STRUCTURE_IDEALE.md)

---

## 🔒 Sécurité & bonnes pratiques

- Audits réguliers (CertiK, Quantstamp, tests de pénétration, vérification formelle ZK)
- Gestion des clés (HSM, keyring-backend, sauvegarde mnémonique)
- Linting, formatage, couverture de tests (voir SECURITY.md)

---

## 💰 Tokenomics Dual-Token

- **LDC** : Token de gouvernance (1 citoyen = 1 vote, burn rate 0.5% par vote)
- **NEO** : Récompenses pour l’engagement citoyen (fact-checking, contributions, anti-désinformation)

---

## 📈 Monitoring, observabilité, testnets

- Monitoring Prometheus, APIs REST/gRPC, logs détaillés
- Guide testnet : [docs/INSTALLATION.md](docs/INSTALLATION.md), [blockchain/x/multichain/README.md](blockchain/x/multichain/README.md)

---

## ❓ FAQ & pédagogie

- [NecessaryKnowledge](NecessaryKnowledge/README.md) : formation, vulgarisation, compliance, sécurité, etc.

---

## 🗺️ Roadmap & architecture

- [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)
- [docs/ROADMAP.md](docs/ROADMAP.md)
- [docs/WHITEPAPER.md](docs/WHITEPAPER.md)

---

## 🧩 Scripts CLI utiles

- Initialisation de la chaîne, comptes, genesis, gentxs : voir scripts/
- Déploiement CosmWasm, analyses NEOMA, ZK-KYC, vote, etc.

---

## 📚 Documentation complète

- [docs/INSTALLATION.md](docs/INSTALLATION.md)
- [docs/SECURITY.md](docs/SECURITY.md)
- [docs/GOVERNANCE.md](docs/GOVERNANCE.md)
- [docs/STRUCTURE_IDEALE.md](docs/STRUCTURE_IDEALE.md)
- [blockchain/x/privacy/README.md](blockchain/x/privacy/README.md)
- [frontend/src/modules/ai/README.md](frontend/src/modules/ai/README.md)
- [blockchain/x/multichain/README.md](blockchain/x/multichain/README.md)
- [frontend/src/modules/multichain/README.md](frontend/src/modules/multichain/README.md)
- [smart-contracts/README.md](smart-contracts/README.md)
- [scripts/README.md](scripts/README.md)

---

## 📚 Base de connaissances NecessaryKnowledge

Retrouvez l’ensemble des guides, exemples, ressources et FAQ pour tous les domaines du projet :
- [Vue d’ensemble et sommaire](../NecessaryKnowledge/README.md)
- [Blockchain & Cryptography](../NecessaryKnowledge/01_Blockchain_Cryptography/)
- [AI & Data Analysis](../NecessaryKnowledge/02_AI_DataAnalysis/)
- [UX/UI Development](../NecessaryKnowledge/03_UX_UI_Development/)
- [Legal & Compliance](../NecessaryKnowledge/04_Legal_Compliance/)
- [Governance & Tokenomics](../NecessaryKnowledge/05_Governance_Tokenomics/)
- [Security & Resilience](../NecessaryKnowledge/06_Security_Resilience/)
- [Communication & Social Networks](../NecessaryKnowledge/07_Communication_SocialNetworks/)
- [Project Management & Community](../NecessaryKnowledge/08_Project_Management/)
- [Infrastructure & Operations](../NecessaryKnowledge/09_Infrastructure_Operations/)
- [Ethics & Philosophy](../NecessaryKnowledge/10_Ethics_Phylosophy/)
- [Education & Adoption](../NecessaryKnowledge/11_Education_Adoption/)
- [Institutional Relations & Diplomacy](../NecessaryKnowledge/12_Institutional_Relations/)
- [Political Economy & Impact Analysis](../NecessaryKnowledge/13_Political_Economy/)
- [Crisis Management & Organizational Resilience](../NecessaryKnowledge/14_Crisis_Management/)
- [Sustainability & Environmental Impact](../NecessaryKnowledge/15_Sustainability/)

---

## 🚦 Quickstart

```sh
# Backend (Go)
cd blockchain && go run main.go

# Frontend (Next.js)
cd frontend && npm install && npm run dev
```

---

## 🤝 Contribuer

- Voir [CONTRIBUTING.md](CONTRIBUTING.md)
- Rejoindre la communauté Discord, Telegram, etc.

---

## 🌐 Réseaux sociaux

- [Medium](https://medium.com/@LibraDemosChain)
- [YouTube](https://www.youtube.com/@LibraDemosChain)
- [TikTok](https://www.tiktok.com/@librademoschain.l)
- [Snapchat](https://www.snapchat.com/add/librademoschain)
- [LinkedIn](https://www.linkedin.com/company/LibraDemos)
- [GitHub](https://github.com/LibraDemosChain)
- [Discord](https://discord.gg/6QfYVPtm)
- [Telegram](https://t.me/+PaesMzbS0xIzYWM0)
- [WhatsApp](https://chat.whatsapp.com/G9CPYwVlEn44rY7ZjNUjDs)
- [Twitter (X)](https://twitter.com/LibraDemosChain)
- [Instagram](https://www.instagram.com/librademoschain?igsh=dW5ibWZ3Zzg5ZjFy&utm_source=qr)

---

> Pour toute question, contactez support@librademoschain.com ou ouvrez une issue sur GitHub.
