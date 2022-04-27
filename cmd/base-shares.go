package cmd

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func RegisterBaseSharesCmd(root *cobra.Command) {
	root.AddCommand(baseSharesCommand)
}

var baseSharesCommand = &cobra.Command{
	Use:   "base-shares",
	Short: "Load and save curent base shares for adventures",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		defer mongoClient.Disconnect(context.Background())
		onChainClient := GetOnChainClient()

		adventures := []constants.Adventure{
			constants.AdventurePond,
			constants.AdventureStream,
			constants.AdventureSwamp,
			constants.AdventureRiver,
			constants.AdventureForest,
			constants.AdventureGreatLake,
		}

		collection := &db.BaseSharesCollection{
			Connection: mongoClient,
		}
		for _, adventure := range adventures {
			totalBaseShares, err := onChainClient.GetTotalBaseSharesByAdventure(adventure)
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
				continue
			}

			err = collection.Insert(models.BaseSharesDocument{
				Adventure:       adventure.String(),
				TotalBaseShares: models.NewBigInt(totalBaseShares),
			})
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
			}
		}
	},
}
