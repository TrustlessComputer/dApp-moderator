// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmpayment

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

// GmpaymentMetaData contains all meta data concerning the Gmpayment contract.
var GmpaymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountGM\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_gmPayments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalGM\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalGM\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userPayment\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gmTokenArg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adminArg\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adminArg\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GmpaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use GmpaymentMetaData.ABI instead.
var GmpaymentABI = GmpaymentMetaData.ABI

// Gmpayment is an auto generated Go binding around an Ethereum contract.
type Gmpayment struct {
	GmpaymentCaller     // Read-only binding to the contract
	GmpaymentTransactor // Write-only binding to the contract
	GmpaymentFilterer   // Log filterer for contract events
}

// GmpaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmpaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmpaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmpaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmpaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmpaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmpaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmpaymentSession struct {
	Contract     *Gmpayment        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmpaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmpaymentCallerSession struct {
	Contract *GmpaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GmpaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmpaymentTransactorSession struct {
	Contract     *GmpaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GmpaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmpaymentRaw struct {
	Contract *Gmpayment // Generic contract binding to access the raw methods on
}

// GmpaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmpaymentCallerRaw struct {
	Contract *GmpaymentCaller // Generic read-only contract binding to access the raw methods on
}

// GmpaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmpaymentTransactorRaw struct {
	Contract *GmpaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmpayment creates a new instance of Gmpayment, bound to a specific deployed contract.
func NewGmpayment(address common.Address, backend bind.ContractBackend) (*Gmpayment, error) {
	contract, err := bindGmpayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gmpayment{GmpaymentCaller: GmpaymentCaller{contract: contract}, GmpaymentTransactor: GmpaymentTransactor{contract: contract}, GmpaymentFilterer: GmpaymentFilterer{contract: contract}}, nil
}

// NewGmpaymentCaller creates a new read-only instance of Gmpayment, bound to a specific deployed contract.
func NewGmpaymentCaller(address common.Address, caller bind.ContractCaller) (*GmpaymentCaller, error) {
	contract, err := bindGmpayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmpaymentCaller{contract: contract}, nil
}

// NewGmpaymentTransactor creates a new write-only instance of Gmpayment, bound to a specific deployed contract.
func NewGmpaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*GmpaymentTransactor, error) {
	contract, err := bindGmpayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmpaymentTransactor{contract: contract}, nil
}

// NewGmpaymentFilterer creates a new log filterer instance of Gmpayment, bound to a specific deployed contract.
func NewGmpaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*GmpaymentFilterer, error) {
	contract, err := bindGmpayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmpaymentFilterer{contract: contract}, nil
}

// bindGmpayment binds a generic wrapper to an already deployed contract.
func bindGmpayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GmpaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gmpayment *GmpaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gmpayment.Contract.GmpaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gmpayment *GmpaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gmpayment.Contract.GmpaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gmpayment *GmpaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gmpayment.Contract.GmpaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gmpayment *GmpaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gmpayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gmpayment *GmpaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gmpayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gmpayment *GmpaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gmpayment.Contract.contract.Transact(opts, method, params...)
}

