package services

import (
	"github.com/Zoncord/zoncord-id/models"
	"github.com/Zoncord/zoncord-id/validation"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHasher(password string) string {
	// Password hasher function
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(passwordHash)
}

func SignUp(email, password1, password2, firstName, lastName string) (string, error) {
	err := validation.PasswordsValidation(password1, password2)
	if err != nil {
		return "", err
	}
	var user models.User
	err = user.Create(
		email,
		password1,
		firstName,
		lastName,
	)
	if err != nil {
		return "", err
	}
	token, err := models.CreateAccessToken(user, 0, "read write")
	return token.Token, err
}
