package main

import (
	"github.com/Zoncord/zoncord-id/routers"
	"github.com/Zoncord/zoncord-id/units/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func initGin() {
	runningMode := os.Getenv("RUNNING_MODE")
	switch runningMode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}
func main() {
	zap.L().Info("work started")
	initGin()
	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}
	r := routers.InitRouters()

	err = r.Run()
	for err != nil {
		err = r.Run()
	}
}
