package models

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"os"
	"strings"
	"time"
)

type RefreshToken struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	ApplicationID uint      `json:"application_id"`
	Token         string    `json:"token"`
	AccessTokenID uint      `json:"access_token"`
	Revoked       time.Time `json:"revoked"`
	RedirectUri   string    `json:"redirect_uri"`
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

func CheckRefreshToken(Token string, clientID uint, clientSecret string, redirectURI string) error {
	// TODO: проверка кода на валидность
	application, err := checkApplication(clientID, clientSecret)
	if err != nil {
		return err
	}
	var refreshToken RefreshToken
	refreshToken.Token = Token
	refreshToken.ApplicationID = application.ID
	err = db.First(&refreshToken).Error
	if err != nil {
		return err
	}
	if strings.Contains(refreshToken.RedirectUri, redirectURI) {
		return fmt.Errorf("redirect_uri is not valid")
	}

	return nil
}
