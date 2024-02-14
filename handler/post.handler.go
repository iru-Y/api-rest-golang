package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iru-Y/api-rest-golang/schemas"
	"github.com/iru-Y/api-rest-golang/ps"
)
	


func PostUser(ctx *gin.Context) {
	
    request := CreateUserRequest{}
    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
    }

    user := schemas.User{
        ROLE:     request.ROLE,
        NAME:     request.NAME,
        PASSWORD: request.PASSWORD,
        EMAIL:    request.EMAIL,
    }

    res := ps.Insert(&user)

    if res.Error != nil {
    }

    responseSucess(ctx, "post user", user)
}
