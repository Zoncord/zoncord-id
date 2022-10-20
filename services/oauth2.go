package services

import (
	"github.com/Zoncord/zoncord-id/deserialization"
	"github.com/Zoncord/zoncord-id/models"
)

func GetAccessToken(accessTokenBody *deserialization.AccessTokenBody) (string, error) {
	// check if application is valid
	applicationID, err := models.GetApplicationIDByCredentials(accessTokenBody.ClientID, accessTokenBody.ClientSecret)
	if err != nil {
		return "", err
	}
	// depending on the type of grant, check the required parameter and check redirectUri
	if accessTokenBody.GrantType == "authorization_code" {
		userID, err := models.GetUserIDFromCode(applicationID, accessTokenBody.Code, accessTokenBody.RedirectUri)
		if err != nil {
			return "", err
		}
		refreshToken, err := models.CreateRefreshToken(userID, applicationID)
		accessToken, err := models.CreateAccessTokenInDB(userID, applicationID, "read write openid", refreshToken)
		if err != nil {
			return "", err
		}
		return accessToken, nil
	}
	if accessTokenBody.GrantType == "refresh_token" {
		// getting the token object and checking its validiry
		refreshToken, err := models.GetRefreshToken(accessTokenBody.RefreshToken)
		if err != nil {
			return "", err
		}
		userID := refreshToken.UserID
		accessToken, err := models.CreateAccessTokenInDB(userID, applicationID, "read write openid", refreshToken)
		if err != nil {
			return "", err
		}
		return accessToken, nil
	}
	return "", models.ErrInvalidGrant
}
