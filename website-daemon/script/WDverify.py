# -*- coding: utf-8 -*-
"""
Created on Wed May 19 00:38:21 2021

@author: galan
"""
#!/usr/local/bin/python
import pickle
import hashlib
import merklelib
import sys

def hashfunc(value):
    return hashlib.sha256(value).hexdigest()

root = sys.argv[1]
proof_path = "../../middle-daemon-website-daemon-storage/Daily Cert Verification/proof.pickle"
with open(proof_path, 'rb') as f:
    proof = pickle.load(f)
    
hash_path = "../storage/Daily Cert Verification/cert_hash.txt"


with open(hash_path) as f:
    contents = f.readlines()

cert_hash = contents[0]

result_file = "../storage/result/result.txt"

text_file = open(result_file, "w+")

if (merklelib.verify_leaf_inclusion(cert_hash, proof, hashfunc, root)):
    n = text_file.write("Present certificate is valid")
else:
    n = text_file.write("Present certificate is not valid")
text_file.close()
