package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
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
	timeToExpire := time.Now().Add(time.Hour * 24)
	// The expiration time increases since id 1 is the master application
	if applicationID == 1 {
		timeToExpire = time.Now().Add(time.Hour * 8760)
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":        userID,
		"application_id": applicationID,
		"expires":        timeToExpire.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString(hmacSampleSecret)
	accessToken := AccessToken{
		UserID:        userID,
		Token:         tokenString,
		ApplicationID: applicationID,
		Expires:       timeToExpire,
		Scope:         scope,
	}
	return accessToken, err
}

func CreateAccessTokenInDB(userID uint, applicationID uint, scope string, refreshToken RefreshToken) (string, error) {
	accessToken, err := CreateAccessToken(userID, applicationID, scope)
	if err != nil {
		return "", err
	}
	accessToken.SourceRefreshToken = refreshToken
	err = db.Create(&accessToken).Error
	if err != nil {
		return "", err
	}
	err = db.Save(&accessToken).Error
	if err != nil {
		return "", err
	}
	return accessToken.Token, nil
}
