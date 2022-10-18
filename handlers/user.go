package handlers

import (
	"github.com/Zoncord/zoncord-id/models"
	"github.com/Zoncord/zoncord-id/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func PostSignIn(c *gin.Context) {
	zap.L().Info("Starting sign in")
	email := c.PostForm("email")
	password := c.PostForm("password")
	accessToken, err := services.SignIn(email, password)
	if err != nil {
		if err == models.DatabaseNotAvailable {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "Something happened to us, we are already working on it",
			})
			zap.L().Error(err.Error())
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": err.Error(),
		})
		zap.L().Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Login",
		"token":   accessToken,
	})
	zap.L().Info("user successfully signed in")
}

func PostSignUp(c *gin.Context) {
	zap.L().Info("starting signing up user")
	token, err := services.SignUp(
		c.PostForm("email"),
		c.PostForm("password1"),
		c.PostForm("password2"),
		c.PostForm("first_name"),
		c.PostForm("last_name"),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		zap.L().Info("Validation error: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Registration",
		"token":   token,
	})
	zap.L().Info("user successfully signed up")
}

func GetCurrentUserData(c *gin.Context) {
	token := c.GetHeader("Authorization")
	user, err := models.GetUserByToken(token)
	if err == models.DatabaseNotAvailable {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Something happened to us, we are already working on it",
		})
		return
	}
	if err == models.InvalidToken {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "Invalid token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
	zap.L().Info("got user data")
}
