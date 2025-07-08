// Service d’intégration IA (exemple, à enrichir)
export function callAI(prompt: string): Promise<string> {
  // Exemple d’appel à une API IA (à adapter)
  return fetch('/api/ai', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ prompt })
  })
    .then(res => res.json())
    .then(data => data.result as string);
}
