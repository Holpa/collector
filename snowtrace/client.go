package snowtrace

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	SnowTraceApiClient struct {
		ApiKey     string
		HttpClient http.Client
	}
)

func NewSnowTraceClient(apiKey string) *SnowTraceApiClient {
	httpClient := http.Client{}

	return &SnowTraceApiClient{
		ApiKey:     apiKey,
		HttpClient: httpClient,
	}
}

// ----------------------------------------
// Response data
// ----------------------------------------

type (
	BlockByTimestampResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
)

func (client *SnowTraceApiClient) GetBlockAtTimestamp(timestamp uint64) (uint64, error) {
	params := url.Values{}
	params.Add("module", "block")
	params.Add("action", "getblocknobytime")
	params.Add("timestamp", fmt.Sprint(timestamp))
	params.Add("closest", "before")
	params.Add("apikey", client.ApiKey)

	url := fmt.Sprintf("%s?%s", constants.SNOWTRACE_ENDPOINT, params.Encode())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	data := BlockByTimestampResponse{}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return 0, err
	}

	if data.Message != "OK" {
		return 0, fmt.Errorf("request failed with message %s", data.Message)
	}

	return strconv.ParseUint(data.Result, 10, 0)
}
