package oauth2

import (
	"gorm.io/gorm"
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
