package models

import "time"

type (
	BaseSharesDocument struct {
		Adventure       string    `bson:"adventure"`
		Updated         time.Time `bson:"updated"`
		TotalBaseShares *BigInt   `bson:"totalBaseShares"`
	}
)
