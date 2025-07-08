# Exemples – Blockchain & Cryptographie

## Exemple 1 : Signature et vérification (Rust)
```rust
use ed25519_dalek::{Keypair, Signature, Signer, Verifier};
let keypair = Keypair::generate(&mut rand::rngs::OsRng);
let message = b"LibraDemosChain";
let signature: Signature = keypair.sign(message);
assert!(keypair.verify(message, &signature).is_ok());
```

## Exemple 2 : Circuit Circom pour vote anonyme
```circom
// vote.circom
component main = Vote();
signal input vote;
signal output hash;
hash <== Poseidon([vote]);
```

## Exemple 3 : Commande CLI pour upload sur IPFS
```sh
ipfs add mon_document.pdf
```

## Exemple 4 : Script de monitoring d’un smart contract (CosmWasm)
```sh
wasmd query wasm contract-state smart <CONTRACT_ADDR> '{"state":{}}'
```
