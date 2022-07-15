package main

import "ZoncordID/routers"

func main() {
	r := routers.InitRouters()
	err := r.Run()
	for err != nil {
		err = r.Run()
	}
}
