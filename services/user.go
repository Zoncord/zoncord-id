package services

import (
	"fmt"
	"github.com/Zoncord/zoncord-id/errors"
	"golang.org/x/crypto/bcrypt"
)

func PasswordComplexityCheck(password string) error {
	// Password complexity check function
	if len(password) < 8 {
		return errors.PasswordTooShort
	}
	if len(password) > 64 {
		return errors.PasswordTooLong
	}
	return nil
}

func PasswordValidation(password1 string, password2 string) (bool, error) {
	// Password validation function
	if password1 != password2 {
		return false, fmt.Errorf("Passwords do not match.")
	}
	return true, nil
}

func PasswordHasher(password string) string {
	// Password hasher function
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(passwordHash)
}
