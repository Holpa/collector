package mongo

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	HOPPERS_COLLECTION = "hoppers"
)

type (
	HoppersCollection struct {
		Connection *mongo.Client
	}
)

func (col *HoppersCollection) GetCollection() *mongo.Collection {
	return GetCollection(col.Connection, HOPPERS_COLLECTION)
}

func (col *HoppersCollection) InsertMany(hoppers []models.HopperDocument) error {
	collection := col.GetCollection()

	data := make([]interface{}, len(hoppers))
	for i, hopper := range hoppers {
		data[i] = hopper
	}

	_, err := collection.InsertMany(
		context.Background(),
		data,
	)
	return err
}

func (col *HoppersCollection) Clear() error {
	collection := col.GetCollection()

	_, err := collection.DeleteMany(
		context.Background(),
		bson.D{},
	)
	return err
}
