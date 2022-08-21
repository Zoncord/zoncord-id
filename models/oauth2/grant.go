package oauth2

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
