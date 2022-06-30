// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controllers

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

// BiddingwarMetaData contains all meta data concerning the Biddingwar contract.
var BiddingwarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Bid\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"ganeIndex\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"bidder\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"bidAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"bidTime\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GameParamChanged\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"bidEndTime\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"maxBidInterval\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"commission\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GameStarted\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"gameIndex\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"startTime\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReceived\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WinnerGetRewards\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"gameIndex\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"winner\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"bid\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"bidEndTime\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"checkGameRunning\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"commission\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"curBidAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"curBidTime\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"gameIndex\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"lastBidder\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"maxBidInterval\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"rewardToWinner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setGameParams\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_bidEndTime\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_maxBidInterval\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_commission\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"start\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"startBidding\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"startTime\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalWinnerAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"toAddress\",\"internalType\":\"address\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// BiddingwarABI is the input ABI used to generate the binding from.
// Deprecated: Use BiddingwarMetaData.ABI instead.
var BiddingwarABI = BiddingwarMetaData.ABI

// Biddingwar is an auto generated Go binding around an Ethereum contract.
type Biddingwar struct {
	BiddingwarCaller     // Read-only binding to the contract
	BiddingwarTransactor // Write-only binding to the contract
	BiddingwarFilterer   // Log filterer for contract events
}

// BiddingwarCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiddingwarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiddingwarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiddingwarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingwarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiddingwarSession struct {
	Contract     *Biddingwar       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiddingwarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiddingwarCallerSession struct {
	Contract *BiddingwarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BiddingwarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiddingwarTransactorSession struct {
	Contract     *BiddingwarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BiddingwarRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiddingwarRaw struct {
	Contract *Biddingwar // Generic contract binding to access the raw methods on
}

// BiddingwarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiddingwarCallerRaw struct {
	Contract *BiddingwarCaller // Generic read-only contract binding to access the raw methods on
}

// BiddingwarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiddingwarTransactorRaw struct {
	Contract *BiddingwarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiddingwar creates a new instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwar(address common.Address, backend bind.ContractBackend) (*Biddingwar, error) {
	contract, err := bindBiddingwar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Biddingwar{BiddingwarCaller: BiddingwarCaller{contract: contract}, BiddingwarTransactor: BiddingwarTransactor{contract: contract}, BiddingwarFilterer: BiddingwarFilterer{contract: contract}}, nil
}

// NewBiddingwarCaller creates a new read-only instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwarCaller(address common.Address, caller bind.ContractCaller) (*BiddingwarCaller, error) {
	contract, err := bindBiddingwar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingwarCaller{contract: contract}, nil
}

// NewBiddingwarTransactor creates a new write-only instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwarTransactor(address common.Address, transactor bind.ContractTransactor) (*BiddingwarTransactor, error) {
	contract, err := bindBiddingwar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingwarTransactor{contract: contract}, nil
}

// NewBiddingwarFilterer creates a new log filterer instance of Biddingwar, bound to a specific deployed contract.
func NewBiddingwarFilterer(address common.Address, filterer bind.ContractFilterer) (*BiddingwarFilterer, error) {
	contract, err := bindBiddingwar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiddingwarFilterer{contract: contract}, nil
}

// bindBiddingwar binds a generic wrapper to an already deployed contract.
func bindBiddingwar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BiddingwarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingwar *BiddingwarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingwar.Contract.BiddingwarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingwar *BiddingwarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.Contract.BiddingwarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingwar *BiddingwarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingwar.Contract.BiddingwarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingwar *BiddingwarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingwar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingwar *BiddingwarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingwar *BiddingwarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingwar.Contract.contract.Transact(opts, method, params...)
}

// BidEndTime is a free data retrieval call binding the contract method 0x4ba9a927.
//
// Solidity: function bidEndTime() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) BidEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "bidEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidEndTime is a free data retrieval call binding the contract method 0x4ba9a927.
//
// Solidity: function bidEndTime() view returns(uint256)
func (_Biddingwar *BiddingwarSession) BidEndTime() (*big.Int, error) {
	return _Biddingwar.Contract.BidEndTime(&_Biddingwar.CallOpts)
}

