package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
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

func CheckCode(code string, clientID uint, clientSecret string, redirectURI string) error {
	// validation of code
	application, err := checkApplication(clientID, clientSecret)
	if err != nil {
		return err
	}
	var grant Grant
	grant.Code = code
	grant.ApplicationID = application.ID
	err = db.First(&grant).Error
	if err != nil {
		return err
	}
	if strings.Contains(grant.RedirectUri, redirectURI) {
		return fmt.Errorf("redirect_uri is not valid")
	}

	return nil
}
