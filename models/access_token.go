package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"os"
	"time"
)

type AccessToken struct {
	gorm.Model
	UserID             uint         `json:"user_id"`
	Token              string       `json:"token"`
	ApplicationID      uint         `json:"application_id"`
	Expires            time.Time    `json:"expires"`
	SourceRefreshToken RefreshToken `json:"source_refresh_token"`
	Scope              string       `json:"scope"`
}

func CreateAccessToken(userID uint, applicationID uint, scope string) (AccessToken, error) {
	// TODO проверка на валидность параметров

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"expires": time.Now().Unix() + 86400,
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString(hmacSampleSecret)
	accessToken := AccessToken{
		UserID:        userID,
		Token:         tokenString,
		ApplicationID: applicationID,
		Expires:       time.Now().Add(time.Hour * 24),
		Scope:         scope,
	}
	err = db.Create(&accessToken).Error
	accessToken.SourceRefreshToken, err = CreateRefreshToken(userID, applicationID, accessToken.ID)
	err = db.Save(&accessToken).Error
	return accessToken, err
}
