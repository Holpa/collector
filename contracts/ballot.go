// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BallotMetaData contains all meta data concerning the Ballot contract.
var BallotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_flyAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veFlyAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotEnoughVeFly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VEFLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_zones\",\"type\":\"address[]\"}],\"name\":\"addZones\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"arrZones\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEmissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"forceUnvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_countRewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"openBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeZone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"setBonusEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_countRewardRate\",\"type\":\"uint256\"}],\"name\":\"setCountRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vefly\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userVeFlyUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vefly\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"zones\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"zonesUserVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"zonesVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BallotABI is the input ABI used to generate the binding from.
// Deprecated: Use BallotMetaData.ABI instead.
var BallotABI = BallotMetaData.ABI

// Ballot is an auto generated Go binding around an Ethereum contract.
type Ballot struct {
	BallotCaller     // Read-only binding to the contract
	BallotTransactor // Write-only binding to the contract
	BallotFilterer   // Log filterer for contract events
}

// BallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotSession struct {
	Contract     *Ballot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotCallerSession struct {
	Contract *BallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotTransactorSession struct {
	Contract     *BallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotRaw struct {
	Contract *Ballot // Generic contract binding to access the raw methods on
}

// BallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotCallerRaw struct {
	Contract *BallotCaller // Generic read-only contract binding to access the raw methods on
}

// BallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotTransactorRaw struct {
	Contract *BallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallot creates a new instance of Ballot, bound to a specific deployed contract.
func NewBallot(address common.Address, backend bind.ContractBackend) (*Ballot, error) {
	contract, err := bindBallot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// NewBallotCaller creates a new read-only instance of Ballot, bound to a specific deployed contract.
func NewBallotCaller(address common.Address, caller bind.ContractCaller) (*BallotCaller, error) {
	contract, err := bindBallot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotCaller{contract: contract}, nil
}

// NewBallotTransactor creates a new write-only instance of Ballot, bound to a specific deployed contract.
func NewBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotTransactor, error) {
	contract, err := bindBallot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotTransactor{contract: contract}, nil
}

// NewBallotFilterer creates a new log filterer instance of Ballot, bound to a specific deployed contract.
func NewBallotFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotFilterer, error) {
	contract, err := bindBallot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotFilterer{contract: contract}, nil
}

