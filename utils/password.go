package utils

import (
	"github.com/iru-Y/api-rest-golang/infra"
	"golang.org/x/crypto/bcrypt"
)

var (
	logger *infra.Logger
)

func generateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Não foi possível gerar o hash do password")
	}
	return string(hashedPassword), nil
}
