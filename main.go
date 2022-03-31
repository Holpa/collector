package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	"github.com/steschwa/hopper-analytics-collector/models"
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

	err = collection.Clear()
	if err != nil {
		return err
	}

	hopperDocuments := make([]models.HopperDocument, len(hoppers))
	for i, hopper := range hoppers {
		hopperDocuments[i] = helpers.HopperToHopperDocument(hopper, rewardsCalculator)
	}
	return collection.InsertMany(hopperDocuments)
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

	err = collection.Clear()
	if err != nil {
		return err
	}

	listingDocuments := make([]models.ListingDocument, len(listings))
	for i, listing := range listings {
		listingDocuments[i] = helpers.ListingToListingDocument(listing)
	}

	return collection.InsertMany(listingDocuments)
}
