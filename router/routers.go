package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/handler"
)

func initializeRoutes (router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET(
			"/test", handler.GetAll,
		)}
}
