package cmd

import (
	"context"
	"log"
	"math/big"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/graph"
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
		defer mongoClient.Disconnect(context.Background())
		onChainClient := GetOnChainClient()

		flySupply, err := onChainClient.GetFlySupply()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
		flySupplyF, _ := utils.ToDecimal(flySupply, 18).Float64()

		// TODO Improve loading burned FLY
		// They are not neccessary sent to 0x0 address but just reduced from total supply (atleast with leveling up)
		// One can listen for the `ERC20(FLY).Transfer` event with the 0x0 address being the recipient
		// Also more investigation of how the Breeding pool handles burning FLY
		flyBurned, err := onChainClient.GetFlyBurned()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
		flyBurnedF, _ := utils.ToDecimal(flyBurned, 18).Float64()

		flyAvailable := big.NewInt(0).Sub(flySupply, flyBurned)
		flyAvailableF, _ := utils.ToDecimal(flyAvailable, 18).Float64()

		transfersGraph := graph.NewTransfersGraphClient()
		totalDeposited, err := transfersGraph.FetchTotalDeposited()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
		totalDepositedF, _ := utils.ToDecimal(totalDeposited.String(), 18).Float64()

		totalWidthdrawn, err := transfersGraph.FetchTotalWithdrawn()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
		totalWidthdrawnF, _ := utils.ToDecimal(totalWidthdrawn.String(), 18).Float64()

		currentStaked := totalDepositedF - totalWidthdrawnF

		free := flyAvailableF - currentStaked

		supplyDocument := models.SupplyDocument{
			Type:      models.FLY_SUPPLY,
			Supply:    flySupplyF,
			Burned:    flyBurnedF,
			Available: flyAvailableF,
			Staked:    currentStaked,
			Free:      free,
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
