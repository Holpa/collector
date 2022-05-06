package coingecko

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	CoinGeckoClient struct {
		HttpClient http.Client
	}
)

func NewCoinGeckoClient() *CoinGeckoClient {
	httpClient := http.Client{}

	return &CoinGeckoClient{
		HttpClient: httpClient,
	}
}

func toStringSlice[T constants.CoinGeckoId | constants.CoinGeckoCurrency](value []T) []string {
	out := make([]string, len(value))
	for i, v := range value {
		out[i] = string(v)
	}

	return out
}

// ----------------------------------------
// Response data
// ----------------------------------------
type (
	CurrentPrice     map[constants.CoinGeckoCurrency]float64
	CurrentPriceData map[constants.CoinGeckoId]CurrentPrice

	HistoricalPrice          []float64
	HistoricalPricesData     []HistoricalPrice
	HistoricalPricesResponse struct {
		Prices HistoricalPricesData `json:"prices"`
	}
)

// ----------------------------------------
// Methods
// ----------------------------------------

func (client *CoinGeckoClient) CurrentPrice(ids []constants.CoinGeckoId, currencies []constants.CoinGeckoCurrency) (CurrentPriceData, error) {
	params := url.Values{}
	params.Add("ids", strings.Join(toStringSlice(ids), ","))
	params.Add("vs_currencies", strings.Join(toStringSlice(currencies), ","))

	url := fmt.Sprintf("%s/simple/price?%s", constants.COINGECKO_ENDPOINT, params.Encode())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return CurrentPriceData{}, err
	}

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return CurrentPriceData{}, err
	}
	defer res.Body.Close()

	data := CurrentPriceData{}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return CurrentPriceData{}, err
	}

	// Validate response
	for _, id := range ids {
		priceData, ok := data[id]
		if !ok {
			return CurrentPriceData{}, fmt.Errorf("missing id %s", id)
		}

		for _, currency := range currencies {
			_, ok := priceData[currency]
			if !ok {
				return CurrentPriceData{}, fmt.Errorf("missing currency %s in id %s", currency, id)
			}
		}
	}

	return data, nil
}

func (client *CoinGeckoClient) HistoricalPrices(id constants.CoinGeckoId, currency constants.CoinGeckoCurrency) (HistoricalPricesData, error) {
	params := url.Values{}
	params.Add("vs_currency", string(currency))

	daysMax := 89
	notBefore := time.Unix(constants.HOPPERS_FLY_TS, 0)
	daysDiff := time.Since(notBefore).Hours() / 24

	days := int(math.Floor(math.Min(float64(daysMax), daysDiff)))

	params.Add("days", fmt.Sprint(days))

	url := fmt.Sprintf("%s/coins/%s/market_chart?%s", constants.COINGECKO_ENDPOINT, string(id), params.Encode())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return HistoricalPricesData{}, err
	}

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return HistoricalPricesData{}, err
	}
	defer res.Body.Close()

	data := HistoricalPricesResponse{}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return HistoricalPricesData{}, err
	}

	return data.Prices, nil
}
