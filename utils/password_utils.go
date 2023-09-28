// utils/token_utils.go
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain text password as input and returns its bcrypt hash.
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password with a cost factor of 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword verifies a plain text password against a hashed password.
func VerifyPassword(hashedPassword, plainPassword string) error {
	// Compare the hashed password with the plain password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
