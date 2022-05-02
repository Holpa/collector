package cmd

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/coingecko"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterMigrateHistoricalPrices(root *cobra.Command) {
	root.AddCommand(migrateHistoricalPricesCommand)
}

var migrateHistoricalPricesCommand = &cobra.Command{
	Use:   "migrate-historical-prices",
	Short: "Migrate (load and save) historical FLY and AVAX prices",
	Run: func(cmd *cobra.Command, args []string) {
		mongoClient := GetMongo()
		defer mongoClient.Disconnect(context.Background())

		ids := []constants.CoinGeckoId{
			constants.COINGECKO_AVAX,
			constants.COINGECKO_FLY,
		}
		currencies := []constants.CoinGeckoCurrency{
			constants.COINGECKO_USD,
			constants.COINGECKO_EUR,
		}

		pricesCollection := &db.PricesCollection{
			Connection: mongoClient,
		}
		pricesCollection.Clear()

		var wg sync.WaitGroup
		for _, coinId := range ids {
			for _, currency := range currencies {
				wg.Add(1)
				go func(mongoClient *mongo.Client, coinId constants.CoinGeckoId, currency constants.CoinGeckoCurrency) {
					defer wg.Done()

					coingeckoClient := coingecko.NewCoinGeckoClient()
					historicalPrices, err := coingeckoClient.HistoricalPrices(coinId, currency)
					if err != nil {
						log.Println(err)
						return
					}

					docs := make([]models.PriceDocument, len(historicalPrices))
					for i, historicalPrice := range historicalPrices {
						tsMs := historicalPrice[0]
						ts := time.Unix(int64(tsMs/1000), 0)

						price := historicalPrice[1]

						docs[i] = models.PriceDocument{
							Coin:      coinId,
							Currency:  currency,
							Price:     price,
							Timestamp: ts,
						}
					}

					pricesCollection := &db.PricesCollection{
						Connection: mongoClient,
					}
					err = pricesCollection.InsertMany(docs)
					if err != nil {
						log.Println(err)
					}
				}(mongoClient, coinId, currency)
			}
		}

		wg.Wait()
	},
}
