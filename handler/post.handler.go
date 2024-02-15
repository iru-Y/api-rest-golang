package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/schemas"
	"gorm.io/gorm"
)

var (
	logger *infra.Logger
	db     *gorm.DB
)

func PostUser(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Erro ao validar os dados, verifique os campos")
		return
	}
	userID, err := uuid.NewRandom()
	if err != nil {
		logger.Errorf("Erro ao gerar UUID para o usuário: %v", err)
		responseError(ctx, http.StatusInternalServerError, "Erro ao criar o usuário")
		return
	}
	user := schemas.User{
		ID:       userID.String(),
		ROLE:     request.ROLE,
		NAME:     request.NAME,
		PASSWORD: request.PASSWORD,
		PASSWORDCONFIRM: request.PASSWORDCONFIRM,
		EMAIL:    request.EMAIL,
	}

	result := db.Create(&user)
	if result.Error != nil {
		logger.Error("Erro ao criar o usuário no banco de dados: %v")
		responseError(ctx, http.StatusInternalServerError, "Erro ao criar o usuário no banco de dados")
		return
	}

	responseSucess(ctx, "Accepted!", user)
}
