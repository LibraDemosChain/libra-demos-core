# Exemples – Infrastructure & Operations

## Exemple 1 : Script Docker Compose
```yaml
version: '3'
services:
  backend:
    build: ./backend
  frontend:
    build: ./frontend
```

## Exemple 2 : Dashboard Grafana
- Visualiser les métriques des nœuds Cosmos.

## Exemple 3 : Script de backup
```sh
tar czf backup.tar.gz ./data
```
