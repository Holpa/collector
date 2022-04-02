package mongo

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MARKETS_COLLECTION = "listings"
)

type (
	MarketsCollection struct {
		Connection *mongo.Client
	}
)

func (col *MarketsCollection) GetCollection() *mongo.Collection {
	return GetCollection(col.Connection, MARKETS_COLLECTION)
}

func (col *MarketsCollection) Clear() error {
	collection := col.GetCollection()

	_, err := collection.DeleteMany(
		context.Background(),
		bson.D{},
	)

	return err
}

func (col *MarketsCollection) InsertMany(listings []models.ListingDocument) error {
	collection := col.GetCollection()

	listingDocuments := make([]interface{}, len(listings))
	for i, listing := range listings {
		listingDocuments[i] = listing
	}

	_, err := collection.InsertMany(
		context.Background(),
		listingDocuments,
	)
	return err
}
