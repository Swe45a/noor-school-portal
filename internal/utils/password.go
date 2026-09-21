package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

const passwordAlphabet = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz23456789"
const passwordLength = 10
const bcryptCost = 12

// GenerateRandomPassword returns a random 10-character alphanumeric password.
func GenerateRandomPassword() (string, error) {
	b := make([]byte, passwordLength)
	alphabetLen := big.NewInt(int64(len(passwordAlphabet)))

	for i := range b {
		n, err := rand.Int(rand.Reader, alphabetLen)
		if err != nil {
			return "", fmt.Errorf("generating password: %w", err)
		}
		b[i] = passwordAlphabet[n.Int64()]
	}

	return string(b), nil
}

// HashPassword hashes a plaintext password with bcrypt at cost 12.
func HashPassword(plaintext string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcryptCost)
	if err != nil {
		return "", fmt.Errorf("hashing password: %w", err)
	}
	return string(hash), nil
}
