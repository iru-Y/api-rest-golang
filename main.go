package main

import (
	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/router"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func main() {
	var err error
	logger := infra.NewLogger("main")
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	logger.Debugf("Database connect success")

	router.InitializeRouter(db)
}