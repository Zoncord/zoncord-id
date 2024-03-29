package handlers

import (
	"net/http"

	"github.com/Zoncord/zoncord-id/pkg/models"
	"github.com/Zoncord/zoncord-id/pkg/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostSignIn(c *gin.Context) {
	// TODO remove business logic
	zap.L().Info("Starting sign in")
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, err := models.CheckAuth(email, password)
	if err == models.DatabaseNotAvailable {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Something happened to us, we are already working on it",
		})
		zap.L().Error(err.Error())
		return
	}
	if err == models.InvalidEmailOrPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "Invalid credentials",
		})
		zap.L().Error(err.Error())
		return
	}
	token, err := models.CreateAccessToken(user.ID, 1, "read write")
	if err != nil {
		zap.L().Error(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Login",
		"token":   token.Token,
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
