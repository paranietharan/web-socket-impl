#!/bin/bash

set -e  # stop if any command fails

echo "Creating server.key and server.crt ..."

# Generate RSA private key
openssl genrsa -out server.key 2048

# Create self-signed certificate
openssl req -new -x509 -sha256 \
    -key server.key \
    -out server.crt \
    -days 3650 \
    -subj "/C=US/ST=State/L=City/O=MyOrg/OU=Dev/CN=localhost"

echo "Done Generated server.key and server.crt"