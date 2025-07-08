#!/bin/bash
# Build complet backend + frontend + smart contracts
set -e
cd $(dirname $0)/..
cd blockchain && go build -o librademosd
cd ../frontend && npm install && npm run build
cd ../smart-contracts && cargo build --release
