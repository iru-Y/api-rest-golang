package router

import (
	"github.com/gin-gonic/gin"
)
func Initialize() {
	
	v1 := gin.Default()
	initializeRoutes(v1)
	v1.Run(":8080")

}
