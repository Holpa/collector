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

func GetHopperRating(adventure constants.Adventure, hopper models.Hopper) int {
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

func CalculatePondRating(hopper models.Hopper) int {
	return hopper.Strength
}
func CalculateStreamRating(hopper models.Hopper) int {
	return hopper.Agility
}
func CalculateSwampRating(hopper models.Hopper) int {
	return hopper.Vitality
}
func CalculateRiverRating(hopper models.Hopper) int {
	if !CanEnterRiver(hopper) {
		return 0
	}

	return hopper.Strength * hopper.Intelligence
}
func CalculateForestRating(hopper models.Hopper) int {
	if !CanEnterForest(hopper) {
		return 0
	}

	return hopper.Agility * hopper.Vitality * hopper.Intelligence
}
func CalculateGreatLakeRating(hopper models.Hopper) int {
	if !CanEnterGreatLake(hopper) {
		return 0
	}

	return hopper.Strength * hopper.Agility * hopper.Vitality * hopper.Intelligence
}
