package ps

import (
	"fmt"

	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	ps *gorm.DB
	err error
)

func ConnectDb() (*gorm.DB, error) {
    config := infra.NewConfig()
    logger := infra.NewLogger("db")
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

    ps, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        logger.Errorf("Error on connect to Database")
        return nil, err
    }
    logger.Infof("Connected Successfully to the Database")
    return ps, nil
}

func Insert(user *schemas.User) error {
    result := ps.Create(user)

    if result.Error != nil {
        return result.Error
    }
    return nil
}
