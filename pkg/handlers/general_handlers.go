package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func MethodNotAllowed(c *gin.Context) {
	zap.L().Info("method not allowed")
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"detail": "Method not allowed",
	})
}
