import React, { useState } from 'react';
import { supabase } from '../services/supabase';

const SupabaseDemo = () => {
  const [user, setUser] = useState<any>(null);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleLogin = async () => {
    setError('');
    const { data, error } = await supabase.auth.signInWithPassword({ email, password });
    if (error) setError(error.message);
    else setUser(data.user);
  };

  return (
    <div>
      <h2>Démo Supabase</h2>
      {user ? (
        <div>
          <p>Connecté en tant que : {user.email}</p>
        </div>
      ) : (
        <div>
          <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
          <input placeholder="Mot de passe" type="password" value={password} onChange={e => setPassword(e.target.value)} />
          <button onClick={handleLogin}>Connexion</button>
          {error && <p style={{ color: 'red' }}>{error}</p>}
        </div>
      )}
    </div>
  );
};

export default SupabaseDemo;
