package handlers

import (
	"github.com/Zoncord/zoncord-id/models"
	"github.com/Zoncord/zoncord-id/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostSignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, err := models.CheckAuth(email, password)
	if err == models.DatabaseNotAvailable {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Something happened to us, we are already working on it",
		})
		return
	}
	if err == models.InvalidEmailOrPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "Invalid credentials",
		})
		return
	}
	token, err := models.CreateAccessToken(user, 0, "read write")
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Login",
		"token":   token.Token,
	})
}

func PostSignUp(c *gin.Context) {
	token, err := services.SignUp(
		c.PostForm("email"),
		c.PostForm("password1"),
		c.PostForm("password2"),
		c.PostForm("first_name"),
		c.PostForm("last_name"),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Registration",
		"token":   token,
	})
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
}
