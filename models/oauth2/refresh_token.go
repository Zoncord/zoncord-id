package oauth2

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RefreshToken struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	ApplicationID uint      `json:"application_id"`
	Token         string    `json:"token"`
	AccessTokenID uint      `json:"access_token"`
	Revoked       time.Time `json:"revoked"`
}
