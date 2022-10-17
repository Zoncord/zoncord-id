package logging

import (
	"go.uber.org/zap"
	"os"
)

func InitLogger() error {
	runningMode := os.Getenv("RUNNING_MODE")
	var logger *zap.Logger
	var err error
	switch runningMode {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	return nil
}
