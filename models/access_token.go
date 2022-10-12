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

func CreateAccessToken(user User, applicationID uint, scope string) (AccessToken, error) {
	// TODO проверка на валидность параметров

	timeToExpire := time.Now().Add(time.Hour * 24)
	// The expiration time increases since id 0 is the master application
	if applicationID == 0 {
		timeToExpire = time.Now().Add(time.Hour * 8760)
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"expires": timeToExpire.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString(hmacSampleSecret)
	accessToken := AccessToken{
		UserID:        user.ID,
		Token:         tokenString,
		ApplicationID: applicationID,
		Expires:       timeToExpire,
		Scope:         scope,
	}
	err = db.Create(&accessToken).Error
	accessToken.SourceRefreshToken, err = CreateRefreshToken(user.ID, applicationID, accessToken.ID)
	err = db.Save(&accessToken).Error
	return accessToken, err
}