// bindBallot binds a generic wrapper to an already deployed contract.
func bindBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.BallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transact(opts, method, params...)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_Ballot *BallotCaller) FLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_Ballot *BallotSession) FLY() (common.Address, error) {
	return _Ballot.Contract.FLY(&_Ballot.CallOpts)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_Ballot *BallotCallerSession) FLY() (common.Address, error) {
	return _Ballot.Contract.FLY(&_Ballot.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x0a661e06.
//
// Solidity: function VEFLY() view returns(address)
func (_Ballot *BallotCaller) VEFLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "VEFLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEFLY is a free data retrieval call binding the contract method 0x0a661e06.
//
// Solidity: function VEFLY() view returns(address)
func (_Ballot *BallotSession) VEFLY() (common.Address, error) {
	return _Ballot.Contract.VEFLY(&_Ballot.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x0a661e06.
//
// Solidity: function VEFLY() view returns(address)
func (_Ballot *BallotCallerSession) VEFLY() (common.Address, error) {
	return _Ballot.Contract.VEFLY(&_Ballot.CallOpts)
}

// ArrZones is a free data retrieval call binding the contract method 0x7cf8e327.
//
// Solidity: function arrZones(uint256 ) view returns(address)
func (_Ballot *BallotCaller) ArrZones(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "arrZones", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArrZones is a free data retrieval call binding the contract method 0x7cf8e327.
//
// Solidity: function arrZones(uint256 ) view returns(address)
func (_Ballot *BallotSession) ArrZones(arg0 *big.Int) (common.Address, error) {
	return _Ballot.Contract.ArrZones(&_Ballot.CallOpts, arg0)
}

// ArrZones is a free data retrieval call binding the contract method 0x7cf8e327.
//
// Solidity: function arrZones(uint256 ) view returns(address)
func (_Ballot *BallotCallerSession) ArrZones(arg0 *big.Int) (common.Address, error) {
	return _Ballot.Contract.ArrZones(&_Ballot.CallOpts, arg0)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_Ballot *BallotCaller) BonusEmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "bonusEmissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_Ballot *BallotSession) BonusEmissionRate() (*big.Int, error) {
	return _Ballot.Contract.BonusEmissionRate(&_Ballot.CallOpts)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_Ballot *BallotCallerSession) BonusEmissionRate() (*big.Int, error) {
	return _Ballot.Contract.BonusEmissionRate(&_Ballot.CallOpts)
}

// CountReward is a free data retrieval call binding the contract method 0xdbb18109.
//
// Solidity: function countReward() view returns(uint256)
func (_Ballot *BallotCaller) CountReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "countReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountReward is a free data retrieval call binding the contract method 0xdbb18109.
//
// Solidity: function countReward() view returns(uint256)
func (_Ballot *BallotSession) CountReward() (*big.Int, error) {
	return _Ballot.Contract.CountReward(&_Ballot.CallOpts)
}

// CountReward is a free data retrieval call binding the contract method 0xdbb18109.
//
// Solidity: function countReward() view returns(uint256)
func (_Ballot *BallotCallerSession) CountReward() (*big.Int, error) {
	return _Ballot.Contract.CountReward(&_Ballot.CallOpts)
}

// CountRewardRate is a free data retrieval call binding the contract method 0xc1f6d2a4.
//
// Solidity: function countRewardRate() view returns(uint256)
func (_Ballot *BallotCaller) CountRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "countRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountRewardRate is a free data retrieval call binding the contract method 0xc1f6d2a4.
//
// Solidity: function countRewardRate() view returns(uint256)
func (_Ballot *BallotSession) CountRewardRate() (*big.Int, error) {
	return _Ballot.Contract.CountRewardRate(&_Ballot.CallOpts)
}

// CountRewardRate is a free data retrieval call binding the contract method 0xc1f6d2a4.
//
// Solidity: function countRewardRate() view returns(uint256)
func (_Ballot *BallotCallerSession) CountRewardRate() (*big.Int, error) {
	return _Ballot.Contract.CountRewardRate(&_Ballot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ballot *BallotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ballot *BallotSession) Owner() (common.Address, error) {
	return _Ballot.Contract.Owner(&_Ballot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ballot *BallotCallerSession) Owner() (common.Address, error) {
	return _Ballot.Contract.Owner(&_Ballot.CallOpts)
}

// RewardSnapshot is a free data retrieval call binding the contract method 0x53a3785b.
//
// Solidity: function rewardSnapshot() view returns(uint256)
func (_Ballot *BallotCaller) RewardSnapshot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "rewardSnapshot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardSnapshot is a free data retrieval call binding the contract method 0x53a3785b.
//
// Solidity: function rewardSnapshot() view returns(uint256)
func (_Ballot *BallotSession) RewardSnapshot() (*big.Int, error) {
	return _Ballot.Contract.RewardSnapshot(&_Ballot.CallOpts)
}

// RewardSnapshot is a free data retrieval call binding the contract method 0x53a3785b.
//
// Solidity: function rewardSnapshot() view returns(uint256)
func (_Ballot *BallotCallerSession) RewardSnapshot() (*big.Int, error) {
	return _Ballot.Contract.RewardSnapshot(&_Ballot.CallOpts)
}

// UserVeFlyUsed is a free data retrieval call binding the contract method 0x45ea8d42.
//
// Solidity: function userVeFlyUsed(address ) view returns(uint256)
func (_Ballot *BallotCaller) UserVeFlyUsed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "userVeFlyUsed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserVeFlyUsed is a free data retrieval call binding the contract method 0x45ea8d42.
//
// Solidity: function userVeFlyUsed(address ) view returns(uint256)
func (_Ballot *BallotSession) UserVeFlyUsed(arg0 common.Address) (*big.Int, error) {
	return _Ballot.Contract.UserVeFlyUsed(&_Ballot.CallOpts, arg0)
}

// UserVeFlyUsed is a free data retrieval call binding the contract method 0x45ea8d42.
//
// Solidity: function userVeFlyUsed(address ) view returns(uint256)
func (_Ballot *BallotCallerSession) UserVeFlyUsed(arg0 common.Address) (*big.Int, error) {
	return _Ballot.Contract.UserVeFlyUsed(&_Ballot.CallOpts, arg0)
}

// Zones is a free data retrieval call binding the contract method 0x8f77b6ee.
//
// Solidity: function zones(address ) view returns(bool)
func (_Ballot *BallotCaller) Zones(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "zones", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Zones is a free data retrieval call binding the contract method 0x8f77b6ee.
//
// Solidity: function zones(address ) view returns(bool)
func (_Ballot *BallotSession) Zones(arg0 common.Address) (bool, error) {
	return _Ballot.Contract.Zones(&_Ballot.CallOpts, arg0)
}

// Zones is a free data retrieval call binding the contract method 0x8f77b6ee.
//
// Solidity: function zones(address ) view returns(bool)
func (_Ballot *BallotCallerSession) Zones(arg0 common.Address) (bool, error) {
	return _Ballot.Contract.Zones(&_Ballot.CallOpts, arg0)
}

// ZonesUserVotes is a free data retrieval call binding the contract method 0xbfd3c604.
//
// Solidity: function zonesUserVotes(address , address ) view returns(uint256)
func (_Ballot *BallotCaller) ZonesUserVotes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "zonesUserVotes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZonesUserVotes is a free data retrieval call binding the contract method 0xbfd3c604.
//
// Solidity: function zonesUserVotes(address , address ) view returns(uint256)
func (_Ballot *BallotSession) ZonesUserVotes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ballot.Contract.ZonesUserVotes(&_Ballot.CallOpts, arg0, arg1)
}

// ZonesUserVotes is a free data retrieval call binding the contract method 0xbfd3c604.
//
// Solidity: function zonesUserVotes(address , address ) view returns(uint256)
func (_Ballot *BallotCallerSession) ZonesUserVotes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ballot.Contract.ZonesUserVotes(&_Ballot.CallOpts, arg0, arg1)
}

// ZonesVotes is a free data retrieval call binding the contract method 0xbb68eb8b.
//
// Solidity: function zonesVotes(address ) view returns(uint256)
func (_Ballot *BallotCaller) ZonesVotes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "zonesVotes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZonesVotes is a free data retrieval call binding the contract method 0xbb68eb8b.
//
// Solidity: function zonesVotes(address ) view returns(uint256)
func (_Ballot *BallotSession) ZonesVotes(arg0 common.Address) (*big.Int, error) {
	return _Ballot.Contract.ZonesVotes(&_Ballot.CallOpts, arg0)
}

// ZonesVotes is a free data retrieval call binding the contract method 0xbb68eb8b.
//
// Solidity: function zonesVotes(address ) view returns(uint256)
func (_Ballot *BallotCallerSession) ZonesVotes(arg0 common.Address) (*big.Int, error) {
	return _Ballot.Contract.ZonesVotes(&_Ballot.CallOpts, arg0)
}

// AddZones is a paid mutator transaction binding the contract method 0x0e8040a4.
//
// Solidity: function addZones(address[] _zones) returns()
func (_Ballot *BallotTransactor) AddZones(opts *bind.TransactOpts, _zones []common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "addZones", _zones)
}

// AddZones is a paid mutator transaction binding the contract method 0x0e8040a4.
//
// Solidity: function addZones(address[] _zones) returns()
func (_Ballot *BallotSession) AddZones(_zones []common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.AddZones(&_Ballot.TransactOpts, _zones)
}

// AddZones is a paid mutator transaction binding the contract method 0x0e8040a4.
//
// Solidity: function addZones(address[] _zones) returns()
func (_Ballot *BallotTransactorSession) AddZones(_zones []common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.AddZones(&_Ballot.TransactOpts, _zones)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xa97b8b4d.
//
// Solidity: function closeBallot() returns()
func (_Ballot *BallotTransactor) CloseBallot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "closeBallot")
}

// CloseBallot is a paid mutator transaction binding the contract method 0xa97b8b4d.
//
// Solidity: function closeBallot() returns()
func (_Ballot *BallotSession) CloseBallot() (*types.Transaction, error) {
	return _Ballot.Contract.CloseBallot(&_Ballot.TransactOpts)
}

// CloseBallot is a paid mutator transaction binding the contract method 0xa97b8b4d.
//
// Solidity: function closeBallot() returns()
func (_Ballot *BallotTransactorSession) CloseBallot() (*types.Transaction, error) {
	return _Ballot.Contract.CloseBallot(&_Ballot.TransactOpts)
}

// Count is a paid mutator transaction binding the contract method 0x06661abd.
//
// Solidity: function count() returns()
func (_Ballot *BallotTransactor) Count(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "count")
}

// Count is a paid mutator transaction binding the contract method 0x06661abd.
//
// Solidity: function count() returns()
func (_Ballot *BallotSession) Count() (*types.Transaction, error) {
	return _Ballot.Contract.Count(&_Ballot.TransactOpts)
}

// Count is a paid mutator transaction binding the contract method 0x06661abd.
//
// Solidity: function count() returns()
func (_Ballot *BallotTransactorSession) Count() (*types.Transaction, error) {
	return _Ballot.Contract.Count(&_Ballot.TransactOpts)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address _user) returns()
func (_Ballot *BallotTransactor) ForceUnvote(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "forceUnvote", _user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address _user) returns()
func (_Ballot *BallotSession) ForceUnvote(_user common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.ForceUnvote(&_Ballot.TransactOpts, _user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address _user) returns()
func (_Ballot *BallotTransactorSession) ForceUnvote(_user common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.ForceUnvote(&_Ballot.TransactOpts, _user)
}

// OpenBallot is a paid mutator transaction binding the contract method 0x3e0990c1.
//
// Solidity: function openBallot(uint256 _countRewardRate, uint256 _bonusEmissionRate) returns()
func (_Ballot *BallotTransactor) OpenBallot(opts *bind.TransactOpts, _countRewardRate *big.Int, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "openBallot", _countRewardRate, _bonusEmissionRate)
}

// OpenBallot is a paid mutator transaction binding the contract method 0x3e0990c1.
//
// Solidity: function openBallot(uint256 _countRewardRate, uint256 _bonusEmissionRate) returns()
func (_Ballot *BallotSession) OpenBallot(_countRewardRate *big.Int, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.OpenBallot(&_Ballot.TransactOpts, _countRewardRate, _bonusEmissionRate)
}

// OpenBallot is a paid mutator transaction binding the contract method 0x3e0990c1.
//
// Solidity: function openBallot(uint256 _countRewardRate, uint256 _bonusEmissionRate) returns()
func (_Ballot *BallotTransactorSession) OpenBallot(_countRewardRate *big.Int, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.OpenBallot(&_Ballot.TransactOpts, _countRewardRate, _bonusEmissionRate)
}

// RemoveZone is a paid mutator transaction binding the contract method 0xfe8e362a.
//
// Solidity: function removeZone(uint256 index) returns()
func (_Ballot *BallotTransactor) RemoveZone(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "removeZone", index)
}

// RemoveZone is a paid mutator transaction binding the contract method 0xfe8e362a.
//
// Solidity: function removeZone(uint256 index) returns()
func (_Ballot *BallotSession) RemoveZone(index *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.RemoveZone(&_Ballot.TransactOpts, index)
}

// RemoveZone is a paid mutator transaction binding the contract method 0xfe8e362a.
//
// Solidity: function removeZone(uint256 index) returns()
func (_Ballot *BallotTransactorSession) RemoveZone(index *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.RemoveZone(&_Ballot.TransactOpts, index)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_Ballot *BallotTransactor) SetBonusEmissionRate(opts *bind.TransactOpts, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "setBonusEmissionRate", _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_Ballot *BallotSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.SetBonusEmissionRate(&_Ballot.TransactOpts, _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_Ballot *BallotTransactorSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.SetBonusEmissionRate(&_Ballot.TransactOpts, _bonusEmissionRate)
}

// SetCountRewardRate is a paid mutator transaction binding the contract method 0xcb9bf6e5.
//
// Solidity: function setCountRewardRate(uint256 _countRewardRate) returns()
func (_Ballot *BallotTransactor) SetCountRewardRate(opts *bind.TransactOpts, _countRewardRate *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "setCountRewardRate", _countRewardRate)
}

// SetCountRewardRate is a paid mutator transaction binding the contract method 0xcb9bf6e5.
//
// Solidity: function setCountRewardRate(uint256 _countRewardRate) returns()
func (_Ballot *BallotSession) SetCountRewardRate(_countRewardRate *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.SetCountRewardRate(&_Ballot.TransactOpts, _countRewardRate)
}

// SetCountRewardRate is a paid mutator transaction binding the contract method 0xcb9bf6e5.
//
// Solidity: function setCountRewardRate(uint256 _countRewardRate) returns()
func (_Ballot *BallotTransactorSession) SetCountRewardRate(_countRewardRate *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.SetCountRewardRate(&_Ballot.TransactOpts, _countRewardRate)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Ballot *BallotTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Ballot *BallotSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.SetOwner(&_Ballot.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Ballot *BallotTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.SetOwner(&_Ballot.TransactOpts, _owner)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address user, uint256 vefly) returns()
func (_Ballot *BallotTransactor) Unvote(opts *bind.TransactOpts, user common.Address, vefly *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "unvote", user, vefly)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address user, uint256 vefly) returns()
func (_Ballot *BallotSession) Unvote(user common.Address, vefly *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Unvote(&_Ballot.TransactOpts, user, vefly)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(address user, uint256 vefly) returns()
func (_Ballot *BallotTransactorSession) Unvote(user common.Address, vefly *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Unvote(&_Ballot.TransactOpts, user, vefly)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address user, uint256 vefly) returns()
func (_Ballot *BallotTransactor) Vote(opts *bind.TransactOpts, user common.Address, vefly *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "vote", user, vefly)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address user, uint256 vefly) returns()
func (_Ballot *BallotSession) Vote(user common.Address, vefly *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, user, vefly)
}

// Vote is a paid mutator transaction binding the contract method 0x5f74bbde.
//
// Solidity: function vote(address user, uint256 vefly) returns()
func (_Ballot *BallotTransactorSession) Vote(user common.Address, vefly *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, user, vefly)
}

// BallotUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the Ballot contract.
type BallotUpdatedOwnerIterator struct {
	Event *BallotUpdatedOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BallotUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotUpdatedOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BallotUpdatedOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BallotUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotUpdatedOwner represents a UpdatedOwner event raised by the Ballot contract.
type BallotUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_Ballot *BallotFilterer) FilterUpdatedOwner(opts *bind.FilterOpts, owner []common.Address) (*BallotUpdatedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ballot.contract.FilterLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &BallotUpdatedOwnerIterator{contract: _Ballot.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_Ballot *BallotFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *BallotUpdatedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ballot.contract.WatchLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotUpdatedOwner)
				if err := _Ballot.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedOwner is a log parse operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_Ballot *BallotFilterer) ParseUpdatedOwner(log types.Log) (*BallotUpdatedOwner, error) {
	event := new(BallotUpdatedOwner)
	if err := _Ballot.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
