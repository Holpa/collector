package helpers

import (
	"math"
	"strings"

	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
)

func HopperToHopperDocument(hopper models.Hopper, rewardsCalculator *RewardsCalculator) models.HopperDocument {
	listingActive := false
	listingPrice := 0.0
	for _, listing := range hopper.Listings {
		if listing.Enabled && !listing.Sold {
			listingActive = true
			val, _ := listing.Price.Float64()
			listingPrice = val * math.Pow(10, -18)
			break
		}
	}

	baseFlyPond := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventurePond, hopper)
	baseFlyStream := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureStream, hopper)
	baseFlySwamp := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureSwamp, hopper)
	baseFlyRiver := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureRiver, hopper)
	baseFlyForest := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureForest, hopper)
	baseFlyGreatLake := rewardsCalculator.CalculateBaseFlyRewards(constants.AdventureGreatLake, hopper)

	adventure, _ := constants.AdventureFromContract(hopper.ChainOwner)
	hopperAdventure := adventure.String()
	if !hopper.Adventure {
		hopperAdventure = ""
	}

	return models.HopperDocument{
		TokenId:           hopper.TokenId,
		Strength:          hopper.Strength,
		Agility:           hopper.Agility,
		Vitality:          hopper.Vitality,
		Intelligence:      hopper.Intelligence,
		Fertility:         hopper.Fertility,
		Level:             hopper.Level,
		Image:             hopper.Image,
		InAdventure:       hopper.Adventure,
		Adventure:         hopperAdventure,
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
		ListingActive:     listingActive,
		ListingPrice:      listingPrice,
		BaseFlyPond:       baseFlyPond,
		BaseFlyStream:     baseFlyStream,
		BaseFlySwamp:      baseFlySwamp,
		BaseFlyRiver:      baseFlyRiver,
		BaseFlyForest:     baseFlyForest,
		BaseFlyGreatLake:  baseFlyGreatLake,
		Owner:             strings.ToLower(hopper.User),
	}
}
