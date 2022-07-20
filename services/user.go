package services

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordComplexityCheck(password string) error {
	// Password complexity check function
	if len(password) < 8 {
		// Password is too short
	}
	if len(password) > 64 {
		// Password is too long
	}
	return nil
}

func PasswordHasher(password string) string {
	// Password hasher function
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(passwordHash)
}
