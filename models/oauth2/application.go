package oauth2

import "github.com/jinzhu/gorm"

type Application struct {
	gorm.Model
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	UserID string `json:"user_id"`
	Name   string `json:"name"`

	RedirectUris           string `json:"redirect_uris"`
	ClientType             string `json:"client_type"`
	AuthorizationGrantType string `json:"authorization_grant_type"`
	Algorithm              string `json:"algorithm"`

	SkipAuthorization bool `json:"skip_authorization"`

	AccessTokens []AccessToken `json:"access_tokens"`
}
