# -*- coding: utf-8 -*-
"""
Created on Tue May 18 22:29:40 2021

@author: galan
"""
#!/usr/bin/env python

import hashlib
from merklelib import MerkleTree
from pathlib import Path
import sys

def hashfunc(value):
    return hashlib.sha256(value).hexdigest()

dom_name = sys.argv[1]

file_name = "../../CA-middle-daemon-storage/Hashlists/" + dom_name + "/hashlist.txt"
Path("../storage/merkle-roots/" + dom_name).mkdir(parents=True, exist_ok=True)
root_file = "../storage/merkle-roots/" + dom_name + "/merkleroot.txt"
hash_file = open(file_name, 'r')

hash_list = []
for line in hash_file:
  hash_list.append(line[:-1])

hash_file.close()

tree = MerkleTree(hash_list, hashfunc)
root = tree.merkle_root

text_file = open(root_file, "w+")
n = text_file.write(root)
text_file.close()
