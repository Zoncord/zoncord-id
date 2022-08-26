package handlers

import (
	"github.com/Zoncord/zoncord-id/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostGrant(c *gin.Context) {

}

func PostAccessToken(c *gin.Context) {
	clientID := c.PostForm("client_id")
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
		err := models.CheckCode(code, clientID, clientSecret, redirectURI)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "Invalid code",
			})
			return
		}
	}

	if grantType == "refresh_token" {
		refreshToken := c.PostForm("refresh_token")
		err := models.CheckRefreshToken(refreshToken, clientID, clientSecret, redirectURI)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "Invalid refresh token",
			})
			return
		}
	}
}
