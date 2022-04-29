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
// Query filters
// ----------------------------------------

type (
	TransfersFilter struct {
		Direction constants.TransferDirection
		User      string
		MethodId  constants.TransferMethodId
	}
)

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *TransfersGraphClient) FetchTransfers(req *graphql.Request) ([]Transfer, error) {
	transfers := []Transfer{}

	queryBeforeTs := time.Now()
	for {
		unixTs := queryBeforeTs.Unix()
		req.Var("before", unixTs)

		res := &TransfersResponse{}
		if err := client.Graph.Run(context.Background(), req, res); err != nil {
			return []Transfer{}, err
		}

		for _, transferGraph := range res.Transfers {
			transfer := parseTransfer(transferGraph)
			transfers = append(transfers, transfer)

			queryBeforeTs = transfer.Timestamp
		}

		if len(res.Transfers) < 1000 {
			break
		}
	}

	return transfers, nil
}

func (client *TransfersGraphClient) FetchDepositTransfers() ([]Transfer, error) {
	req := graphql.NewRequest(GET_STAKED_TRANSFERS)
	req.Var("methodId", constants.METHOD_ID_FLY_STAKE_DEPOSIT)
	return client.FetchTransfers(req)
}

func (client *TransfersGraphClient) FetchTotalDeposited() (decimal.Decimal, error) {
	depositTransfers, err := client.FetchDepositTransfers()
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	total := decimal.NewFromInt(0)
	for _, transfer := range depositTransfers {
		amountDec, err := decimal.NewFromString(transfer.Amount.String())
		if err != nil {
			return decimal.NewFromInt(0), err
		}
		total = total.Add(amountDec)
	}

	return total, nil
}

func (client *TransfersGraphClient) FetchWithdrawTransfers() ([]Transfer, error) {
	req := graphql.NewRequest(GET_STAKED_TRANSFERS)
	req.Var("methodId", constants.METHOD_ID_FLY_STAKE_WITHDRAW)
	return client.FetchTransfers(req)
}

func (client *TransfersGraphClient) FetchTotalWithdrawn() (decimal.Decimal, error) {
	withdrawTransfers, err := client.FetchWithdrawTransfers()
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	total := decimal.NewFromInt(0)
	for _, transfer := range withdrawTransfers {
		amountDec, err := decimal.NewFromString(transfer.Amount.String())
		if err != nil {
			return decimal.NewFromInt(0), err
		}
		total = total.Add(amountDec)
	}

	return total, nil
}

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
	return client.FetchTransfers(req)
}
