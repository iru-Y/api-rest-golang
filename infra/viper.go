package infra

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	URI string `mapstructure:"uri"`
}

func ViperInitializer () error {
	viper.SetConfigFile("viper-test.yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de configuração:", err)
		return nil
	}

	var mongoDBConfig MongoDBConfig

	err = viper.UnmarshalKey("mongodb", &mongoDBConfig)
	if err != nil {
		fmt.Println("Erro ao unmarshal as configurações do MongoDB:", err)
		return nil
	}

	fmt.Println("String de Conexão do MongoDB:", mongoDBConfig.URI)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBConfig.URI))
	if err != nil {
		fmt.Println("Erro ao conectar ao MongoDB:", err)
		return nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Erro ao pingar o MongoDB:", err)
		return nil
	}

	fmt.Println("Conexão ao MongoDB estabelecida com sucesso.")
	return nil

}