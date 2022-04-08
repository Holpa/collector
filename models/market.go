package models

import (
	"math/big"
	"time"
)

type (
	Listing struct {
		Id        string
		Enabled   bool
		Sold      bool
		Price     *big.Float
		Buyer     string
		Timestamp time.Time
		HopperId  string
	}

	ListingDocument struct {
		Id        string    `bson:"string"`
		Enabled   bool      `bson:"enabled"`
		Sold      bool      `bson:"sold"`
		Price     float64   `bson:"price"`
		Buyer     string    `bson:"buyer"`
		Timestamp time.Time `bson:"timestamp"`
		HopperId  string    `bson:"hopperId"`
	}
)
