package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/schemas"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(DB *gorm.DB) AuthHandler {
	return AuthHandler{DB}
}
func (acc *AuthHandler) SignUpUser(ctx *gin.Context) {
	var payload *schemas.User

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		logger.Error(http.StatusBadRequest, err.Error())
		responseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if payload.PASSWORD != payload.PASSWORDCONFIRM {
		responseError(ctx, http.StatusBadRequest, "Passwords do not match")
		return
	}
}
