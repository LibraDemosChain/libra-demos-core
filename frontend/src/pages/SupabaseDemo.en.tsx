import React, { useState } from 'react';
import { supabase } from '../services/supabase';

const SupabaseDemoEn = () => {
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
      <h2>Supabase Demo</h2>
      {user ? (
        <div>
          <p>Logged in as: {user.email}</p>
        </div>
      ) : (
        <div>
          <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
          <input placeholder="Password" type="password" value={password} onChange={e => setPassword(e.target.value)} />
          <button onClick={handleLogin}>Login</button>
          {error && <p style={{ color: 'red' }}>{error}</p>}
        </div>
      )}
    </div>
  );
};

export default SupabaseDemoEn;
