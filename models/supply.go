package models

import "time"

type (
	SupplyType string

	SupplyDocument struct {
		Type      SupplyType `bson:"type"`
		Timestamp time.Time  `bson:"timestamp"`
		Supply    *BigInt    `bson:"supply"`
	}
)

const (
	FLY_SUPPLY SupplyType = "fly"
)
