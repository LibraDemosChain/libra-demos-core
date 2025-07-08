# Module ZK-KYC

Ce module gère la vérification d’identité confidentielle via Zero-Knowledge Proofs.

## Exemples d’utilisation CLI

```sh
# Soumettre une preuve ZK-KYC
ldcd tx zkkyc submit-kyc-proof --proof-data "0x..." --public-inputs "..." --kyc-level 1 --from citizen1

# Vérifier le statut KYC
ldcd query zkkyc verification ldc1...
```

## Ressources
- [docs/SECURITY.md](../../../docs/SECURITY.md)
- [docs/STRUCTURE_IDEALE.md](../../../docs/STRUCTURE_IDEALE.md)
