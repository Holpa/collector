package constants

const (
	BASE_FLY_EMISSIONS  = 85_000
	BONUS_FLY_EMISSIONS = 50_000
)

type FlyEmissions float64

const (
	FlyEmissionsPond      FlyEmissions = 0.05
	FlyEmissionsStream    FlyEmissions = 0.05
	FlyEmissionsSwamp     FlyEmissions = 0.05
	FlyEmissionsRiver     FlyEmissions = 0.33
	FlyEmissionsForest    FlyEmissions = 0.3
	FlyEmissionsGreatLake FlyEmissions = 0.22
)
