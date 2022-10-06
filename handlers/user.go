package handlers

import (
	"fmt"
	"github.com/Zoncord/zoncord-id/errors"
	"github.com/Zoncord/zoncord-id/models"
	"github.com/Zoncord/zoncord-id/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultSignIn(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"detail": "Method not allowed",
	})
}

func PostSignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	err := models.CheckAuth(email, password)
	if err == errors.DatabaseNotAvailable {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Something happened to us, we are already working on it",
		})
		return
	}
	if err == errors.InvalidEmailOrPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"detail": "Invalid credentials",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Login",
	})
}

func DefaultSignUp(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"detail": "Method not allowed",
	})
}

func PostSignUp(c *gin.Context) {
	// create user
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")
	//
	err := services.PasswordComplexityCheck(password1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": fmt.Errorf("password is too short or too long"),
		})
	}

	if password1 != password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Passwords do not match.",
		})
	}
	err = models.CreateUser(c.PostForm("email"), password1, c.PostForm("first_name"), c.PostForm("last_name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": fmt.Errorf("email already exists"),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful Registration",
	})
}
