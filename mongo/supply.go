package mongo

import (
	"context"
	"time"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SUPPLY_COLLECTION = "supplies"
)

type (
	SuppliesCollection struct {
		Client *MongoDbClient
	}
)

func (col *SuppliesCollection) GetCollection() *mongo.Collection {
	return col.Client.Database.Collection(SUPPLY_COLLECTION)
}

func (col *SuppliesCollection) Insert(supply models.SupplyDocument) error {
	collection := col.GetCollection()

	supply.Timestamp = time.Now()

	_, err := collection.InsertOne(
		context.Background(),
		supply,
	)
	return err
}

func (col *SuppliesCollection) InsertMany(supplies []models.SupplyDocument) error {
	collection := col.GetCollection()

	data := make([]interface{}, len(supplies))
	for i, supply := range supplies {
		data[i] = supply
	}

	_, err := collection.InsertMany(
		context.Background(),
		data,
	)
	return err
}
