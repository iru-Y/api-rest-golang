package infra

import (
	"context"
	"time"

	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *driver.Client
	collection *driver.Collection
)


func InitializeMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := driver.Connect(ctx, options.Client().ApplyURI(config.ConnectionString))
	if err != nil {
		return err
	}
	client.Database("teste")
	if err != nil {
		return err
	}

	return nil
}

func CloseConnection() {
	if client != nil {
		logger.Errorf("lascouse")
		client.Disconnect(context.Background())
	}
}

func Insert(document interface{}) error {
	_, err := collection.InsertOne(context.TODO(), document)
	return err
}
