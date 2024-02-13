package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/schemas"
)

var (
	logger = infra.NewLogger("Entrando no User Handler")
)

func GetAllUser(ctx *gin.Context) {

}

func PostUser(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Bad request for postUser")
	}

	user := schemas.User{
		ROLE:     request.ROLE,
		NAME:     request.NAME,
		PASSWORD: request.PASSWORD,
		EMAIL:    request.EMAIL,
	}

	if err := infra.Insert(&user).Error; err != nil {
		logger.Errorf("error on create user: %v", err)
		responseError(ctx, http.StatusInternalServerError, "error creating user on database")
		return
	}
}
