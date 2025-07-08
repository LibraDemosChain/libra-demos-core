import React, { useState } from 'react';
import { callAI } from '../services/ai';

const AIDemo = () => {
  const [prompt, setPrompt] = useState('');
  const [result, setResult] = useState('');
  const [loading, setLoading] = useState(false);

  const handleCallAI = async () => {
    setLoading(true);
    const res = await callAI(prompt);
    setResult(res);
    setLoading(false);
  };

  return (
    <div>
      <h2>Démo IA</h2>
      <input placeholder="Prompt" value={prompt} onChange={e => setPrompt(e.target.value)} />
      <button onClick={handleCallAI} disabled={loading}>Appeler l’IA</button>
      {loading && <p>Chargement...</p>}
      {result && <p>Résultat : {result}</p>}
    </div>
  );
};

export default AIDemo;
