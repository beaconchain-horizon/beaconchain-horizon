<div align="center">

```
  ██╗  ██╗ ██████╗ ██████╗ ██╗███████╗ ██████╗ ███╗   ██╗
  ██║  ██║██╔═══██╗██╔══██╗██║╚══███╔╝██╔═══██╗████╗  ██║
  ███████║██║   ██║██████╔╝██║  ███╔╝ ██║   ██║██╔██╗ ██║
  ██╔══██║██║   ██║██╔══██╗██║ ███╔╝  ██║   ██║██║╚██╗██║
  ██║  ██║╚██████╔╝██║  ██║██║███████╗╚██████╔╝██║ ╚████║
  ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═══╝
```

### ▸ Offline-First Tamper-Evident Ledger ◂

🌍 **GLOBAL · INTEROPERABLE · SOVEREIGN-NEUTRAL** 🌍

🇺🇸 🇬🇧 🇨🇭 🇩🇪 🇫🇷 🇯🇵 🇨🇳 🇷🇺 🇮🇷 🇮🇳 🇧🇷 🇨🇦
🇦🇺 🇰🇷 🇮🇹 🇪🇸 🇸🇪 🇸🇬 🇦🇪 🇹🇷 🇸🇦 🇶🇦 🇿🇦 🇳🇱

![Version](https://img.shields.io/badge/Version-1.4-0ea5e9?style=for-the-badge)
![Date](https://img.shields.io/badge/Date-2026--10--16-38bdf8?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-PRODUCTION-22c55e?style=for-the-badge)

</div>

---

## What It Is

Horizon Core is an offline-first, tamper-evident ledger for banking, industrial, and air-gapped environments. It produces cryptographically signed records that any third party can verify without trust in the vendor.

---

## Core Capabilities

| Domain | Use Case |
|:-------|:---------|
| 🏦 Banking | Inter-bank settlement · dual-control · immutable audit |
| ⚙️ Industry | SCADA anchoring · sensor provenance · regulatory reporting |
| 🔒 Air-Gap | Full offline operation for critical infrastructure |

---

## Verified Performance

| Metric | Value | Method |
|:-------|:------|:-------|
| Ledger Throughput | **523,560 TPS** | Batch 100K · 191 ms |
| Consensus Latency | **94 ms / cycle** | 7 parallel verifiers |
| Register Round-Trip | **167 ms** | End-to-end |
| Data Loss | **ZERO** | WAL + fsync |
| Success Rate | **100%** | All batches |

*Measured on Windows 8-core laptop, 2026-10-16.*

---

## International Standards — Referenced

| Standard | Issuer | Region | Reference |
|:---------|:-------|:------:|:----------|
| **SEC 2 v2** — secp256k1 curve | SECG | 🇺🇸 USA · 🌍 Global | [secg.org/sec2-v2.pdf](https://www.secg.org/sec2-v2.pdf) |
| **RFC 6979** — Deterministic ECDSA | IETF | 🌍 International | [datatracker.ietf.org](https://datatracker.ietf.org/doc/html/rfc6979) |
| **BIP-62** — Low-S enforcement | Bitcoin Core | 🌍 International | [github/bitcoin/bips](https://github.com/bitcoin/bips/blob/master/bip-0062.mediawiki) |
| **FIPS 180-4** — SHA-256 | NIST | 🇺🇸 USA | [csrc.nist.gov](https://csrc.nist.gov/publications/detail/fips/180/4/final) |
| **ISO/IEC 25010** — Software quality | ISO | 🇨🇭 Switzerland · 🌍 | [iso25000.com](https://iso25000.com/index.php/en/iso-25000-standards/iso-25010) |
| **ISO 8583** — Financial transactions | ISO | 🇨🇭 Switzerland · 🌍 | [iso.org/standard](https://www.iso.org/standard/79451.html) |
| **ISO 20022** — Financial messaging | ISO | 🇨🇭 Switzerland · 🌍 | [iso20022.org](https://www.iso20022.org/) |
| **Basel III** — Banking regulation | BIS | 🇨🇭 Switzerland · 🌍 | [bis.org/basel_framework](https://www.bis.org/basel_framework/) |
| **GDPR** — Data protection | European Union | 🇪🇺 EU | [eur-lex.europa.eu](https://eur-lex.europa.eu/eli/reg/2016/679/oj) |

*Standards referenced by design — not certifications claimed. Each is publicly verifiable at the link provided.*

---

## Independent Verification

Every claim here can be reproduced on any machine.

### Step 1 — Clone source

```bash
git clone https://github.com/beaconchain-horizon/horizon-core
cd horizon-core
```

### Step 2 — Start the ecosystem

```bash
start-all-silent.bat
```

### Step 3 — Verify chain integrity

```bash
curl -s localhost:8080/api/v1/chain/info
curl -s "localhost:8080/api/v1/blocks?limit=10"
```

### Step 4 — Run TPS benchmark

```bash
bash tps-report.sh
```

Expected: Ledger peak > 500,000 TPS.

### Step 5 — Run security test suite

```bash
./pentest.exe -url http://localhost:8080
```

Expected: 192 / 192 attack vectors blocked.

### Step 6 — Verify signatures independently

Each block contains a 64-byte raw R ‖ S signature. Verify with any secp256k1 library. No vendor tooling required.

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

---

## Roadmap

| Phase | Capability | Status |
|:-----:|:-----------|:------:|
| 1 | Single-node Ledger | DONE |
| 2 | Air-Gap Mode | DONE |
| 3 | Dual Control | DONE |
| 4 | Industrial Anchor | DONE |
| 5 | Multi-node Consensus | DONE |
| 6 | HSM Integration | Q4 2026 |
| 7 | ISO 27001 | Q2 2027 |
| 8 | FIPS Certification | Q3 2027 |

---

## License

Commercial. For inquiries, contact the maintainers.

---

## Contact

**Organization:** `github.com/beaconchain-horizon`

**Security Repository:** `github.com/beaconchain-horizon/horizon-security-assessment`

---

<div align="center">

### 🌍 Offline · Verifiable · Industrial-Grade 🌍

**🌅 © 2026 HORIZON 🌅**

</div>

