package main

import (
	"crypto/rand"
	"encoding/hex"
)

func generateUniqueURL() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
