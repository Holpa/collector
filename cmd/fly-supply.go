package cmd

import (
	"log"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/models"
	"github.com/steschwa/hopper-analytics-collector/mongo"
	"github.com/steschwa/hopper-analytics-collector/utils"
)

func RegisterFlySupplyCmd(root *cobra.Command) {
	root.AddCommand(flySupplyCommand)
}

const STEP_SECONDS = 60 * 60 * 4

type supplyIntermediate struct {
	Supply         float64
	Burned         float64
	StakeDeposited float64
	StakeWithdrawn float64
}
type graphWithMutation func(client *graph.TransfersGraphClient, from time.Time, to time.Time, intermediate *supplyIntermediate)

var flySupplyCommand = &cobra.Command{
	Use:   "fly-supply",
	Short: "Load and save FLY supply",
	Run: func(cmd *cobra.Command, args []string) {
		dbClient := GetMongo()
		defer dbClient.Disconnect()

		suppliesCol := &mongo.FlySuppliesCollection{
			Client: dbClient,
		}

		latestSupplies, err := suppliesCol.FindLatest()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		from := time.Unix(constants.HOPPERS_FLY_TS, 0)
		latestSupply := models.FlySupplyDocument{}
		if len(latestSupplies) >= 1 {
			doc := latestSupplies[0]
			latestSupply = doc
			from = doc.Timestamp
		}

		diffSeconds := time.Since(from).Seconds()
		iterations := int(diffSeconds / STEP_SECONDS)

		log.Printf("Querying %d ranges\n", iterations)

		graphClient := graph.NewTransfersGraphClient()

		operations := []graphWithMutation{
			mintsOperation,
			burnsOperation,
			stakeDepositOperation,
			stakeWithdrawOperation,
		}

		for i := 0; i < iterations; i++ {
			intermediate := &supplyIntermediate{}
			var wg sync.WaitGroup

			start := from.Add(time.Second * STEP_SECONDS * time.Duration(i))
			end := start.Add(time.Second * STEP_SECONDS)

			log.Printf("%d) %s - %s\n", i, start.Format(time.RFC3339), end.Format(time.RFC3339))

			for _, operation := range operations {
				wg.Add(1)
				go func(operation graphWithMutation, from time.Time, to time.Time) {
					defer wg.Done()
					operation(graphClient, from, to, intermediate)
				}(operation, start, end)
			}

			wg.Wait()

			available := math.Max(0, intermediate.Supply-intermediate.Burned)
			staked := math.Max(0, intermediate.StakeDeposited-intermediate.StakeWithdrawn)
			free := math.Max(0, available-staked)

			latestSupply.Timestamp = end
			latestSupply.Supply += intermediate.Supply
			latestSupply.Burned += intermediate.Burned
			latestSupply.Available += available
			latestSupply.Staked += staked
			latestSupply.Free += free

			err = suppliesCol.Insert(latestSupply)
			if err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}
		}
	},
}

func mintsOperation(client *graph.TransfersGraphClient, from time.Time, to time.Time, intermediate *supplyIntermediate) {
	transfers, err := client.FetchMints(from, to)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
		return
	}

	intermediate.Supply = sumTransfers(transfers)
}
func burnsOperation(client *graph.TransfersGraphClient, from time.Time, to time.Time, intermediate *supplyIntermediate) {
	transfers, err := client.FetchBurns(from, to)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
		return
	}

	intermediate.Burned = sumTransfers(transfers)
}
func stakeDepositOperation(client *graph.TransfersGraphClient, from time.Time, to time.Time, intermediate *supplyIntermediate) {
	transfers, err := client.FetchDeposits(from, to)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
		return
	}

	intermediate.StakeDeposited = sumTransfers(transfers)
}
func stakeWithdrawOperation(client *graph.TransfersGraphClient, from time.Time, to time.Time, intermediate *supplyIntermediate) {
	transfers, err := client.FetchWithdraws(from, to)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
		return
	}

	intermediate.StakeWithdrawn = sumTransfers(transfers)
}

func sumTransfers(transfers []graph.Transfer) float64 {
	total := big.NewInt(0)

	for _, transfer := range transfers {
		total = big.NewInt(0).Add(total, transfer.Amount)
	}

	totalDec := utils.ToDecimal(total, 18)
	totalF, _ := totalDec.Float64()

	return totalF
}
