package mongo

import (
	"context"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (col *FlySuppliesCollection) InsertMany(supplies []models.FlySupplyDocument) error {
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

func (col *FlySuppliesCollection) FindLatest() ([]models.FlySupplyDocument, error) {
	collection := col.GetCollection()

	limit := int64(1)

	cursor, err := collection.Find(
		context.Background(),
		bson.D{{}},
		&options.FindOptions{
			Sort: bson.D{{
				Key:   "timestamp",
				Value: -1,
			}},
			Limit: &limit,
		},
	)
	if err != nil {
		return []models.FlySupplyDocument{}, err
	}
	defer cursor.Close(context.Background())

	docs := []models.FlySupplyDocument{}
	if err = cursor.All(context.Background(), &docs); err != nil {
		return []models.FlySupplyDocument{}, err
	}

	return docs, nil
}
