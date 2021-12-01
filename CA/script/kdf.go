package main

import (
	"crypto/sha1"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func kdf(password string, salt string) [32]uint8 {
	// generate the derived key
	// base on the input parameter
	// password to stretch
	// salt to use
	// number of iteration to use while generating the new key
	// length of the new key to be generated
	// hash function to be used while deriving the new key
	derivedKey := pbkdf2.Key([]byte(password), []byte(salt), 10, 128, sha1.New)

	// compute the hash function of the derived key to make it stronger
	hashDerivedKey := sha256.Sum256(derivedKey)

	return hashDerivedKey
}
