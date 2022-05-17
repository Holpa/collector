package helpers

import (
	"fmt"
	"math"
	"strings"

	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
)

func HopperToHopperDocument(hopper models.Hopper, rewardsCalculator *RewardsCalculator) models.HopperDocument {
	marketPrice := 0.0
	for _, listing := range hopper.Listings {
		if listing.Enabled && !listing.Sold {
			val, _ := listing.Price.Float64()
			marketPrice = val * math.Pow(10, -18)
			break
		}
	}

	baseFlyPond := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventurePond, hopper)
	baseFlyStream := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureStream, hopper)
	baseFlySwamp := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureSwamp, hopper)
	baseFlyRiver := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureRiver, hopper)
	baseFlyForest := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureForest, hopper)
	baseFlyGreatLake := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureGreatLake, hopper)

	var activity models.HopperActivity = models.HopperActivityIdle

	chainOwnerLower := strings.ToLower(hopper.ChainOwner)
	switch chainOwnerLower {
	case constants.ADVENTURE_POND_CONTRACT:
		activity = models.HopperActivityPond
	case constants.ADVENTURE_STREAM_CONTRACT:
		activity = models.HopperActivityStream
	case constants.ADVENTURE_SWAMP_CONTRACT:
		activity = models.HopperActivitySwamp
	case constants.ADVENTURE_RIVER_CONTRACT:
		activity = models.HopperActivityRiver
	case constants.ADVENTURE_FOREST_CONTRACT:
		activity = models.HopperActivityForest
	case constants.ADVENTURE_GREAT_LAKE_CONTRACT:
		activity = models.HopperActivityGreatLake
	case constants.BREEDING_CONTRACT_V1:
		activity = models.HopperActivityBreeding
	case constants.BREEDING_CONTRACT_V2:
		activity = models.HopperActivityBreeding
	case constants.MARKETPLACE_CONTRACT:
		activity = models.HopperActivityMarketplace
	}

	image := fmt.Sprintf("https://hoppers.mypinata.cloud/ipfs/QmPWaQSContemvXU21KvLQNhtgqtSgQbYVFAfz8fn4QdsL/%s.png", hopper.TokenId)

	return models.HopperDocument{
		TokenId:           hopper.TokenId,
		Strength:          hopper.Strength,
		Agility:           hopper.Agility,
		Vitality:          hopper.Vitality,
		Intelligence:      hopper.Intelligence,
		Fertility:         hopper.Fertility,
		Level:             hopper.Level,
		Image:             image,
		Activity:          activity,
		CanEnterPond:      true,
		CanEnterStream:    true,
		CanEnterSwamp:     true,
		CanEnterRiver:     CanEnterRiver(hopper),
		CanEnterForest:    CanEnterForest(hopper),
		CanEnterGreatLake: CanEnterGreatLake(hopper),
		RatingPond:        CalculatePondRating(hopper),
		RatingStream:      CalculateStreamRating(hopper),
		RatingSwamp:       CalculateSwampRating(hopper),
		RatingRiver:       CalculateRiverRating(hopper),
		RatingForest:      CalculateForestRating(hopper),
		RatingGreatLake:   CalculateGreatLakeRating(hopper),
		MarketPrice:       marketPrice,
		BaseFlyPond:       baseFlyPond,
		BaseFlyStream:     baseFlyStream,
		BaseFlySwamp:      baseFlySwamp,
		BaseFlyRiver:      baseFlyRiver,
		BaseFlyForest:     baseFlyForest,
		BaseFlyGreatLake:  baseFlyGreatLake,
		Owner:             strings.ToLower(hopper.User),
	}
}
