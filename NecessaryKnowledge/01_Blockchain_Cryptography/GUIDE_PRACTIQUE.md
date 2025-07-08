# Guide Pratique – Blockchain & Cryptographie

Ce guide propose des tutoriels concrets, des bonnes pratiques et des exemples pour appliquer la cryptographie et la blockchain dans LibraDemosChain.

## 1. Générer une paire de clés (Cosmos SDK)
```sh
# Générer une clé pour un validateur Cosmos
$ gaiad keys add mon_validateur
```

## 2. Déployer un smart contract CosmWasm
```sh
# Compiler le contrat
$ cargo wasm
# Déployer sur testnet Cosmos
$ wasmd tx wasm store mon_contrat.wasm --from wallet --chain-id=CHAIN_ID
```

## 3. Créer une preuve ZK avec Circom
```sh
# Compiler un circuit Circom
$ circom vote.circom --r1cs --wasm --sym
# Générer une preuve
$ snarkjs groth16 fullprove input.json vote.r1cs vote.wasm zkey_final.zkey proof.json public.json
```

## 4. Archiver un document sur Arweave
```sh
# Uploader un fichier
$ arweave deploy mon_fichier.pdf --wallet wallet.json
```

## Bonnes pratiques
- Toujours auditer les smart contracts (Slither, Mythril).
- Utiliser des wallets multi-signatures pour la gouvernance.
- Privilégier les primitives cryptographiques éprouvées (Poseidon, Ed25519).

## Erreurs fréquentes
- Mauvaise gestion des clés privées : privilégier le stockage chiffré et la sauvegarde hors-ligne.
- Oublier de vérifier la compatibilité IBC lors de l’intégration multichain.

## Aller plus loin
- Voir [Resources.md](./Resources.md) pour des liens avancés.
