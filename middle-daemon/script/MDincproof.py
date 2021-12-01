# -*- coding: utf-8 -*-
"""
Created on Wed May 19 00:21:57 2021

@author: galan
"""
#!/usr/bin/env python

import hashlib
from merklelib import MerkleTree
import pickle
import sys

def hashfunc(value):
    return hashlib.sha256(value).hexdigest()

dom_name = sys.argv[1]
day_num = sys.argv[2]
file_name = "../../CA-middle-daemon-storage/Hashlists/" + dom_name + "/hashlist.txt"

hash_file = open(file_name, 'r')

hash_list = []
for line in hash_file:
  hash_list.append(line[:-1])

hash_file.close()

tree = MerkleTree(hash_list, hashfunc)

# Write root
root = tree.merkle_root
root_file = "../../middle-daemon-website-daemon-storage/Daily Cert Verification/merkleroot.txt"

text_file = open(root_file, "w+")
n = text_file.write(root)
text_file.close()

proof = tree.get_proof(hash_list[int(day_num)])

filename = "../../middle-daemon-website-daemon-storage/Daily Cert Verification/proof.pickle"

with open(filename, 'wb') as f:
    pickle.dump(proof, f)