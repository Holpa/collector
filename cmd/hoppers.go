package cmd

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func RegisterHoppersCmd(root *cobra.Command) {
	root.AddCommand(hoppersCommand)
}

var hoppersCommand = &cobra.Command{
	Use:   "hoppers",
	Short: "Load and save a snapshot of all hoppers",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		onChainClient := GetOnChainClient()

		graph := graph.NewHoppersGraphClient()

		hoppers, err := graph.FetchAllHoppers()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		rewardsCalculator := helpers.NewRewardsCalculator(onChainClient)

		collection := &db.HoppersCollection{
			Connection: mongoClient,
		}

		err = collection.Clear()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		hopperDocuments := make([]models.HopperDocument, len(hoppers))
		for i, hopper := range hoppers {
			hopperDocuments[i] = helpers.HopperToHopperDocument(hopper, rewardsCalculator)
		}

		err = collection.InsertMany(hopperDocuments)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
	},
}
