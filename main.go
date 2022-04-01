package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"sync"

	"github.com/shopspring/decimal"
	"github.com/steschwa/hopper-analytics-collector/constants"
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
		loadAndSaveVotes,
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

func loadAndSaveVotes(mongoClient *mongo.Client) error {
	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		return err
	}

	adventures := []constants.Adventure{
		constants.AdventurePond,
		constants.AdventureStream,
		constants.AdventureSwamp,
		constants.AdventureRiver,
		constants.AdventureForest,
		constants.AdventureGreatLake,
	}

	totalVotes := big.NewInt(0)
	voteDocuments := []models.VoteDocument{}
	for _, adventure := range adventures {
		votes, err := onChainClient.GetVotesByAdventure(adventure)
		if err != nil {
			log.Println(err)
			continue
		}

		totalVotes = big.NewInt(0).Add(totalVotes, votes)
		voteDocuments = append(voteDocuments, models.VoteDocument{
			Adventure: adventure.String(),
			Votes:     models.NewBigInt(votes),
		})
	}

	collection := &db.VotesCollection{
		Connection: mongoClient,
	}

	for _, voteDocument := range voteDocuments {
		patchedVoteDocument := models.VoteDocument(voteDocument)

		votes, err := decimal.NewFromString(voteDocument.Votes.Int().String())
		if err != nil {
			log.Println(err)
			continue
		}

		totalShareStr, err := decimal.NewFromString(totalVotes.String())
		if err != nil {
			log.Println(err)
			continue
		}
		share, _ := votes.Div(totalShareStr).Float64()

		patchedVoteDocument.TotalVotes = models.NewBigInt(totalVotes)
		patchedVoteDocument.VotesShare = share

		err = collection.Upsert(patchedVoteDocument)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
