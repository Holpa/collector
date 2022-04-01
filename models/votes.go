package models

import (
	"time"
)

type (
	VoteDocument struct {
		Adventure  string    `bson:"adventure"`
		Updated    time.Time `bson:"updated"`
		Votes      *BigInt   `bson:"votes"`
		VotesShare float64   `bson:"votesShare"`
		TotalVotes *BigInt   `bson:"totalVotes"`
	}
)
