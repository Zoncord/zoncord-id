package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Application struct {
	gorm.Model
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	UserID uint   `json:"user_id"`
	Name   string `json:"name"`

	RedirectUris           string `json:"redirect_uris"`
	ClientType             string `json:"client_type"`
	AuthorizationGrantType string `json:"authorization_grant_type"`
	Algorithm              string `json:"algorithm"`

	SkipAuthorization bool `json:"skip_authorization"`

	AccessTokens  []AccessToken  `json:"access_tokens"`
	RefreshTokens []RefreshToken `json:"refresh_tokens"`
	Grants        []Grant        `json:"grants"`
}

func GetApplicationIDByCredentials(clientID string, clientSecret string) (uint, error) {
	// check if application exists
	var application Application
	err := db.Where("client_id = ? AND client_secret = ?", clientID, clientSecret).First(&application).Error
	if errors.Is(err, logger.ErrRecordNotFound) {
		return 0, ErrInvalidCredentials
	}
	if err != nil {
		zap.L().Error("Error while getting application by credentials", zap.Error(err))
		return 0, ErrInternalServerError
	}
	return application.ID, nil
}

func GetApplicationIDByClientID(clientID string, redirectUri string) (uint, error) {
	// check if application exists
	var application Application
	err := db.Where("client_id = ?", clientID).First(&application).Error
	if errors.Is(err, logger.ErrRecordNotFound) {
		return 0, ErrInvalidCredentials
	}
	if err != nil {
		zap.L().Error("Error while getting application by clientID", zap.Error(err))
		return 0, ErrInternalServerError
	}
	// check if redirectUri is valid
	if !strings.Contains(application.RedirectUris, redirectUri) {
		return 0, ErrInvalidRedirectUri
	}
	return application.ID, nil
}