// GmPayments is a free data retrieval call binding the contract method 0xa76d8448.
//
// Solidity: function _gmPayments(address ) view returns(uint256)
func (_Gmpayment *GmpaymentCaller) GmPayments(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gmpayment.contract.Call(opts, &out, "_gmPayments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GmPayments is a free data retrieval call binding the contract method 0xa76d8448.
//
// Solidity: function _gmPayments(address ) view returns(uint256)
func (_Gmpayment *GmpaymentSession) GmPayments(arg0 common.Address) (*big.Int, error) {
	return _Gmpayment.Contract.GmPayments(&_Gmpayment.CallOpts, arg0)
}

// GmPayments is a free data retrieval call binding the contract method 0xa76d8448.
//
// Solidity: function _gmPayments(address ) view returns(uint256)
func (_Gmpayment *GmpaymentCallerSession) GmPayments(arg0 common.Address) (*big.Int, error) {
	return _Gmpayment.Contract.GmPayments(&_Gmpayment.CallOpts, arg0)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address admin)
func (_Gmpayment *GmpaymentCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gmpayment.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address admin)
func (_Gmpayment *GmpaymentSession) GetAdmin() (common.Address, error) {
	return _Gmpayment.Contract.GetAdmin(&_Gmpayment.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address admin)
func (_Gmpayment *GmpaymentCallerSession) GetAdmin() (common.Address, error) {
	return _Gmpayment.Contract.GetAdmin(&_Gmpayment.CallOpts)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_Gmpayment *GmpaymentCaller) GetMessageHash(opts *bind.CallOpts, user common.Address, totalGM *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Gmpayment.contract.Call(opts, &out, "getMessageHash", user, totalGM)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_Gmpayment *GmpaymentSession) GetMessageHash(user common.Address, totalGM *big.Int) ([32]byte, error) {
	return _Gmpayment.Contract.GetMessageHash(&_Gmpayment.CallOpts, user, totalGM)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_Gmpayment *GmpaymentCallerSession) GetMessageHash(user common.Address, totalGM *big.Int) ([32]byte, error) {
	return _Gmpayment.Contract.GetMessageHash(&_Gmpayment.CallOpts, user, totalGM)
}

// GetUserPayment is a free data retrieval call binding the contract method 0x34aac996.
//
// Solidity: function getUserPayment(address user) view returns(uint256 userPayment)
func (_Gmpayment *GmpaymentCaller) GetUserPayment(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gmpayment.contract.Call(opts, &out, "getUserPayment", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserPayment is a free data retrieval call binding the contract method 0x34aac996.
//
// Solidity: function getUserPayment(address user) view returns(uint256 userPayment)
func (_Gmpayment *GmpaymentSession) GetUserPayment(user common.Address) (*big.Int, error) {
	return _Gmpayment.Contract.GetUserPayment(&_Gmpayment.CallOpts, user)
}

// GetUserPayment is a free data retrieval call binding the contract method 0x34aac996.
//
// Solidity: function getUserPayment(address user) view returns(uint256 userPayment)
func (_Gmpayment *GmpaymentCallerSession) GetUserPayment(user common.Address) (*big.Int, error) {
	return _Gmpayment.Contract.GetUserPayment(&_Gmpayment.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gmpayment *GmpaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gmpayment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gmpayment *GmpaymentSession) Owner() (common.Address, error) {
	return _Gmpayment.Contract.Owner(&_Gmpayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gmpayment *GmpaymentCallerSession) Owner() (common.Address, error) {
	return _Gmpayment.Contract.Owner(&_Gmpayment.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x8f0bc152.
//
// Solidity: function claim(address user, uint256 totalGM, bytes signature) returns()
func (_Gmpayment *GmpaymentTransactor) Claim(opts *bind.TransactOpts, user common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "claim", user, totalGM, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x8f0bc152.
//
// Solidity: function claim(address user, uint256 totalGM, bytes signature) returns()
func (_Gmpayment *GmpaymentSession) Claim(user common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _Gmpayment.Contract.Claim(&_Gmpayment.TransactOpts, user, totalGM, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x8f0bc152.
//
// Solidity: function claim(address user, uint256 totalGM, bytes signature) returns()
func (_Gmpayment *GmpaymentTransactorSession) Claim(user common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _Gmpayment.Contract.Claim(&_Gmpayment.TransactOpts, user, totalGM, signature)
}

// EmergencyWithdrawERC20 is a paid mutator transaction binding the contract method 0xdf29b982.
//
// Solidity: function emergencyWithdrawERC20(address token, uint256 amount) returns()
func (_Gmpayment *GmpaymentTransactor) EmergencyWithdrawERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "emergencyWithdrawERC20", token, amount)
}

// EmergencyWithdrawERC20 is a paid mutator transaction binding the contract method 0xdf29b982.
//
// Solidity: function emergencyWithdrawERC20(address token, uint256 amount) returns()
func (_Gmpayment *GmpaymentSession) EmergencyWithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gmpayment.Contract.EmergencyWithdrawERC20(&_Gmpayment.TransactOpts, token, amount)
}

// EmergencyWithdrawERC20 is a paid mutator transaction binding the contract method 0xdf29b982.
//
// Solidity: function emergencyWithdrawERC20(address token, uint256 amount) returns()
func (_Gmpayment *GmpaymentTransactorSession) EmergencyWithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gmpayment.Contract.EmergencyWithdrawERC20(&_Gmpayment.TransactOpts, token, amount)
}

// EmergencyWithdrawETH is a paid mutator transaction binding the contract method 0x6b792c4b.
//
// Solidity: function emergencyWithdrawETH(uint256 amount) returns()
func (_Gmpayment *GmpaymentTransactor) EmergencyWithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "emergencyWithdrawETH", amount)
}

// EmergencyWithdrawETH is a paid mutator transaction binding the contract method 0x6b792c4b.
//
// Solidity: function emergencyWithdrawETH(uint256 amount) returns()
func (_Gmpayment *GmpaymentSession) EmergencyWithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Gmpayment.Contract.EmergencyWithdrawETH(&_Gmpayment.TransactOpts, amount)
}

// EmergencyWithdrawETH is a paid mutator transaction binding the contract method 0x6b792c4b.
//
// Solidity: function emergencyWithdrawETH(uint256 amount) returns()
func (_Gmpayment *GmpaymentTransactorSession) EmergencyWithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Gmpayment.Contract.EmergencyWithdrawETH(&_Gmpayment.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address gmTokenArg, address adminArg) returns()
func (_Gmpayment *GmpaymentTransactor) Initialize(opts *bind.TransactOpts, gmTokenArg common.Address, adminArg common.Address) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "initialize", gmTokenArg, adminArg)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address gmTokenArg, address adminArg) returns()
func (_Gmpayment *GmpaymentSession) Initialize(gmTokenArg common.Address, adminArg common.Address) (*types.Transaction, error) {
	return _Gmpayment.Contract.Initialize(&_Gmpayment.TransactOpts, gmTokenArg, adminArg)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address gmTokenArg, address adminArg) returns()
func (_Gmpayment *GmpaymentTransactorSession) Initialize(gmTokenArg common.Address, adminArg common.Address) (*types.Transaction, error) {
	return _Gmpayment.Contract.Initialize(&_Gmpayment.TransactOpts, gmTokenArg, adminArg)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gmpayment *GmpaymentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gmpayment *GmpaymentSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gmpayment.Contract.RenounceOwnership(&_Gmpayment.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gmpayment *GmpaymentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gmpayment.Contract.RenounceOwnership(&_Gmpayment.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address adminArg) returns()
func (_Gmpayment *GmpaymentTransactor) SetAdmin(opts *bind.TransactOpts, adminArg common.Address) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "setAdmin", adminArg)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address adminArg) returns()
func (_Gmpayment *GmpaymentSession) SetAdmin(adminArg common.Address) (*types.Transaction, error) {
	return _Gmpayment.Contract.SetAdmin(&_Gmpayment.TransactOpts, adminArg)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address adminArg) returns()
func (_Gmpayment *GmpaymentTransactorSession) SetAdmin(adminArg common.Address) (*types.Transaction, error) {
	return _Gmpayment.Contract.SetAdmin(&_Gmpayment.TransactOpts, adminArg)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gmpayment *GmpaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gmpayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gmpayment *GmpaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gmpayment.Contract.TransferOwnership(&_Gmpayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gmpayment *GmpaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gmpayment.Contract.TransferOwnership(&_Gmpayment.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gmpayment *GmpaymentTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gmpayment.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gmpayment *GmpaymentSession) Receive() (*types.Transaction, error) {
	return _Gmpayment.Contract.Receive(&_Gmpayment.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gmpayment *GmpaymentTransactorSession) Receive() (*types.Transaction, error) {
	return _Gmpayment.Contract.Receive(&_Gmpayment.TransactOpts)
}

// GmpaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gmpayment contract.
type GmpaymentOwnershipTransferredIterator struct {
	Event *GmpaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GmpaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmpaymentOwnershipTransferred)
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
		it.Event = new(GmpaymentOwnershipTransferred)
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
func (it *GmpaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmpaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmpaymentOwnershipTransferred represents a OwnershipTransferred event raised by the Gmpayment contract.
type GmpaymentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gmpayment *GmpaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GmpaymentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gmpayment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GmpaymentOwnershipTransferredIterator{contract: _Gmpayment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gmpayment *GmpaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GmpaymentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gmpayment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmpaymentOwnershipTransferred)
				if err := _Gmpayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gmpayment *GmpaymentFilterer) ParseOwnershipTransferred(log types.Log) (*GmpaymentOwnershipTransferred, error) {
	event := new(GmpaymentOwnershipTransferred)
	if err := _Gmpayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmpaymentPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the Gmpayment contract.
type GmpaymentPaidIterator struct {
	Event *GmpaymentPaid // Event containing the contract specifics and raw log

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
func (it *GmpaymentPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmpaymentPaid)
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
		it.Event = new(GmpaymentPaid)
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
func (it *GmpaymentPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmpaymentPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmpaymentPaid represents a Paid event raised by the Gmpayment contract.
type GmpaymentPaid struct {
	User     common.Address
	AmountGM *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address indexed user, uint256 amountGM)
func (_Gmpayment *GmpaymentFilterer) FilterPaid(opts *bind.FilterOpts, user []common.Address) (*GmpaymentPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gmpayment.contract.FilterLogs(opts, "Paid", userRule)
	if err != nil {
		return nil, err
	}
	return &GmpaymentPaidIterator{contract: _Gmpayment.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address indexed user, uint256 amountGM)
func (_Gmpayment *GmpaymentFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *GmpaymentPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gmpayment.contract.WatchLogs(opts, "Paid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmpaymentPaid)
				if err := _Gmpayment.contract.UnpackLog(event, "Paid", log); err != nil {
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

// ParsePaid is a log parse operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address indexed user, uint256 amountGM)
func (_Gmpayment *GmpaymentFilterer) ParsePaid(log types.Log) (*GmpaymentPaid, error) {
	event := new(GmpaymentPaid)
	if err := _Gmpayment.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
