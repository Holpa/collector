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
		ChainOwner   string
		User         string
		Listings     []Listing
	}

	HopperActivity string

	HopperDocument struct {
		TokenId           string         `bson:"tokenId"`
		Strength          int            `bson:"strength"`
		Agility           int            `bson:"agility"`
		Vitality          int            `bson:"vitality"`
		Intelligence      int            `bson:"intelligence"`
		Fertility         int            `bson:"fertility"`
		Level             int            `bson:"level"`
		Image             string         `bson:"image"`
		Activity          HopperActivity `bson:"activity"`
		CanEnterPond      bool           `bson:"canEnterPond"`
		CanEnterStream    bool           `bson:"canEnterStream"`
		CanEnterSwamp     bool           `bson:"canEnterSwamp"`
		CanEnterRiver     bool           `bson:"canEnterRiver"`
		CanEnterForest    bool           `bson:"canEnterForest"`
		CanEnterGreatLake bool           `bson:"canEnterGreatLake"`
		RatingPond        float64        `bson:"ratingPond"`
		RatingStream      float64        `bson:"ratingStream"`
		RatingSwamp       float64        `bson:"ratingSwamp"`
		RatingRiver       float64        `bson:"ratingRiver"`
		RatingForest      float64        `bson:"ratingForest"`
		RatingGreatLake   float64        `bson:"ratingGreatLake"`
		MarketPrice       float64        `bson:"marketPrice"`
		BaseFlyPond       float64        `bson:"baseFlyPond"`
		BaseFlyStream     float64        `bson:"baseFlyStream"`
		BaseFlySwamp      float64        `bson:"baseFlySwamp"`
		BaseFlyRiver      float64        `bson:"baseFlyRiver"`
		BaseFlyForest     float64        `bson:"baseFlyForest"`
		BaseFlyGreatLake  float64        `bson:"baseFlyGreatLake"`
		Owner             string         `bson:"owner"`
	}
)

const (
	HopperActivityIdle        HopperActivity = "idle"
	HopperActivityPond        HopperActivity = "pond"
	HopperActivityStream      HopperActivity = "stream"
	HopperActivitySwamp       HopperActivity = "swamp"
	HopperActivityRiver       HopperActivity = "river"
	HopperActivityForest      HopperActivity = "forest"
	HopperActivityGreatLake   HopperActivity = "great-lake"
	HopperActivityMarketplace HopperActivity = "marketplace"
	HopperActivityBreeding    HopperActivity = "breeding"
)
