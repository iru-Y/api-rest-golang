package ps

import (
	"fmt"

	"github.com/iru-Y/api-rest-golang/infra"
	"github.com/iru-Y/api-rest-golang/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB) {
	config := infra.NewConfig()
	logger := infra.NewLogger("db")
    
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
	if err != nil {
		logger.Errorf("Error on connect to Database")
		return nil
	}
    err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil
	}
	logger.Infof("Connected Successfully to the Database")
	return db
}
