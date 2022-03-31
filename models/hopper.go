package models

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
		RatingPond        float64 `bson:"ratingPond"`
		RatingStream      float64 `bson:"ratingStream"`
		RatingSwamp       float64 `bson:"ratingSwamp"`
		RatingRiver       float64 `bson:"ratingRiver"`
		RatingForest      float64 `bson:"ratingForest"`
		RatingGreatLake   float64 `bson:"ratingGreatLake"`
		ListingActive     bool    `bson:"listingActive"`
		ListingPrice      float64 `bson:"listingPrice"`
		BaseFlyPond       float64 `bson:"baseFlyPond"`
		BaseFlyStream     float64 `bson:"baseFlyStream"`
		BaseFlySwamp      float64 `bson:"baseFlySwamp"`
		BaseFlyRiver      float64 `bson:"baseFlyRiver"`
		BaseFlyForest     float64 `bson:"baseFlyForest"`
		BaseFlyGreatLake  float64 `bson:"baseFlyGreatLake"`
	}
)
