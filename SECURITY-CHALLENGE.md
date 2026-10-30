# 🚨 Horizon Core - The Atomic License Challenge

We challenge white-hat hackers, security auditors, and the developer community to test the robustness of our **Atomic License Enforcement** under concurrent stress.

## The Goal
Bypass the strict volume limit of the Trial License (1,000,000 TX) or crash the Switch node using Race Conditions, API flooding, or payload spoofing.

## Our Defense Mechanism (Fail-Closed)
We rely on atomic SQL reservations executed *before* the transaction processes. 

Here is the exact terminal proof from our live 7-node network:

[1] Verifying Horizon Angel Network (7 Nodes Consensus)
  ✅ Node :9001 -> Law Hash: 3ac449e02f8e6410...
  ...
  🏆 PROOF: All 7 nodes share IDENTICAL law_hash.

[2] Verifying Security & Fail-Closed Mechanism
  SQLi...... BLOCKED (400)
  AuthBypass BLOCKED (401)
  XSS....... BLOCKED (401)
  🏆 PROOF: System is Fail-Closed.

[3] Verifying Atomic License Logic (Go 1.24)
  SQL: db.Model(&License{}).Where("status = ? AND used + ? <= volume", "active", count).UpdateColumn(...)
  🏆 PROOF: Atomic Reservation enforced BEFORE execution.

## How to Participate
1. Download the Trial Package from our GitHub Releases.
2. Run it in an isolated Docker environment.
3. If you find a bypass, open an Issue with your curl payload.
