package handler

import (
	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/ps"
	
)

var (
	err *error
)

func InitializeRoutes() {
	logger = infra.NewLogger("handler")
	
	db = ps.ConnectDb() 
	if db != nil {
		logger.Error("Error ao inicializar as rotas", err)
	}
}
