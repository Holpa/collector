package cmd

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func RegisterMarketsCmd(root *cobra.Command) {
	root.AddCommand(marketsCommand)
}

var marketsCommand = &cobra.Command{
	Use:   "markets",
	Short: "Load and save a snapshot of all market listings",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		defer mongoClient.Disconnect(context.Background())

		graph := graph.NewMarketsGraphClient()

		listings, err := graph.FetchAllListings()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		collection := &db.MarketsCollection{
			Connection: mongoClient,
		}

		err = collection.Clear()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		listingDocuments := make([]models.ListingDocument, len(listings))
		for i, listing := range listings {
			listingDocuments[i] = helpers.ListingToListingDocument(listing)
		}

		err = collection.InsertMany(listingDocuments)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
	},
}
