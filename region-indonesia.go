package main

import (
	"wec-region-indonesia/router"
	"wec-region-indonesia/utils"
)

func main() {

	go utils.Generate()
	router.SetupRoutes()

}
