import React, { createContext, useContext, useState, ReactNode } from 'react';

type Language = 'fr' | 'en';

const LanguageContext = createContext<{ lang: Language; setLang: (l: Language) => void }>({ lang: 'fr', setLang: () => {} });

export const LanguageProvider = ({ children }: { children: ReactNode }) => {
  const [lang, setLang] = useState<Language>('fr');
  return <LanguageContext.Provider value={{ lang, setLang }}>{children}</LanguageContext.Provider>;
};

export const useLanguage = () => useContext(LanguageContext);
