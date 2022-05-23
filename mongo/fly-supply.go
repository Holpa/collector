package mongo

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	FLY_SUPPLIES_COLLECTION = "fly-supplies"
)

type (
	FlySuppliesCollection struct {
		Client *MongoDbClient
	}
)

func (col *FlySuppliesCollection) GetCollection() *mongo.Collection {
	return col.Client.Database.Collection(FLY_SUPPLIES_COLLECTION)
}

func (col *FlySuppliesCollection) Insert(supply models.FlySupplyDocument) error {
	collection := col.GetCollection()

	_, err := collection.InsertOne(
		context.Background(),
		supply,
	)
	return err
}

func (col *FlySuppliesCollection) Clear() error {
	collection := col.GetCollection()

	_, err := collection.DeleteMany(
		context.Background(),
		bson.D{},
	)
	return err
}
