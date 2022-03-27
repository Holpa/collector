package helpers

import (
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/models"
)

type (
	RewardsCalculator struct {
		OnChainClient *contracts.OnChainClient
	}
)

func (calculator *RewardsCalculator) CalculateBaseEShare(adventure constants.Adventure, hopper models.Hopper) float64 {
	if !CanEnter(adventure, hopper) {
		return 0
	}

	adventureBaseShares, err := calculator.OnChainClient.GetTotalBaseShares(adventure)
	if err != nil {
		return 0
	}
	hopperBaseShares := CalculateBaseShare(adventure, hopper)

	return float64(hopperBaseShares) / float64(adventureBaseShares.Uint64())
}

func (calculator *RewardsCalculator) CalculateBaseFlyRewards(adventure constants.Adventure, hopper models.Hopper) float64 {
	if !CanEnter(adventure, hopper) {
		return 0
	}

	eshare := calculator.CalculateBaseEShare(adventure, hopper)
	adventureRewards := GetAdventureRewards(adventure)

	return eshare * adventureRewards
}

func GetAdventureRewards(adventure constants.Adventure) float64 {
	adventureRewardsShare := 0.0

	switch adventure {
	case constants.AdventurePond:
		adventureRewardsShare = float64(constants.FlyEmissionsPond)
	case constants.AdventureStream:
		adventureRewardsShare = float64(constants.FlyEmissionsStream)
	case constants.AdventureSwamp:
		adventureRewardsShare = float64(constants.FlyEmissionsSwamp)
	case constants.AdventureRiver:
		adventureRewardsShare = float64(constants.FlyEmissionsRiver)
	case constants.AdventureForest:
		adventureRewardsShare = float64(constants.FlyEmissionsForest)
	case constants.AdventureGreatLake:
		adventureRewardsShare = float64(constants.FlyEmissionsGreatLake)
	}

	return constants.BASE_FLY_EMISSIONS * adventureRewardsShare
}
