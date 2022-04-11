package contracts

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	OnChainClient struct {
		Connection *ethclient.Client
		Cache      *cache.Cache
	}

	ZoneContract interface {
		TotalBaseShare(opts *bind.CallOpts) (*big.Int, error)
		TotalVeShare(opts *bind.CallOpts) (*big.Int, error)
		BaseSharesBalance(opts *bind.CallOpts, address common.Address) (*big.Int, error)
		VeSharesBalance(opts *bind.CallOpts, address common.Address) (*big.Int, error)
		UserMaxFlyGeneration(opts *bind.CallOpts, address common.Address) (*big.Int, error)

		GetUserGeneratedFly(opts *bind.CallOpts, account common.Address, totalUserBaseShares *big.Int) (*big.Int, *big.Int, error)
		GetUserBonusGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error)
	}
)

func NewOnChainClient() (*OnChainClient, error) {
	client, err := ethclient.Dial(constants.AVAX_RPC)
	if err != nil {
		return &OnChainClient{}, err
	}

	return &OnChainClient{
		Connection: client,
		Cache:      cache.New(time.Minute*1, time.Second*30),
	}, nil
}

// ----------------------------------------
// Adventure callers getter
// ----------------------------------------

func (client *OnChainClient) getPondCaller() (*AdventurePondCaller, error) {
	return NewAdventurePondCaller(common.HexToAddress(constants.ADVENTURE_POND_CONTRACT), client.Connection)
}
func (client *OnChainClient) getStreamCaller() (*AdventureStreamCaller, error) {
	return NewAdventureStreamCaller(common.HexToAddress(constants.ADVENTURE_STREAM_CONTRACT), client.Connection)
}
func (client *OnChainClient) getSwampCaller() (*AdventureSwampCaller, error) {
	return NewAdventureSwampCaller(common.HexToAddress(constants.ADVENTURE_SWAMP_CONTRACT), client.Connection)
}
func (client *OnChainClient) getRiverCaller() (*AdventureRiverCaller, error) {
	return NewAdventureRiverCaller(common.HexToAddress(constants.ADVENTURE_RIVER_CONTRACT), client.Connection)
}
func (client *OnChainClient) getForestCaller() (*AdventureForestCaller, error) {
	return NewAdventureForestCaller(common.HexToAddress(constants.ADVENTURE_FOREST_CONTRACT), client.Connection)
}
func (client *OnChainClient) getGreatLakeCaller() (*AdventureGreatLakeCaller, error) {
	return NewAdventureGreatLakeCaller(common.HexToAddress(constants.ADVENTURE_GREAT_LAKE_CONTRACT), client.Connection)
}
func (client *OnChainClient) getBallotCaller() (*BallotCaller, error) {
	return NewBallotCaller(common.HexToAddress(constants.BALLOT_CONTRACT), client.Connection)
}
func (client *OnChainClient) getFlyCaller() (*FlyCaller, error) {
	return NewFlyCaller(common.HexToAddress(constants.FLY_CONTRACT), client.Connection)
}

func (client *OnChainClient) getAdventureCaller(adventure constants.Adventure) (ZoneContract, error) {
	switch adventure {
	case constants.AdventurePond:
		return client.getPondCaller()
	case constants.AdventureStream:
		return client.getStreamCaller()
	case constants.AdventureSwamp:
		return client.getSwampCaller()
	case constants.AdventureRiver:
		return client.getRiverCaller()
	case constants.AdventureForest:
		return client.getForestCaller()
	case constants.AdventureGreatLake:
		return client.getGreatLakeCaller()
	}

	// should never happen
	return nil, errors.New("unknown adventure")
}

func getContractByAdventure(adventure constants.Adventure) string {
	switch adventure {
	case constants.AdventurePond:
		return constants.ADVENTURE_POND_CONTRACT
	case constants.AdventureStream:
		return constants.ADVENTURE_STREAM_CONTRACT
	case constants.AdventureSwamp:
		return constants.ADVENTURE_SWAMP_CONTRACT
	case constants.AdventureRiver:
		return constants.ADVENTURE_RIVER_CONTRACT
	case constants.AdventureForest:
		return constants.ADVENTURE_FOREST_CONTRACT
	case constants.AdventureGreatLake:
		return constants.ADVENTURE_GREAT_LAKE_CONTRACT
	}

	return ""
}

// ----------------------------------------
// Contract read wrappers
// ----------------------------------------

