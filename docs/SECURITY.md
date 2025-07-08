# Sécurité LibraDemosChain

## Principes

- Privacy by design (ZK, Anoma, Namada)
- Sécurité des smart contracts (CosmWasm)
- Audit régulier (CertiK, Quantstamp, tests de pénétration, vérification formelle ZK)
- Gouvernance sécurisée (vote anonyme, slashing, rotation des validateurs)

## Bonnes pratiques

- Ne jamais partager vos clés privées
- Utiliser des wallets sécurisés (HSM, keyring-backend)
- Sauvegarder les mnémoniques hors-ligne
- Vérifier les smart contracts avant déploiement (linting, formatage, couverture de tests)

## Audit

Voir [SECURITY_AUDIT_CHECKLIST.md](../SECURITY_AUDIT_CHECKLIST.md)

## Standards de code

- Linting : `golangci-lint`, `gofumpt`
- Formatage : `prettier`, `gofumpt`
- Tests : Couverture minimale 80%
- Documentation : Godoc pour toutes les fonctions publiques
