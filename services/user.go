package services

import (
	"fmt"
	"github.com/Zoncord/zoncord-id/errors"
	"golang.org/x/crypto/bcrypt"
)

func PasswordNumberValidation(password string) error {
	for _, letter := range password {
		if rune('0') <= letter && letter <= rune('9') {
			return nil
		}
	}
	return fmt.Errorf("password must include number")
}

const PasswordMinLength int = 10
const PasswordMaxLength int = 64

func PasswordLengthValidation(password string) error {
	if len(password) < PasswordMinLength {
		return errors.PasswordTooShort
	}
	if len(password) > PasswordMaxLength {
		return errors.PasswordTooLong
	}
	return nil
}

func PasswordEquivalencyValidation(password1 string, password2 string) error {
	if password1 != password2 {
		return fmt.Errorf("passwords don't match")
	}
	return nil
}

func PasswordsValidation(password1 string, password2 string) error {
	// Password validation function
	err := PasswordEquivalencyValidation(password1, password2)
	if err != nil {
		return err
	}
	err = PasswordLengthValidation(password1)
	if err != nil {
		return err
	}
	err = PasswordNumberValidation(password1)
	if err != nil {
		return err
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
