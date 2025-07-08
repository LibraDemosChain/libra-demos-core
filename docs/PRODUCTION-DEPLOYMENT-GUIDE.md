# PRODUCTION DEPLOYMENT GUIDE - LibraDemosChain

[🇫🇷 Version française](../../LibraDemosChain-Copilote/LibraDemosChain-Copilote/PRODUCTION-DEPLOYMENT-GUIDE.md)

## SYSTEM REQUIREMENTS

### Minimum Server
- **CPU**: 8 cores (16 threads)
- **RAM**: 32 GB
- **Storage**: 1 TB SSD NVMe
- **Network**: 1 Gbps
- **OS**: Ubuntu 22.04 LTS

### Recommended Server
- **CPU**: 16 cores (32 threads)
- **RAM**: 64 GB
- **Storage**: 2 TB SSD NVMe RAID1
- **Network**: 10 Gbps
- **OS**: Ubuntu 22.04 LTS

## ENVIRONMENT VARIABLES

### Required
```bash
# Database
DATABASE_URL=postgresql://user:password@localhost:5432/librademos

# Security
SESSION_SECRET=your-256-bit-secret-key
JWT_SECRET=your-jwt-secret-key

# Email
GMAIL_USER=your-email@gmail.com
GMAIL_APP_PASSWORD=your-app-password
```

### AI (Optional but recommended)
```bash
# OpenAI (for advanced NEOMA)
OPENAI_API_KEY=sk-your-openai-key

# Anthropic (for advanced NEOMA)
# ...
```

...existing content...
