# Module Privacy (Confidentialité)

Ce module implémente la privacy avancée pour LibraDemosChain : ZK, Anoma, Namada, Ring Signatures, shielded transactions.

## Exemples d’utilisation CLI

```sh
# Activer la privacy dans la config
go run main.go --privacy

# Soumettre une transaction anonyme
ldcd tx privacy send --from citizen1 --to citizen2 --amount 100ldc --zk-proof "0x..."
```

## Ressources
- [docs/SECURITY.md](../../../docs/SECURITY.md)
- [docs/STRUCTURE_IDEALE.md](../../../docs/STRUCTURE_IDEALE.md)
