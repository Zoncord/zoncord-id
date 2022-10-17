package handlers

import (
	"github.com/Zoncord/zoncord-id/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func PostGrant(c *gin.Context) {

}

func PostAccessToken(c *gin.Context) {
	zap.L().Error("starting PostAccessToken")
	clientID, err := strconv.ParseUint(c.PostForm("client_id"), 10, 32)
	clientID32 := uint(clientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid client_id",
		})
	}
	clientSecret := c.PostForm("client_secret")
	redirectURI := c.PostForm("redirect_uri")
	grantType := c.PostForm("grant_type")
	if (grantType != "authorization_code") && (grantType != "refresh_token") {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid grant type",
		})
		return
	}

	if grantType == "authorization_code" {
		code := c.PostForm("code")
		err := models.CheckCode(code, clientID32, clientSecret, redirectURI)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "Invalid code",
			})
			return
		}
	}

	if grantType == "refresh_token" {
		refreshToken := c.PostForm("refresh_token")
		err := models.CheckRefreshToken(refreshToken, clientID32, clientSecret, redirectURI)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "Invalid refresh token",
			})
			return
		}
	}
	zap.L().Error("end PostAccessToken")
	//return models.CreateAccessToken()
	//	TODO: return access token
}
