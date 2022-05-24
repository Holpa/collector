package graph

import (
	"context"
	"time"

	"github.com/machinebox/graphql"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	FlyGraphClient struct {
		Graph *graphql.Client
	}
)

func NewFlyGraphClient() *FlyGraphClient {
	return &FlyGraphClient{
		Graph: graphql.NewClient(constants.CUSTOM_GRAPH_URL),
	}
}

// ----------------------------------------
// Queries
// ----------------------------------------

var MINTS_DAY_DATAS = `
query {
	flyMintsDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1000
	) {
		id
		startTimestamp
		minted
	}
}`

var BURNS_DAY_DATAS = `
query {
	flyBurnsDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1000
	) {
		id
		startTimestamp
		burned
	}
}`

var CIRCULATING_DAY_DATAS = `
query {
	flyCirculatingDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1000
	) {
		id
		startTimestamp
		circulating
	}
}`

var STAKE_DEPOSIT_DAY_DATAS = `
query {
	flyStakeDepositDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1000
	) {
		id
		startTimestamp
		deposited
	}
}`

var STAKE_WITHDRAW_DAY_DATAS = `
query {
	flyStakeWithdrawDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1000
	) {
		id
		startTimestamp
		withdrawn
	}
}`

// ----------------------------------------
// Graph responses
// ----------------------------------------

type (
	DayDataGraph struct {
		Id             string `json:"id"`
		StartTimestamp string `json:"startTimestamp"`
	}

	MintDayDataGraph struct {
		DayDataGraph
		Minted string `json:"minted"`
	}
	MintsDayDataResponse struct {
		FlyMintsDayDatas []MintDayDataGraph `json:"flyMintsDayDatas"`
	}

	BurnDayDataGraph struct {
		DayDataGraph
		Burned string `json:"burned"`
	}
	BurnsDayDataResponse struct {
		FlyBurnsDayDatas []BurnDayDataGraph `json:"flyBurnsDayDatas"`
	}

	CirculatingDayDataGraph struct {
		DayDataGraph
		Circulating string `json:"circulating"`
	}
	CirculatingDayDataResponse struct {
		FlyCirculatingDayDatas []CirculatingDayDataGraph `json:"flyCirculatingDayDatas"`
	}

	StakeDepositDayDataGraph struct {
		DayDataGraph
		Deposited string `json:"deposited"`
	}
	StakeDepositDayDataResponse struct {
		FlyStakeDepositDayDatas []StakeDepositDayDataGraph `json:"flyStakeDepositDayDatas"`
	}

	StakeWithdrawDayDataGraph struct {
		DayDataGraph
		Withdrawn string `json:"withdrawn"`
	}
	StakeWithdrawDayDataResponse struct {
		FlyStakeWithdrawDayDatas []StakeWithdrawDayDataGraph `json:"flyStakeWithdrawDayDatas"`
	}
)

// ----------------------------------------
// Graph response converters
// ----------------------------------------

func parseMintDayData(graph MintDayDataGraph) MintDayData {
	return MintDayData{
		DayData: DayData{
			Id:             ParseUInt(graph.Id),
			StartTimestamp: time.Unix(int64(ParseUInt(graph.StartTimestamp)), 0),
		},
		Minted: ParseBigInt(graph.Minted),
	}
}
func parseBurnDayData(graph BurnDayDataGraph) BurnDayData {
	return BurnDayData{
		DayData: DayData{
			Id:             ParseUInt(graph.Id),
			StartTimestamp: time.Unix(int64(ParseUInt(graph.StartTimestamp)), 0),
		},
		Burned: ParseBigInt(graph.Burned),
	}
}
func parseCirculatingDayData(graph CirculatingDayDataGraph) CirculatingDayData {
	return CirculatingDayData{
		DayData: DayData{
			Id:             ParseUInt(graph.Id),
			StartTimestamp: time.Unix(int64(ParseUInt(graph.StartTimestamp)), 0),
		},
		Circulating: ParseBigInt(graph.Circulating),
	}
}
func parseStakeDepositDayData(graph StakeDepositDayDataGraph) StakeDepositDayData {
	return StakeDepositDayData{
		DayData: DayData{
			Id:             ParseUInt(graph.Id),
			StartTimestamp: time.Unix(int64(ParseUInt(graph.StartTimestamp)), 0),
		},
		Deposited: ParseBigInt(graph.Deposited),
	}
}
func parseStakeWithdrawDayData(graph StakeWithdrawDayDataGraph) StakeWithdrawDayData {
	return StakeWithdrawDayData{
		DayData: DayData{
			Id:             ParseUInt(graph.Id),
			StartTimestamp: time.Unix(int64(ParseUInt(graph.StartTimestamp)), 0),
		},
		Withdrawn: ParseBigInt(graph.Withdrawn),
	}
}

// ----------------------------------------
// Query functions
// ----------------------------------------

func (client *FlyGraphClient) FetchMintsDayData() ([]MintDayData, error) {
	req := graphql.NewRequest(MINTS_DAY_DATAS)

	res := MintsDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return []MintDayData{}, err
	}

	dayDatas := make([]MintDayData, len(res.FlyMintsDayDatas))
	for i, graph := range res.FlyMintsDayDatas {
		dayDatas[i] = parseMintDayData(graph)
	}

	return dayDatas, nil
}

func (client *FlyGraphClient) FetchBurnsDayData() ([]BurnDayData, error) {
	req := graphql.NewRequest(BURNS_DAY_DATAS)

	res := BurnsDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return []BurnDayData{}, err
	}

	dayDatas := make([]BurnDayData, len(res.FlyBurnsDayDatas))
	for i, graph := range res.FlyBurnsDayDatas {
		dayDatas[i] = parseBurnDayData(graph)
	}

	return dayDatas, nil
}

func (client *FlyGraphClient) FetchCirculatingDayData() ([]CirculatingDayData, error) {
	req := graphql.NewRequest(CIRCULATING_DAY_DATAS)

	res := CirculatingDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return []CirculatingDayData{}, err
	}

	dayDatas := make([]CirculatingDayData, len(res.FlyCirculatingDayDatas))
	for i, graph := range res.FlyCirculatingDayDatas {
		dayDatas[i] = parseCirculatingDayData(graph)
	}

	return dayDatas, nil
}

func (client *FlyGraphClient) FetchStakeDepositDayData() ([]StakeDepositDayData, error) {
	req := graphql.NewRequest(STAKE_DEPOSIT_DAY_DATAS)

	res := StakeDepositDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return []StakeDepositDayData{}, err
	}

	dayDatas := make([]StakeDepositDayData, len(res.FlyStakeDepositDayDatas))
	for i, graph := range res.FlyStakeDepositDayDatas {
		dayDatas[i] = parseStakeDepositDayData(graph)
	}

	return dayDatas, nil
}

func (client *FlyGraphClient) FetchStakeWithdrawDayData() ([]StakeWithdrawDayData, error) {
	req := graphql.NewRequest(STAKE_WITHDRAW_DAY_DATAS)

	res := StakeWithdrawDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return []StakeWithdrawDayData{}, err
	}

	dayDatas := make([]StakeWithdrawDayData, len(res.FlyStakeWithdrawDayDatas))
	for i, graph := range res.FlyStakeWithdrawDayDatas {
		dayDatas[i] = parseStakeWithdrawDayData(graph)
	}

	return dayDatas, nil
}
