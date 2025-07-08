import os
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
DOC_EXTS = ['.md']
BILINGUAL_SUFFIX = '.en.md'

report = []

for dirpath, _, filenames in os.walk(ROOT):
    files = [f for f in filenames if any(f.endswith(ext) for ext in DOC_EXTS)]
    for f in files:
        if f.endswith(BILINGUAL_SUFFIX):
            base = f[:-len(BILINGUAL_SUFFIX)] + '.md'
            if base not in files:
                report.append(f"Missing FR version for: {os.path.join(dirpath, f)}")
        elif f.endswith('.md') and not f.endswith(BILINGUAL_SUFFIX):
            en = f[:-3] + BILINGUAL_SUFFIX
            if en not in files:
                report.append(f"Missing EN version for: {os.path.join(dirpath, f)}")

if report:
    print("Bilingual documentation desynchronization detected:")
    for line in report:
        print("-", line)
else:
    print("All documentation files are synchronized (FR/EN).")
