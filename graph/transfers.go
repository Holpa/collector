package graph

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	TransfersGraphClient struct {
		Graph *graphql.Client
	}

	TransfersFilter struct {
		Direction constants.TransferDirection
		User      string
		MethodId  constants.TransferMethodId
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

var MINTS_QUERY = fmt.Sprintf(`
query($after: Int!, $before: Int!) {
	transfers(
		where: {
			from: "%s",
			timestamp_gt: $after,
			timestamp_lt: $before
		}
		orderBy: timestamp
		orderDirection: asc
		first: 1000
	) {
		from
		to
		methodId
		contract
		amount
		timestamp
	}
}`, constants.NULL_ADDRESS)

var BURNS_QUERY = fmt.Sprintf(`
query($after: Int!, $before: Int!) {
	transfers(
		where: {
			to: "%s",
			timestamp_gt: $after,
			timestamp_lt: $before
		}
		orderBy: timestamp
		orderDirection: asc
		first: 1000
	) {
		from
		to
		methodId
		contract
		amount
		timestamp
	}
}`, constants.NULL_ADDRESS)

var STAKE_TRANSFERS = fmt.Sprintf(`
query($after: Int!, $before: Int!, $methodId: String!) {
	transfers(
		where: {
			contract: "%s",
			methodId: $methodId,
			timestamp_gt: $after,
			timestamp_lt: $before
		},
		orderBy: timestamp,
		orderDirection: asc,
		first: 1000
	) {
		from
		to
		methodId
		contract
		amount
		timestamp
	}
}`, constants.VE_FLY_CONTRACT)

// ----------------------------------------
// Graph responses
// ----------------------------------------

type (
	TransferGraph struct {
		From      string `json:"from"`
		To        string `json:"to"`
		MethodId  string `json:"methodId"`
		Contract  string `json:"contract"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	}
	TransfersResponse struct {
		Transfers []TransferGraph `json:"transfers"`
	}
	Transfer struct {
		From      string
		To        string
		MethodId  string
		Contract  string
		Amount    *big.Int
		Timestamp time.Time
	}
)

// ----------------------------------------
// Graph request filters
// ----------------------------------------

type (
	TransfersClaimedFilter struct {
		Adventure constants.Adventure
		User      string
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------

func parseTransfer(transferGraph TransferGraph) Transfer {
	return Transfer{
		From:      transferGraph.From,
		To:        transferGraph.To,
		MethodId:  transferGraph.MethodId,
		Contract:  transferGraph.Contract,
		Amount:    ParseBigInt(transferGraph.Amount),
		Timestamp: time.Unix(int64(ParseUInt(transferGraph.Timestamp)), 0),
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *TransfersGraphClient) FetchTransfers(req *graphql.Request, from time.Time, to time.Time) ([]Transfer, error) {
	transfers := []Transfer{}

	queryAfterTs := from
	for {
		req.Var("after", queryAfterTs.Unix())
		req.Var("before", to.Unix())

		res := &TransfersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return []Transfer{}, err
		}

		maxTimestamp := time.Unix(0, 0)
		for _, transferGraph := range res.Transfers {
			transfer := parseTransfer(transferGraph)
			transfers = append(transfers, transfer)

			queryAfterTs = transfer.Timestamp

			if transfer.Timestamp.After(maxTimestamp) {
				maxTimestamp = transfer.Timestamp
			}
		}

		if len(res.Transfers) < 1000 {
			break
		}

		if maxTimestamp.After(to) {
			break
		}
	}

	return transfers, nil
}

// REFACTOR Why? Ugly + hard to grasp 🤡
// Maybe refactor so the caller has to use `FetchTransfers` directly and create query on it's own
func (client *TransfersGraphClient) FetchFilteredTransfers(filter TransfersFilter) ([]Transfer, error) {
	parameters := map[string]string{
		"$before": "Int!",
	}
	requestVars := map[string]string{}
	where := map[string]string{
		"timestamp_lt": "$before",
	}

	if filter.User != "" {
		if filter.Direction == constants.TransferDirectionToUser {
			parameters["$to"] = "String!"
			requestVars["to"] = filter.User
			where["to"] = "$to"
		} else if filter.Direction == constants.TransferDirectionFromUser {
			parameters["$from"] = "String!"
			requestVars["from"] = filter.User
			where["from"] = "$from"
		}
	}
	if filter.MethodId != constants.METHOD_ID_ANY {
		parameters["$methodId"] = "String!"
		requestVars["methodId"] = string(filter.MethodId)
		where["methodId"] = "$methodId"
	}

	urlifiedQuery := ""
	for key, value := range parameters {
		urlifiedQuery += fmt.Sprintf("%s: %s,", key, value)
	}

	urlifiedWhere := ""
	for key, value := range where {
		urlifiedWhere += fmt.Sprintf("%s: %s,", key, value)
	}

	query := fmt.Sprintf(`
		query(%s) {
			transfers(
				where: {%s},
				orderBy: timestamp,
				orderDirection: desc,
				first: 1000
			) {
				from
				to
				methodId
				contract
				amount
				timestamp
			}
		}
	`, urlifiedQuery, urlifiedWhere)

	req := graphql.NewRequest(query)

	for key, value := range requestVars {
		req.Var(key, value)
	}
	return client.FetchTransfers(req, time.Unix(0, 0), time.Now())
}

func (client *TransfersGraphClient) FetchDeposits(from time.Time, to time.Time) ([]Transfer, error) {
	req := graphql.NewRequest(STAKE_TRANSFERS)
	req.Var("methodId", constants.METHOD_ID_FLY_STAKE_DEPOSIT)
	return client.FetchTransfers(req, from, to)
}

func (client *TransfersGraphClient) FetchWithdraws(from time.Time, to time.Time) ([]Transfer, error) {
	req := graphql.NewRequest(STAKE_TRANSFERS)
	req.Var("methodId", constants.METHOD_ID_FLY_STAKE_WITHDRAW)
	return client.FetchTransfers(req, from, to)
}

func (client *TransfersGraphClient) FetchMints(from time.Time, to time.Time) ([]Transfer, error) {
	req := graphql.NewRequest(MINTS_QUERY)
	return client.FetchTransfers(req, from, to)
}

func (client *TransfersGraphClient) FetchBurns(from time.Time, to time.Time) ([]Transfer, error) {
	req := graphql.NewRequest(BURNS_QUERY)
	return client.FetchTransfers(req, from, to)
}
