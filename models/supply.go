package models

import "time"

type (
	SupplyType string

	SupplyDocument struct {
		Type      SupplyType `bson:"type"`
		Timestamp time.Time  `bson:"timestamp"`
		Supply    float64    `bson:"supply"`
		Burned    float64    `bson:"burned"`
		Available float64    `bson:"available"`
		Staked    float64    `bson:"staked"`
		Free      float64    `bson:"free"`
	}
)

const (
	FLY_SUPPLY SupplyType = "fly"
)
