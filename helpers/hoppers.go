package helpers

import (
	"github.com/steschwa/hopper-analytics-collector/models"
)

func HopperToHopperDocument(hopper models.Hopper) models.HopperDocument {
	return models.HopperDocument{
		TokenId:         hopper.TokenId,
		Strength:        hopper.Strength,
		Agility:         hopper.Agility,
		Vitality:        hopper.Vitality,
		Intelligence:    hopper.Intelligence,
		Fertility:       hopper.Fertility,
		Level:           hopper.Level,
		Image:           hopper.Image,
		Adventure:       hopper.Adventure,
		Market:          hopper.Market,
		RatingPond:      CalculatePondRating(hopper),
		RatingStream:    CalculateStreamRating(hopper),
		RatingSwamp:     CalculateSwampRating(hopper),
		RatingRiver:     CalculateRiverRating(hopper),
		RatingForest:    CalculateForestRating(hopper),
		RatingGreatLake: CalculateGreatLakeRating(hopper),
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
	}
}
