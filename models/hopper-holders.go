package models

import "time"

type (
	HopperHoldersDocument struct {
		Holders   uint      `bson:"holders"`
		Timestamp time.Time `bson:"timestamp"`
	}
)
