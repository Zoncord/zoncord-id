package deserialization

type AccessTokenBody struct {
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
	GrantType    string `json:"grant_type" binding:"required,oneof=authorization_code refresh_token"`
	RefreshToken string `json:"refresh_token" binding:"required_if=GrantType refresh_token"`
	Code         string `json:"code" binding:"required_if=GrantType authorization_code"`
	RedirectUri  string `json:"redirect_uri" binding:"required"`
}

type GrantBody struct {
	Token        string `header:"Authorization" binding:"required"`
	ClientID     string `json:"client_id" binding:"required"`
	ResponseType string `json:"response_type" binding:"required,oneof=code"`
	RedirectUri  string `json:"redirect_uri" binding:"required"`
}
