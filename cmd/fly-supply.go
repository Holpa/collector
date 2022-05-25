package cmd

import (
	"log"
	"math"
	"sort"
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
var teamStart, _ = time.Parse(time.RFC3339, "2022-04-10T18:00:00.00Z")
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

		graphClient := graph.NewFlyGraphClient()

		mints, err := graphClient.FetchMintsDayData()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		burns, err := graphClient.FetchBurnsDayData()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		circulating, err := graphClient.FetchCirculatingDayData()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		stakeDeposits, err := graphClient.FetchStakeDepositDayData()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		stakeWithdraws, err := graphClient.FetchStakeWithdrawDayData()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}

		lowestMintDayId, foundLowestMintDayId := findLowestDayId(mints)
		lowestBurnDayId, foundLowestBurnDayId := findLowestDayId(burns)
		lowestDayIdChecks := []uint{}
		if foundLowestMintDayId {
			lowestDayIdChecks = append(lowestDayIdChecks, lowestMintDayId)
		}
		if foundLowestBurnDayId {
			lowestDayIdChecks = append(lowestDayIdChecks, lowestBurnDayId)
		}

		highestBurnDayId, foundHighestBurnDayId := findHighestDayId(burns)

		minDayId := utils.Min(lowestDayIdChecks...)
		maxDayId := minDayId
		if foundHighestBurnDayId {
			maxDayId = highestBurnDayId
		}

		mintedAccumulator := 0.0
		burnedAccumulator := 0.0
		circulatingPointer := 0.0
		stakeDepositAccumulator := 0.0
		stakeWithdrawAccumulator := 0.0

		for dayId := minDayId; dayId <= maxDayId; dayId++ {
			timestamp := getTimestamp(dayId)

			minted := getDayDataAmount(mints, dayId)
			burned := getDayDataAmount(burns, dayId)
			circulating := getDayDataAmount(circulating, dayId)
			stakeDeposit := getDayDataAmount(stakeDeposits, dayId)
			stakeWithdraw := getDayDataAmount(stakeWithdraws, dayId)

			mintedAccumulator += minted
			burnedAccumulator += burned
			// Because some days may not have any FLY transfers at all
			// we just reuse the previous circulating supply
			if circulating > 0 {
				circulatingPointer = circulating
			}
			stakeDepositAccumulator += stakeDeposit
			stakeWithdrawAccumulator += stakeWithdraw

			actualCirculating := math.Max(0, circulatingPointer-getVestedFlyDecrement(timestamp)+getVestedFlyIncrement(timestamp))
			staked := math.Max(0, stakeDepositAccumulator-stakeWithdrawAccumulator)
			free := math.Max(0, actualCirculating-staked)

			doc := models.FlySupplyDocument{
				Timestamp:   timestamp,
				Minted:      mintedAccumulator,
				Burned:      burnedAccumulator - getBurnedFlyDecrement(timestamp),
				Circulating: actualCirculating,
				Staked:      staked,
				Free:        free,
			}
			if err = flySupplyCol.Insert(doc); err != nil {
				sentry.CaptureException(err)
				log.Fatalln(err)
			}
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
func getVestedFlyDecrement(now time.Time) float64 {
	decrement := OLD_TEAM_FLY + MARKETING_FLY + APA_FLY

	if now.After(teamStart) {
		decrement += TEAM_FLY
	}

	return float64(decrement)
}

func getBurnedFlyDecrement(now time.Time) float64 {
	decrement := 0.0

	if now.After(teamStart) {
		decrement += 8_000_000
	}

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

func getDayDataAmount[T ~[]I, I graph.DayDataInterface](dayDatas T, dayId uint) float64 {
	for _, data := range dayDatas {
		if data.DayId() == dayId {
			return data.AmountFloat()
		}
	}
	return 0
}

func findLowestDayId[T ~[]I, I graph.DayDataInterface](dayDatas T) (uint, bool) {
	if len(dayDatas) == 0 {
		return 0, false
	}

	copied := make([]graph.DayDataInterface, len(dayDatas))
	for i, dayData := range dayDatas {
		copied[i] = dayData
	}

	sort.SliceStable(copied, func(i, j int) bool {
		return copied[i].DayId() < copied[j].DayId()
	})

	return copied[0].DayId(), true
}
func findHighestDayId[T ~[]I, I graph.DayDataInterface](dayDatas T) (uint, bool) {
	if len(dayDatas) == 0 {
		return 0, false
	}

	copied := make([]graph.DayDataInterface, len(dayDatas))
	for i, dayData := range dayDatas {
		copied[i] = dayData
	}

	sort.SliceStable(copied, func(i, j int) bool {
		return copied[i].DayId() > copied[j].DayId()
	})

	return copied[0].DayId(), true
}

func getTimestamp(dayId uint) time.Time {
	secondsPerDay := 24 * 60 * 60
	return time.Unix(int64(secondsPerDay)*int64(dayId), 0)
}
