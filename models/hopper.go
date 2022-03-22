package models

import "math/big"

type (
	Hopper struct {
		TokenId      string
		Strength     int
		Agility      int
		Vitality     int
		Intelligence int
		Fertility    int
		Level        int
		Image        string
		Adventure    bool
		Listings     []Listing
	}

	Listing struct {
		Enabled bool
		Sold    bool
		Price   *big.Float
	}

	HopperDocument struct {
		TokenId           string  `bson:"tokenId"`
		Strength          int     `bson:"strength"`
		Agility           int     `bson:"agility"`
		Vitality          int     `bson:"vitality"`
		Intelligence      int     `bson:"intelligence"`
		Fertility         int     `bson:"fertility"`
		Level             int     `bson:"level"`
		Image             string  `bson:"image"`
		Adventure         bool    `bson:"adventure"`
		CanEnterPond      bool    `bson:"canEnterPond"`
		CanEnterStream    bool    `bson:"canEnterStream"`
		CanEnterSwamp     bool    `bson:"canEnterSwamp"`
		CanEnterRiver     bool    `bson:"canEnterRiver"`
		CanEnterForest    bool    `bson:"canEnterForest"`
		CanEnterGreatLake bool    `bson:"canEnterGreatLake"`
		RatingPond        int     `bson:"ratingPond"`
		RatingStream      int     `bson:"ratingStream"`
		RatingSwamp       int     `bson:"ratingSwamp"`
		RatingRiver       int     `bson:"ratingRiver"`
		RatingForest      int     `bson:"ratingForest"`
		RatingGreatLake   int     `bson:"ratingGreatLake"`
		ListingActive     bool    `bson:"listingActive"`
		ListingPrice      float64 `bson:"listingPrice"`
	}
)
