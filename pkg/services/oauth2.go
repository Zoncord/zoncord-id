package services

import (
	"github.com/Zoncord/zoncord-id/pkg/deserialization"
	"github.com/Zoncord/zoncord-id/pkg/models"
)

func GetCode(codeBody *deserialization.GrantBody) (string, error) {
	// check user credentials
	userID, err := models.GetUserIDByAccessToken(codeBody.Token)
	if err != nil {
		return "", err
	}
	// check if application and redirectUri is valid
	applicaionID, err := models.GetApplicationIDByClientID(codeBody.ClientID, codeBody.RedirectUri)
	if err != nil {
		return "", err
	}
	// create grant
	code, err := models.CreateGrant(userID, applicaionID, codeBody.RedirectUri, "read write openid")
	return code, err
}

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
		// delete grant
		err = models.DeleteGrant(accessTokenBody.Code)
		if err != nil {
			return "", err
		}
		// create access token
		refreshToken, err := models.CreateRefreshToken(userID, applicationID)
		if err != nil {
			return "", err
		}
		accessToken, err := models.CreateAccessTokenInDB(userID, applicationID, "read write openid", refreshToken)
		if err != nil {
			return "", err
		}
		return accessToken, nil
	}
	if accessTokenBody.GrantType == "refresh_token" {
		// getting the token object and checking its validiry
		refreshToken, err := models.GetRefreshToken(applicationID, accessTokenBody.RefreshToken)
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
