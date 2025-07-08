# Guide Testnet LibraDemosChain

## Rejoindre un testnet Cosmos/Namada

- [Testnets Cosmos](https://github.com/cosmos/testnets)
- [Guide Testnet Namada](https://docs.namada.net/operators/testnet.html)

## Lancer un testnet local

```sh
# Initialiser la chaîne
ldcd init "validator-name" --chain-id librademos-1

# Générer les clés
ldcd keys add validator
ldcd keys add citizen1

# Configurer la genesis
ldcd genesis add-genesis-account validator 1000000000ldc
ldcd genesis add-genesis-account citizen1 1000000neo

# Créer la transaction de genèse du validateur
ldcd genesis gentx validator 1000000ldc --chain-id librademos-1

# Finaliser la genesis
ldcd genesis collect-gentxs

# Démarrer le nœud
ldcd start
```

## Connexion IBC

Voir [blockchain/x/multichain/README.md](../blockchain/x/multichain/README.md)
