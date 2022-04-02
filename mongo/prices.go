package mongo

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	PRICES_COLLECTION = "prices"
)

type (
	PricesCollection struct {
		Connection *mongo.Client
	}
)

func (col *PricesCollection) GetCollection() *mongo.Collection {
	return GetCollection(col.Connection, PRICES_COLLECTION)
}

func (col *PricesCollection) InsertMany(prices []models.PriceDocument) error {
	collection := col.GetCollection()

	priceDocuments := make([]interface{}, len(prices))
	for i, price := range prices {
		priceDocuments[i] = price
	}

	_, err := collection.InsertMany(
		context.Background(),
		priceDocuments,
	)
	return err
}
