package cmd

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func RegisterVotesCmd(root *cobra.Command) {
	root.AddCommand(votesCommand)
}

var votesCommand = &cobra.Command{
	Use:   "votes",
	Short: "Load and save curent votes / veShare for adventures",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		onChainClient := GetOnChainClient()

		adventures := []constants.Adventure{
			constants.AdventurePond,
			constants.AdventureStream,
			constants.AdventureSwamp,
			constants.AdventureRiver,
			constants.AdventureForest,
			constants.AdventureGreatLake,
		}

		collection := &db.VotesCollection{
			Connection: mongoClient,
		}

		for _, adventure := range adventures {
			votes, err := onChainClient.GetTotalVotesByAdventure(adventure)
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
				continue
			}
			veShare, err := onChainClient.GetTotalVeShareByAdventure(adventure)
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
				continue
			}

			voteDocument := models.VoteDocument{
				Adventure: adventure.String(),
				Votes:     models.NewBigInt(votes),
				VeShare:   models.NewBigInt(veShare),
			}

			err = collection.Insert(voteDocument)
			if err != nil {
				log.Println(err)
				sentry.CaptureException(err)
			}
		}
	},
}
