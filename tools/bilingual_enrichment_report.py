import os
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
KEYWORDS = [
    'blockchain', 'cryptography', 'privacy', 'ZK', 'DAO', 'governance', 'AI', 'KYC', 'tokenomics',
    'Cosmos', 'Namada', 'Anoma', 'NEOMA', 'smart contract', 'resilience', 'compliance', 'multichain',
]

report = []

for dirpath, _, filenames in os.walk(ROOT):
    for f in filenames:
        if f.endswith('.en.md'):
            path = os.path.join(dirpath, f)
            with open(path, encoding='utf-8') as file:
                content = file.read().lower()
                if any(kw in content for kw in KEYWORDS):
                    report.append(path)

if report:
    print("Technical/high-value bilingual docs to review for enrichment:")
    for line in report:
        print("-", line)
else:
    print("No technical/high-value modules detected for enrichment.")
