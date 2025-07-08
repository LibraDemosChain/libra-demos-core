import React, { useState } from 'react';
import { SigningStargateClient } from '@cosmjs/stargate';

const CosmosDemoEn = () => {
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
      <h2>Cosmos Demo</h2>
      <input placeholder="Cosmos address" value={address} onChange={e => setAddress(e.target.value)} />
      <button onClick={handleCheckBalance}>Check balance</button>
      {balance && <p>Balance: {balance}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
};

export default CosmosDemoEn;
