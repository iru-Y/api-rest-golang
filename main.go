package main

import (
	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/router"
)

var (
	logger *infra.Logger
)

func main() {
	logger := infra.NewLogger("main")
	err := infra.Init()
	if err != nil {
		logger.Errorf("config inicialization error: %v", err)
		return
	}
	infra.InitializeMongo()
	router.Initialize()
}
