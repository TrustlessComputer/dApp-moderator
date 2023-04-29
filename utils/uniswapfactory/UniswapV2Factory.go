// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapfactory

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

// UniswapfactoryMetaData contains all meta data concerning the Uniswapfactory contract.
var UniswapfactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pairImplementation\",\"type\":\"address\"}],\"name\":\"setPairImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapfactoryMetaData.ABI instead.
var UniswapfactoryABI = UniswapfactoryMetaData.ABI

// Uniswapfactory is an auto generated Go binding around an Ethereum contract.
type Uniswapfactory struct {
	UniswapfactoryCaller     // Read-only binding to the contract
	UniswapfactoryTransactor // Write-only binding to the contract
	UniswapfactoryFilterer   // Log filterer for contract events
}

// UniswapfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapfactorySession struct {
	Contract     *Uniswapfactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapfactoryCallerSession struct {
	Contract *UniswapfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapfactoryTransactorSession struct {
	Contract     *UniswapfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapfactoryRaw struct {
	Contract *Uniswapfactory // Generic contract binding to access the raw methods on
}

// UniswapfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapfactoryCallerRaw struct {
	Contract *UniswapfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapfactoryTransactorRaw struct {
	Contract *UniswapfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapfactory creates a new instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactory(address common.Address, backend bind.ContractBackend) (*Uniswapfactory, error) {
	contract, err := bindUniswapfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapfactory{UniswapfactoryCaller: UniswapfactoryCaller{contract: contract}, UniswapfactoryTransactor: UniswapfactoryTransactor{contract: contract}, UniswapfactoryFilterer: UniswapfactoryFilterer{contract: contract}}, nil
}

// NewUniswapfactoryCaller creates a new read-only instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactoryCaller(address common.Address, caller bind.ContractCaller) (*UniswapfactoryCaller, error) {
	contract, err := bindUniswapfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryCaller{contract: contract}, nil
}

// NewUniswapfactoryTransactor creates a new write-only instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapfactoryTransactor, error) {
	contract, err := bindUniswapfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryTransactor{contract: contract}, nil
}

// NewUniswapfactoryFilterer creates a new log filterer instance of Uniswapfactory, bound to a specific deployed contract.
func NewUniswapfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapfactoryFilterer, error) {
	contract, err := bindUniswapfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryFilterer{contract: contract}, nil
}

