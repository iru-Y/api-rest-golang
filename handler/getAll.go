package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	 ctx.JSON(
		http.StatusOK, gin.H  {
				"message" : "entrou no GetAll",
			},
		
	 )
}
