package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	MongoDbClient struct {
		Client   *mongo.Client
		Database *mongo.Database
		Name     string
	}
)

func NewMongoDbClient(databaseName string) *MongoDbClient {
	return &MongoDbClient{
		Name: databaseName,
	}
}

func (dbClient *MongoDbClient) Connect(uri string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	dbClient.Client = client
	dbClient.Database = dbClient.Client.Database(dbClient.Name)

	return nil
}

func (dbClient *MongoDbClient) Disconnect() error {
	if dbClient.Client == nil {
		return fmt.Errorf("not connected")
	}

	return dbClient.Client.Disconnect(context.Background())
}
