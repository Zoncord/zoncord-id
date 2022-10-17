package services

import (
	"github.com/Zoncord/zoncord-id/models"
	"github.com/Zoncord/zoncord-id/validation"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHasher(password string) string {
	// Password hasher function
	zap.L().Error("starting hashing password")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	zap.L().Error("password successfully hashed")
	return string(passwordHash)
}

func SignUp(email, password1, password2, firstName, lastName string) (string, error) {
	zap.L().Error("starting sign up")
	err := validation.PasswordsValidation(password1, password2)
	if err != nil {
		zap.L().Error(err.Error())
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
		zap.L().Error(err.Error())
		return "", err
	}
	token, err := models.CreateAccessToken(user, 1, "read write")
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}
	zap.L().Error("successfully signed up")
	return token.Token, nil
}
