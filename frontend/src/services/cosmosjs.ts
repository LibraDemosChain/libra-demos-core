// Exemple d’intégration CosmJS (Cosmos SDK)
import { SigningStargateClient } from '@cosmjs/stargate';

export async function connectCosmos(rpcUrl: string, mnemonic: string) {
  // À adapter : gestion du wallet, sécurité, etc.
  const client = await SigningStargateClient.connectWithSigner(rpcUrl, mnemonic);
  return client;
}
