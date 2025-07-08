# Examples – Blockchain & Cryptography

[🇫🇷 Version française](./EXEMPLES.md)

## Example 1: Signing and verification (Rust)
```rust
use ed25519_dalek::{Keypair, Signature, Signer, Verifier};
let keypair = Keypair::generate(&mut rand::rngs::OsRng);
let message = b"LibraDemosChain";
let signature: Signature = keypair.sign(message);
assert!(keypair.verify(message, &signature).is_ok());
```

## Example 2: Circom circuit for anonymous voting
```circom
// vote.circom
component main = Vote();
signal input vote;
signal output hash;
hash <== Poseidon([vote]);
```

## Example 3: CLI command to upload to IPFS
```sh
ipfs add my_document.pdf
```

## Example 4: Monitoring script for a smart contract (CosmWasm)
```sh
wasmd query wasm contract-state smart <CONTRACT_ADDR> '{"state":{}}'
```
