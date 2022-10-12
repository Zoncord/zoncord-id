package main

import (
	"github.com/Zoncord/zoncord-id/routers"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := routers.InitRouters()
	err := r.Run()
	for err != nil {
		err = r.Run()
	}
}
