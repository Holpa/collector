package helpers

import (
	"math"
	"math/big"

	"github.com/steschwa/hopper-analytics-collector/models"
)

func HopperToHopperDocument(hopper models.Hopper) models.HopperDocument {
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

	return models.HopperDocument{
		TokenId:           hopper.TokenId,
		Strength:          hopper.Strength,
		Agility:           hopper.Agility,
		Vitality:          hopper.Vitality,
		Intelligence:      hopper.Intelligence,
		Fertility:         hopper.Fertility,
		Level:             hopper.Level,
		Image:             hopper.Image,
		Adventure:         hopper.Adventure,
		Market:            hopper.Market,
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
	}
}

func HopperDocumentToHopper(hopperDocument models.HopperDocument) models.Hopper {
	return models.Hopper{
		TokenId:      hopperDocument.TokenId,
		Strength:     hopperDocument.Strength,
		Agility:      hopperDocument.Agility,
		Vitality:     hopperDocument.Vitality,
		Intelligence: hopperDocument.Intelligence,
		Fertility:    hopperDocument.Fertility,
		Level:        hopperDocument.Level,
		Image:        hopperDocument.Image,
		Adventure:    hopperDocument.Adventure,
		Market:       hopperDocument.Market,
		Listings: []models.Listing{{
			Enabled: hopperDocument.ListingActive,
			Sold:    false,
			Price:   big.NewFloat(hopperDocument.ListingPrice),
		}},
	}
}
