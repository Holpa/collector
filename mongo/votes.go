package mongo

import (
	"context"
	"time"

	"github.com/steschwa/hopper-analytics-collector/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	VOTES_COLLECTION = "votes"
)

type (
	VotesCollection struct {
		Connection *mongo.Client
	}
)

func (col *VotesCollection) GetCollection() *mongo.Collection {
	return GetCollection(col.Connection, VOTES_COLLECTION)
}

func (col *VotesCollection) Upsert(vote models.VoteDocument) error {
	collection := col.GetCollection()

	vote.Updated = time.Now()

	upsert := true
	_, err := collection.ReplaceOne(
		context.Background(),
		bson.D{{Key: "adventure", Value: vote.Adventure}},
		vote,
		&options.ReplaceOptions{
			Upsert: &upsert,
		},
	)

	return err
}
