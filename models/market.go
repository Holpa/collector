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
		Seller    string
		Timestamp time.Time
		HopperId  string
	}

	ListingDocument struct {
		Id        string    `bson:"id"`
		Enabled   bool      `bson:"enabled"`
		Sold      bool      `bson:"sold"`
		Price     float64   `bson:"price"`
		Buyer     string    `bson:"buyer"`
		Seller    string    `bson:"seller"`
		Timestamp time.Time `bson:"timestamp"`
		HopperId  string    `bson:"hopperId"`
	}
)
