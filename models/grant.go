package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Grant struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	Code          string    `json:"code"`
	ApplicationID uint      `json:"application_id"`
	Expires       time.Time `json:"expires"`
	RedirectUri   string    `json:"redirect_uri"`
	Scope         string    `json:"scope"`
}

func CheckCode(code string, clientID string, clientSecret string, redirectURI string) error {
	// TODO: проверка кода на валидность
	return nil
}

func CheckRefreshToken(refreshToken string, clientID string, clientSecret string, redirectURI string) error {
	// TODO: проверка кода на валидность
	return nil
}
