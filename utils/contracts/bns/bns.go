// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bns

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
	_ = abi.ConvertType
)

// BnsMetaData contains all meta data concerning the Bns contract.
var BnsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyUpgraded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientRegistrationFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PfpTooLarge\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"filename\",\"type\":\"string\"}],\"name\":\"PfpUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ResolverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"afterUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNames\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPfp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"map\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPfpFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minRegistrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"namesLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_names\",\"type\":\"bytes[]\"}],\"name\":\"registerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setMinRegistrationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_filename\",\"type\":\"string\"}],\"name\":\"setPfp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BnsABI is the input ABI used to generate the binding from.
// Deprecated: Use BnsMetaData.ABI instead.
var BnsABI = BnsMetaData.ABI

// Bns is an auto generated Go binding around an Ethereum contract.
type Bns struct {
	BnsCaller     // Read-only binding to the contract
	BnsTransactor // Write-only binding to the contract
	BnsFilterer   // Log filterer for contract events
}

// BnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BnsSession struct {
	Contract     *Bns              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BnsCallerSession struct {
	Contract *BnsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BnsTransactorSession struct {
	Contract     *BnsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BnsRaw struct {
	Contract *Bns // Generic contract binding to access the raw methods on
}

// BnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BnsCallerRaw struct {
	Contract *BnsCaller // Generic read-only contract binding to access the raw methods on
}

// BnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BnsTransactorRaw struct {
	Contract *BnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBns creates a new instance of Bns, bound to a specific deployed contract.
func NewBns(address common.Address, backend bind.ContractBackend) (*Bns, error) {
	contract, err := bindBns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bns{BnsCaller: BnsCaller{contract: contract}, BnsTransactor: BnsTransactor{contract: contract}, BnsFilterer: BnsFilterer{contract: contract}}, nil
}

// NewBnsCaller creates a new read-only instance of Bns, bound to a specific deployed contract.
func NewBnsCaller(address common.Address, caller bind.ContractCaller) (*BnsCaller, error) {
	contract, err := bindBns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BnsCaller{contract: contract}, nil
}

// NewBnsTransactor creates a new write-only instance of Bns, bound to a specific deployed contract.
func NewBnsTransactor(address common.Address, transactor bind.ContractTransactor) (*BnsTransactor, error) {
	contract, err := bindBns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BnsTransactor{contract: contract}, nil
}

// NewBnsFilterer creates a new log filterer instance of Bns, bound to a specific deployed contract.
func NewBnsFilterer(address common.Address, filterer bind.ContractFilterer) (*BnsFilterer, error) {
	contract, err := bindBns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BnsFilterer{contract: contract}, nil
}

// bindBns binds a generic wrapper to an already deployed contract.
func bindBns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BnsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bns *BnsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bns.Contract.BnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bns *BnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bns.Contract.BnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bns *BnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bns.Contract.BnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bns *BnsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bns *BnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bns *BnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bns.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bns *BnsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bns *BnsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bns.Contract.BalanceOf(&_Bns.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bns *BnsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bns.Contract.BalanceOf(&_Bns.CallOpts, owner)
}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Bns *BnsCaller) CurrentId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "currentId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Bns *BnsSession) CurrentId() (*big.Int, error) {
	return _Bns.Contract.CurrentId(&_Bns.CallOpts)
}

