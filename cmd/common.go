package cmd

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_URI_ENV = "MONGO_URI"
)

func GetMongo() *mongo.Client {
	mongoUri := os.Getenv(MONGO_URI_ENV)
	if mongoUri == "" {
		log.Fatalf("Missing enviroment variable %s\n", MONGO_URI_ENV)
	}

	mongoClient, err := db.Connect(mongoUri)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
	}

	return mongoClient
}

func GetOnChainClient() *contracts.OnChainClient {
	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
	}

	return onChainClient
}
