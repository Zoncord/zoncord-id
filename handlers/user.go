package handlers

import (
	"ZoncordID/services"
	"fmt"
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
	c.JSON(http.StatusOK, gin.H{
		"email":    email,
		"password": password,
	})
}

func DefaultSignUp(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"detail": "Method not allowed",
	})
}

func PostSignUp(c *gin.Context) {
	//
	email := c.PostForm("email")
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

	c.JSON(http.StatusOK, gin.H{
		"email":    email,
		"password": password1,
	})
}
