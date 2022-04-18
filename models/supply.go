package models

import "time"

type (
	SupplyType string

	LegacySupplyDocument struct {
		Type      SupplyType `bson:"type"`
		Timestamp time.Time  `bson:"timestamp"`
		Supply    *BigInt    `bson:"supply"`
	}

	SupplyDocument struct {
		Type      SupplyType `bson:"type"`
		Timestamp time.Time  `bson:"timestamp"`
		Supply    float64    `bson:"supply"`
	}
)

const (
	FLY_SUPPLY SupplyType = "fly"
)
