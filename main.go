package main

import (
	"github.com/Zoncord/zoncord-id/routers"
	"github.com/Zoncord/zoncord-id/units/logger"
	"go.uber.org/zap"
)

func main() {
	zap.L().Info("work started")
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
