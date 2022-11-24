package main

import (
	"api-region/router"
	"api-region/utils"
)

func main() {

	go utils.Generate()
	router.SetupRoutes()

}
