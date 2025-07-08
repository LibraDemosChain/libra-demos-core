// Hook de gestion du wallet utilisateur (exemple, à enrichir)
import { useState } from 'react';

export function useWallet() {
  const [address, setAddress] = useState<string | null>(null);
  // Ajoutez ici la logique de connexion, déconnexion, récupération du solde, etc.
  return { address, setAddress };
}