func (client *OnChainClient) GetTotalBaseShares(adventure constants.Adventure) (*big.Int, error) {
	cacheKey := fmt.Sprintf("%s.total-base-shares", adventure)

	if total, found := client.Cache.Get(cacheKey); found {
		return total.(*big.Int), nil
	}

	caller, err := client.getAdventureCaller(adventure)
	if err != nil {
		return big.NewInt(0), err
	}

	totalBaseShares, err := caller.TotalBaseShare(nil)
	if err != nil {
		return big.NewInt(0), err
	}

	client.Cache.Set(cacheKey, totalBaseShares, cache.DefaultExpiration)
	return totalBaseShares, nil
}

func (client *OnChainClient) GetVotesByAdventure(adventure constants.Adventure) (*big.Int, error) {
	cacheKey := fmt.Sprintf("%s.total-votes", adventure)

	if total, found := client.Cache.Get(cacheKey); found {
		return total.(*big.Int), nil
	}

	caller, err := client.getBallotCaller()
	if err != nil {
		return big.NewInt(0), err
	}

	adventureContract := getContractByAdventure(adventure)

	votes, err := caller.ZonesVotes(nil, common.HexToAddress(adventureContract))
	if err != nil {
		return big.NewInt(0), err
	}

	client.Cache.Set(cacheKey, votes, cache.DefaultExpiration)
	return votes, nil
}

func (client *OnChainClient) GetFlySupply() (*big.Int, error) {
	caller, err := client.getFlyCaller()
	if err != nil {
		return big.NewInt(0), err
	}

	return caller.TotalSupply(nil)
}

func (client *OnChainClient) GetUserBaseSharesBalance(adventure constants.Adventure, user string) (*big.Int, error) {
	cacheKey := fmt.Sprintf("%s.%s.baseshares-balance", adventure, user)

	if baseSharesBalance, found := client.Cache.Get(cacheKey); found {
		return baseSharesBalance.(*big.Int), nil
	}

	caller, err := client.getAdventureCaller(adventure)
	if err != nil {
		return big.NewInt(0), err
	}

	baseSharesBalance, err := caller.BaseSharesBalance(nil, common.HexToAddress(user))
	if err != nil {
		return big.NewInt(0), err
	}

	client.Cache.Set(cacheKey, baseSharesBalance, cache.DefaultExpiration)
	return baseSharesBalance, nil
}

func (client *OnChainClient) GetUserVeShareBalance(adventure constants.Adventure, user string) (*big.Int, error) {
	cacheKey := fmt.Sprintf("%s.%s.veshares-balance", adventure, user)

	if veShareBalance, found := client.Cache.Get(cacheKey); found {
		return veShareBalance.(*big.Int), nil
	}

	caller, err := client.getAdventureCaller(adventure)
	if err != nil {
		return big.NewInt(0), err
	}

	veShareBalance, err := caller.VeSharesBalance(nil, common.HexToAddress(user))
	if err != nil {
		return big.NewInt(0), err
	}

	client.Cache.Set(cacheKey, veShareBalance, cache.DefaultExpiration)
	return veShareBalance, nil
}

func (client *OnChainClient) GetUserMaxFlyGeneration(adventure constants.Adventure, user string) (*big.Int, error) {
	cacheKey := fmt.Sprintf("%s.%s.user-max-fly-generation", adventure, user)

	if maxFlyGeneration, found := client.Cache.Get(cacheKey); found {
		return maxFlyGeneration.(*big.Int), nil
	}

	caller, err := client.getAdventureCaller(adventure)
	if err != nil {
		return big.NewInt(0), err
	}

	maxFlyGeneration, err := caller.UserMaxFlyGeneration(nil, common.HexToAddress(user))
	if err != nil {
		return big.NewInt(0), err
	}

	client.Cache.Set(cacheKey, maxFlyGeneration, cache.DefaultExpiration)
	return maxFlyGeneration, nil
}

func (client *OnChainClient) GetUserGeneratedFly(adventure constants.Adventure, user string) (*big.Int, *big.Int, error) {
	caller, err := client.getAdventureCaller(adventure)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	baseSharesBalance, err := client.GetUserBaseSharesBalance(adventure, user)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return caller.GetUserGeneratedFly(nil, common.HexToAddress(user), baseSharesBalance)
}

func (client *OnChainClient) GetUserBonusGeneratedFly(adventure constants.Adventure, user string) (*big.Int, *big.Int, error) {
	caller, err := client.getAdventureCaller(adventure)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	baseSharesBalance, err := client.GetUserVeShareBalance(adventure, user)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return caller.GetUserBonusGeneratedFly(nil, common.HexToAddress(user), baseSharesBalance)
}
