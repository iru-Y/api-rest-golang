package routers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()

	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("127.0.0.1:" + port)

}
