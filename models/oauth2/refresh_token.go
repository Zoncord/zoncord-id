package oauth2

import "github.com/jinzhu/gorm"

type RefreshToken struct {
	gorm.Model
	UserID      string      `json:"user_id"`
	Application Application `json:"application"`
}
