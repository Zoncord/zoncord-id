package models

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type RefreshToken struct {
	gorm.Model
	UserID        uint
	ApplicationID uint
	Token         string
	AccessTokenID uint
	Revoked       time.Time
}

func CreateRefreshToken(userID uint, applicationID uint) (RefreshToken, error) {
	// create refresh token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":        userID,
		"application_id": applicationID,
		"created_at":     time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		zap.L().Error("Error while signing refresh token", zap.Error(err))
		return RefreshToken{}, ErrInternalServerError
	}

	refreshToken := RefreshToken{
		UserID:        userID,
		ApplicationID: applicationID,
		Token:         tokenString,
	}

	return refreshToken, nil
}

func GetRefreshToken(applicationID uint, refreshToken string) (RefreshToken, error) {
	var token RefreshToken
	err := db.First("application_id = ? AND token = ?", applicationID, refreshToken).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return RefreshToken{}, ErrInvalidGrant
	}
	if err != nil {
		zap.L().Error("Error while getting refresh token", zap.Error(err))
		return RefreshToken{}, ErrInternalServerError
	}

	return token, nil
}
