package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
