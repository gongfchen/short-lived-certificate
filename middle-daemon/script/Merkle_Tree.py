# -*- coding: utf-8 -*-
"""
Created on Sun May 16 23:19:06 2021

@author: galan
"""

import string
import hashlib
import merklelib
from merklelib import MerkleTree

def hashfunc(value):
    return hashlib.sha256(value).hexdigest()

data = list(string.ascii_letters)

tree = MerkleTree(data, hashfunc)

proof = tree.get_proof('A')

if tree.verify_leaf_inclusion('A', proof):
    print('A is in the tree')
else:
    print('A is not in the tree')
    
    
MR = tree.merkle_root


kappa = merklelib.verify_leaf_inclusion('A', proof, hashfunc, MR)


from merklelib import utils

keepo = hashfunc(b'\x00' + utils.to_string('a'))