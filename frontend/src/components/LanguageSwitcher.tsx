import React, { useState } from 'react';

const LanguageSwitcher = () => {
  const [lang, setLang] = useState<'fr' | 'en'>('fr');
  return (
    <div>
      <button onClick={() => setLang('fr')}>FR</button>
      <button onClick={() => setLang('en')}>EN</button>
      <div style={{ marginTop: 8 }}>
        {lang === 'fr' ? 'Interface en français' : 'Interface in English'}
      </div>
    </div>
  );
};

export default LanguageSwitcher;
