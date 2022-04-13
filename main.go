package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/steschwa/hopper-analytics-collector/coingecko"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	"github.com/steschwa/hopper-analytics-collector/graph"
	"github.com/steschwa/hopper-analytics-collector/helpers"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/getsentry/sentry-go"
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

	initSentry()
	defer sentry.Flush(2 * time.Second)

	mongoClient, err := db.Connect(mongoUri)
	if err != nil {
		log.Fatalln(err)
	}
	defer mongoClient.Disconnect(context.Background())

	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	operations := []Operation{
		loadAndSaveHoppers(onChainClient),
		loadAndSaveMarketListings,
		loadAndSaveVotes(onChainClient),
		loadAndSavePrices,
		loadAndSaveSupplies(onChainClient),
		loadAndSaveBaseShares(onChainClient),
	}

	for _, operation := range operations {
		wg.Add(1)
		go func(mongoClient *mongo.Client, operation Operation) {
			defer wg.Done()

			err := operation(mongoClient)
			if err != nil {
				sentry.CaptureException(err)
			}
		}(mongoClient, operation)
	}

	wg.Wait()
}

func initSentry() {
	env := os.Getenv("ENV")

	if env == "" {
		env = "production"
	}

	if env == "production" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         "https://0997d0f7af464bd29da229b2c9f39c05@o1202748.ingest.sentry.io/6328151",
			Environment: "production",
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func loadAndSaveHoppers(onChainClient *contracts.OnChainClient) Operation {
	return func(mongoClient *mongo.Client) error {
		graph := graph.NewHoppersGraphClient()

		hoppers, err := graph.FetchAllHoppers()
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

func loadAndSavePrices(mongoClient *mongo.Client) error {
	coinGeckoClient := coingecko.NewCoinGeckoClient()

	ids := []constants.CoinGeckoId{
		constants.COINGECKO_AVAX,
		constants.COINGECKO_FLY,
	}
	currencies := []constants.CoinGeckoCurrency{
		constants.COINGECKO_USD,
		constants.COINGECKO_EUR,
	}

	prices, err := coinGeckoClient.CurrentPrice(ids, currencies)
	if err != nil {
		return err
	}

	priceDocuments := []models.PriceDocument{}
	for coin, priceData := range prices {
		for currency, price := range priceData {
			priceDocuments = append(priceDocuments, models.PriceDocument{
				Coin:      coin,
				Currency:  currency,
				Price:     price,
				Timestamp: time.Now(),
			})
		}
	}

	pricesCollection := &db.PricesCollection{
		Connection: mongoClient,
	}
	err = pricesCollection.Clear()
	if err != nil {
		return err
	}

	return pricesCollection.InsertMany(priceDocuments)
}

func loadAndSaveSupplies(onChainClient *contracts.OnChainClient) Operation {
	return func(mongoClient *mongo.Client) error {
		flySupply, err := onChainClient.GetFlySupply()
		if err != nil {
			return err
		}

		supplyDocument := models.SupplyDocument{
			Type:   models.FLY_SUPPLY,
			Supply: models.NewBigInt(flySupply),
		}

		collection := &db.SuppliesCollection{
			Connection: mongoClient,
		}
		return collection.Insert(supplyDocument)
	}
}

func loadAndSaveVotes(onChainClient *contracts.OnChainClient) Operation {
	return func(mongoClient *mongo.Client) error {
		adventures := []constants.Adventure{
			constants.AdventurePond,
			constants.AdventureStream,
			constants.AdventureSwamp,
			constants.AdventureRiver,
			constants.AdventureForest,
			constants.AdventureGreatLake,
		}

		collection := &db.VotesCollection{
			Connection: mongoClient,
		}

		for _, adventure := range adventures {
			votes, err := onChainClient.GetTotalVotesByAdventure(adventure)
			if err != nil {
				sentry.CaptureException(err)
				continue
			}
			veShare, err := onChainClient.GetTotalVeShareByAdventure(adventure)
			if err != nil {
				sentry.CaptureException(err)
				continue
			}

			voteDocument := models.VoteDocument{
				Adventure: adventure.String(),
				Votes:     models.NewBigInt(votes),
				VeShare:   models.NewBigInt(veShare),
			}

			err = collection.Insert(voteDocument)
			if err != nil {
				sentry.CaptureException(err)
			}
		}

		return nil
	}
}

func loadAndSaveBaseShares(onChainClient *contracts.OnChainClient) Operation {
	return func(mongoClient *mongo.Client) error {
		adventures := []constants.Adventure{
			constants.AdventurePond,
			constants.AdventureStream,
			constants.AdventureSwamp,
			constants.AdventureRiver,
			constants.AdventureForest,
			constants.AdventureGreatLake,
		}

		collection := &db.BaseSharesCollection{
			Connection: mongoClient,
		}
		for _, adventure := range adventures {
			totalBaseShares, err := onChainClient.GetTotalBaseSharesByAdventure(adventure)
			if err != nil {
				sentry.CaptureException(err)
				continue
			}

			collection.Insert(models.BaseSharesDocument{
				Adventure:       adventure.String(),
				TotalBaseShares: models.NewBigInt(totalBaseShares),
			})
		}

		return nil
	}
}
