package mongo

import (
	"context"
	"time"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	HOPPER_HOLDERS_COLLECTION = "hopper-holders"
)

type (
	HopperHoldersCollection struct {
		Connection *mongo.Client
	}
)

func (col *HopperHoldersCollection) GetCollection() *mongo.Collection {
	return GetCollection(col.Connection, HOPPER_HOLDERS_COLLECTION)
}

func (col *HopperHoldersCollection) Insert(hopperHolders models.HopperHoldersDocument) error {
	collection := col.GetCollection()

	hopperHolders.Timestamp = time.Now()

	_, err := collection.InsertOne(context.Background(), hopperHolders)

	return err
}
