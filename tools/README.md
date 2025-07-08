# LibraDemosChain Documentation Tools

## check_bilingual_sync.py
Script to check that every documentation file (.md) has a synchronized English (.en.md) and French (.md) version. Reports missing or unsynchronized files. Integrates with CI.

## bilingual_enrichment_report.py
Lists all English documentation files containing technical or high-value keywords (blockchain, privacy, AI, DAO, etc.) for targeted human review and enrichment.

---

**Usage:**

```sh
python tools/check_bilingual_sync.py
python tools/bilingual_enrichment_report.py
```

---

These tools help maintain a professional, fully bilingual, and up-to-date documentation base for LibraDemosChain.
