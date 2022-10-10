package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"detail": "Method not allowed",
	})
}
