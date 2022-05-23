package cmd

import (
	"log"
	"math/big"
	"time"

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

const (
	LP_FLY         = 20_000_000
	OLD_TEAM_FLY   = 10_000_000
	TEAM_FLY       = 10_000_000
	ROCKET_JOE_FLY = 2_500_000
	MARKETING_FLY  = 2_500_000
	APA_FLY        = 1_000_000
)

var lpStart, _ = time.Parse(time.RFC3339, "2022-03-12T18:00:00.00Z")
var teamStart, _ = time.Parse(time.RFC3339, "2022-04-11T18:00:00.00Z")
var marketingStart, _ = time.Parse(time.RFC3339, "2022-03-11T18:00:00.00Z")
var apaStart, _ = time.Parse(time.RFC3339, "2022-03-11T18:00:00.00Z")

var flySupplyCommand = &cobra.Command{
	Use:   "fly-supply",
	Short: "Load and save FLY supply",
	Run: func(cmd *cobra.Command, args []string) {
		dbClient := GetMongo()
		defer dbClient.Disconnect()

		flySupplyCol := &db.FlySuppliesCollection{
			Client: dbClient,
		}
		err := flySupplyCol.Clear()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		now := time.Now()
		startTimestamp := time.Unix(1646697600, 0) // 2022-03-08T00:00:00

		graphClient := graph.NewFlyGraphClient()

		mintedAccumulator := 0.0
		burnedAccumulator := 0.0
		stakeDepositAccumulator := 0.0
		stakeWithdrawAccumulator := 0.0

		for {
			if startTimestamp.After(now) {
				break
			}

			mintedB, err := loadLatestMinted(graphClient, startTimestamp)
			if err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}
			mintedF, _ := utils.ToDecimal(mintedB, 18).Float64()

			burnedB, err := loadLatestBurned(graphClient, startTimestamp)
			if err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}
			burnedF, _ := utils.ToDecimal(burnedB, 18).Float64()

			stakeDepositB, err := loadLatestStakeDeposit(graphClient, startTimestamp)
			if err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}
			stakeDepositF, _ := utils.ToDecimal(stakeDepositB, 18).Float64()

			stakeWithdrawB, err := loadLatestStakeWithdraw(graphClient, startTimestamp)
			if err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}
			stakeWithdrawF, _ := utils.ToDecimal(stakeWithdrawB, 18).Float64()

			mintedAccumulator += mintedF
			burnedAccumulator += burnedF
			stakeDepositAccumulator += stakeDepositF
			stakeWithdrawAccumulator += stakeWithdrawF

			minted := mintedAccumulator - getVestedFlyDecrement() + getVestedFlyIncrement(startTimestamp)
			burned := burnedAccumulator
			circulating := minted - burned
			staked := stakeDepositAccumulator - stakeWithdrawAccumulator
			free := circulating - staked

			doc := models.FlySupplyDocument{
				Timestamp:   startTimestamp,
				Minted:      minted,
				Burned:      burned,
				Circulating: circulating,
				Staked:      staked,
				Free:        free,
			}
			if err = flySupplyCol.Insert(doc); err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}

			log.Printf("Date: %s, Minted: %.0f, Burned: %0.f, Staked: %0.f\n", startTimestamp, minted, burned, staked)

			startTimestamp = startTimestamp.Add(time.Hour * 24)
		}
	},
}

func getVestedFlyIncrement(now time.Time) float64 {
	lp := calculateVestedFly(now, LP_FLY, lpStart, time.Hour*24*365*2)
	team := calculateVestedFly(now, TEAM_FLY, teamStart, time.Hour*24*365*2)
	marketing := calculateVestedFly(now, MARKETING_FLY, marketingStart, time.Hour*24*28*6)
	apa := calculateVestedFly(now, APA_FLY, apaStart, time.Hour*24*28*6)

	return lp + team + marketing + apa
}
func getVestedFlyDecrement() float64 {
	decrement := 0.0 + OLD_TEAM_FLY + MARKETING_FLY + APA_FLY
	return decrement
}

func calculateVestedFly(now time.Time, vestedFly float64, vestingStart time.Time, vestedDuration time.Duration) float64 {
	if vestingStart.After(now) {
		return 0
	}

	end := vestingStart.Add(vestedDuration)
	if end.Before(now) {
		return vestedFly
	}

	progress := now.Sub(vestingStart)
	alreadyVested := progress.Seconds() / vestedDuration.Seconds()

	return vestedFly * alreadyVested
}

func loadLatestMinted(graphClient *graph.FlyGraphClient, startTimestamp time.Time) (*big.Int, error) {
	dayData, err := graphClient.FetchMintsDayData(startTimestamp)
	if err != nil {
		return big.NewInt(0), err
	}
	return dayData.Minted, nil
}
func loadLatestBurned(graphClient *graph.FlyGraphClient, startTimestamp time.Time) (*big.Int, error) {
	dayData, err := graphClient.FetchBurnsDayData(startTimestamp)
	if err != nil {
		return big.NewInt(0), err
	}
	return dayData.Burned, nil
}
func loadLatestStakeDeposit(graphClient *graph.FlyGraphClient, startTimestamp time.Time) (*big.Int, error) {
	dayData, err := graphClient.FetchStakeDepositDayData(startTimestamp)
	if err != nil {
		return big.NewInt(0), err
	}
	return dayData.Deposited, nil
}
func loadLatestStakeWithdraw(graphClient *graph.FlyGraphClient, startTimestamp time.Time) (*big.Int, error) {
	dayData, err := graphClient.FetchStakeWithdrawDayData(startTimestamp)
	if err != nil {
		return big.NewInt(0), err
	}
	return dayData.Withdrawn, nil
}
