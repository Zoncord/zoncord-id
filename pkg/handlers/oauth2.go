package handlers

import (
	"net/http"

	"github.com/Zoncord/zoncord-id/pkg/deserialization"
	"github.com/Zoncord/zoncord-id/pkg/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostAuthorize(c *gin.Context) {
	// data validation and deserialization
	authorizationData := deserialization.GrantBody{}
	err := c.ShouldBindJSON(&authorizationData)
	if err != nil {
		// detailed output of validation errors
		errs := deserialization.GetDetailedErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		zap.L().Warn("validation failed", zap.Any("authorizationData", authorizationData), zap.Any("errors", errs))
		return
	}
	zap.L().Info("validation successful", zap.Any("authorizationData", authorizationData))

	// business logic execution
	code, err := services.GetCode(&authorizationData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"detail": "grant successfully created",
		"code":   code,
	})
}

func PostAccessToken(c *gin.Context) {
	// data validation and deserialization
	authorizationData := deserialization.AccessTokenBody{}
	err := c.ShouldBindJSON(&authorizationData)
	if err != nil {
		// detailed output of validation errors
		errs := deserialization.GetDetailedErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		zap.L().Warn("validation failed", zap.Any("authorizationData", authorizationData), zap.Any("errors", errs))
		return
	}
	zap.L().Info("validation successful", zap.Any("authorizationData", authorizationData))

	// business logic execution
	accessToken, err := services.GetAccessToken(&authorizationData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"detail":       "token successfully granted",
		"access_token": accessToken,
	})
}
