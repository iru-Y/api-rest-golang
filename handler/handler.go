package handler

import (
	"github.com/iru-Y/api-rest-golang/infra"
)

var (
	logger *infra.Logger
)

func InitializeHandler(){
	logger = infra.NewLogger("handler")
}
