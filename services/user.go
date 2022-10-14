package services

import (
	"github.com/Zoncord/zoncord-id/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// TODO: move validation functions to validation/rules.go file
func PasswordNumberValidation(password string) error {
	for _, letter := range password {
		if rune('0') <= letter && letter <= rune('9') {
			return nil
		}
	}
	return errors.PasswordMustIncludeNumber
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

func PasswordsEquivalencyValidation(password1 string, password2 string) error {
	if password1 != password2 {
		return errors.PasswordsDontMatch
	}
	return nil
}

func PasswordValidation(password string) error {
	err := PasswordLengthValidation(password)
	if err != nil {
		return err
	}
	err = PasswordNumberValidation(password)
	if err != nil {
		return err
	}
	return nil
}

func PasswordsValidation(password1 string, password2 string) error {
	// Password validation function
	err := PasswordsEquivalencyValidation(password1, password2)
	if err != nil {
		return err
	}
	err = PasswordValidation(password1)
	if err != nil {
		return err
	}
	return nil
}

func EmailValidation(email string) error {
	emailRegex := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if !emailRegex.MatchString(email) {
		return errors.InvalidEmailFormat
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
