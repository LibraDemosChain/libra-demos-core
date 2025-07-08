# Monitoring & Observabilité LibraDemosChain

## Prometheus & Metrics

- Le nœud expose des métriques Prometheus sur :26660/metrics
- Exemple :
```sh
curl localhost:26660/metrics
```

## APIs REST, gRPC & Tendermint

- API REST (LCD) :
```sh
curl localhost:1317/cosmos/gov/v1/proposals
```
- API gRPC :
```sh
grpcurl -plaintext localhost:9090 list
```
- API Tendermint RPC :
```sh
curl localhost:26657/status
```

## Logs & supervision

- Les logs détaillés sont accessibles via :
```sh
ldcd start --log_level info
```
- Intégration possible avec Grafana, Loki, etc.
