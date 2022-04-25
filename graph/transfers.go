package graph

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/machinebox/graphql"
	"github.com/shopspring/decimal"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

const (
	METHOD_ID_DEPOSIT  = "0xb6b55f25"
	METHOD_ID_WITHDRAW = "0x2e1a7d4d"
)

type (
	TransfersGraphClient struct {
		Graph *graphql.Client
	}
)

func NewTransfersGraphClient() *TransfersGraphClient {
	return &TransfersGraphClient{
		Graph: graphql.NewClient(constants.FLY_TRANSFERS_GRAPH_URL),
	}
}

// ----------------------------------------
// Queries
// ----------------------------------------

var GET_STAKED_TRANSFERS = fmt.Sprintf(`
query($before: Int!, $methodId: String!) {
	transfers(
		where: {
			contract: "%s",
			methodId: $methodId,
			timestamp_lt: $before
		},
		orderBy: timestamp,
		orderDirection: desc,
		first: 1000
	) {
		amount
		timestamp
	}
}`, constants.VE_FLY_CONTRACT)

// ----------------------------------------
// Graph responses
// ----------------------------------------

type (
	TransferGraph struct {
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	}

	TransfersResponse struct {
		Transfers []TransferGraph `json:"transfers"`
	}

	Transfer struct {
		Amount    *big.Int
		Timestamp time.Time
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------

func parseTransfer(transferGraph TransferGraph) Transfer {
	return Transfer{
		Amount:    ParseBigInt(transferGraph.Amount),
		Timestamp: time.Unix(int64(ParseUInt(transferGraph.Timestamp)), 0),
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *TransfersGraphClient) FetchTotalDeposited() (decimal.Decimal, error) {
	total := decimal.NewFromInt(0)

	queryBeforeTs := time.Now()
	for {
		unixTs := queryBeforeTs.Unix()
		req := graphql.NewRequest(GET_STAKED_TRANSFERS)
		req.Var("before", unixTs)
		req.Var("methodId", METHOD_ID_DEPOSIT)

		res := &TransfersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return decimal.NewFromInt(0), err
		}

		for _, transferGraph := range res.Transfers {
			transfer := parseTransfer(transferGraph)

			amountDec, err := decimal.NewFromString(transfer.Amount.String())
			if err != nil {
				return decimal.NewFromInt(0), err
			}
			total = total.Add(amountDec)
			queryBeforeTs = transfer.Timestamp
		}

		if len(res.Transfers) < 1000 {
			break
		}
	}

	return total, nil
}

func (client *TransfersGraphClient) FetchTotalWithdrawn() (decimal.Decimal, error) {
	total := decimal.NewFromInt(0)

	queryBeforeTs := time.Now()
	for {
		unixTs := queryBeforeTs.Unix()
		req := graphql.NewRequest(GET_STAKED_TRANSFERS)
		req.Var("before", unixTs)
		req.Var("methodId", METHOD_ID_WITHDRAW)

		res := &TransfersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return decimal.NewFromInt(0), err
		}

		for _, transferGraph := range res.Transfers {
			transfer := parseTransfer(transferGraph)

			amountDec, err := decimal.NewFromString(transfer.Amount.String())
			if err != nil {
				return decimal.NewFromInt(0), err
			}
			total = total.Add(amountDec)
			queryBeforeTs = transfer.Timestamp
		}

		if len(res.Transfers) < 1000 {
			break
		}
	}

	return total, nil
}
