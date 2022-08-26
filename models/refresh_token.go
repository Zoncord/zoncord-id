package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

type RefreshToken struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	ApplicationID uint      `json:"application_id"`
	Token         string    `json:"token"`
	AccessTokenID uint      `json:"access_token"`
	Revoked       time.Time `json:"revoked"`
}

func CreateRefreshToken(userID uint, applicationID uint, accessTokenID uint) (RefreshToken, error) {
	// TODO проверка на валидность параметров

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         userID,
		"access_token_id": accessTokenID,
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString(hmacSampleSecret)

	refreshToken := RefreshToken{
		UserID:        userID,
		ApplicationID: applicationID,
		Token:         tokenString,
		AccessTokenID: accessTokenID,
	}
	err = db.Create(&refreshToken).Error
	return refreshToken, err
}
