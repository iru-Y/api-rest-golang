package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iru-Y/api-rest-golang/schemas"
	"github.com/spf13/viper"
	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *driver.Client
	collection *driver.Collection
)


func InitializeMongo() error {
	viper.SetConfigFile("viper-test.yml")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo de configuração: %v", err)
	}

	var mongoDBConfig struct {
		URI string `mapstructure:"uri"`
	}

	err = viper.UnmarshalKey("mongodb", &mongoDBConfig)
	if err != nil {
		return fmt.Errorf("erro ao unmarshal as configurações do MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err = driver.Connect(ctx, options.Client().ApplyURI(mongoDBConfig.URI))
	if err != nil {
		return fmt.Errorf("erro ao conectar ao MongoDB: %v", err)
	}

	db := client.Database("teste")
	collection = db.Collection("teste")

	return nil
}

func CloseConnection() {
	if client != nil {
		logger.Errorf("lascouse")
		client.Disconnect(context.Background())
	}
}

func Insert(user *schemas.User) error {
	user.ID = uuid.New().String()
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
