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

## 🛠️ How to Test (For Hackers)
Since our architecture is Offline-First and Air-Gap ready, there is no public API to attack. You must attack it locally:

1. **Download the Trial Package:**
   Go to: https://github.com/beaconchain-horizon/horizon-core/releases
   Download the latest `trial` zip file.

2. **Run the Isolated Environment:**
   Extract the zip. Run the `start.bat` (Windows) or use the provided Docker setup.
   The Backend will run on port `8080` and the Switch on port `8091`.

3. **Launch Your Attack:**
   Use `curl`, Postman, or your custom scripts to bombard the Switch (`/api/v1/tx/batch`) with concurrent requests. Try to bypass the atomic SQL check or crash the server.

4. **Report a Vulnerability:**
   If you successfully bypass the limit or crash the node, open an Issue with your `curl` payload and terminal output. We will reward verified critical vulnerabilities.
