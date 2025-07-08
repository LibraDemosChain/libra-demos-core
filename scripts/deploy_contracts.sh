#!/bin/bash
# Déploiement des smart contracts CosmWasm
set -e
cd $(dirname $0)/..
cd smart-contracts
cargo build --release
# Ajoutez ici la commande de déploiement adaptée à votre réseau
# ex: wasmd tx wasm store target/wasm32-unknown-unknown/release/*.wasm --from wallet --chain-id chainid
