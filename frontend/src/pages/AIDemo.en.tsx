import React, { useState } from 'react';
import { callAI } from '../services/ai';

const AIDemoEn = () => {
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
      <h2>AI Demo</h2>
      <input placeholder="Prompt" value={prompt} onChange={e => setPrompt(e.target.value)} />
      <button onClick={handleCallAI} disabled={loading}>Call AI</button>
      {loading && <p>Loading...</p>}
      {result && <p>Result: {result}</p>}
    </div>
  );
};

export default AIDemoEn;
