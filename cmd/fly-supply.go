package cmd

import (
	"log"
	"math"
	"math/big"
	"strings"
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

const (
	STEP_SECONDS         = 3600 * 4 // 4 hours
	YEAR_SECONDS         = 3600 * 24 * 365
	ROCKET_JOE_FLY       = 2_500_000
	ROCKET_JOE_RELEASE   = 1647104400
	LP_FLY               = 20_000_000
	LP_FLY_LOCKED_AMOUNT = 14_000_000
	TEAM_FLY             = 10_000_000
	MARKETING_FLY        = 2_500_000
	APA_FLY              = 1_000_000
)

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
		iterations := int(math.Max(1, diffSeconds/STEP_SECONDS))

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

			log.Printf("%d) %s - %s\n", i, start, end)

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
	filteredTransfers := filterTransfersByBlacklists(transfers)

	intermediate.Supply = sumTransfers(filteredTransfers) +
		getRocketJoeIncrement(from, to) +
		getLpSupplyIncrement(to) +
		getTeamSupplyIncrement(to) +
		getMarketingSupplyIncrement() +
		getAPASupplyIncrement()
}
func burnsOperation(client *graph.TransfersGraphClient, from time.Time, to time.Time, intermediate *supplyIntermediate) {
	transfers, err := client.FetchBurns(from, to)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
		return
	}
	filteredTransfers := filterTransfersByBlacklists(transfers)

	intermediate.Burned = sumTransfers(filteredTransfers)
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

func getRocketJoeIncrement(rangeFrom time.Time, rangeTo time.Time) float64 {
	rocketJoeLaunch := time.Unix(ROCKET_JOE_RELEASE, 0)

	if rangeFrom.Before(rocketJoeLaunch) && rangeTo.After(rocketJoeLaunch) {
		return ROCKET_JOE_FLY
	}

	return 0
}

func getLpSupplyIncrement(at time.Time) float64 {
	lockedUntil := time.Unix(constants.HOPPERS_FLY_TS, 0).Add(time.Hour * 24 * 365)

	availableFly := float64(LP_FLY - LP_FLY_LOCKED_AMOUNT)
	if lockedUntil.Before(at) {
		availableFly = LP_FLY_LOCKED_AMOUNT
	}

	return STEP_SECONDS * (availableFly / YEAR_SECONDS)
}

func getTeamSupplyIncrement(at time.Time) float64 {
	lockedUntil := time.Unix(constants.HOPPERS_FLY_TS, 0).Add(time.Hour * 24 * 30)

	if lockedUntil.After(at) {
		return 0
	}

	return STEP_SECONDS * (TEAM_FLY / (YEAR_SECONDS * 2))
}

func getMarketingSupplyIncrement() float64 {
	return STEP_SECONDS * MARKETING_FLY / (YEAR_SECONDS / 2)
}

func getAPASupplyIncrement() float64 {
	return STEP_SECONDS * APA_FLY / (YEAR_SECONDS / 2)
}

var mintsTrxBlacklist = []string{
	"0xb8dcbf2dc6edb5a97b7ea16cb75a8d100245674cda8b0799e7f25682ba8c2a06",
	"0x029fc5e120d6d2c3be8cf18b6f5516508f906a277485b5d6ac1cea72e14fa379",
	"0xfdfd4d5a96afad8743dd853dae13f4f3ad93750a0bfe0d97e7aad58bec011600",
	"0x4d4c3634ded08aaa7018ba87bc0d8fa3d39eb84d0743933ea6bb2055192e8435",
	"0xeb939bb2ba201928997e07b3f4a2e8d45e3867a35e60a55e183473cc04ec43ad",
	"0x7d89ddfd22afe943b631145c6df56a04baa32f15ed942792d2bc15522b496fe1",
	"0x448435d513b0b6b9da214e9d9a530309dfcf548229d8aa7ea2da47a53ba69fd7",
	"0x4570ffd5bc1ea5b67d6525794e9276d6056a2795030c6ba7108c9a1451be198a",
}
var burnsTrxBlacklist = []string{
	"0xce1c35772732713430e800629487e695e9a3ec3aca41dc2692796991e69906d2",
	"0xc723dfa300c6b797bbca206818d40ab028ba9cd8b6584fb1dbf856015ba0a0f1",
	"0x74be7eaf80f816b9befba4741dfcc2068bd333e5a8d2ef1b48f16e83a0efba54",
	"0xefe36dff105fbe23d0af79e6c06ef830421bc8fa4a148ec8a79b747cc6fb0e0e",
	"0x55b83e092a835dfb2d778eadc08f1c5c5b12b0e4d3946357d639a6980f5944c0",
	"0x034a3375f0be44157c1a8fa7a0cda3c4e4cf09b17b6a4977495e3a6265fef97b",
	"0xe5d65156dafcb32c57c05fd310dc9a83f7574e79bd03b2cd34ddd738b316e8b9",
	"0x5e398bfc47072d53d8b6b9e84c91f342f9c011b982b66d8278632b61728cf7a8",
	"0x94bc8ee6af210816112d31745f8e7c82678526f28acd3646db721e20183dba33",
	"0x384120d64febccc9f54ad423c67cae6cff8d9fce2e1f48ae2fdf2fa79f4a584d",
	"0x97acca48430c5eaebb1a62b634db0961ec886c4fd6805bec821a47ffade10c2b",
}

func isBlacklisted(transfer graph.Transfer) bool {
	transactionLower := strings.ToLower(transfer.Transaction)
	for _, trx := range mintsTrxBlacklist {
		if strings.ToLower(trx) == transactionLower {
			return true
		}
	}

	for _, trx := range burnsTrxBlacklist {
		if strings.ToLower(trx) == transactionLower {
			return true
		}
	}

	return false
}

func filterTransfersByBlacklists(transfers []graph.Transfer) []graph.Transfer {
	filtered := []graph.Transfer{}

	for _, transfer := range transfers {
		if !isBlacklisted(transfer) {
			filtered = append(filtered, transfer)
		}
	}

	return filtered
}
