package mongo

import (
	"context"
	"time"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	VOTES_COLLECTION = "votes"
)

type (
	VotesCollection struct {
		Client *MongoDbClient
	}
)

func (col *VotesCollection) GetCollection() *mongo.Collection {
	return col.Client.Database.Collection(VOTES_COLLECTION)
}

func (col *VotesCollection) Insert(vote models.VoteDocument) error {
	collection := col.GetCollection()

	vote.Updated = time.Now()

	_, err := collection.InsertOne(
		context.Background(),
		vote,
	)

	return err
}
