#!/bin/bash
cd "$(dirname "$0")"
SIG=$(grep -oE '"signature": "[0-9a-f]{128}"' certificate.json | cut -d'"' -f4)
PUB=$(grep -oE '"authority_pubkey": "[0-9a-f]{130}"' certificate.json | cut -d'"' -f4)
HASH=$(grep -v '"signature"' certificate.json | tr -d '\n' | sha256sum | cut -d' ' -f1)
echo "sig:  ${SIG:0:32}... (len=${#SIG})"
echo "pub:  ${PUB:0:32}..."
echo "hash: $HASH"
python3 - <<PYEOF
from ecdsa import VerifyingKey, SECP256k1
from ecdsa.util import sigdecode_string
try:
    sig = bytes.fromhex("$SIG")
    h   = bytes.fromhex("$HASH")
    vk  = VerifyingKey.from_string(bytes.fromhex("${PUB:2}"), curve=SECP256k1)
    vk.verify(sig, h, hashfunc=None, sigdecode=sigdecode_string)
    print("✓✓✓ CERTIFICATE VALID")
except Exception as e:
    print(f"✗ {e}")
PYEOF
