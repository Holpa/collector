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
		Market       bool
	}

	HopperDocument struct {
		TokenId         string `bson:"tokenId"`
		Strength        int    `bson:"strength"`
		Agility         int    `bson:"agility"`
		Vitality        int    `bson:"vitality"`
		Intelligence    int    `bson:"intelligence"`
		Fertility       int    `bson:"fertility"`
		Level           int    `bson:"level"`
		Image           string `bson:"image"`
		Adventure       bool   `bson:"adventure"`
		Market          bool   `bson:"market"`
		RatingPond      int    `bson:"ratingPond"`
		RatingStream    int    `bson:"ratingStream"`
		RatingSwamp     int    `bson:"ratingSwamp"`
		RatingRiver     int    `bson:"ratingRiver"`
		RatingForest    int    `bson:"ratingForest"`
		RatingGreatLake int    `bson:"ratingGreatLake"`
	}
)
