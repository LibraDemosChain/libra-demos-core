# Module DAO & Gouvernance

Ce module gère la gouvernance décentralisée, les DAOs et le vote anonyme.

## Exemples d’utilisation CLI

```sh
# Créer une proposition
ldcd tx governance create-proposal --title "Amélioration du protocole" --description "Proposition d'amélioration..." --proposal-type "parameter_change" --voting-period 604800 --deposit 1000ldc --from citizen1

# Voter sur une proposition
ldcd tx governance vote 1 yes --zk-proof "0x..." --from citizen1

# Consulter les propositions
ldcd query governance proposals
```

## Ressources
- [docs/GOVERNANCE.md](../../../docs/GOVERNANCE.md)
- [docs/STRUCTURE_IDEALE.md](../../../docs/STRUCTURE_IDEALE.md)
