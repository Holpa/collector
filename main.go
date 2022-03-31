package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_URI_ENV = "MONGO_URI"
)

type (
	Operation func(mongoClient *mongo.Client) error
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

	var wg sync.WaitGroup

	operations := []Operation{
		loadAndSaveHoppers,
		loadAndSaveMarketListings,
	}

	for _, operation := range operations {
		wg.Add(1)
		go func(mongoClient *mongo.Client, operation Operation) {
			defer wg.Done()

			err := operation(mongoClient)
			if err != nil {
				log.Println(err)
			}
		}(mongoClient, operation)
	}

	wg.Wait()
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

func loadAndSaveMarketListings(mongoClient *mongo.Client) error {
	graph := graph.NewMarketsGraphClient()

	listings, err := graph.FetchAllListings()
	if err != nil {
		return err
	}

	collection := &db.MarketsCollection{
		Connection: mongoClient,
	}
	for _, listing := range listings {
		err = collection.Upsert(helpers.ListingToListingDocument(listing))
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
