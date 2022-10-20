package models

import (
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
	err := db.First("application_id = ? AND code = ?", applicationID, code).Error
	if err == gorm.ErrRecordNotFound {
		return 0, ErrInvalidGrant
	}
	if err != nil {
		zap.L().Error("Error while checking code", zap.Error(err))
		return 0, ErrInternalServerError
	}

	// check if redirect_uri is valid
	if strings.Contains(grant.RedirectUri, redirectUri) {
		return 0, ErrInvalidRedirectUri
	}

	return grant.UserID, nil
}
