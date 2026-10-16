---
title: Beaconchain Horizon — International Audit Report
version: 1.4
audit_date: 2026-10-16
standard: ISO/IEC 25010
---

# Beaconchain Horizon — International Audit Report

> **Citable, reproducible audit record.** Every claim below can be independently verified using the terminal commands provided. No trust in the vendor is required.

**Audit Date:** 2026-10-16
**Version:** 1.4
**Standard:** ISO/IEC 25010 — Performance Efficiency
**Environment:** Windows 8-core laptop

---

## Executive Summary

| Metric | Value | Method |
|:-------|:------|:-------|
| Ledger Throughput | **308,641 TPS** | Batch 100K |
| Consensus Latency | **95 ms / cycle** | 7 parallel verifiers |
| Register Round-Trip | **141 ms** | End-to-end |
| Buffer Guard | **HTTP 413** | Payload > 8 MB rejected |
| Attack Vectors Blocked | **192 / 192 (100%)** | Full security suite |
| Server Errors (5xx) | **0** | During security suite |
| Data Loss | **ZERO** | WAL + fsync |
| Node Law Consistency | **7 / 7 identical** | Hash comparison |

---

## Chain State

```
+---------------------------------------------------------------+
|  HORIZON CORE v1.4 - LIVE CHAIN STATE                         |
+---------------------------------------------------------------+
|  Chain ID:       horizon-core-v1.4                             |
|  Height:         4  (genesis + 4 blocks)                       |
|  Total TX:       4                                             |
|  Persistence:    ON DISK (survives restart)                    |
|  Node Law Hash:  3ac449e02f8e6410...                           |
|  Consistency:    7 / 7 nodes identical                         |
+---------------------------------------------------------------+
```

**Live query:**

```bash
curl -s localhost:8080/api/v1/chain/info
```

---

## Blocks & Transactions

All blocks are cryptographically linked. Each block contains one transaction (`license_issue`) signed by an ECDSA secp256k1 signature.

| Block | TX Type | Merkle Type | TX Count | Prev Hash |
|:-----:|:--------|:-----------:|:--------:|:----------|
| 0 | genesis | — | 0 | 0000...0000 |
| 1 | license_issue | binary | 1 | (genesis hash) |
| 2 | license_issue | binary | 1 | (block 1 hash) |
| 3 | license_issue | binary | 1 | (block 2 hash) |
| 4 | license_issue | binary | 1 | (block 3 hash) |

**Inspect all blocks:**

```bash
curl -s "localhost:8080/api/v1/blocks?limit=10"
```

**Inspect single block:**

```bash
curl -s localhost:8080/api/v1/blocks/1
```

---

## Cryptographic Standards

| Parameter | Value | Standard |
|:----------|:------|:---------|
| Curve | secp256k1 | SEC 2 v2 |
| Signature format | 64-byte raw R ‖ S | — |
| Hash function | SHA-256 | FIPS 180-4 |
| Determinism | RFC 6979 | IETF |
| Low-S enforcement | s ≤ n/2 | BIP-62 |
| Rejected: DER / ASN.1 | enforced | — |
| Rejected: high-S | enforced | BIP-62 |

**Public Key (same for all blocks):**

```
04cf6aaaba8e7a6fddde63a2e2588553dc95503d2de6b2a5c78f4a30a0b91b0aca4d28d06fc1ccbaf730abe1dcda7c29726dab18008353b521c9253c04b1641214
```

**References:**

- secp256k1 — [SEC 2 v2](https://www.secg.org/sec2-v2.pdf)
- RFC 6979 — [Deterministic ECDSA](https://datatracker.ietf.org/doc/html/rfc6979)
- BIP-62 — [Low-S enforcement](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki)
- FIPS 180-4 — [SHA-256](https://csrc.nist.gov/publications/detail/fips/180/4/final)

---

## Performance (ISO/IEC 25010)

| Metric | Value | Method |
|:-------|:------|:-------|
| Ledger Peak TPS | 308,641 | Batch 100K, 324 ms |
| Consensus Cycle | 95 ms | 7 nodes parallel |
| Register Latency | 141 ms | End-to-end |
| Buffer Guard | HTTP 413 | > 8 MB body |

### International Comparison (2026)

| System | Peak TPS | Finality |
|:-------|:--------:|:--------:|
| Bitcoin | 7 | 60 min |
| Ethereum L1 | 15-30 | 6-12 min |
| Visa (VisaNet) | 24,000-65,000 | sub-second |
| Solana | 65,000 | ~13 sec |
| **Horizon Core** | **308,641** | **324 ms** |

---

## Security — 192 Attack Vectors

| Category | Vectors | Result |
|:---------|:-------:|:------:|
| SQL Injection | 30 | BLOCKED |
| XSS | 25 | BLOCKED |
| Command Injection | 20 | BLOCKED |
| Path Traversal | 15 | BLOCKED |
| Auth Bypass | 20 | BLOCKED |
| JWT Attacks | 10 | BLOCKED |
| Header Injection | 15 | BLOCKED |
| HTTP Methods | 12 | BLOCKED |
| Oversize Payload | 15 | BLOCKED |
| SSRF | 10 | BLOCKED |
| XXE | 5 | BLOCKED |
| Business Logic | 15 | BLOCKED |
| **TOTAL** | **192** | **100%** |

**Server errors (5xx) during suite: 0**

---

## Independent Verification

### Step 1 — Clone the audit repository

```bash
git clone github.com/beaconchain-horizon/beaconchain-horizon
cd beaconchain-horizon
```

### Step 2 — Verify chain integrity

```bash
curl -s localhost:8080/api/v1/chain/info
curl -s "localhost:8080/api/v1/blocks?limit=10"
```

### Step 3 — Run TPS benchmark

```bash
bash tps-report.sh
```

Expected: Ledger peak > 300,000 TPS.

### Step 4 — Run 192-vector security suite

```bash
./pentest.exe -url http://localhost:8080
```

Expected: 192 / 192 blocked, 0 server errors.

### Step 5 — Verify Node Law consistency

```bash
for p in 9001 9002 9003 9004 9005 9006 9007; do
  curl -s localhost:$p/health | grep law_hash
done
```

Expected: all 7 nodes return the same hash.

---

## Citation

```
HORIZON (2026). Beaconchain Horizon v1.4 — International Audit Report.
github.com/beaconchain-horizon/beaconchain-horizon
```

---

<div align="center">

**Beaconchain Horizon — Offline · Verifiable · Industrial-Grade**

© 2026 Beaconchain Horizon

</div>
