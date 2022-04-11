package graph

import (
	"context"
	"time"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
)

type (
	MarketsGraphClient struct {
		Graph *graphql.Client
	}
)

func NewMarketsGraphClient() *MarketsGraphClient {
	return &MarketsGraphClient{
		Graph: graphql.NewClient(constants.HOPPERS_GRAPH_URL),
	}
}

// ----------------------------------------
// Queries
// ----------------------------------------

const GET_LISTINGS_QUERY = `
query($skip: Int!) {
	listings(
		first: 1000,
		skip: $skip,
		orderBy: timestamp,
		orderDirection: asc
	) {
		id
		price
		timestamp
		enabled
		sold
		buyer
		owner
		hopper {
			tokenId
		}
	}
}`

// ----------------------------------------
// Graph responses
// ----------------------------------------

type (
	MarketListingGraph struct {
		Id        string       `json:"id"`
		Price     string       `json:"price"`
		Enabled   bool         `json:"enabled"`
		Sold      bool         `json:"sold"`
		Buyer     string       `json:"buyer"`
		Owner     string       `json:"owner"`
		Timestamp string       `json:"timestamp"`
		Hopper    MarketHopper `json:"hopper"`
	}

	MarketHopper struct {
		TokenId string `json:"tokenId"`
	}

	MarketsResponse struct {
		Listings []MarketListingGraph `json:"listings"`
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------

func parseMarketListing(listingGraph MarketListingGraph) models.Listing {
	return models.Listing{
		Id:        listingGraph.Id,
		Enabled:   listingGraph.Enabled,
		Sold:      listingGraph.Sold,
		Price:     ParseBigFloat(listingGraph.Price),
		Buyer:     listingGraph.Buyer,
		Seller:    listingGraph.Owner,
		Timestamp: time.Unix(int64(ParseInt(listingGraph.Timestamp)), 0),
		HopperId:  listingGraph.Hopper.TokenId,
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *MarketsGraphClient) FetchAllListings() ([]models.Listing, error) {
	listings := make([]models.Listing, 0)

	for skip := 0; true; skip += 1000 {
		req := graphql.NewRequest(GET_LISTINGS_QUERY)
		req.Var("skip", skip)

		res := &MarketsResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return []models.Listing{}, err
		}

		for _, listing := range res.Listings {
			listings = append(listings, parseMarketListing(listing))
		}

		if len(res.Listings) < 1000 {
			break
		}
	}

	return listings, nil
}
