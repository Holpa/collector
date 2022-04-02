package models

import (
	"time"

	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	PriceDocument struct {
		Coin      constants.CoinGeckoId       `bson:"coin"`
		Currency  constants.CoinGeckoCurrency `bson:"currency"`
		Price     float64                     `bson:"price"`
		Timestamp time.Time                   `bson:"timestamp"`
	}
)
