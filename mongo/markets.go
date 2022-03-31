package mongo

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (col *MarketsCollection) Upsert(listing models.ListingDocument) error {
	collection := col.GetCollection()

	upsert := true
	_, err := collection.ReplaceOne(
		context.Background(),
		bson.D{{Key: "id", Value: listing.Id}},
		listing,
		&options.ReplaceOptions{Upsert: &upsert},
	)

	return err
}
