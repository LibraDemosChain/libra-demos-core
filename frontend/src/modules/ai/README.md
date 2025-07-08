# Module IA NEOMA

Ce module intègre l’intelligence artificielle NEOMA pour la gouvernance, la modération et l’analyse anti-désinformation.

## Fonctionnalités principales
- Analyse automatique des propositions et votes
- Détection de désinformation et modération intelligente
- Recommandations pour la gouvernance DAO

## Installation & configuration
- Dépendances IA à installer (voir `package.json`)
- Configuration : `frontend/src/modules/ai/config.js`

## Développement
- Code source : `frontend/src/modules/ai/`
- Tests : `tests/ai/`

## Exemples d’utilisation CLI

```sh
# Analyser du contenu
ldcd tx neoma analyze-content --content-hash "sha256hash" --content-type "article" --content "Contenu à analyser..." --from citizen1

# Consulter une analyse
ldcd query neoma analysis sha256hash
```

## Ressources
- [docs/GOVERNANCE.md](../../../docs/GOVERNANCE.md)
- [docs/STRUCTURE_IDEALE.md](../../../docs/STRUCTURE_IDEALE.md)
