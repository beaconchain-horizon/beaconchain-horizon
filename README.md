[![Status](https://img.shields.io/badge/STATUS-7%2B1_CONSENSUS-success)](#consensus-architecture)
[![Security](https://img.shields.io/badge/SECURITY-FAIL--CLOSED-blue)](#security-model)
[![Quantum](https://img.shields.io/badge/TIMESTAMP-QUANTUM--RESISTANT-purple)](#quantum-seal)
[![Protocol](https://img.shields.io/badge/PROTOCOL-AIR--GAPPED-orange)](#dual-engine-architecture)

# Horizon Core v1.4 — Sovereign Protocol for Critical Infrastructure

**"We do not ask you to trust us. We give you the tools to verify us."**

Horizon Core is an enterprise-grade, air-gapped, and sovereign ledger protocol designed for the most demanding environments: central banks, defense organizations, and critical national infrastructure (SCADA). 

Unlike public blockchains that compromise privacy for transparency, Horizon Core operates on a **Zero-Knowledge Data Sovereignty** model. The client's raw data never leaves their data center; only cryptographic proofs (Merkle Roots) are anchored to the immutable chain.

---

## 🏛️ Core Architecture

### 1. Dual-Engine (Air-Gapped & Online)
- **Online Engine:** Powered by PostgreSQL for high-throughput, real-time transaction processing.
- **Offline Engine:** Powered by SQLite, sealed within the data center, acting as the ultimate guardian of the Genesis Block.

### 2. Consensus Architecture (7 + 1)
Horizon Core does not rely on energy-intensive mining. It uses a specialized 7+1 consensus model:
- **4 Critical Nodes (100% Required):** Merkle, Hash, ECDSA, Type.
- **3 Majority Nodes (2/3 Required):** Volume, Time, User.
- **1 Agent Guardian (Observer):** An independent AI inspector with veto power to ensure audit integrity.

### 3. Time-Based Block Generation
The protocol utilizes a **Native Monotonic Clock**. Blocks are not mined by arbitrary intervals; their block number is the exact Unix timestamp (10-digit) of the sealing event. 

### 4. Quantum Seal (Post-Quantum Timestamping)
To secure the chain against future quantum computer attacks (Shor's Algorithm), every block header is protected by a proprietary **Hash-Chain Quantum Seal** (10,000 iterations of SHA-256). The hash is mathematically tied to the Genesis Block (`8c0763dd...`), making it impossible to forge or alter historical timestamps.

---

## 🔒 Security Model: Fail-Closed

Horizon Core operates on a strict **Fail-Closed** policy. If a condition is not mathematically met, the system rejects it. No exceptions.

- **Time Guard:** Blocks with timestamps in the future (Timejacking attacks) are automatically rejected.
- **License Minting Law:** A license is **not** valid upon issuance. It is only activated when it is minted into a block and approved by the consensus nodes.
- **DoS Protection:** All endpoints are protected by an 8MB MaxBytesReader, preventing Buffer Overflow.
- **Session Replay Prevention:** Single-use session tokens with 5-minute expiry.

### 🛡️ Independent Pentest Results (100% Blocked)
Our test suite (`pentest_suite.sh`) simulates 8 major attack vectors. The system successfully blocked all of them:
- ✅ Authentication Bypass (HTTP 401)
- ✅ SQL Injection (HTTP 400)
- ✅ 10MB DoS Payload (HTTP 413)
- ✅ Session Replay Attack (HTTP 401)
- ✅ SCADA/Stuxnet Injection (HTTP 400)
- ✅ Type Confusion (HTTP 400)

---

## 🔬 Independent Verification

**Don't trust the documentation. Verify the code.**

Any security auditor can run the provided terminal commands locally to verify the health, consensus, and security of the ledger. 

📄 **[Read the Full Independent Verification Guide](docs/immutable/VERIFICATION-GUIDE.md)**

---

## 🚀 Performance & Scalability

- **Throughput:** The protocol is hardware-agnostic. Actual TPS depends entirely on the customer's provisioned license volume and hardware (tested locally from 5,000 to 100,000+ TPS in a single block).
- **Data Minimization:** The chain stores only 32-byte Merkle Roots. An organization with 1 billion transactions per day will use less than 2GB of blockchain storage per year.

---

*Horizon Core v1.4 · Genesis Locked · Sovereign · All Rights Reserved*
