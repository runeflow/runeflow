package util

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomString creates a random 32-byte string
func RandomString() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
