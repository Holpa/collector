package mongo

import (
	"context"
	"time"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	HOPPERS_ACTIVITY_COLLECTION = "hoppers-activity"
)

type (
	HoppersActivityCollection struct {
		Client *MongoDbClient
	}
)

func (col *HoppersActivityCollection) GetCollection() *mongo.Collection {
	return col.Client.Database.Collection(HOPPERS_ACTIVITY_COLLECTION)
}

func (col *HoppersActivityCollection) Insert(doc models.HoppersActivityDocument) error {
	collection := col.GetCollection()

	doc.Timestamp = time.Now()

	_, err := collection.InsertOne(context.Background(), doc)

	return err
}
