package cmd

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterHoppersActivityCmd(root *cobra.Command) {
	root.AddCommand(hoppersActivityCommand)
}

var hoppersActivityCommand = &cobra.Command{
	Use:   "hoppers-activity",
	Short: "Save a snapshot of what all Hoppers are doing",
	Run: func(cmd *cobra.Command, args []string) {
		dbClient := GetMongo()
		defer dbClient.Disconnect()

		hoppersCollection := &db.HoppersCollection{
			Client: dbClient,
		}

		cursor, err := hoppersCollection.GetCollection().Find(context.Background(), bson.D{})
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		hoppers := []models.HopperDocument{}
		if err = cursor.All(context.Background(), &hoppers); err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		doc := models.HoppersActivityDocument{}
		for _, hopper := range hoppers {
			if hopper.Activity == models.HopperActivityPond {
				doc.Pond++
			} else if hopper.Activity == models.HopperActivityStream {
				doc.Stream++
			} else if hopper.Activity == models.HopperActivitySwamp {
				doc.Swamp++
			} else if hopper.Activity == models.HopperActivityRiver {
				doc.River++
			} else if hopper.Activity == models.HopperActivityForest {
				doc.Forest++
			} else if hopper.Activity == models.HopperActivityGreatLake {
				doc.GreatLake++
			} else if hopper.Activity == models.HopperActivityBreeding {
				doc.Breeding++
			} else if hopper.Activity == models.HopperActivityMarketplace {
				doc.Marketplace++
			} else if hopper.Activity == models.HopperActivityIdle {
				doc.Idle++
			}
		}

		doc.Adventure = doc.Pond + doc.Stream + doc.Swamp + doc.River + doc.Forest + doc.GreatLake

		hoppersActivityCollection := &db.HoppersActivityCollection{
			Client: dbClient,
		}

		if err = hoppersActivityCollection.Insert(doc); err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
	},
}
