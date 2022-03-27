package main

import (
	"context"
	"log"
	"os"

	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_URI_ENV = "MONGO_URI"
)

func main() {
	mongoUri := os.Getenv(MONGO_URI_ENV)
	if mongoUri == "" {
		log.Fatalf("Missing enviroment variable %s\n", MONGO_URI_ENV)
	}

	mongoClient, err := db.Connect(mongoUri)
	if err != nil {
		log.Fatalln(err)
	}
	defer mongoClient.Disconnect(context.Background())

	loadAndSaveHoppers(mongoClient)
}

func loadAndSaveHoppers(mongoClient *mongo.Client) error {
	graph := graph.NewHoppersGraphClient()

	hoppers, err := graph.FetchAllHoppers()
	if err != nil {
		return err
	}

	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		return err
	}
	rewardsCalculator := helpers.NewRewardsCalculator(onChainClient)

	collection := &db.HoppersCollection{
		Connection: mongoClient,
	}
	for _, hopper := range hoppers {
		err = collection.Upsert(helpers.HopperToHopperDocument(hopper, rewardsCalculator))
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
