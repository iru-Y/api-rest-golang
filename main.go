package main

import (
	"github.com/iru-Y/api-rest-golang/config"
	"github.com/iru-Y/api-rest-golang/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	router.Initialize()

}
