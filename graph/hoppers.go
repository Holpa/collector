package graph

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-collector/constants"
	"github.com/steschwa/hopper-analytics-collector/models"
)

type (
	HoppersGraphClient struct {
		Graph *graphql.Client
	}
)

func NewHoppersGraphClient() *HoppersGraphClient {
	return &HoppersGraphClient{
		Graph: graphql.NewClient(constants.HOPPERS_GRAPH_URL),
	}
}

// ----------------------------------------
// Queries
// ----------------------------------------

const GET_HOPPERS_QUERY = `
query($skip: Int!) {
	hopperNFTs(
		first: 1000,
		skip: $skip,
		orderBy: tokenId,
		orderDirection: asc
	) {
		tokenId
		strength
		agility
		vitality
		intelligence
		fertility
		level
		adventure
		image
		chainOwner
		listings {
			enabled
			sold
			price
		}
	}
}`

// ----------------------------------------
// Graph responses
// ----------------------------------------

type (
	HopperGraph struct {
		TokenId      string               `json:"tokenId"`
		Strength     string               `json:"strength"`
		Agility      string               `json:"agility"`
		Vitality     string               `json:"vitality"`
		Intelligence string               `json:"intelligence"`
		Fertility    string               `json:"fertility"`
		Level        string               `json:"level"`
		Adventure    bool                 `json:"adventure"`
		Image        string               `json:"image"`
		ChainOwner   string               `json:"chainOwner"`
		Listings     []HopperListingGraph `json:"listings"`
	}

	HopperListingGraph struct {
		Enabled bool   `json:"enabled"`
		Sold    bool   `json:"sold"`
		Price   string `json:"price"`
	}

	HoppersResponse struct {
		HopperNFTs []HopperGraph `json:"hopperNFTs"`
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------

func parseHopper(hopperGraph HopperGraph) models.Hopper {
	listings := make([]models.Listing, len(hopperGraph.Listings))
	for i, listing := range hopperGraph.Listings {
		listings[i] = parseHopperListing(listing)
	}

	return models.Hopper{
		TokenId:      hopperGraph.TokenId,
		Strength:     ParseInt(hopperGraph.Strength),
		Agility:      ParseInt(hopperGraph.Agility),
		Vitality:     ParseInt(hopperGraph.Vitality),
		Intelligence: ParseInt(hopperGraph.Intelligence),
		Fertility:    ParseInt(hopperGraph.Fertility),
		Level:        ParseInt(hopperGraph.Level),
		Adventure:    hopperGraph.Adventure,
		ChainOwner:   hopperGraph.ChainOwner,
		Image:        hopperGraph.Image,
		Listings:     listings,
	}
}

func parseHopperListing(listingGraph HopperListingGraph) models.Listing {
	return models.Listing{
		Enabled: listingGraph.Enabled,
		Sold:    listingGraph.Sold,
		Price:   ParseBigFloat(listingGraph.Price),
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *HoppersGraphClient) FetchAllHoppers() ([]models.Hopper, error) {
	hoppers := make([]models.Hopper, 0)

	for i := 0; i <= constants.HOPPERS_TOTAL_SUPPLY; i += 1000 {
		req := graphql.NewRequest(GET_HOPPERS_QUERY)
		req.Var("skip", i)

		res := &HoppersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return []models.Hopper{}, err
		}

		for _, hopper := range res.HopperNFTs {
			hoppers = append(hoppers, parseHopper(hopper))
		}
	}

	return hoppers, nil
}
