package models

import "time"

type (
	FlySupplyDocument struct {
		Timestamp time.Time `bson:"timestamp"`
		Supply    float64   `bson:"supply"`
		Burned    float64   `bson:"burned"`
		// Supply - Burned
		Available float64 `bson:"available"`
		Staked    float64 `bson:"staked"`
		// Available - Staked
		Free float64 `bson:"free"`
	}
)
