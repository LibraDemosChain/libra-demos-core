# Practical Guide – Blockchain & Cryptography

[🇫🇷 Version française](./GUIDE_PRACTIQUE.md)

This guide provides concrete tutorials, best practices, and examples for applying cryptography and blockchain in LibraDemosChain.

## 1. Generate a key pair (Cosmos SDK)
```sh
# Generate a key for a Cosmos validator
$ gaiad keys add my_validator
```

## 2. Deploy a CosmWasm smart contract
```sh
# Compile the contract
$ cargo wasm
# Deploy on Cosmos testnet
$ wasmd tx wasm store my_contract.wasm --from wallet --chain-id=CHAIN_ID
```

## 3. Create a ZK proof with Circom
```sh
# Compile a Circom circuit
$ circom vote.circom --r1cs --wasm --sym
# Generate a proof
$ snarkjs groth16 fullprove input.json vote.r1cs vote.wasm zkey_final.zkey proof.json public.json
```

## 4. Archive a document on Arweave
```sh
# Upload a file
$ arweave deploy my_file.pdf --wallet wallet.json
```

## Best practices
- Always audit smart contracts (Slither, Mythril).
- Use multi-signature wallets for governance.
- Prefer proven cryptographic primitives (Poseidon, Ed25519).

## Common mistakes
- Poor private key management: prefer encrypted storage and offline backup.
- Forgetting to check IBC compatibility when integrating multichain.
