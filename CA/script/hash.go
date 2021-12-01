package main

import (
	"crypto/sha256"
)

// Use Sha256 to hash an input string
func hash(s string) []uint8 {
	data := []byte(s)
	h := sha256.Sum256(data)
	return h[:]
}
