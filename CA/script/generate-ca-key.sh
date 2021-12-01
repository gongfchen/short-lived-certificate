#!/bin/bash

CA_KEY_PATH="../storage/root-certificate/ca_key.pem"
CA_CERT_PATH="../storage/root-certificate/ca_cert.pem"

openssl genrsa -out $CA_KEY_PATH 2048
openssl req -new -x509 -sha256 -key $CA_KEY_PATH -out $CA_CERT_PATH -days 365
