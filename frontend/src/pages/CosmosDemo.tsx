import React, { useState } from 'react';
import { SigningStargateClient } from '@cosmjs/stargate';

const CosmosDemo = () => {
  const [rpcUrl, setRpcUrl] = useState('https://rpc.cosmos.directory/cosmoshub');
  const [address, setAddress] = useState('');
  const [balance, setBalance] = useState<string | null>(null);
  const [error, setError] = useState('');

  const handleCheckBalance = async () => {
    setError('');
    try {
      const client = await SigningStargateClient.connect(rpcUrl);
      const result = await client.getAllBalances(address);
      setBalance(result.length > 0 ? result[0].amount + ' ' + result[0].denom : '0');
    } catch (e: any) {
      setError(e.message);
    }
  };

  return (
    <div>
      <h2>Démo Cosmos</h2>
      <input placeholder="Adresse Cosmos" value={address} onChange={e => setAddress(e.target.value)} />
      <button onClick={handleCheckBalance}>Vérifier le solde</button>
      {balance && <p>Solde : {balance}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
};

export default CosmosDemo;
