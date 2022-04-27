package cmd

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AggregatedHopperHolders struct {
		User         string `bson:"user"`
		HoppersCount uint   `bson:"count"`
	}
)

func RegisterHopperHoldersCmd(root *cobra.Command) {
	root.AddCommand(hopperHoldersCommand)
}

var hopperHoldersCommand = &cobra.Command{
	Use:   "hopper-holders",
	Short: "Save current hopper holders count",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		defer mongoClient.Disconnect(context.Background())

		hoppersCollection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		cursor, err := hoppersCollection.GetCollection().Aggregate(
			context.Background(),
			getHopperHoldersAggregationPipeline(),
		)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		aggregates := []AggregatedHopperHolders{}
		if err = cursor.All(context.Background(), &aggregates); err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		hopperHoldersCollection := &db.HopperHoldersCollection{
			Connection: mongoClient,
		}
		doc := models.HopperHoldersDocument{
			Holders: uint(len(aggregates)),
		}

		if err = hopperHoldersCollection.Insert(doc); err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
	},
}

func getHopperHoldersAggregationPipeline() mongo.Pipeline {
	return mongo.Pipeline{
		getHopperHoldersAggregationGrouping(),
		getHopperHoldersAggregationProjection(),
	}
}
func getHopperHoldersAggregationGrouping() bson.D {
	return bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id": "$owner",
			"count": bson.M{
				"$count": bson.D{},
			},
		},
	}}
}
func getHopperHoldersAggregationProjection() bson.D {
	return bson.D{{
		Key: "$project",
		Value: bson.M{
			"user":  "$_id",
			"count": 1,
		},
	}}
}
