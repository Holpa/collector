package graph

import (
	"context"
	"math/big"
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
query latestMintsDayDatas(
	$startTimestamp: BigInt!
) {
	flyMintsDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1
		where: {
			startTimestamp: $startTimestamp
		}
	) {
		id
		startTimestamp
		minted
	}
}`

var BURNS_DAY_DATAS = `
query latestBurnsDayDatas(
	$startTimestamp: BigInt!
) {
	flyBurnsDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1
		where: {
			startTimestamp: $startTimestamp
		}
	) {
		id
		startTimestamp
		burned
	}
}`

var STAKE_DEPOSIT_DAY_DATAS = `
query latestStakeDepositDayDatas(
	$startTimestamp: BigInt!
) {
	flyStakeDepositDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1
		where: {
			startTimestamp: $startTimestamp
		}
	) {
		id
		startTimestamp
		deposited
	}
}`

var STAKE_WITHDRAW_DAY_DATAS = `
query latestWithdrawDayDatas(
	$startTimestamp: BigInt!
) {
	flyStakeWithdrawDayDatas(
		orderBy: id
		orderDirection: asc
		first: 1
		where: {
			startTimestamp: $startTimestamp
		}
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
	DayData struct {
		Id             uint
		StartTimestamp time.Time
	}

	MintDayDataGraph struct {
		DayDataGraph
		Minted string `json:"minted"`
	}
	MintDayData struct {
		DayData
		Minted *big.Int
	}
	MintsDayDataResponse struct {
		FlyMintsDayDatas []MintDayDataGraph `json:"flyMintsDayDatas"`
	}

	BurnDayDataGraph struct {
		DayDataGraph
		Burned string `json:"burned"`
	}
	BurnDayData struct {
		DayData
		Burned *big.Int
	}
	BurnsDayDataResponse struct {
		FlyBurnsDayDatas []BurnDayDataGraph `json:"flyBurnsDayDatas"`
	}

	StakeDepositDayDataGraph struct {
		DayDataGraph
		Deposited string `json:"deposited"`
	}
	StakeDepositDayData struct {
		DayData
		Deposited *big.Int
	}
	StakeDepositDayDataResponse struct {
		FlyStakeDepositDayDatas []StakeDepositDayDataGraph `json:"flyStakeDepositDayDatas"`
	}

	StakeWithdrawDayDataGraph struct {
		DayDataGraph
		Withdrawn string `json:"withdrawn"`
	}
	StakeWithdrawDayData struct {
		DayData
		Withdrawn *big.Int
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

func (client *FlyGraphClient) FetchMintsDayData(startTimestamp time.Time) (MintDayData, error) {
	req := graphql.NewRequest(MINTS_DAY_DATAS)
	req.Var("startTimestamp", startTimestamp.Unix())

	res := MintsDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return MintDayData{}, err
	}

	dayDatas := make([]MintDayData, len(res.FlyMintsDayDatas))
	for i, graph := range res.FlyMintsDayDatas {
		dayDatas[i] = parseMintDayData(graph)
	}

	if len(dayDatas) == 0 {
		return MintDayData{}, nil
	}

	return dayDatas[0], nil
}

func (client *FlyGraphClient) FetchBurnsDayData(startTimestamp time.Time) (BurnDayData, error) {
	req := graphql.NewRequest(BURNS_DAY_DATAS)
	req.Var("startTimestamp", startTimestamp.Unix())

	res := BurnsDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return BurnDayData{}, err
	}

	dayDatas := make([]BurnDayData, len(res.FlyBurnsDayDatas))
	for i, graph := range res.FlyBurnsDayDatas {
		dayDatas[i] = parseBurnDayData(graph)
	}

	if len(dayDatas) == 0 {
		return BurnDayData{}, nil
	}

	return dayDatas[0], nil
}

func (client *FlyGraphClient) FetchStakeDepositDayData(startTimestamp time.Time) (StakeDepositDayData, error) {
	req := graphql.NewRequest(STAKE_DEPOSIT_DAY_DATAS)
	req.Var("startTimestamp", startTimestamp.Unix())

	res := StakeDepositDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return StakeDepositDayData{}, err
	}

	dayDatas := make([]StakeDepositDayData, len(res.FlyStakeDepositDayDatas))
	for i, graph := range res.FlyStakeDepositDayDatas {
		dayDatas[i] = parseStakeDepositDayData(graph)
	}

	if len(dayDatas) == 0 {
		return StakeDepositDayData{}, nil
	}

	return dayDatas[0], nil
}

func (client *FlyGraphClient) FetchStakeWithdrawDayData(startTimestamp time.Time) (StakeWithdrawDayData, error) {
	req := graphql.NewRequest(STAKE_WITHDRAW_DAY_DATAS)
	req.Var("startTimestamp", startTimestamp.Unix())

	res := StakeWithdrawDayDataResponse{}
	if err := client.Graph.Run(context.Background(), req, &res); err != nil {
		return StakeWithdrawDayData{}, err
	}

	dayDatas := make([]StakeWithdrawDayData, len(res.FlyStakeWithdrawDayDatas))
	for i, graph := range res.FlyStakeWithdrawDayDatas {
		dayDatas[i] = parseStakeWithdrawDayData(graph)
	}

	if len(dayDatas) == 0 {
		return StakeWithdrawDayData{}, nil
	}

	return dayDatas[0], nil
}