// bindUniswapfactory binds a generic wrapper to an already deployed contract.
func bindUniswapfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapfactory *UniswapfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapfactory.Contract.UniswapfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapfactory *UniswapfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.UniswapfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapfactory *UniswapfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.UniswapfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapfactory *UniswapfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapfactory *UniswapfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapfactory *UniswapfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Uniswapfactory *UniswapfactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Uniswapfactory *UniswapfactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Uniswapfactory.Contract.AllPairs(&_Uniswapfactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Uniswapfactory *UniswapfactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Uniswapfactory.Contract.AllPairs(&_Uniswapfactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapfactory *UniswapfactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapfactory *UniswapfactorySession) AllPairsLength() (*big.Int, error) {
	return _Uniswapfactory.Contract.AllPairsLength(&_Uniswapfactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapfactory *UniswapfactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _Uniswapfactory.Contract.AllPairsLength(&_Uniswapfactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapfactory *UniswapfactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapfactory *UniswapfactorySession) FeeTo() (common.Address, error) {
	return _Uniswapfactory.Contract.FeeTo(&_Uniswapfactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapfactory *UniswapfactoryCallerSession) FeeTo() (common.Address, error) {
	return _Uniswapfactory.Contract.FeeTo(&_Uniswapfactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapfactory *UniswapfactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapfactory *UniswapfactorySession) FeeToSetter() (common.Address, error) {
	return _Uniswapfactory.Contract.FeeToSetter(&_Uniswapfactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapfactory *UniswapfactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _Uniswapfactory.Contract.FeeToSetter(&_Uniswapfactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Uniswapfactory *UniswapfactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Uniswapfactory *UniswapfactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Uniswapfactory.Contract.GetPair(&_Uniswapfactory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Uniswapfactory *UniswapfactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Uniswapfactory.Contract.GetPair(&_Uniswapfactory.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapfactory *UniswapfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapfactory *UniswapfactorySession) Owner() (common.Address, error) {
	return _Uniswapfactory.Contract.Owner(&_Uniswapfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Uniswapfactory *UniswapfactoryCallerSession) Owner() (common.Address, error) {
	return _Uniswapfactory.Contract.Owner(&_Uniswapfactory.CallOpts)
}

// PairImplementation is a free data retrieval call binding the contract method 0x71f3c596.
//
// Solidity: function pairImplementation() view returns(address)
func (_Uniswapfactory *UniswapfactoryCaller) PairImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapfactory.contract.Call(opts, &out, "pairImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairImplementation is a free data retrieval call binding the contract method 0x71f3c596.
//
// Solidity: function pairImplementation() view returns(address)
func (_Uniswapfactory *UniswapfactorySession) PairImplementation() (common.Address, error) {
	return _Uniswapfactory.Contract.PairImplementation(&_Uniswapfactory.CallOpts)
}

// PairImplementation is a free data retrieval call binding the contract method 0x71f3c596.
//
// Solidity: function pairImplementation() view returns(address)
func (_Uniswapfactory *UniswapfactoryCallerSession) PairImplementation() (common.Address, error) {
	return _Uniswapfactory.Contract.PairImplementation(&_Uniswapfactory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapfactory *UniswapfactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapfactory *UniswapfactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.CreatePair(&_Uniswapfactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapfactory *UniswapfactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.CreatePair(&_Uniswapfactory.TransactOpts, tokenA, tokenB)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _feeToSetter) returns()
func (_Uniswapfactory *UniswapfactoryTransactor) Initialize(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "initialize", _feeToSetter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _feeToSetter) returns()
func (_Uniswapfactory *UniswapfactorySession) Initialize(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.Initialize(&_Uniswapfactory.TransactOpts, _feeToSetter)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _feeToSetter) returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) Initialize(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.Initialize(&_Uniswapfactory.TransactOpts, _feeToSetter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Uniswapfactory *UniswapfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Uniswapfactory *UniswapfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Uniswapfactory.Contract.RenounceOwnership(&_Uniswapfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Uniswapfactory.Contract.RenounceOwnership(&_Uniswapfactory.TransactOpts)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Uniswapfactory *UniswapfactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Uniswapfactory *UniswapfactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.SetFeeTo(&_Uniswapfactory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.SetFeeTo(&_Uniswapfactory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Uniswapfactory *UniswapfactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Uniswapfactory *UniswapfactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.SetFeeToSetter(&_Uniswapfactory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.SetFeeToSetter(&_Uniswapfactory.TransactOpts, _feeToSetter)
}

// SetPairImplementation is a paid mutator transaction binding the contract method 0x6b262edc.
//
// Solidity: function setPairImplementation(address _pairImplementation) returns()
func (_Uniswapfactory *UniswapfactoryTransactor) SetPairImplementation(opts *bind.TransactOpts, _pairImplementation common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "setPairImplementation", _pairImplementation)
}

// SetPairImplementation is a paid mutator transaction binding the contract method 0x6b262edc.
//
// Solidity: function setPairImplementation(address _pairImplementation) returns()
func (_Uniswapfactory *UniswapfactorySession) SetPairImplementation(_pairImplementation common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.SetPairImplementation(&_Uniswapfactory.TransactOpts, _pairImplementation)
}

// SetPairImplementation is a paid mutator transaction binding the contract method 0x6b262edc.
//
// Solidity: function setPairImplementation(address _pairImplementation) returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) SetPairImplementation(_pairImplementation common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.SetPairImplementation(&_Uniswapfactory.TransactOpts, _pairImplementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Uniswapfactory *UniswapfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Uniswapfactory *UniswapfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.TransferOwnership(&_Uniswapfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Uniswapfactory *UniswapfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Uniswapfactory.Contract.TransferOwnership(&_Uniswapfactory.TransactOpts, newOwner)
}

// UniswapfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Uniswapfactory contract.
type UniswapfactoryOwnershipTransferredIterator struct {
	Event *UniswapfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniswapfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapfactoryOwnershipTransferred)
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
		it.Event = new(UniswapfactoryOwnershipTransferred)
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
func (it *UniswapfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Uniswapfactory contract.
type UniswapfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Uniswapfactory *UniswapfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniswapfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Uniswapfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryOwnershipTransferredIterator{contract: _Uniswapfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Uniswapfactory *UniswapfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniswapfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Uniswapfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapfactoryOwnershipTransferred)
				if err := _Uniswapfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Uniswapfactory *UniswapfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*UniswapfactoryOwnershipTransferred, error) {
	event := new(UniswapfactoryOwnershipTransferred)
	if err := _Uniswapfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapfactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Uniswapfactory contract.
type UniswapfactoryPairCreatedIterator struct {
	Event *UniswapfactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *UniswapfactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapfactoryPairCreated)
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
		it.Event = new(UniswapfactoryPairCreated)
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
func (it *UniswapfactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapfactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapfactoryPairCreated represents a PairCreated event raised by the Uniswapfactory contract.
type UniswapfactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapfactory *UniswapfactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*UniswapfactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Uniswapfactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &UniswapfactoryPairCreatedIterator{contract: _Uniswapfactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapfactory *UniswapfactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *UniswapfactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Uniswapfactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapfactoryPairCreated)
				if err := _Uniswapfactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapfactory *UniswapfactoryFilterer) ParsePairCreated(log types.Log) (*UniswapfactoryPairCreated, error) {
	event := new(UniswapfactoryPairCreated)
	if err := _Uniswapfactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
