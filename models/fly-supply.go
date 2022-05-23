package models

import "time"

type (
	FlySupplyDocument struct {
		Timestamp time.Time `bson:"timestamp"`
		Minted    float64   `bson:"minted"`
		Burned    float64   `bson:"burned"`
		// Supply - Burned
		Circulating float64 `bson:"circulating"`
		Staked      float64 `bson:"staked"`
		// Circulating - Staked
		Free float64 `bson:"free"`
	}
)
