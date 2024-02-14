package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/handler"
	"gorm.io/gorm"
)

func InitializeRouter(db *gorm.DB) {
	r := gin.Default()
	handler.InitializeRoutes()
	initializeRoutes(r)
	r.Run(":8080")
}
