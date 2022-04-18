package cmd

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"github.com/steschwa/hopper-analytics-collector/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterMigrateSupplyCmd(root *cobra.Command) {
	root.AddCommand(migrateSupplyCommand)
}

var migrateSupplyCommand = &cobra.Command{
	Use:   "migrate-supply",
	Short: "Migrate legacy supply to new schema",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()

		collection := &db.SuppliesCollection{
			Connection: mongoClient,
		}

		legacyFilter := bson.D{{
			Key: "supply.v",
			Value: bson.M{
				"$exists": true,
			},
		}}

		cursor, err := collection.GetCollection().Find(
			context.Background(),
			legacyFilter,
		)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		legacySupplies := []models.LegacySupplyDocument{}
		if err := cursor.All(context.Background(), &legacySupplies); err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		migrated := make([]models.SupplyDocument, len(legacySupplies))
		for i, legacy := range legacySupplies {
			supply, _ := utils.ToDecimal(legacy.Supply.Int(), 18).Float64()

			migrated[i] = models.SupplyDocument{
				Type:      legacy.Type,
				Timestamp: legacy.Timestamp,
				Supply:    supply,
			}
		}

		_, err = collection.GetCollection().DeleteMany(
			context.Background(),
			legacyFilter,
		)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		if err = collection.InsertMany(migrated); err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		log.Printf("Migrated %d / %d supply documents\n", len(migrated), len(legacySupplies))
	},
}

// db.supplies.find({"supply.v": {"$exists": true}})
