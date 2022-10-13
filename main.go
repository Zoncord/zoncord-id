package main

import (
	"github.com/Zoncord/zoncord-id/routers"
)

func main() {
	r := routers.InitRouters()
	err := r.Run()
	for err != nil {
		err = r.Run()
	}
}
