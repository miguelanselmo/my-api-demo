package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash returns the bcrypt hash of the password
func Hash(value string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash: %w", err)
	}
	return string(hashedValue), nil
}

// CheckHash checks if the provided password is correct or not
func CheckHash(value string, hashedValue string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
}