// BidEndTime is a free data retrieval call binding the contract method 0x4ba9a927.
//
// Solidity: function bidEndTime() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) BidEndTime() (*big.Int, error) {
	return _Biddingwar.Contract.BidEndTime(&_Biddingwar.CallOpts)
}

// CheckGameRunning is a free data retrieval call binding the contract method 0x2190dfdf.
//
// Solidity: function checkGameRunning() view returns(bool)
func (_Biddingwar *BiddingwarCaller) CheckGameRunning(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "checkGameRunning")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckGameRunning is a free data retrieval call binding the contract method 0x2190dfdf.
//
// Solidity: function checkGameRunning() view returns(bool)
func (_Biddingwar *BiddingwarSession) CheckGameRunning() (bool, error) {
	return _Biddingwar.Contract.CheckGameRunning(&_Biddingwar.CallOpts)
}

// CheckGameRunning is a free data retrieval call binding the contract method 0x2190dfdf.
//
// Solidity: function checkGameRunning() view returns(bool)
func (_Biddingwar *BiddingwarCallerSession) CheckGameRunning() (bool, error) {
	return _Biddingwar.Contract.CheckGameRunning(&_Biddingwar.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingwar *BiddingwarSession) Commission() (*big.Int, error) {
	return _Biddingwar.Contract.Commission(&_Biddingwar.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) Commission() (*big.Int, error) {
	return _Biddingwar.Contract.Commission(&_Biddingwar.CallOpts)
}

// CurBidAmount is a free data retrieval call binding the contract method 0xa700001b.
//
// Solidity: function curBidAmount() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) CurBidAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "curBidAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurBidAmount is a free data retrieval call binding the contract method 0xa700001b.
//
// Solidity: function curBidAmount() view returns(uint256)
func (_Biddingwar *BiddingwarSession) CurBidAmount() (*big.Int, error) {
	return _Biddingwar.Contract.CurBidAmount(&_Biddingwar.CallOpts)
}

// CurBidAmount is a free data retrieval call binding the contract method 0xa700001b.
//
// Solidity: function curBidAmount() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) CurBidAmount() (*big.Int, error) {
	return _Biddingwar.Contract.CurBidAmount(&_Biddingwar.CallOpts)
}

// CurBidTime is a free data retrieval call binding the contract method 0x53a49c92.
//
// Solidity: function curBidTime() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) CurBidTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "curBidTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurBidTime is a free data retrieval call binding the contract method 0x53a49c92.
//
// Solidity: function curBidTime() view returns(uint256)
func (_Biddingwar *BiddingwarSession) CurBidTime() (*big.Int, error) {
	return _Biddingwar.Contract.CurBidTime(&_Biddingwar.CallOpts)
}

// CurBidTime is a free data retrieval call binding the contract method 0x53a49c92.
//
// Solidity: function curBidTime() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) CurBidTime() (*big.Int, error) {
	return _Biddingwar.Contract.CurBidTime(&_Biddingwar.CallOpts)
}

// GameIndex is a free data retrieval call binding the contract method 0x5654a341.
//
// Solidity: function gameIndex() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) GameIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "gameIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameIndex is a free data retrieval call binding the contract method 0x5654a341.
//
// Solidity: function gameIndex() view returns(uint256)
func (_Biddingwar *BiddingwarSession) GameIndex() (*big.Int, error) {
	return _Biddingwar.Contract.GameIndex(&_Biddingwar.CallOpts)
}

// GameIndex is a free data retrieval call binding the contract method 0x5654a341.
//
// Solidity: function gameIndex() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) GameIndex() (*big.Int, error) {
	return _Biddingwar.Contract.GameIndex(&_Biddingwar.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Biddingwar *BiddingwarCaller) LastBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "lastBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Biddingwar *BiddingwarSession) LastBidder() (common.Address, error) {
	return _Biddingwar.Contract.LastBidder(&_Biddingwar.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Biddingwar *BiddingwarCallerSession) LastBidder() (common.Address, error) {
	return _Biddingwar.Contract.LastBidder(&_Biddingwar.CallOpts)
}

