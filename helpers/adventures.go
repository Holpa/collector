package helpers

import (
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
)

func CanEnter(adventure constants.Adventure, hopper models.Hopper) bool {
	switch adventure {
	case constants.AdventurePond, constants.AdventureStream, constants.AdventureSwamp:
		return true
	case constants.AdventureRiver:
		return CanEnterRiver(hopper)
	case constants.AdventureForest:
		return CanEnterForest(hopper)
	case constants.AdventureGreatLake:
		return CanEnterGreatLake(hopper)
	default:
		return false
	}
}

func CanEnterRiver(hopper models.Hopper) bool {
	return hopper.Strength >= 5 && hopper.Intelligence >= 5
}
func CanEnterForest(hopper models.Hopper) bool {
	return hopper.Agility >= 5 && hopper.Vitality >= 5 && hopper.Intelligence >= 5
}
func CanEnterGreatLake(hopper models.Hopper) bool {
	return hopper.Strength >= 5 && hopper.Agility >= 5 && hopper.Vitality >= 5 && hopper.Intelligence >= 5
}

func GetHopperRating(adventure constants.Adventure, hopper models.Hopper) float64 {
	switch adventure {
	case constants.AdventurePond:
		return CalculatePondRating(hopper)
	case constants.AdventureStream:
		return CalculateStreamRating(hopper)
	case constants.AdventureSwamp:
		return CalculateSwampRating(hopper)
	case constants.AdventureRiver:
		return CalculateRiverRating(hopper)
	case constants.AdventureForest:
		return CalculateForestRating(hopper)
	case constants.AdventureGreatLake:
		return CalculateGreatLakeRating(hopper)
	default:
		return 0
	}
}

func CalculatePondRating(hopper models.Hopper) float64 {
	return float64(hopper.Strength) / constants.MAX_RATING_POND
}
func CalculateStreamRating(hopper models.Hopper) float64 {
	return float64(hopper.Agility) / constants.MAX_RATING_STREAM
}
func CalculateSwampRating(hopper models.Hopper) float64 {
	return float64(hopper.Vitality) / constants.MAX_RATING_SWAMP
}
func CalculateRiverRating(hopper models.Hopper) float64 {
	if !CanEnterRiver(hopper) {
		return 0
	}

	return float64(hopper.Strength*hopper.Intelligence) / constants.MAX_RATING_RIVER
}
func CalculateForestRating(hopper models.Hopper) float64 {
	if !CanEnterForest(hopper) {
		return 0
	}

	return float64(hopper.Agility*hopper.Vitality*hopper.Intelligence) / constants.MAX_RATING_FOREST
}
func CalculateGreatLakeRating(hopper models.Hopper) float64 {
	if !CanEnterGreatLake(hopper) {
		return 0
	}

	return float64(hopper.Strength*hopper.Agility*hopper.Vitality*hopper.Intelligence) / constants.MAX_RATING_GREAT_LAKE
}

func CalculateBaseShare(adventure constants.Adventure, hopper models.Hopper) int {
	switch adventure {
	case constants.AdventurePond:
		return CalculatePondBaseShare(hopper)
	case constants.AdventureStream:
		return CalculateStreamBaseShare(hopper)
	case constants.AdventureSwamp:
		return CalculateSwampBaseShare(hopper)
	case constants.AdventureRiver:
		return CalculateRiverBaseShare(hopper)
	case constants.AdventureForest:
		return CalculateForestBaseShare(hopper)
	case constants.AdventureGreatLake:
		return CalculateGreatLakeBaseShare(hopper)
	default:
		return 0
	}
}
func CalculatePondBaseShare(hopper models.Hopper) int {
	return hopper.Strength
}
func CalculateStreamBaseShare(hopper models.Hopper) int {
	return hopper.Agility
}
func CalculateSwampBaseShare(hopper models.Hopper) int {
	return hopper.Vitality
}
func CalculateRiverBaseShare(hopper models.Hopper) int {
	if !CanEnterRiver(hopper) {
		return 0
	}

	return hopper.Strength * hopper.Intelligence
}
func CalculateForestBaseShare(hopper models.Hopper) int {
	if !CanEnterForest(hopper) {
		return 0
	}

	return hopper.Agility * hopper.Vitality * hopper.Intelligence
}
func CalculateGreatLakeBaseShare(hopper models.Hopper) int {
	if !CanEnterGreatLake(hopper) {
		return 0
	}

	return hopper.Strength * hopper.Agility * hopper.Vitality * hopper.Intelligence
}
