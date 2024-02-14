package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeRoutes()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/post", handler.PostUser)
	}
}