// MaxBidInterval is a free data retrieval call binding the contract method 0xc29bea35.
//
// Solidity: function maxBidInterval() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) MaxBidInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "maxBidInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBidInterval is a free data retrieval call binding the contract method 0xc29bea35.
//
// Solidity: function maxBidInterval() view returns(uint256)
func (_Biddingwar *BiddingwarSession) MaxBidInterval() (*big.Int, error) {
	return _Biddingwar.Contract.MaxBidInterval(&_Biddingwar.CallOpts)
}

// MaxBidInterval is a free data retrieval call binding the contract method 0xc29bea35.
//
// Solidity: function maxBidInterval() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) MaxBidInterval() (*big.Int, error) {
	return _Biddingwar.Contract.MaxBidInterval(&_Biddingwar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwar *BiddingwarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwar *BiddingwarSession) Owner() (common.Address, error) {
	return _Biddingwar.Contract.Owner(&_Biddingwar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Biddingwar *BiddingwarCallerSession) Owner() (common.Address, error) {
	return _Biddingwar.Contract.Owner(&_Biddingwar.CallOpts)
}

// StartBidding is a free data retrieval call binding the contract method 0x7d7b39e4.
//
// Solidity: function startBidding() view returns(bool)
func (_Biddingwar *BiddingwarCaller) StartBidding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "startBidding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StartBidding is a free data retrieval call binding the contract method 0x7d7b39e4.
//
// Solidity: function startBidding() view returns(bool)
func (_Biddingwar *BiddingwarSession) StartBidding() (bool, error) {
	return _Biddingwar.Contract.StartBidding(&_Biddingwar.CallOpts)
}

// StartBidding is a free data retrieval call binding the contract method 0x7d7b39e4.
//
// Solidity: function startBidding() view returns(bool)
func (_Biddingwar *BiddingwarCallerSession) StartBidding() (bool, error) {
	return _Biddingwar.Contract.StartBidding(&_Biddingwar.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Biddingwar *BiddingwarSession) StartTime() (*big.Int, error) {
	return _Biddingwar.Contract.StartTime(&_Biddingwar.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) StartTime() (*big.Int, error) {
	return _Biddingwar.Contract.StartTime(&_Biddingwar.CallOpts)
}

// TotalWinnerAmount is a free data retrieval call binding the contract method 0x28535057.
//
// Solidity: function totalWinnerAmount() view returns(uint256)
func (_Biddingwar *BiddingwarCaller) TotalWinnerAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingwar.contract.Call(opts, &out, "totalWinnerAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWinnerAmount is a free data retrieval call binding the contract method 0x28535057.
//
// Solidity: function totalWinnerAmount() view returns(uint256)
func (_Biddingwar *BiddingwarSession) TotalWinnerAmount() (*big.Int, error) {
	return _Biddingwar.Contract.TotalWinnerAmount(&_Biddingwar.CallOpts)
}

// TotalWinnerAmount is a free data retrieval call binding the contract method 0x28535057.
//
// Solidity: function totalWinnerAmount() view returns(uint256)
func (_Biddingwar *BiddingwarCallerSession) TotalWinnerAmount() (*big.Int, error) {
	return _Biddingwar.Contract.TotalWinnerAmount(&_Biddingwar.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingwar *BiddingwarTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingwar *BiddingwarSession) Bid() (*types.Transaction, error) {
	return _Biddingwar.Contract.Bid(&_Biddingwar.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingwar *BiddingwarTransactorSession) Bid() (*types.Transaction, error) {
	return _Biddingwar.Contract.Bid(&_Biddingwar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Biddingwar *BiddingwarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Biddingwar *BiddingwarSession) RenounceOwnership() (*types.Transaction, error) {
	return _Biddingwar.Contract.RenounceOwnership(&_Biddingwar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Biddingwar *BiddingwarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Biddingwar.Contract.RenounceOwnership(&_Biddingwar.TransactOpts)
}

// RewardToWinner is a paid mutator transaction binding the contract method 0x582598be.
//
// Solidity: function rewardToWinner() returns()
func (_Biddingwar *BiddingwarTransactor) RewardToWinner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "rewardToWinner")
}

// RewardToWinner is a paid mutator transaction binding the contract method 0x582598be.
//
// Solidity: function rewardToWinner() returns()
func (_Biddingwar *BiddingwarSession) RewardToWinner() (*types.Transaction, error) {
	return _Biddingwar.Contract.RewardToWinner(&_Biddingwar.TransactOpts)
}

// RewardToWinner is a paid mutator transaction binding the contract method 0x582598be.
//
// Solidity: function rewardToWinner() returns()
func (_Biddingwar *BiddingwarTransactorSession) RewardToWinner() (*types.Transaction, error) {
	return _Biddingwar.Contract.RewardToWinner(&_Biddingwar.TransactOpts)
}

// SetGameParams is a paid mutator transaction binding the contract method 0x1fe6e884.
//
// Solidity: function setGameParams(uint256 _bidEndTime, uint256 _maxBidInterval, uint256 _commission) returns()
func (_Biddingwar *BiddingwarTransactor) SetGameParams(opts *bind.TransactOpts, _bidEndTime *big.Int, _maxBidInterval *big.Int, _commission *big.Int) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "setGameParams", _bidEndTime, _maxBidInterval, _commission)
}

// SetGameParams is a paid mutator transaction binding the contract method 0x1fe6e884.
//
// Solidity: function setGameParams(uint256 _bidEndTime, uint256 _maxBidInterval, uint256 _commission) returns()
func (_Biddingwar *BiddingwarSession) SetGameParams(_bidEndTime *big.Int, _maxBidInterval *big.Int, _commission *big.Int) (*types.Transaction, error) {
	return _Biddingwar.Contract.SetGameParams(&_Biddingwar.TransactOpts, _bidEndTime, _maxBidInterval, _commission)
}

// SetGameParams is a paid mutator transaction binding the contract method 0x1fe6e884.
//
// Solidity: function setGameParams(uint256 _bidEndTime, uint256 _maxBidInterval, uint256 _commission) returns()
func (_Biddingwar *BiddingwarTransactorSession) SetGameParams(_bidEndTime *big.Int, _maxBidInterval *big.Int, _commission *big.Int) (*types.Transaction, error) {
	return _Biddingwar.Contract.SetGameParams(&_Biddingwar.TransactOpts, _bidEndTime, _maxBidInterval, _commission)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Biddingwar *BiddingwarTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Biddingwar *BiddingwarSession) Start() (*types.Transaction, error) {
	return _Biddingwar.Contract.Start(&_Biddingwar.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Biddingwar *BiddingwarTransactorSession) Start() (*types.Transaction, error) {
	return _Biddingwar.Contract.Start(&_Biddingwar.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Biddingwar *BiddingwarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Biddingwar *BiddingwarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Biddingwar.Contract.TransferOwnership(&_Biddingwar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Biddingwar *BiddingwarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Biddingwar.Contract.TransferOwnership(&_Biddingwar.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address toAddress) returns()
func (_Biddingwar *BiddingwarTransactor) Withdraw(opts *bind.TransactOpts, toAddress common.Address) (*types.Transaction, error) {
	return _Biddingwar.contract.Transact(opts, "withdraw", toAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address toAddress) returns()
func (_Biddingwar *BiddingwarSession) Withdraw(toAddress common.Address) (*types.Transaction, error) {
	return _Biddingwar.Contract.Withdraw(&_Biddingwar.TransactOpts, toAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address toAddress) returns()
func (_Biddingwar *BiddingwarTransactorSession) Withdraw(toAddress common.Address) (*types.Transaction, error) {
	return _Biddingwar.Contract.Withdraw(&_Biddingwar.TransactOpts, toAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Biddingwar *BiddingwarTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingwar.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Biddingwar *BiddingwarSession) Receive() (*types.Transaction, error) {
	return _Biddingwar.Contract.Receive(&_Biddingwar.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Biddingwar *BiddingwarTransactorSession) Receive() (*types.Transaction, error) {
	return _Biddingwar.Contract.Receive(&_Biddingwar.TransactOpts)
}

// BiddingwarBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Biddingwar contract.
type BiddingwarBidIterator struct {
	Event *BiddingwarBid // Event containing the contract specifics and raw log

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
func (it *BiddingwarBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarBid)
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
		it.Event = new(BiddingwarBid)
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
func (it *BiddingwarBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarBid represents a Bid event raised by the Biddingwar contract.
type BiddingwarBid struct {
	GaneIndex *big.Int
	Bidder    common.Address
	BidAmount *big.Int
	BidTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x3138d8d517460c959fb333d4e8d87ea984f1cf15d6742c02e2955dd27a622b70.
//
// Solidity: event Bid(uint256 ganeIndex, address bidder, uint256 bidAmount, uint256 bidTime)
func (_Biddingwar *BiddingwarFilterer) FilterBid(opts *bind.FilterOpts) (*BiddingwarBidIterator, error) {

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "Bid")
	if err != nil {
		return nil, err
	}
	return &BiddingwarBidIterator{contract: _Biddingwar.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0x3138d8d517460c959fb333d4e8d87ea984f1cf15d6742c02e2955dd27a622b70.
//
// Solidity: event Bid(uint256 ganeIndex, address bidder, uint256 bidAmount, uint256 bidTime)
func (_Biddingwar *BiddingwarFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *BiddingwarBid) (event.Subscription, error) {

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "Bid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarBid)
				if err := _Biddingwar.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0x3138d8d517460c959fb333d4e8d87ea984f1cf15d6742c02e2955dd27a622b70.
//
// Solidity: event Bid(uint256 ganeIndex, address bidder, uint256 bidAmount, uint256 bidTime)
func (_Biddingwar *BiddingwarFilterer) ParseBid(log types.Log) (*BiddingwarBid, error) {
	event := new(BiddingwarBid)
	if err := _Biddingwar.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarGameParamChangedIterator is returned from FilterGameParamChanged and is used to iterate over the raw logs and unpacked data for GameParamChanged events raised by the Biddingwar contract.
type BiddingwarGameParamChangedIterator struct {
	Event *BiddingwarGameParamChanged // Event containing the contract specifics and raw log

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
func (it *BiddingwarGameParamChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarGameParamChanged)
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
		it.Event = new(BiddingwarGameParamChanged)
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
func (it *BiddingwarGameParamChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarGameParamChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarGameParamChanged represents a GameParamChanged event raised by the Biddingwar contract.
type BiddingwarGameParamChanged struct {
	BidEndTime     *big.Int
	MaxBidInterval *big.Int
	Commission     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGameParamChanged is a free log retrieval operation binding the contract event 0xf0f3b998cdacbc18800a2aedf4265a2949a7f78caa822645f52f46e0dab94bae.
//
// Solidity: event GameParamChanged(uint256 bidEndTime, uint256 maxBidInterval, uint256 commission)
func (_Biddingwar *BiddingwarFilterer) FilterGameParamChanged(opts *bind.FilterOpts) (*BiddingwarGameParamChangedIterator, error) {

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "GameParamChanged")
	if err != nil {
		return nil, err
	}
	return &BiddingwarGameParamChangedIterator{contract: _Biddingwar.contract, event: "GameParamChanged", logs: logs, sub: sub}, nil
}

// WatchGameParamChanged is a free log subscription operation binding the contract event 0xf0f3b998cdacbc18800a2aedf4265a2949a7f78caa822645f52f46e0dab94bae.
//
// Solidity: event GameParamChanged(uint256 bidEndTime, uint256 maxBidInterval, uint256 commission)
func (_Biddingwar *BiddingwarFilterer) WatchGameParamChanged(opts *bind.WatchOpts, sink chan<- *BiddingwarGameParamChanged) (event.Subscription, error) {

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "GameParamChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarGameParamChanged)
				if err := _Biddingwar.contract.UnpackLog(event, "GameParamChanged", log); err != nil {
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

// ParseGameParamChanged is a log parse operation binding the contract event 0xf0f3b998cdacbc18800a2aedf4265a2949a7f78caa822645f52f46e0dab94bae.
//
// Solidity: event GameParamChanged(uint256 bidEndTime, uint256 maxBidInterval, uint256 commission)
func (_Biddingwar *BiddingwarFilterer) ParseGameParamChanged(log types.Log) (*BiddingwarGameParamChanged, error) {
	event := new(BiddingwarGameParamChanged)
	if err := _Biddingwar.contract.UnpackLog(event, "GameParamChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarGameStartedIterator is returned from FilterGameStarted and is used to iterate over the raw logs and unpacked data for GameStarted events raised by the Biddingwar contract.
type BiddingwarGameStartedIterator struct {
	Event *BiddingwarGameStarted // Event containing the contract specifics and raw log

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
func (it *BiddingwarGameStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarGameStarted)
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
		it.Event = new(BiddingwarGameStarted)
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
func (it *BiddingwarGameStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarGameStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarGameStarted represents a GameStarted event raised by the Biddingwar contract.
type BiddingwarGameStarted struct {
	GameIndex *big.Int
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGameStarted is a free log retrieval operation binding the contract event 0xa2dfaa3573b86004ec9d4d4abe5068d942579c59da4da8cd0d9e45e71a393bc1.
//
// Solidity: event GameStarted(uint256 gameIndex, uint256 startTime)
func (_Biddingwar *BiddingwarFilterer) FilterGameStarted(opts *bind.FilterOpts) (*BiddingwarGameStartedIterator, error) {

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "GameStarted")
	if err != nil {
		return nil, err
	}
	return &BiddingwarGameStartedIterator{contract: _Biddingwar.contract, event: "GameStarted", logs: logs, sub: sub}, nil
}

// WatchGameStarted is a free log subscription operation binding the contract event 0xa2dfaa3573b86004ec9d4d4abe5068d942579c59da4da8cd0d9e45e71a393bc1.
//
// Solidity: event GameStarted(uint256 gameIndex, uint256 startTime)
func (_Biddingwar *BiddingwarFilterer) WatchGameStarted(opts *bind.WatchOpts, sink chan<- *BiddingwarGameStarted) (event.Subscription, error) {

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "GameStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarGameStarted)
				if err := _Biddingwar.contract.UnpackLog(event, "GameStarted", log); err != nil {
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

// ParseGameStarted is a log parse operation binding the contract event 0xa2dfaa3573b86004ec9d4d4abe5068d942579c59da4da8cd0d9e45e71a393bc1.
//
// Solidity: event GameStarted(uint256 gameIndex, uint256 startTime)
func (_Biddingwar *BiddingwarFilterer) ParseGameStarted(log types.Log) (*BiddingwarGameStarted, error) {
	event := new(BiddingwarGameStarted)
	if err := _Biddingwar.contract.UnpackLog(event, "GameStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Biddingwar contract.
type BiddingwarOwnershipTransferredIterator struct {
	Event *BiddingwarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BiddingwarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarOwnershipTransferred)
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
		it.Event = new(BiddingwarOwnershipTransferred)
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
func (it *BiddingwarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarOwnershipTransferred represents a OwnershipTransferred event raised by the Biddingwar contract.
type BiddingwarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Biddingwar *BiddingwarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BiddingwarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BiddingwarOwnershipTransferredIterator{contract: _Biddingwar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Biddingwar *BiddingwarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BiddingwarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarOwnershipTransferred)
				if err := _Biddingwar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Biddingwar *BiddingwarFilterer) ParseOwnershipTransferred(log types.Log) (*BiddingwarOwnershipTransferred, error) {
	event := new(BiddingwarOwnershipTransferred)
	if err := _Biddingwar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the Biddingwar contract.
type BiddingwarPaymentReceivedIterator struct {
	Event *BiddingwarPaymentReceived // Event containing the contract specifics and raw log

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
func (it *BiddingwarPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarPaymentReceived)
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
		it.Event = new(BiddingwarPaymentReceived)
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
func (it *BiddingwarPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarPaymentReceived represents a PaymentReceived event raised by the Biddingwar contract.
type BiddingwarPaymentReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Biddingwar *BiddingwarFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*BiddingwarPaymentReceivedIterator, error) {

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &BiddingwarPaymentReceivedIterator{contract: _Biddingwar.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Biddingwar *BiddingwarFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *BiddingwarPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarPaymentReceived)
				if err := _Biddingwar.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_Biddingwar *BiddingwarFilterer) ParsePaymentReceived(log types.Log) (*BiddingwarPaymentReceived, error) {
	event := new(BiddingwarPaymentReceived)
	if err := _Biddingwar.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingwarWinnerGetRewardsIterator is returned from FilterWinnerGetRewards and is used to iterate over the raw logs and unpacked data for WinnerGetRewards events raised by the Biddingwar contract.
type BiddingwarWinnerGetRewardsIterator struct {
	Event *BiddingwarWinnerGetRewards // Event containing the contract specifics and raw log

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
func (it *BiddingwarWinnerGetRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingwarWinnerGetRewards)
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
		it.Event = new(BiddingwarWinnerGetRewards)
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
func (it *BiddingwarWinnerGetRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingwarWinnerGetRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingwarWinnerGetRewards represents a WinnerGetRewards event raised by the Biddingwar contract.
type BiddingwarWinnerGetRewards struct {
	GameIndex *big.Int
	Winner    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWinnerGetRewards is a free log retrieval operation binding the contract event 0xa01f6ab616422928cdc7d8f9041f197cc98d746c958e0224d32ebbf39bedf0e6.
//
// Solidity: event WinnerGetRewards(uint256 gameIndex, address winner, uint256 amount)
func (_Biddingwar *BiddingwarFilterer) FilterWinnerGetRewards(opts *bind.FilterOpts) (*BiddingwarWinnerGetRewardsIterator, error) {

	logs, sub, err := _Biddingwar.contract.FilterLogs(opts, "WinnerGetRewards")
	if err != nil {
		return nil, err
	}
	return &BiddingwarWinnerGetRewardsIterator{contract: _Biddingwar.contract, event: "WinnerGetRewards", logs: logs, sub: sub}, nil
}

// WatchWinnerGetRewards is a free log subscription operation binding the contract event 0xa01f6ab616422928cdc7d8f9041f197cc98d746c958e0224d32ebbf39bedf0e6.
//
// Solidity: event WinnerGetRewards(uint256 gameIndex, address winner, uint256 amount)
func (_Biddingwar *BiddingwarFilterer) WatchWinnerGetRewards(opts *bind.WatchOpts, sink chan<- *BiddingwarWinnerGetRewards) (event.Subscription, error) {

	logs, sub, err := _Biddingwar.contract.WatchLogs(opts, "WinnerGetRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingwarWinnerGetRewards)
				if err := _Biddingwar.contract.UnpackLog(event, "WinnerGetRewards", log); err != nil {
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

// ParseWinnerGetRewards is a log parse operation binding the contract event 0xa01f6ab616422928cdc7d8f9041f197cc98d746c958e0224d32ebbf39bedf0e6.
//
// Solidity: event WinnerGetRewards(uint256 gameIndex, address winner, uint256 amount)
func (_Biddingwar *BiddingwarFilterer) ParseWinnerGetRewards(log types.Log) (*BiddingwarWinnerGetRewards, error) {
	event := new(BiddingwarWinnerGetRewards)
	if err := _Biddingwar.contract.UnpackLog(event, "WinnerGetRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
