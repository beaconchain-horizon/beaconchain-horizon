# Horizon Core v1.4 — Independent Verification Guide

**Philosophy:** We do not ask you to trust us. We give you the tools to verify us.

Any security auditor, bank, or independent reviewer can run the commands below on their own machine and confirm every claim made by Horizon Core.

---

## Prerequisites
- OS: Windows 10/11 (Git Bash) or Linux
- Go: 1.21+
- PostgreSQL on `localhost:5432`
- USB Key at `/d/mykey.txt` (64 hex chars)

---

## Test 1 — Health Check (12 Services)
**Expected:** All 10 ports (Switch, Offline Switch, 7 Angels, Guardian) return `200`.

```bash
cd ~/Beaconchain/horizon-core
export ADMIN_TOKEN="bench0000000000000000000000000000000000000000000000000000000bench"
for port in 8091 8092 9001 9002 9003 9004 9005 9006 9007 9008; do
  code=$(curl -s -o /dev/null -w "%{http_code}" --max-time 2 http://localhost:$port/health 2>/dev/null)
  echo "  :$port → $code"
done
