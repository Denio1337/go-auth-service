package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// Get hashed string from string
func Hash(raw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(encrypt(raw)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Compare raw with hash
func CompareWithHash(raw string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(encrypt(raw)))
	return err == nil
}

// Encrypt raw string with SHA256 to bypass bcrypt length constraint
func encrypt(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
