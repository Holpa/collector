package cmd

import (
	"log"
	"math/big"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"github.com/steschwa/hopper-analytics-collector/utils"
)

func RegisterFlySupplyCmd(root *cobra.Command) {
	root.AddCommand(flySupplyCommand)
}

var flySupplyCommand = &cobra.Command{
	Use:   "fly-supply",
	Short: "Load and save current FLY supply",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		onChainClient := GetOnChainClient()

		flySupply, err := onChainClient.GetFlySupply()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
		flySupplyF, _ := utils.ToDecimal(flySupply, 18).Float64()

		flyBurned, err := onChainClient.GetFlyBurned()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
		flyBurnedF, _ := utils.ToDecimal(flyBurned, 18).Float64()

		flyAvailable := big.NewInt(0).Sub(flySupply, flyBurned)
		flyAvailableF, _ := utils.ToDecimal(flyAvailable, 18).Float64()

		supplyDocument := models.SupplyDocument{
			Type:      models.FLY_SUPPLY,
			Supply:    flySupplyF,
			Burned:    flyBurnedF,
			Available: flyAvailableF,
		}

		collection := &db.SuppliesCollection{
			Connection: mongoClient,
		}
		err = collection.Insert(supplyDocument)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
	},
}