// CurrentId is a free data retrieval call binding the contract method 0xe00dd161.
//
// Solidity: function currentId() view returns(uint256)
func (_Bns *BnsCallerSession) CurrentId() (*big.Int, error) {
	return _Bns.Contract.CurrentId(&_Bns.CallOpts)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(bytes[])
func (_Bns *BnsCaller) GetAllNames(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "getAllNames")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(bytes[])
func (_Bns *BnsSession) GetAllNames() ([][]byte, error) {
	return _Bns.Contract.GetAllNames(&_Bns.CallOpts)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(bytes[])
func (_Bns *BnsCallerSession) GetAllNames() ([][]byte, error) {
	return _Bns.Contract.GetAllNames(&_Bns.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bns *BnsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bns *BnsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bns.Contract.GetApproved(&_Bns.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bns *BnsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bns.Contract.GetApproved(&_Bns.CallOpts, tokenId)
}

// GetPfp is a free data retrieval call binding the contract method 0xed4b9ff5.
//
// Solidity: function getPfp(uint256 tokenId) view returns(bytes)
func (_Bns *BnsCaller) GetPfp(opts *bind.CallOpts, tokenId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "getPfp", tokenId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPfp is a free data retrieval call binding the contract method 0xed4b9ff5.
//
// Solidity: function getPfp(uint256 tokenId) view returns(bytes)
func (_Bns *BnsSession) GetPfp(tokenId *big.Int) ([]byte, error) {
	return _Bns.Contract.GetPfp(&_Bns.CallOpts, tokenId)
}

// GetPfp is a free data retrieval call binding the contract method 0xed4b9ff5.
//
// Solidity: function getPfp(uint256 tokenId) view returns(bytes)
func (_Bns *BnsCallerSession) GetPfp(tokenId *big.Int) ([]byte, error) {
	return _Bns.Contract.GetPfp(&_Bns.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bns *BnsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bns *BnsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bns.Contract.IsApprovedForAll(&_Bns.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bns *BnsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bns.Contract.IsApprovedForAll(&_Bns.CallOpts, owner, operator)
}

// MinPfpFee is a free data retrieval call binding the contract method 0x68ba7004.
//
// Solidity: function minPfpFee() view returns(uint256)
func (_Bns *BnsCaller) MinPfpFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "minPfpFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPfpFee is a free data retrieval call binding the contract method 0x68ba7004.
//
// Solidity: function minPfpFee() view returns(uint256)
func (_Bns *BnsSession) MinPfpFee() (*big.Int, error) {
	return _Bns.Contract.MinPfpFee(&_Bns.CallOpts)
}

// MinPfpFee is a free data retrieval call binding the contract method 0x68ba7004.
//
// Solidity: function minPfpFee() view returns(uint256)
func (_Bns *BnsCallerSession) MinPfpFee() (*big.Int, error) {
	return _Bns.Contract.MinPfpFee(&_Bns.CallOpts)
}

// MinRegistrationFee is a free data retrieval call binding the contract method 0x9706828a.
//
// Solidity: function minRegistrationFee() view returns(uint256)
func (_Bns *BnsCaller) MinRegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "minRegistrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinRegistrationFee is a free data retrieval call binding the contract method 0x9706828a.
//
// Solidity: function minRegistrationFee() view returns(uint256)
func (_Bns *BnsSession) MinRegistrationFee() (*big.Int, error) {
	return _Bns.Contract.MinRegistrationFee(&_Bns.CallOpts)
}

// MinRegistrationFee is a free data retrieval call binding the contract method 0x9706828a.
//
// Solidity: function minRegistrationFee() view returns(uint256)
func (_Bns *BnsCallerSession) MinRegistrationFee() (*big.Int, error) {
	return _Bns.Contract.MinRegistrationFee(&_Bns.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bns *BnsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bns *BnsSession) Name() (string, error) {
	return _Bns.Contract.Name(&_Bns.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bns *BnsCallerSession) Name() (string, error) {
	return _Bns.Contract.Name(&_Bns.CallOpts)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(bytes)
func (_Bns *BnsCaller) Names(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(bytes)
func (_Bns *BnsSession) Names(arg0 *big.Int) ([]byte, error) {
	return _Bns.Contract.Names(&_Bns.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(bytes)
func (_Bns *BnsCallerSession) Names(arg0 *big.Int) ([]byte, error) {
	return _Bns.Contract.Names(&_Bns.CallOpts, arg0)
}

// NamesLen is a free data retrieval call binding the contract method 0x5cc3cd11.
//
// Solidity: function namesLen() view returns(uint256)
func (_Bns *BnsCaller) NamesLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "namesLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NamesLen is a free data retrieval call binding the contract method 0x5cc3cd11.
//
// Solidity: function namesLen() view returns(uint256)
func (_Bns *BnsSession) NamesLen() (*big.Int, error) {
	return _Bns.Contract.NamesLen(&_Bns.CallOpts)
}

// NamesLen is a free data retrieval call binding the contract method 0x5cc3cd11.
//
// Solidity: function namesLen() view returns(uint256)
func (_Bns *BnsCallerSession) NamesLen() (*big.Int, error) {
	return _Bns.Contract.NamesLen(&_Bns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bns *BnsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bns *BnsSession) Owner() (common.Address, error) {
	return _Bns.Contract.Owner(&_Bns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bns *BnsCallerSession) Owner() (common.Address, error) {
	return _Bns.Contract.Owner(&_Bns.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bns *BnsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bns *BnsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bns.Contract.OwnerOf(&_Bns.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bns *BnsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bns.Contract.OwnerOf(&_Bns.CallOpts, tokenId)
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes ) view returns(bool)
func (_Bns *BnsCaller) Registered(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "registered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes ) view returns(bool)
func (_Bns *BnsSession) Registered(arg0 []byte) (bool, error) {
	return _Bns.Contract.Registered(&_Bns.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes ) view returns(bool)
func (_Bns *BnsCallerSession) Registered(arg0 []byte) (bool, error) {
	return _Bns.Contract.Registered(&_Bns.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(uint256)
func (_Bns *BnsCaller) Registry(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "registry", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(uint256)
func (_Bns *BnsSession) Registry(arg0 []byte) (*big.Int, error) {
	return _Bns.Contract.Registry(&_Bns.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(uint256)
func (_Bns *BnsCallerSession) Registry(arg0 []byte) (*big.Int, error) {
	return _Bns.Contract.Registry(&_Bns.CallOpts, arg0)
}

// Resolver is a free data retrieval call binding the contract method 0x108eaa4e.
//
// Solidity: function resolver(uint256 ) view returns(address)
func (_Bns *BnsCaller) Resolver(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "resolver", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x108eaa4e.
//
// Solidity: function resolver(uint256 ) view returns(address)
func (_Bns *BnsSession) Resolver(arg0 *big.Int) (common.Address, error) {
	return _Bns.Contract.Resolver(&_Bns.CallOpts, arg0)
}

// Resolver is a free data retrieval call binding the contract method 0x108eaa4e.
//
// Solidity: function resolver(uint256 ) view returns(address)
func (_Bns *BnsCallerSession) Resolver(arg0 *big.Int) (common.Address, error) {
	return _Bns.Contract.Resolver(&_Bns.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bns *BnsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bns *BnsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bns.Contract.SupportsInterface(&_Bns.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bns *BnsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bns.Contract.SupportsInterface(&_Bns.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bns *BnsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bns *BnsSession) Symbol() (string, error) {
	return _Bns.Contract.Symbol(&_Bns.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bns *BnsCallerSession) Symbol() (string, error) {
	return _Bns.Contract.Symbol(&_Bns.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bns *BnsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Bns.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bns *BnsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bns.Contract.TokenURI(&_Bns.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bns *BnsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bns.Contract.TokenURI(&_Bns.CallOpts, tokenId)
}

// AfterUpgrade is a paid mutator transaction binding the contract method 0xce184325.
//
// Solidity: function afterUpgrade() returns()
func (_Bns *BnsTransactor) AfterUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "afterUpgrade")
}

// AfterUpgrade is a paid mutator transaction binding the contract method 0xce184325.
//
// Solidity: function afterUpgrade() returns()
func (_Bns *BnsSession) AfterUpgrade() (*types.Transaction, error) {
	return _Bns.Contract.AfterUpgrade(&_Bns.TransactOpts)
}

// AfterUpgrade is a paid mutator transaction binding the contract method 0xce184325.
//
// Solidity: function afterUpgrade() returns()
func (_Bns *BnsTransactorSession) AfterUpgrade() (*types.Transaction, error) {
	return _Bns.Contract.AfterUpgrade(&_Bns.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bns *BnsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bns *BnsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.Approve(&_Bns.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bns *BnsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.Approve(&_Bns.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Bns *BnsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Bns *BnsSession) Initialize() (*types.Transaction, error) {
	return _Bns.Contract.Initialize(&_Bns.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Bns *BnsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Bns.Contract.Initialize(&_Bns.TransactOpts)
}

// Map is a paid mutator transaction binding the contract method 0x376a2b69.
//
// Solidity: function map(uint256 tokenId, address to) returns()
func (_Bns *BnsTransactor) Map(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "map", tokenId, to)
}

// Map is a paid mutator transaction binding the contract method 0x376a2b69.
//
// Solidity: function map(uint256 tokenId, address to) returns()
func (_Bns *BnsSession) Map(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Bns.Contract.Map(&_Bns.TransactOpts, tokenId, to)
}

// Map is a paid mutator transaction binding the contract method 0x376a2b69.
//
// Solidity: function map(uint256 tokenId, address to) returns()
func (_Bns *BnsTransactorSession) Map(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Bns.Contract.Map(&_Bns.TransactOpts, tokenId, to)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address owner, bytes name) payable returns(uint256)
func (_Bns *BnsTransactor) Register(opts *bind.TransactOpts, owner common.Address, name []byte) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "register", owner, name)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address owner, bytes name) payable returns(uint256)
func (_Bns *BnsSession) Register(owner common.Address, name []byte) (*types.Transaction, error) {
	return _Bns.Contract.Register(&_Bns.TransactOpts, owner, name)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address owner, bytes name) payable returns(uint256)
func (_Bns *BnsTransactorSession) Register(owner common.Address, name []byte) (*types.Transaction, error) {
	return _Bns.Contract.Register(&_Bns.TransactOpts, owner, name)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x27201c97.
//
// Solidity: function registerBatch(address owner, bytes[] _names) returns()
func (_Bns *BnsTransactor) RegisterBatch(opts *bind.TransactOpts, owner common.Address, _names [][]byte) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "registerBatch", owner, _names)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x27201c97.
//
// Solidity: function registerBatch(address owner, bytes[] _names) returns()
func (_Bns *BnsSession) RegisterBatch(owner common.Address, _names [][]byte) (*types.Transaction, error) {
	return _Bns.Contract.RegisterBatch(&_Bns.TransactOpts, owner, _names)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x27201c97.
//
// Solidity: function registerBatch(address owner, bytes[] _names) returns()
func (_Bns *BnsTransactorSession) RegisterBatch(owner common.Address, _names [][]byte) (*types.Transaction, error) {
	return _Bns.Contract.RegisterBatch(&_Bns.TransactOpts, owner, _names)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bns *BnsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bns *BnsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bns.Contract.RenounceOwnership(&_Bns.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bns *BnsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bns.Contract.RenounceOwnership(&_Bns.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bns *BnsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bns *BnsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.SafeTransferFrom(&_Bns.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bns *BnsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.SafeTransferFrom(&_Bns.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bns *BnsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bns *BnsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bns.Contract.SafeTransferFrom0(&_Bns.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bns *BnsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bns.Contract.SafeTransferFrom0(&_Bns.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bns *BnsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bns *BnsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bns.Contract.SetApprovalForAll(&_Bns.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bns *BnsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bns.Contract.SetApprovalForAll(&_Bns.TransactOpts, operator, approved)
}

// SetMinRegistrationFee is a paid mutator transaction binding the contract method 0x06301a9d.
//
// Solidity: function setMinRegistrationFee(uint256 fee) returns()
func (_Bns *BnsTransactor) SetMinRegistrationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "setMinRegistrationFee", fee)
}

// SetMinRegistrationFee is a paid mutator transaction binding the contract method 0x06301a9d.
//
// Solidity: function setMinRegistrationFee(uint256 fee) returns()
func (_Bns *BnsSession) SetMinRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.SetMinRegistrationFee(&_Bns.TransactOpts, fee)
}

// SetMinRegistrationFee is a paid mutator transaction binding the contract method 0x06301a9d.
//
// Solidity: function setMinRegistrationFee(uint256 fee) returns()
func (_Bns *BnsTransactorSession) SetMinRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.SetMinRegistrationFee(&_Bns.TransactOpts, fee)
}

// SetPfp is a paid mutator transaction binding the contract method 0xc83cae6a.
//
// Solidity: function setPfp(uint256 tokenId, bytes b, string _filename) payable returns()
func (_Bns *BnsTransactor) SetPfp(opts *bind.TransactOpts, tokenId *big.Int, b []byte, _filename string) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "setPfp", tokenId, b, _filename)
}

// SetPfp is a paid mutator transaction binding the contract method 0xc83cae6a.
//
// Solidity: function setPfp(uint256 tokenId, bytes b, string _filename) payable returns()
func (_Bns *BnsSession) SetPfp(tokenId *big.Int, b []byte, _filename string) (*types.Transaction, error) {
	return _Bns.Contract.SetPfp(&_Bns.TransactOpts, tokenId, b, _filename)
}

// SetPfp is a paid mutator transaction binding the contract method 0xc83cae6a.
//
// Solidity: function setPfp(uint256 tokenId, bytes b, string _filename) payable returns()
func (_Bns *BnsTransactorSession) SetPfp(tokenId *big.Int, b []byte, _filename string) (*types.Transaction, error) {
	return _Bns.Contract.SetPfp(&_Bns.TransactOpts, tokenId, b, _filename)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bns *BnsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bns *BnsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.TransferFrom(&_Bns.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bns *BnsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bns.Contract.TransferFrom(&_Bns.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bns *BnsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bns.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bns *BnsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bns.Contract.TransferOwnership(&_Bns.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bns *BnsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bns.Contract.TransferOwnership(&_Bns.TransactOpts, newOwner)
}

// BnsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bns contract.
type BnsApprovalIterator struct {
	Event *BnsApproval // Event containing the contract specifics and raw log

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
func (it *BnsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsApproval)
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
		it.Event = new(BnsApproval)
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
func (it *BnsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsApproval represents a Approval event raised by the Bns contract.
type BnsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bns *BnsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BnsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BnsApprovalIterator{contract: _Bns.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bns *BnsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BnsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsApproval)
				if err := _Bns.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bns *BnsFilterer) ParseApproval(log types.Log) (*BnsApproval, error) {
	event := new(BnsApproval)
	if err := _Bns.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bns contract.
type BnsApprovalForAllIterator struct {
	Event *BnsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BnsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsApprovalForAll)
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
		it.Event = new(BnsApprovalForAll)
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
func (it *BnsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsApprovalForAll represents a ApprovalForAll event raised by the Bns contract.
type BnsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bns *BnsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BnsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BnsApprovalForAllIterator{contract: _Bns.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bns *BnsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BnsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsApprovalForAll)
				if err := _Bns.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bns *BnsFilterer) ParseApprovalForAll(log types.Log) (*BnsApprovalForAll, error) {
	event := new(BnsApprovalForAll)
	if err := _Bns.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bns contract.
type BnsInitializedIterator struct {
	Event *BnsInitialized // Event containing the contract specifics and raw log

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
func (it *BnsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsInitialized)
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
		it.Event = new(BnsInitialized)
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
func (it *BnsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsInitialized represents a Initialized event raised by the Bns contract.
type BnsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bns *BnsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BnsInitializedIterator, error) {

	logs, sub, err := _Bns.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BnsInitializedIterator{contract: _Bns.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bns *BnsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BnsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bns.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsInitialized)
				if err := _Bns.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bns *BnsFilterer) ParseInitialized(log types.Log) (*BnsInitialized, error) {
	event := new(BnsInitialized)
	if err := _Bns.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Bns contract.
type BnsNameRegisteredIterator struct {
	Event *BnsNameRegistered // Event containing the contract specifics and raw log

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
func (it *BnsNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsNameRegistered)
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
		it.Event = new(BnsNameRegistered)
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
func (it *BnsNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsNameRegistered represents a NameRegistered event raised by the Bns contract.
type BnsNameRegistered struct {
	Name []byte
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x804e89f7acfccaf194f6757b177acb1ab1828a7c01ef308f9f8e4bbd2fb63873.
//
// Solidity: event NameRegistered(bytes name, uint256 indexed id)
func (_Bns *BnsFilterer) FilterNameRegistered(opts *bind.FilterOpts, id []*big.Int) (*BnsNameRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "NameRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return &BnsNameRegisteredIterator{contract: _Bns.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x804e89f7acfccaf194f6757b177acb1ab1828a7c01ef308f9f8e4bbd2fb63873.
//
// Solidity: event NameRegistered(bytes name, uint256 indexed id)
func (_Bns *BnsFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *BnsNameRegistered, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "NameRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsNameRegistered)
				if err := _Bns.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x804e89f7acfccaf194f6757b177acb1ab1828a7c01ef308f9f8e4bbd2fb63873.
//
// Solidity: event NameRegistered(bytes name, uint256 indexed id)
func (_Bns *BnsFilterer) ParseNameRegistered(log types.Log) (*BnsNameRegistered, error) {
	event := new(BnsNameRegistered)
	if err := _Bns.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bns contract.
type BnsOwnershipTransferredIterator struct {
	Event *BnsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BnsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsOwnershipTransferred)
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
		it.Event = new(BnsOwnershipTransferred)
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
func (it *BnsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsOwnershipTransferred represents a OwnershipTransferred event raised by the Bns contract.
type BnsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bns *BnsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BnsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BnsOwnershipTransferredIterator{contract: _Bns.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bns *BnsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BnsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsOwnershipTransferred)
				if err := _Bns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bns *BnsFilterer) ParseOwnershipTransferred(log types.Log) (*BnsOwnershipTransferred, error) {
	event := new(BnsOwnershipTransferred)
	if err := _Bns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsPfpUpdatedIterator is returned from FilterPfpUpdated and is used to iterate over the raw logs and unpacked data for PfpUpdated events raised by the Bns contract.
type BnsPfpUpdatedIterator struct {
	Event *BnsPfpUpdated // Event containing the contract specifics and raw log

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
func (it *BnsPfpUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsPfpUpdated)
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
		it.Event = new(BnsPfpUpdated)
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
func (it *BnsPfpUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsPfpUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsPfpUpdated represents a PfpUpdated event raised by the Bns contract.
type BnsPfpUpdated struct {
	Id       *big.Int
	Filename string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPfpUpdated is a free log retrieval operation binding the contract event 0x03b5cd0dc45f63f640f448d0037b63cc9b7e0ed566c7cf1a29fc44d18ad9e931.
//
// Solidity: event PfpUpdated(uint256 indexed id, string filename)
func (_Bns *BnsFilterer) FilterPfpUpdated(opts *bind.FilterOpts, id []*big.Int) (*BnsPfpUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "PfpUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &BnsPfpUpdatedIterator{contract: _Bns.contract, event: "PfpUpdated", logs: logs, sub: sub}, nil
}

// WatchPfpUpdated is a free log subscription operation binding the contract event 0x03b5cd0dc45f63f640f448d0037b63cc9b7e0ed566c7cf1a29fc44d18ad9e931.
//
// Solidity: event PfpUpdated(uint256 indexed id, string filename)
func (_Bns *BnsFilterer) WatchPfpUpdated(opts *bind.WatchOpts, sink chan<- *BnsPfpUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "PfpUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsPfpUpdated)
				if err := _Bns.contract.UnpackLog(event, "PfpUpdated", log); err != nil {
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

// ParsePfpUpdated is a log parse operation binding the contract event 0x03b5cd0dc45f63f640f448d0037b63cc9b7e0ed566c7cf1a29fc44d18ad9e931.
//
// Solidity: event PfpUpdated(uint256 indexed id, string filename)
func (_Bns *BnsFilterer) ParsePfpUpdated(log types.Log) (*BnsPfpUpdated, error) {
	event := new(BnsPfpUpdated)
	if err := _Bns.contract.UnpackLog(event, "PfpUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsResolverUpdatedIterator is returned from FilterResolverUpdated and is used to iterate over the raw logs and unpacked data for ResolverUpdated events raised by the Bns contract.
type BnsResolverUpdatedIterator struct {
	Event *BnsResolverUpdated // Event containing the contract specifics and raw log

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
func (it *BnsResolverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsResolverUpdated)
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
		it.Event = new(BnsResolverUpdated)
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
func (it *BnsResolverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsResolverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsResolverUpdated represents a ResolverUpdated event raised by the Bns contract.
type BnsResolverUpdated struct {
	Id   *big.Int
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResolverUpdated is a free log retrieval operation binding the contract event 0x3c055b8ebff34805cc0c0216cfd96bfe2613c75002ce000f6268324952541d06.
//
// Solidity: event ResolverUpdated(uint256 indexed id, address indexed addr)
func (_Bns *BnsFilterer) FilterResolverUpdated(opts *bind.FilterOpts, id []*big.Int, addr []common.Address) (*BnsResolverUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "ResolverUpdated", idRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &BnsResolverUpdatedIterator{contract: _Bns.contract, event: "ResolverUpdated", logs: logs, sub: sub}, nil
}

// WatchResolverUpdated is a free log subscription operation binding the contract event 0x3c055b8ebff34805cc0c0216cfd96bfe2613c75002ce000f6268324952541d06.
//
// Solidity: event ResolverUpdated(uint256 indexed id, address indexed addr)
func (_Bns *BnsFilterer) WatchResolverUpdated(opts *bind.WatchOpts, sink chan<- *BnsResolverUpdated, id []*big.Int, addr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "ResolverUpdated", idRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsResolverUpdated)
				if err := _Bns.contract.UnpackLog(event, "ResolverUpdated", log); err != nil {
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

// ParseResolverUpdated is a log parse operation binding the contract event 0x3c055b8ebff34805cc0c0216cfd96bfe2613c75002ce000f6268324952541d06.
//
// Solidity: event ResolverUpdated(uint256 indexed id, address indexed addr)
func (_Bns *BnsFilterer) ParseResolverUpdated(log types.Log) (*BnsResolverUpdated, error) {
	event := new(BnsResolverUpdated)
	if err := _Bns.contract.UnpackLog(event, "ResolverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BnsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bns contract.
type BnsTransferIterator struct {
	Event *BnsTransfer // Event containing the contract specifics and raw log

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
func (it *BnsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BnsTransfer)
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
		it.Event = new(BnsTransfer)
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
func (it *BnsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BnsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BnsTransfer represents a Transfer event raised by the Bns contract.
type BnsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bns *BnsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BnsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bns.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BnsTransferIterator{contract: _Bns.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bns *BnsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BnsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bns.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BnsTransfer)
				if err := _Bns.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bns *BnsFilterer) ParseTransfer(log types.Log) (*BnsTransfer, error) {
	event := new(BnsTransfer)
	if err := _Bns.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
