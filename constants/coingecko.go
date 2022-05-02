package constants

import "strings"

type (
	CoinGeckoId       string
	CoinGeckoCurrency string
)

const (
	COINGECKO_ENDPOINT = "https://api.coingecko.com/api/v3"
)

const (
	COINGECKO_AVAX CoinGeckoId = "avalanche-2"
	COINGECKO_FLY  CoinGeckoId = "hoppers-game"
)

const (
	COINGECKO_USD CoinGeckoCurrency = "usd"
	COINGECKO_EUR CoinGeckoCurrency = "eur"
)

func CoingeckoIdFromString(value string, fallback CoinGeckoId) CoinGeckoId {
	lowerCased := strings.ToLower(value)

	switch lowerCased {
	case "avax":
		return COINGECKO_AVAX
	case "fly":
		return COINGECKO_FLY
	default:
		return fallback
	}
}
func (coingeckoId CoinGeckoId) String() string {
	switch coingeckoId {
	case COINGECKO_AVAX:
		return "avax"
	case COINGECKO_FLY:
		return "fly"
	default:
		return "unknown-id"
	}
}

func CoingeckoCurrenyFromString(value string) CoinGeckoCurrency {
	lowerCased := strings.ToLower(value)

	switch lowerCased {
	case "usd":
		return COINGECKO_USD
	case "eur":
		return COINGECKO_EUR
	default:
		return COINGECKO_USD
	}
}
func (coingeckoCurrency CoinGeckoCurrency) String() string {
	switch coingeckoCurrency {
	case COINGECKO_USD:
		return "usd"
	case COINGECKO_EUR:
		return "eur"
	default:
		return "unknown-currency"
	}
}
