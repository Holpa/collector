package cmd

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/steschwa/hopper-analytics-collector/coingecko"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
	db "github.com/steschwa/hopper-analytics-collector/mongo"
)

func RegisterMigrateHistoricalPrices(root *cobra.Command) {
	root.AddCommand(migrateHistoricalPricesCommand)
}

var migrateHistoricalPricesCommand = &cobra.Command{
	Use:   "migrate-historical-prices",
	Short: "Migrate (load and save) historical FLY and AVAX prices",
	Run: func(cmd *cobra.Command, args []string) {
		dbClient := GetMongo()
		defer dbClient.Disconnect()

		ids := []constants.CoinGeckoId{
			constants.COINGECKO_AVAX,
			constants.COINGECKO_FLY,
		}
		currencies := []constants.CoinGeckoCurrency{
			constants.COINGECKO_USD,
			constants.COINGECKO_EUR,
		}

		pricesCollection := &db.PricesCollection{
			Client: dbClient,
		}
		pricesCollection.Clear()

		var wg sync.WaitGroup
		for _, coinId := range ids {
			for _, currency := range currencies {
				wg.Add(1)
				go func(dbClient *db.MongoDbClient, coinId constants.CoinGeckoId, currency constants.CoinGeckoCurrency) {
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
						Client: dbClient,
					}
					err = pricesCollection.InsertMany(docs)
					if err != nil {
						log.Println(err)
					}
				}(dbClient, coinId, currency)
			}
		}

		wg.Wait()
	},
}
