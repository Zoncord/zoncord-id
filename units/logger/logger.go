package logging

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func InitLogger() error {
	runningMode := os.Getenv("RUNNING_MODE")
	var logger *zap.Logger
	var err error
	switch runningMode {
	case "dev":
		gin.SetMode(gin.ReleaseMode)
		logger, err = zap.NewDevelopment() // TODO: move to another file
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		logger, err = zap.NewProduction()
	default:
		gin.SetMode(gin.DebugMode)
		logger = zap.NewExample()
	}

	if err != nil {
		return err
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	return nil
}
