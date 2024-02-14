package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/schemas"
	"github.com/iru-Y/api-rest-golang/infra"
	"gorm.io/gorm"
)

var (
    logger *infra.Logger
	db *gorm.DB
)

func PostUser(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Erro ao validar os dados, verifique os campos")
		return
	}

	user := schemas.User{
		ROLE:     request.ROLE,
		NAME:     request.NAME,
		PASSWORD: request.PASSWORD,
		EMAIL:    request.EMAIL,
	}

	result := db.Create(&user);
    if result.Error != nil {
		logger.Errorf("Erro ao criar o usuário no banco de dados: %v")
		responseError(ctx, http.StatusInternalServerError, "Erro ao criar o usuário no banco de dados")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": user})
}
