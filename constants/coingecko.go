package constants

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
