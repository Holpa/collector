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
		Connection *mongo.Client
	}
)

func (col *SuppliesCollection) GetCollection() *mongo.Collection {
	return GetCollection(col.Connection, SUPPLY_COLLECTION)
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
