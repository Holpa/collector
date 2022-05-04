package cmd

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/steschwa/hopper-analytics-collector/contracts"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

const (
	MONGO_URI_ENV = "MONGO_URI"
	MONGO_DB_ENV  = "MONGO_DB"
)

func GetMongo() *db.MongoDbClient {
	mongoUri := os.Getenv(MONGO_URI_ENV)
	if mongoUri == "" {
		log.Fatalf("Missing enviroment variable %s\n", MONGO_URI_ENV)
	}

	databaseName := os.Getenv(MONGO_DB_ENV)
	if databaseName == "" {
		log.Fatalf("Missing environment variable %s\n", MONGO_DB_ENV)
	}

	client := db.NewMongoDbClient(databaseName)
	err := client.Connect(mongoUri)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
	}

	return client
}

func GetOnChainClient() *contracts.OnChainClient {
	onChainClient, err := contracts.NewOnChainClient()
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalln(err)
	}

	return onChainClient
}
