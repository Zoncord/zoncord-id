package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
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

func GetUserIDFromCode(applicationID uint, code string, redirectUri string) (uint, error) {
	// check if code is valid
	var grant Grant
	err := db.Where("application_id = ? AND code = ?", applicationID, code).First(&grant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrInvalidGrant
	}
	if err != nil {
		zap.L().Error("Error while checking code", zap.Error(err))
		return 0, ErrInternalServerError
	}

	// check if redirect_uri is valid
	if !strings.Contains(grant.RedirectUri, redirectUri) {
		zap.L().Error("Error while checking redirect_uri", zap.Error(err), zap.Any("grant", grant), zap.String("redirect_uri", redirectUri))
		return 0, ErrInvalidRedirectUri
	}

	return grant.UserID, nil
}

func CreateGrant(userID uint, applicationID uint, redirectUri string, scope string) (string, error) {
	// create grant
	grant := Grant{
		UserID:        userID,
		Code:          RandStringBytesMaskImprSrcUnsafe(32),
		ApplicationID: applicationID,
		Expires:       time.Now().Add(10 * time.Minute),
		RedirectUri:   redirectUri,
		Scope:         scope,
	}
	err := db.Create(&grant).Error
	if err != nil {
		zap.L().Error("Error while creating grant", zap.Error(err))
		return "", ErrInternalServerError
	}

	return grant.Code, nil
}

func DeleteGrant(code string) error {
	err := db.Where("code = ?", code).Delete(&Grant{}).Error
	if err != nil {
		zap.L().Error("Error while deleting grant", zap.Error(err))
		return ErrInternalServerError
	}

	return nil
}
