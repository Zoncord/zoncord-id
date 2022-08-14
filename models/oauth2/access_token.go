package oauth2

import (
	"gorm.io/gorm"
	"time"
)

type AccessToken struct {
	gorm.Model
	UserID             string      `json:"user_id"`
	Token              string      `json:"token"`
	Application        Application `json:"application"`
	Expires            time.Time   `json:"expires"`
	SourceRefreshToken string      `json:"source_refresh_token"`
}
