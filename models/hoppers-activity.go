package models

import "time"

type (
	HoppersActivityDocument struct {
		Adventure   uint      `bson:"adventure"`
		Pond        uint      `bson:"pond"`
		Stream      uint      `bson:"stream"`
		Swamp       uint      `bson:"swamp"`
		River       uint      `bson:"river"`
		Forest      uint      `bson:"forest"`
		GreatLake   uint      `bson:"greatLake"`
		Breeding    uint      `bson:"breeding"`
		Marketplace uint      `bson:"marketplace"`
		Idle        uint      `bson:"idle"`
		Timestamp   time.Time `bson:"timestamp"`
	}
)
