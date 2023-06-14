// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package soul_contract

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

// SoulContractMetaData contains all meta data concerning the SoulContract contract.
var SoulContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"}],\"name\":\"AuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionClaimBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"}],\"name\":\"AuctionMinBidIncrementPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"AuctionReservePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"}],\"name\":\"AuctionTimeBufferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"Reserve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_auctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bfs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_bidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gmToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_mintAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_minted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_paramsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_randomizerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_script\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_signerMint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"batchMint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"biddable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdm\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBfs\",\"type\":\"address\"}],\"name\":\"changeBfs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBrc20\",\"type\":\"address\"}],\"name\":\"changeBrc20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"changeParamAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"changeRandomizerAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newScript\",\"type\":\"string\"}],\"name\":\"changeScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdd\",\"type\":\"address\"}],\"name\":\"changeSignerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"claimBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalGM\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paramsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"randomizerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bfs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signerMint\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalGM\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p5jsScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenHTML\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"variableScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"web3Script\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SoulContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SoulContractMetaData.ABI instead.
var SoulContractABI = SoulContractMetaData.ABI

// SoulContract is an auto generated Go binding around an Ethereum contract.
type SoulContract struct {
	SoulContractCaller     // Read-only binding to the contract
	SoulContractTransactor // Write-only binding to the contract
	SoulContractFilterer   // Log filterer for contract events
}

// SoulContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SoulContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SoulContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SoulContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoulContractSession struct {
	Contract     *SoulContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SoulContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SoulContractCallerSession struct {
	Contract *SoulContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SoulContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SoulContractTransactorSession struct {
	Contract     *SoulContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SoulContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SoulContractRaw struct {
	Contract *SoulContract // Generic contract binding to access the raw methods on
}

// SoulContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SoulContractCallerRaw struct {
	Contract *SoulContractCaller // Generic read-only contract binding to access the raw methods on
}

// SoulContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SoulContractTransactorRaw struct {
	Contract *SoulContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSoulContract creates a new instance of SoulContract, bound to a specific deployed contract.
func NewSoulContract(address common.Address, backend bind.ContractBackend) (*SoulContract, error) {
	contract, err := bindSoulContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SoulContract{SoulContractCaller: SoulContractCaller{contract: contract}, SoulContractTransactor: SoulContractTransactor{contract: contract}, SoulContractFilterer: SoulContractFilterer{contract: contract}}, nil
}

// NewSoulContractCaller creates a new read-only instance of SoulContract, bound to a specific deployed contract.
func NewSoulContractCaller(address common.Address, caller bind.ContractCaller) (*SoulContractCaller, error) {
	contract, err := bindSoulContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SoulContractCaller{contract: contract}, nil
}

// NewSoulContractTransactor creates a new write-only instance of SoulContract, bound to a specific deployed contract.
func NewSoulContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SoulContractTransactor, error) {
	contract, err := bindSoulContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SoulContractTransactor{contract: contract}, nil
}

// NewSoulContractFilterer creates a new log filterer instance of SoulContract, bound to a specific deployed contract.
func NewSoulContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SoulContractFilterer, error) {
	contract, err := bindSoulContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SoulContractFilterer{contract: contract}, nil
}

// bindSoulContract binds a generic wrapper to an already deployed contract.
func bindSoulContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SoulContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoulContract *SoulContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SoulContract.Contract.SoulContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoulContract *SoulContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoulContract.Contract.SoulContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoulContract *SoulContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoulContract.Contract.SoulContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoulContract *SoulContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SoulContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoulContract *SoulContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoulContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoulContract *SoulContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoulContract.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x01bc45c9.
//
// Solidity: function _admin() view returns(address)
func (_SoulContract *SoulContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x01bc45c9.
//
// Solidity: function _admin() view returns(address)
func (_SoulContract *SoulContractSession) Admin() (common.Address, error) {
	return _SoulContract.Contract.Admin(&_SoulContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0x01bc45c9.
//
// Solidity: function _admin() view returns(address)
func (_SoulContract *SoulContractCallerSession) Admin() (common.Address, error) {
	return _SoulContract.Contract.Admin(&_SoulContract.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x17dc6bf6.
//
// Solidity: function _auctions(uint256 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage)
func (_SoulContract *SoulContractCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId                   *big.Int
	Erc20Token                common.Address
	Amount                    *big.Int
	StartTime                 *big.Int
	EndTime                   *big.Int
	Bidder                    common.Address
	Settled                   bool
	TimeBuffer                *big.Int
	ReservePrice              *big.Int
	MinBidIncrementPercentage *big.Int
}, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_auctions", arg0)

	outstruct := new(struct {
		TokenId                   *big.Int
		Erc20Token                common.Address
		Amount                    *big.Int
		StartTime                 *big.Int
		EndTime                   *big.Int
		Bidder                    common.Address
		Settled                   bool
		TimeBuffer                *big.Int
		ReservePrice              *big.Int
		MinBidIncrementPercentage *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Erc20Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Settled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.TimeBuffer = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ReservePrice = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MinBidIncrementPercentage = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x17dc6bf6.
//
// Solidity: function _auctions(uint256 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage)
func (_SoulContract *SoulContractSession) Auctions(arg0 *big.Int) (struct {
	TokenId                   *big.Int
	Erc20Token                common.Address
	Amount                    *big.Int
	StartTime                 *big.Int
	EndTime                   *big.Int
	Bidder                    common.Address
	Settled                   bool
	TimeBuffer                *big.Int
	ReservePrice              *big.Int
	MinBidIncrementPercentage *big.Int
}, error) {
	return _SoulContract.Contract.Auctions(&_SoulContract.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x17dc6bf6.
//
// Solidity: function _auctions(uint256 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage)
func (_SoulContract *SoulContractCallerSession) Auctions(arg0 *big.Int) (struct {
	TokenId                   *big.Int
	Erc20Token                common.Address
	Amount                    *big.Int
	StartTime                 *big.Int
	EndTime                   *big.Int
	Bidder                    common.Address
	Settled                   bool
	TimeBuffer                *big.Int
	ReservePrice              *big.Int
	MinBidIncrementPercentage *big.Int
}, error) {
	return _SoulContract.Contract.Auctions(&_SoulContract.CallOpts, arg0)
}

// Bfs is a free data retrieval call binding the contract method 0x33348979.
//
// Solidity: function _bfs() view returns(address)
func (_SoulContract *SoulContractCaller) Bfs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_bfs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bfs is a free data retrieval call binding the contract method 0x33348979.
//
// Solidity: function _bfs() view returns(address)
func (_SoulContract *SoulContractSession) Bfs() (common.Address, error) {
	return _SoulContract.Contract.Bfs(&_SoulContract.CallOpts)
}

// Bfs is a free data retrieval call binding the contract method 0x33348979.
//
// Solidity: function _bfs() view returns(address)
func (_SoulContract *SoulContractCallerSession) Bfs() (common.Address, error) {
	return _SoulContract.Contract.Bfs(&_SoulContract.CallOpts)
}

// Bidders is a free data retrieval call binding the contract method 0x2d61bd09.
//
// Solidity: function _bidders(uint256 , address ) view returns(uint256)
func (_SoulContract *SoulContractCaller) Bidders(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_bidders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bidders is a free data retrieval call binding the contract method 0x2d61bd09.
//
// Solidity: function _bidders(uint256 , address ) view returns(uint256)
func (_SoulContract *SoulContractSession) Bidders(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SoulContract.Contract.Bidders(&_SoulContract.CallOpts, arg0, arg1)
}

// Bidders is a free data retrieval call binding the contract method 0x2d61bd09.
//
// Solidity: function _bidders(uint256 , address ) view returns(uint256)
func (_SoulContract *SoulContractCallerSession) Bidders(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SoulContract.Contract.Bidders(&_SoulContract.CallOpts, arg0, arg1)
}

// GmToken is a free data retrieval call binding the contract method 0x9c6f35f9.
//
// Solidity: function _gmToken() view returns(address)
func (_SoulContract *SoulContractCaller) GmToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_gmToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmToken is a free data retrieval call binding the contract method 0x9c6f35f9.
//
// Solidity: function _gmToken() view returns(address)
func (_SoulContract *SoulContractSession) GmToken() (common.Address, error) {
	return _SoulContract.Contract.GmToken(&_SoulContract.CallOpts)
}

// GmToken is a free data retrieval call binding the contract method 0x9c6f35f9.
//
// Solidity: function _gmToken() view returns(address)
func (_SoulContract *SoulContractCallerSession) GmToken() (common.Address, error) {
	return _SoulContract.Contract.GmToken(&_SoulContract.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0x22f4596f.
//
// Solidity: function _maxSupply() view returns(uint256)
func (_SoulContract *SoulContractCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0x22f4596f.
//
// Solidity: function _maxSupply() view returns(uint256)
func (_SoulContract *SoulContractSession) MaxSupply() (*big.Int, error) {
	return _SoulContract.Contract.MaxSupply(&_SoulContract.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0x22f4596f.
//
// Solidity: function _maxSupply() view returns(uint256)
func (_SoulContract *SoulContractCallerSession) MaxSupply() (*big.Int, error) {
	return _SoulContract.Contract.MaxSupply(&_SoulContract.CallOpts)
}

// MintAt is a free data retrieval call binding the contract method 0xcc22e004.
//
// Solidity: function _mintAt(uint256 ) view returns(uint256)
func (_SoulContract *SoulContractCaller) MintAt(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_mintAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintAt is a free data retrieval call binding the contract method 0xcc22e004.
//
// Solidity: function _mintAt(uint256 ) view returns(uint256)
func (_SoulContract *SoulContractSession) MintAt(arg0 *big.Int) (*big.Int, error) {
	return _SoulContract.Contract.MintAt(&_SoulContract.CallOpts, arg0)
}

// MintAt is a free data retrieval call binding the contract method 0xcc22e004.
//
// Solidity: function _mintAt(uint256 ) view returns(uint256)
func (_SoulContract *SoulContractCallerSession) MintAt(arg0 *big.Int) (*big.Int, error) {
	return _SoulContract.Contract.MintAt(&_SoulContract.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x7de77ecc.
//
// Solidity: function _minted(address ) view returns(uint256)
func (_SoulContract *SoulContractCaller) Minted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_minted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x7de77ecc.
//
// Solidity: function _minted(address ) view returns(uint256)
func (_SoulContract *SoulContractSession) Minted(arg0 common.Address) (*big.Int, error) {
	return _SoulContract.Contract.Minted(&_SoulContract.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x7de77ecc.
//
// Solidity: function _minted(address ) view returns(uint256)
func (_SoulContract *SoulContractCallerSession) Minted(arg0 common.Address) (*big.Int, error) {
	return _SoulContract.Contract.Minted(&_SoulContract.CallOpts, arg0)
}

// ParamsAddress is a free data retrieval call binding the contract method 0xadfc7dae.
//
// Solidity: function _paramsAddress() view returns(address)
func (_SoulContract *SoulContractCaller) ParamsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_paramsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParamsAddress is a free data retrieval call binding the contract method 0xadfc7dae.
//
// Solidity: function _paramsAddress() view returns(address)
func (_SoulContract *SoulContractSession) ParamsAddress() (common.Address, error) {
	return _SoulContract.Contract.ParamsAddress(&_SoulContract.CallOpts)
}

// ParamsAddress is a free data retrieval call binding the contract method 0xadfc7dae.
//
// Solidity: function _paramsAddress() view returns(address)
func (_SoulContract *SoulContractCallerSession) ParamsAddress() (common.Address, error) {
	return _SoulContract.Contract.ParamsAddress(&_SoulContract.CallOpts)
}

// RandomizerAddr is a free data retrieval call binding the contract method 0x66215eb4.
//
// Solidity: function _randomizerAddr() view returns(address)
func (_SoulContract *SoulContractCaller) RandomizerAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_randomizerAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomizerAddr is a free data retrieval call binding the contract method 0x66215eb4.
//
// Solidity: function _randomizerAddr() view returns(address)
func (_SoulContract *SoulContractSession) RandomizerAddr() (common.Address, error) {
	return _SoulContract.Contract.RandomizerAddr(&_SoulContract.CallOpts)
}

// RandomizerAddr is a free data retrieval call binding the contract method 0x66215eb4.
//
// Solidity: function _randomizerAddr() view returns(address)
func (_SoulContract *SoulContractCallerSession) RandomizerAddr() (common.Address, error) {
	return _SoulContract.Contract.RandomizerAddr(&_SoulContract.CallOpts)
}

// Script is a free data retrieval call binding the contract method 0x8a016249.
//
// Solidity: function _script() view returns(string)
func (_SoulContract *SoulContractCaller) Script(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_script")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Script is a free data retrieval call binding the contract method 0x8a016249.
//
// Solidity: function _script() view returns(string)
func (_SoulContract *SoulContractSession) Script() (string, error) {
	return _SoulContract.Contract.Script(&_SoulContract.CallOpts)
}

// Script is a free data retrieval call binding the contract method 0x8a016249.
//
// Solidity: function _script() view returns(string)
func (_SoulContract *SoulContractCallerSession) Script() (string, error) {
	return _SoulContract.Contract.Script(&_SoulContract.CallOpts)
}

// SignerMint is a free data retrieval call binding the contract method 0xbd896843.
//
// Solidity: function _signerMint() view returns(address)
func (_SoulContract *SoulContractCaller) SignerMint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "_signerMint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerMint is a free data retrieval call binding the contract method 0xbd896843.
//
// Solidity: function _signerMint() view returns(address)
func (_SoulContract *SoulContractSession) SignerMint() (common.Address, error) {
	return _SoulContract.Contract.SignerMint(&_SoulContract.CallOpts)
}

// SignerMint is a free data retrieval call binding the contract method 0xbd896843.
//
// Solidity: function _signerMint() view returns(address)
func (_SoulContract *SoulContractCallerSession) SignerMint() (common.Address, error) {
	return _SoulContract.Contract.SignerMint(&_SoulContract.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_SoulContract *SoulContractCaller) Available(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "available", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_SoulContract *SoulContractSession) Available(tokenId *big.Int) (bool, error) {
	return _SoulContract.Contract.Available(&_SoulContract.CallOpts, tokenId)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_SoulContract *SoulContractCallerSession) Available(tokenId *big.Int) (bool, error) {
	return _SoulContract.Contract.Available(&_SoulContract.CallOpts, tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SoulContract *SoulContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SoulContract *SoulContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SoulContract.Contract.BalanceOf(&_SoulContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SoulContract *SoulContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SoulContract.Contract.BalanceOf(&_SoulContract.CallOpts, owner)
}

// Biddable is a free data retrieval call binding the contract method 0x68a17bd8.
//
// Solidity: function biddable(uint256 tokenId) view returns(bool)
func (_SoulContract *SoulContractCaller) Biddable(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "biddable", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Biddable is a free data retrieval call binding the contract method 0x68a17bd8.
//
// Solidity: function biddable(uint256 tokenId) view returns(bool)
func (_SoulContract *SoulContractSession) Biddable(tokenId *big.Int) (bool, error) {
	return _SoulContract.Contract.Biddable(&_SoulContract.CallOpts, tokenId)
}

// Biddable is a free data retrieval call binding the contract method 0x68a17bd8.
//
// Solidity: function biddable(uint256 tokenId) view returns(bool)
func (_SoulContract *SoulContractCallerSession) Biddable(tokenId *big.Int) (bool, error) {
	return _SoulContract.Contract.Biddable(&_SoulContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SoulContract *SoulContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SoulContract *SoulContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SoulContract.Contract.GetApproved(&_SoulContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SoulContract *SoulContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SoulContract.Contract.GetApproved(&_SoulContract.CallOpts, tokenId)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_SoulContract *SoulContractCaller) GetMessageHash(opts *bind.CallOpts, user common.Address, totalGM *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "getMessageHash", user, totalGM)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_SoulContract *SoulContractSession) GetMessageHash(user common.Address, totalGM *big.Int) ([32]byte, error) {
	return _SoulContract.Contract.GetMessageHash(&_SoulContract.CallOpts, user, totalGM)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_SoulContract *SoulContractCallerSession) GetMessageHash(user common.Address, totalGM *big.Int) ([32]byte, error) {
	return _SoulContract.Contract.GetMessageHash(&_SoulContract.CallOpts, user, totalGM)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SoulContract *SoulContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SoulContract *SoulContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SoulContract.Contract.IsApprovedForAll(&_SoulContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SoulContract *SoulContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SoulContract.Contract.IsApprovedForAll(&_SoulContract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SoulContract *SoulContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SoulContract *SoulContractSession) Name() (string, error) {
	return _SoulContract.Contract.Name(&_SoulContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SoulContract *SoulContractCallerSession) Name() (string, error) {
	return _SoulContract.Contract.Name(&_SoulContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SoulContract *SoulContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SoulContract *SoulContractSession) Owner() (common.Address, error) {
	return _SoulContract.Contract.Owner(&_SoulContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SoulContract *SoulContractCallerSession) Owner() (common.Address, error) {
	return _SoulContract.Contract.Owner(&_SoulContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SoulContract *SoulContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SoulContract *SoulContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SoulContract.Contract.OwnerOf(&_SoulContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SoulContract *SoulContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SoulContract.Contract.OwnerOf(&_SoulContract.CallOpts, tokenId)
}

// P5jsScript is a free data retrieval call binding the contract method 0x74973a41.
//
// Solidity: function p5jsScript() view returns(string result)
func (_SoulContract *SoulContractCaller) P5jsScript(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "p5jsScript")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// P5jsScript is a free data retrieval call binding the contract method 0x74973a41.
//
// Solidity: function p5jsScript() view returns(string result)
func (_SoulContract *SoulContractSession) P5jsScript() (string, error) {
	return _SoulContract.Contract.P5jsScript(&_SoulContract.CallOpts)
}

// P5jsScript is a free data retrieval call binding the contract method 0x74973a41.
//
// Solidity: function p5jsScript() view returns(string result)
func (_SoulContract *SoulContractCallerSession) P5jsScript() (string, error) {
	return _SoulContract.Contract.P5jsScript(&_SoulContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SoulContract *SoulContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SoulContract *SoulContractSession) Paused() (bool, error) {
	return _SoulContract.Contract.Paused(&_SoulContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SoulContract *SoulContractCallerSession) Paused() (bool, error) {
	return _SoulContract.Contract.Paused(&_SoulContract.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 projectId, uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_SoulContract *SoulContractCaller) RoyaltyInfo(opts *bind.CallOpts, projectId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "royaltyInfo", projectId, _salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 projectId, uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_SoulContract *SoulContractSession) RoyaltyInfo(projectId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _SoulContract.Contract.RoyaltyInfo(&_SoulContract.CallOpts, projectId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 projectId, uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_SoulContract *SoulContractCallerSession) RoyaltyInfo(projectId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _SoulContract.Contract.RoyaltyInfo(&_SoulContract.CallOpts, projectId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SoulContract *SoulContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SoulContract *SoulContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SoulContract.Contract.SupportsInterface(&_SoulContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SoulContract *SoulContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SoulContract.Contract.SupportsInterface(&_SoulContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SoulContract *SoulContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SoulContract *SoulContractSession) Symbol() (string, error) {
	return _SoulContract.Contract.Symbol(&_SoulContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SoulContract *SoulContractCallerSession) Symbol() (string, error) {
	return _SoulContract.Contract.Symbol(&_SoulContract.CallOpts)
}

// TokenHTML is a free data retrieval call binding the contract method 0xeb29c232.
//
// Solidity: function tokenHTML(bytes32 seed, uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractCaller) TokenHTML(opts *bind.CallOpts, seed [32]byte, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "tokenHTML", seed, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenHTML is a free data retrieval call binding the contract method 0xeb29c232.
//
// Solidity: function tokenHTML(bytes32 seed, uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractSession) TokenHTML(seed [32]byte, tokenId *big.Int) (string, error) {
	return _SoulContract.Contract.TokenHTML(&_SoulContract.CallOpts, seed, tokenId)
}

// TokenHTML is a free data retrieval call binding the contract method 0xeb29c232.
//
// Solidity: function tokenHTML(bytes32 seed, uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractCallerSession) TokenHTML(seed [32]byte, tokenId *big.Int) (string, error) {
	return _SoulContract.Contract.TokenHTML(&_SoulContract.CallOpts, seed, tokenId)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 tokenId) view returns(bytes32)
func (_SoulContract *SoulContractCaller) TokenIdToHash(opts *bind.CallOpts, tokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "tokenIdToHash", tokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 tokenId) view returns(bytes32)
func (_SoulContract *SoulContractSession) TokenIdToHash(tokenId *big.Int) ([32]byte, error) {
	return _SoulContract.Contract.TokenIdToHash(&_SoulContract.CallOpts, tokenId)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 tokenId) view returns(bytes32)
func (_SoulContract *SoulContractCallerSession) TokenIdToHash(tokenId *big.Int) ([32]byte, error) {
	return _SoulContract.Contract.TokenIdToHash(&_SoulContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SoulContract.Contract.TokenURI(&_SoulContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SoulContract.Contract.TokenURI(&_SoulContract.CallOpts, tokenId)
}

// VariableScript is a free data retrieval call binding the contract method 0x32e32326.
//
// Solidity: function variableScript(bytes32 seed, uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractCaller) VariableScript(opts *bind.CallOpts, seed [32]byte, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "variableScript", seed, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VariableScript is a free data retrieval call binding the contract method 0x32e32326.
//
// Solidity: function variableScript(bytes32 seed, uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractSession) VariableScript(seed [32]byte, tokenId *big.Int) (string, error) {
	return _SoulContract.Contract.VariableScript(&_SoulContract.CallOpts, seed, tokenId)
}

// VariableScript is a free data retrieval call binding the contract method 0x32e32326.
//
// Solidity: function variableScript(bytes32 seed, uint256 tokenId) view returns(string result)
func (_SoulContract *SoulContractCallerSession) VariableScript(seed [32]byte, tokenId *big.Int) (string, error) {
	return _SoulContract.Contract.VariableScript(&_SoulContract.CallOpts, seed, tokenId)
}

// Web3Script is a free data retrieval call binding the contract method 0xa0d41e11.
//
// Solidity: function web3Script() view returns(string result)
func (_SoulContract *SoulContractCaller) Web3Script(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulContract.contract.Call(opts, &out, "web3Script")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Web3Script is a free data retrieval call binding the contract method 0xa0d41e11.
//
// Solidity: function web3Script() view returns(string result)
func (_SoulContract *SoulContractSession) Web3Script() (string, error) {
	return _SoulContract.Contract.Web3Script(&_SoulContract.CallOpts)
}

// Web3Script is a free data retrieval call binding the contract method 0xa0d41e11.
//
// Solidity: function web3Script() view returns(string result)
func (_SoulContract *SoulContractCallerSession) Web3Script() (string, error) {
	return _SoulContract.Contract.Web3Script(&_SoulContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.Approve(&_SoulContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.Approve(&_SoulContract.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x624de2a9.
//
// Solidity: function batchMint(address to, uint256 n, bytes signatures) payable returns(uint256[])
func (_SoulContract *SoulContractTransactor) BatchMint(opts *bind.TransactOpts, to common.Address, n *big.Int, signatures []byte) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "batchMint", to, n, signatures)
}

// BatchMint is a paid mutator transaction binding the contract method 0x624de2a9.
//
// Solidity: function batchMint(address to, uint256 n, bytes signatures) payable returns(uint256[])
func (_SoulContract *SoulContractSession) BatchMint(to common.Address, n *big.Int, signatures []byte) (*types.Transaction, error) {
	return _SoulContract.Contract.BatchMint(&_SoulContract.TransactOpts, to, n, signatures)
}

// BatchMint is a paid mutator transaction binding the contract method 0x624de2a9.
//
// Solidity: function batchMint(address to, uint256 n, bytes signatures) payable returns(uint256[])
func (_SoulContract *SoulContractTransactorSession) BatchMint(to common.Address, n *big.Int, signatures []byte) (*types.Transaction, error) {
	return _SoulContract.Contract.BatchMint(&_SoulContract.TransactOpts, to, n, signatures)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdm) returns()
func (_SoulContract *SoulContractTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdm common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeAdmin", newAdm)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdm) returns()
func (_SoulContract *SoulContractSession) ChangeAdmin(newAdm common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeAdmin(&_SoulContract.TransactOpts, newAdm)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdm) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeAdmin(newAdm common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeAdmin(&_SoulContract.TransactOpts, newAdm)
}

// ChangeBfs is a paid mutator transaction binding the contract method 0x2b7d13de.
//
// Solidity: function changeBfs(address newBfs) returns()
func (_SoulContract *SoulContractTransactor) ChangeBfs(opts *bind.TransactOpts, newBfs common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeBfs", newBfs)
}

// ChangeBfs is a paid mutator transaction binding the contract method 0x2b7d13de.
//
// Solidity: function changeBfs(address newBfs) returns()
func (_SoulContract *SoulContractSession) ChangeBfs(newBfs common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeBfs(&_SoulContract.TransactOpts, newBfs)
}

// ChangeBfs is a paid mutator transaction binding the contract method 0x2b7d13de.
//
// Solidity: function changeBfs(address newBfs) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeBfs(newBfs common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeBfs(&_SoulContract.TransactOpts, newBfs)
}

// ChangeBrc20Token is a paid mutator transaction binding the contract method 0x93c6cdc0.
//
// Solidity: function changeBrc20Token(address newBrc20) returns()
func (_SoulContract *SoulContractTransactor) ChangeBrc20Token(opts *bind.TransactOpts, newBrc20 common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeBrc20Token", newBrc20)
}

// ChangeBrc20Token is a paid mutator transaction binding the contract method 0x93c6cdc0.
//
// Solidity: function changeBrc20Token(address newBrc20) returns()
func (_SoulContract *SoulContractSession) ChangeBrc20Token(newBrc20 common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeBrc20Token(&_SoulContract.TransactOpts, newBrc20)
}

// ChangeBrc20Token is a paid mutator transaction binding the contract method 0x93c6cdc0.
//
// Solidity: function changeBrc20Token(address newBrc20) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeBrc20Token(newBrc20 common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeBrc20Token(&_SoulContract.TransactOpts, newBrc20)
}

// ChangeParamAddr is a paid mutator transaction binding the contract method 0xfebfd6c3.
//
// Solidity: function changeParamAddr(address newAddr) returns()
func (_SoulContract *SoulContractTransactor) ChangeParamAddr(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeParamAddr", newAddr)
}

// ChangeParamAddr is a paid mutator transaction binding the contract method 0xfebfd6c3.
//
// Solidity: function changeParamAddr(address newAddr) returns()
func (_SoulContract *SoulContractSession) ChangeParamAddr(newAddr common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeParamAddr(&_SoulContract.TransactOpts, newAddr)
}

// ChangeParamAddr is a paid mutator transaction binding the contract method 0xfebfd6c3.
//
// Solidity: function changeParamAddr(address newAddr) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeParamAddr(newAddr common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeParamAddr(&_SoulContract.TransactOpts, newAddr)
}

// ChangeRandomizerAddr is a paid mutator transaction binding the contract method 0x1ca9741b.
//
// Solidity: function changeRandomizerAddr(address newAddr) returns()
func (_SoulContract *SoulContractTransactor) ChangeRandomizerAddr(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeRandomizerAddr", newAddr)
}

// ChangeRandomizerAddr is a paid mutator transaction binding the contract method 0x1ca9741b.
//
// Solidity: function changeRandomizerAddr(address newAddr) returns()
func (_SoulContract *SoulContractSession) ChangeRandomizerAddr(newAddr common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeRandomizerAddr(&_SoulContract.TransactOpts, newAddr)
}

// ChangeRandomizerAddr is a paid mutator transaction binding the contract method 0x1ca9741b.
//
// Solidity: function changeRandomizerAddr(address newAddr) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeRandomizerAddr(newAddr common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeRandomizerAddr(&_SoulContract.TransactOpts, newAddr)
}

// ChangeScript is a paid mutator transaction binding the contract method 0x140f0e07.
//
// Solidity: function changeScript(string newScript) returns()
func (_SoulContract *SoulContractTransactor) ChangeScript(opts *bind.TransactOpts, newScript string) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeScript", newScript)
}

// ChangeScript is a paid mutator transaction binding the contract method 0x140f0e07.
//
// Solidity: function changeScript(string newScript) returns()
func (_SoulContract *SoulContractSession) ChangeScript(newScript string) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeScript(&_SoulContract.TransactOpts, newScript)
}

// ChangeScript is a paid mutator transaction binding the contract method 0x140f0e07.
//
// Solidity: function changeScript(string newScript) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeScript(newScript string) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeScript(&_SoulContract.TransactOpts, newScript)
}

// ChangeSignerMint is a paid mutator transaction binding the contract method 0xf3f83ddd.
//
// Solidity: function changeSignerMint(address newAdd) returns()
func (_SoulContract *SoulContractTransactor) ChangeSignerMint(opts *bind.TransactOpts, newAdd common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "changeSignerMint", newAdd)
}

// ChangeSignerMint is a paid mutator transaction binding the contract method 0xf3f83ddd.
//
// Solidity: function changeSignerMint(address newAdd) returns()
func (_SoulContract *SoulContractSession) ChangeSignerMint(newAdd common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeSignerMint(&_SoulContract.TransactOpts, newAdd)
}

// ChangeSignerMint is a paid mutator transaction binding the contract method 0xf3f83ddd.
//
// Solidity: function changeSignerMint(address newAdd) returns()
func (_SoulContract *SoulContractTransactorSession) ChangeSignerMint(newAdd common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.ChangeSignerMint(&_SoulContract.TransactOpts, newAdd)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactor) ClaimBid(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "claimBid", tokenId)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 tokenId) returns()
func (_SoulContract *SoulContractSession) ClaimBid(tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.ClaimBid(&_SoulContract.TransactOpts, tokenId)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactorSession) ClaimBid(tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.ClaimBid(&_SoulContract.TransactOpts, tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xd5563f31.
//
// Solidity: function createAuction(uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactor) CreateAuction(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "createAuction", tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xd5563f31.
//
// Solidity: function createAuction(uint256 tokenId) returns()
func (_SoulContract *SoulContractSession) CreateAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.CreateAuction(&_SoulContract.TransactOpts, tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xd5563f31.
//
// Solidity: function createAuction(uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactorSession) CreateAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.CreateAuction(&_SoulContract.TransactOpts, tokenId)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 tokenId, uint256 amount) payable returns()
func (_SoulContract *SoulContractTransactor) CreateBid(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "createBid", tokenId, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 tokenId, uint256 amount) payable returns()
func (_SoulContract *SoulContractSession) CreateBid(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.CreateBid(&_SoulContract.TransactOpts, tokenId, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 tokenId, uint256 amount) payable returns()
func (_SoulContract *SoulContractTransactorSession) CreateBid(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.CreateBid(&_SoulContract.TransactOpts, tokenId, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x94f5a66e.
//
// Solidity: function initialize(string name, string symbol, address admin, address paramsAddress, address randomizerAddr, address gmToken, address bfs, address signerMint) returns()
func (_SoulContract *SoulContractTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, admin common.Address, paramsAddress common.Address, randomizerAddr common.Address, gmToken common.Address, bfs common.Address, signerMint common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "initialize", name, symbol, admin, paramsAddress, randomizerAddr, gmToken, bfs, signerMint)
}

// Initialize is a paid mutator transaction binding the contract method 0x94f5a66e.
//
// Solidity: function initialize(string name, string symbol, address admin, address paramsAddress, address randomizerAddr, address gmToken, address bfs, address signerMint) returns()
func (_SoulContract *SoulContractSession) Initialize(name string, symbol string, admin common.Address, paramsAddress common.Address, randomizerAddr common.Address, gmToken common.Address, bfs common.Address, signerMint common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.Initialize(&_SoulContract.TransactOpts, name, symbol, admin, paramsAddress, randomizerAddr, gmToken, bfs, signerMint)
}

// Initialize is a paid mutator transaction binding the contract method 0x94f5a66e.
//
// Solidity: function initialize(string name, string symbol, address admin, address paramsAddress, address randomizerAddr, address gmToken, address bfs, address signerMint) returns()
func (_SoulContract *SoulContractTransactorSession) Initialize(name string, symbol string, admin common.Address, paramsAddress common.Address, randomizerAddr common.Address, gmToken common.Address, bfs common.Address, signerMint common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.Initialize(&_SoulContract.TransactOpts, name, symbol, admin, paramsAddress, randomizerAddr, gmToken, bfs, signerMint)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address to, uint256 totalGM, bytes signature) payable returns(uint256 tokenId)
func (_SoulContract *SoulContractTransactor) Mint(opts *bind.TransactOpts, to common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "mint", to, totalGM, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address to, uint256 totalGM, bytes signature) payable returns(uint256 tokenId)
func (_SoulContract *SoulContractSession) Mint(to common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _SoulContract.Contract.Mint(&_SoulContract.TransactOpts, to, totalGM, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address to, uint256 totalGM, bytes signature) payable returns(uint256 tokenId)
func (_SoulContract *SoulContractTransactorSession) Mint(to common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _SoulContract.Contract.Mint(&_SoulContract.TransactOpts, to, totalGM, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoulContract *SoulContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoulContract *SoulContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SoulContract.Contract.RenounceOwnership(&_SoulContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoulContract *SoulContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SoulContract.Contract.RenounceOwnership(&_SoulContract.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.SafeTransferFrom(&_SoulContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.SafeTransferFrom(&_SoulContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SoulContract *SoulContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SoulContract *SoulContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SoulContract.Contract.SafeTransferFrom0(&_SoulContract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SoulContract *SoulContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SoulContract.Contract.SafeTransferFrom0(&_SoulContract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SoulContract *SoulContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SoulContract *SoulContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SoulContract.Contract.SetApprovalForAll(&_SoulContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SoulContract *SoulContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SoulContract.Contract.SetApprovalForAll(&_SoulContract.TransactOpts, operator, approved)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x2e993611.
//
// Solidity: function settleAuction(uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactor) SettleAuction(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "settleAuction", tokenId)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x2e993611.
//
// Solidity: function settleAuction(uint256 tokenId) returns()
func (_SoulContract *SoulContractSession) SettleAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.SettleAuction(&_SoulContract.TransactOpts, tokenId)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x2e993611.
//
// Solidity: function settleAuction(uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactorSession) SettleAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.SettleAuction(&_SoulContract.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.TransferFrom(&_SoulContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SoulContract *SoulContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SoulContract.Contract.TransferFrom(&_SoulContract.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoulContract *SoulContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SoulContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoulContract *SoulContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.TransferOwnership(&_SoulContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoulContract *SoulContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoulContract.Contract.TransferOwnership(&_SoulContract.TransactOpts, newOwner)
}

// SoulContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SoulContract contract.
type SoulContractApprovalIterator struct {
	Event *SoulContractApproval // Event containing the contract specifics and raw log

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
func (it *SoulContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractApproval)
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
		it.Event = new(SoulContractApproval)
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
func (it *SoulContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractApproval represents a Approval event raised by the SoulContract contract.
type SoulContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SoulContractApprovalIterator, error) {

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

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractApprovalIterator{contract: _SoulContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SoulContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractApproval)
				if err := _SoulContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SoulContract *SoulContractFilterer) ParseApproval(log types.Log) (*SoulContractApproval, error) {
	event := new(SoulContractApproval)
	if err := _SoulContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SoulContract contract.
type SoulContractApprovalForAllIterator struct {
	Event *SoulContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SoulContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractApprovalForAll)
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
		it.Event = new(SoulContractApprovalForAll)
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
func (it *SoulContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractApprovalForAll represents a ApprovalForAll event raised by the SoulContract contract.
type SoulContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SoulContract *SoulContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SoulContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractApprovalForAllIterator{contract: _SoulContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SoulContract *SoulContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SoulContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractApprovalForAll)
				if err := _SoulContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_SoulContract *SoulContractFilterer) ParseApprovalForAll(log types.Log) (*SoulContractApprovalForAll, error) {
	event := new(SoulContractApprovalForAll)
	if err := _SoulContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for AuctionBid events raised by the SoulContract contract.
type SoulContractAuctionBidIterator struct {
	Event *SoulContractAuctionBid // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionBid)
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
		it.Event = new(SoulContractAuctionBid)
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
func (it *SoulContractAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionBid represents a AuctionBid event raised by the SoulContract contract.
type SoulContractAuctionBid struct {
	TokenId  *big.Int
	Sender   common.Address
	Value    *big.Int
	Extended bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed tokenId, address sender, uint256 value, bool extended)
func (_SoulContract *SoulContractFilterer) FilterAuctionBid(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulContractAuctionBidIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionBidIterator{contract: _SoulContract.contract, event: "AuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed tokenId, address sender, uint256 value, bool extended)
func (_SoulContract *SoulContractFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionBid, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionBid)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionBid", log); err != nil {
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

// ParseAuctionBid is a log parse operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed tokenId, address sender, uint256 value, bool extended)
func (_SoulContract *SoulContractFilterer) ParseAuctionBid(log types.Log) (*SoulContractAuctionBid, error) {
	event := new(SoulContractAuctionBid)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionClaimBidIterator is returned from FilterAuctionClaimBid and is used to iterate over the raw logs and unpacked data for AuctionClaimBid events raised by the SoulContract contract.
type SoulContractAuctionClaimBidIterator struct {
	Event *SoulContractAuctionClaimBid // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionClaimBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionClaimBid)
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
		it.Event = new(SoulContractAuctionClaimBid)
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
func (it *SoulContractAuctionClaimBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionClaimBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionClaimBid represents a AuctionClaimBid event raised by the SoulContract contract.
type SoulContractAuctionClaimBid struct {
	TokenId *big.Int
	Sender  common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionClaimBid is a free log retrieval operation binding the contract event 0x63ad079d900b92497dd34dccc37cb42c76aff5798f6084bd35af77d8cb5473fe.
//
// Solidity: event AuctionClaimBid(uint256 indexed tokenId, address sender, uint256 value)
func (_SoulContract *SoulContractFilterer) FilterAuctionClaimBid(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulContractAuctionClaimBidIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionClaimBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionClaimBidIterator{contract: _SoulContract.contract, event: "AuctionClaimBid", logs: logs, sub: sub}, nil
}

// WatchAuctionClaimBid is a free log subscription operation binding the contract event 0x63ad079d900b92497dd34dccc37cb42c76aff5798f6084bd35af77d8cb5473fe.
//
// Solidity: event AuctionClaimBid(uint256 indexed tokenId, address sender, uint256 value)
func (_SoulContract *SoulContractFilterer) WatchAuctionClaimBid(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionClaimBid, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionClaimBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionClaimBid)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionClaimBid", log); err != nil {
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

// ParseAuctionClaimBid is a log parse operation binding the contract event 0x63ad079d900b92497dd34dccc37cb42c76aff5798f6084bd35af77d8cb5473fe.
//
// Solidity: event AuctionClaimBid(uint256 indexed tokenId, address sender, uint256 value)
func (_SoulContract *SoulContractFilterer) ParseAuctionClaimBid(log types.Log) (*SoulContractAuctionClaimBid, error) {
	event := new(SoulContractAuctionClaimBid)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionClaimBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionClosedIterator is returned from FilterAuctionClosed and is used to iterate over the raw logs and unpacked data for AuctionClosed events raised by the SoulContract contract.
type SoulContractAuctionClosedIterator struct {
	Event *SoulContractAuctionClosed // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionClosed)
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
		it.Event = new(SoulContractAuctionClosed)
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
func (it *SoulContractAuctionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionClosed represents a AuctionClosed event raised by the SoulContract contract.
type SoulContractAuctionClosed struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionClosed is a free log retrieval operation binding the contract event 0xac4a907ec29adcc56774b757ecb1e1b4d597374fc9386107d05e2670259df7d3.
//
// Solidity: event AuctionClosed(uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) FilterAuctionClosed(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulContractAuctionClosedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionClosed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionClosedIterator{contract: _SoulContract.contract, event: "AuctionClosed", logs: logs, sub: sub}, nil
}

// WatchAuctionClosed is a free log subscription operation binding the contract event 0xac4a907ec29adcc56774b757ecb1e1b4d597374fc9386107d05e2670259df7d3.
//
// Solidity: event AuctionClosed(uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) WatchAuctionClosed(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionClosed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionClosed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionClosed)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionClosed", log); err != nil {
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

// ParseAuctionClosed is a log parse operation binding the contract event 0xac4a907ec29adcc56774b757ecb1e1b4d597374fc9386107d05e2670259df7d3.
//
// Solidity: event AuctionClosed(uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) ParseAuctionClosed(log types.Log) (*SoulContractAuctionClosed, error) {
	event := new(SoulContractAuctionClosed)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SoulContract contract.
type SoulContractAuctionCreatedIterator struct {
	Event *SoulContractAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionCreated)
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
		it.Event = new(SoulContractAuctionCreated)
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
func (it *SoulContractAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionCreated represents a AuctionCreated event raised by the SoulContract contract.
type SoulContractAuctionCreated struct {
	TokenId   *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed tokenId, uint256 startTime, uint256 endTime)
func (_SoulContract *SoulContractFilterer) FilterAuctionCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulContractAuctionCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionCreatedIterator{contract: _SoulContract.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed tokenId, uint256 startTime, uint256 endTime)
func (_SoulContract *SoulContractFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionCreated)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed tokenId, uint256 startTime, uint256 endTime)
func (_SoulContract *SoulContractFilterer) ParseAuctionCreated(log types.Log) (*SoulContractAuctionCreated, error) {
	event := new(SoulContractAuctionCreated)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionExtendedIterator is returned from FilterAuctionExtended and is used to iterate over the raw logs and unpacked data for AuctionExtended events raised by the SoulContract contract.
type SoulContractAuctionExtendedIterator struct {
	Event *SoulContractAuctionExtended // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionExtended)
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
		it.Event = new(SoulContractAuctionExtended)
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
func (it *SoulContractAuctionExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionExtended represents a AuctionExtended event raised by the SoulContract contract.
type SoulContractAuctionExtended struct {
	TokenId *big.Int
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionExtended is a free log retrieval operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed tokenId, uint256 endTime)
func (_SoulContract *SoulContractFilterer) FilterAuctionExtended(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulContractAuctionExtendedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionExtended", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionExtendedIterator{contract: _SoulContract.contract, event: "AuctionExtended", logs: logs, sub: sub}, nil
}

// WatchAuctionExtended is a free log subscription operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed tokenId, uint256 endTime)
func (_SoulContract *SoulContractFilterer) WatchAuctionExtended(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionExtended, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionExtended", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionExtended)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
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

// ParseAuctionExtended is a log parse operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed tokenId, uint256 endTime)
func (_SoulContract *SoulContractFilterer) ParseAuctionExtended(log types.Log) (*SoulContractAuctionExtended, error) {
	event := new(SoulContractAuctionExtended)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionMinBidIncrementPercentageUpdatedIterator is returned from FilterAuctionMinBidIncrementPercentageUpdated and is used to iterate over the raw logs and unpacked data for AuctionMinBidIncrementPercentageUpdated events raised by the SoulContract contract.
type SoulContractAuctionMinBidIncrementPercentageUpdatedIterator struct {
	Event *SoulContractAuctionMinBidIncrementPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionMinBidIncrementPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionMinBidIncrementPercentageUpdated)
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
		it.Event = new(SoulContractAuctionMinBidIncrementPercentageUpdated)
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
func (it *SoulContractAuctionMinBidIncrementPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionMinBidIncrementPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionMinBidIncrementPercentageUpdated represents a AuctionMinBidIncrementPercentageUpdated event raised by the SoulContract contract.
type SoulContractAuctionMinBidIncrementPercentageUpdated struct {
	MinBidIncrementPercentage *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterAuctionMinBidIncrementPercentageUpdated is a free log retrieval operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_SoulContract *SoulContractFilterer) FilterAuctionMinBidIncrementPercentageUpdated(opts *bind.FilterOpts) (*SoulContractAuctionMinBidIncrementPercentageUpdatedIterator, error) {

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionMinBidIncrementPercentageUpdatedIterator{contract: _SoulContract.contract, event: "AuctionMinBidIncrementPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionMinBidIncrementPercentageUpdated is a free log subscription operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_SoulContract *SoulContractFilterer) WatchAuctionMinBidIncrementPercentageUpdated(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionMinBidIncrementPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionMinBidIncrementPercentageUpdated)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
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

// ParseAuctionMinBidIncrementPercentageUpdated is a log parse operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_SoulContract *SoulContractFilterer) ParseAuctionMinBidIncrementPercentageUpdated(log types.Log) (*SoulContractAuctionMinBidIncrementPercentageUpdated, error) {
	event := new(SoulContractAuctionMinBidIncrementPercentageUpdated)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionReservePriceUpdatedIterator is returned from FilterAuctionReservePriceUpdated and is used to iterate over the raw logs and unpacked data for AuctionReservePriceUpdated events raised by the SoulContract contract.
type SoulContractAuctionReservePriceUpdatedIterator struct {
	Event *SoulContractAuctionReservePriceUpdated // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionReservePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionReservePriceUpdated)
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
		it.Event = new(SoulContractAuctionReservePriceUpdated)
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
func (it *SoulContractAuctionReservePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionReservePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionReservePriceUpdated represents a AuctionReservePriceUpdated event raised by the SoulContract contract.
type SoulContractAuctionReservePriceUpdated struct {
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionReservePriceUpdated is a free log retrieval operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_SoulContract *SoulContractFilterer) FilterAuctionReservePriceUpdated(opts *bind.FilterOpts) (*SoulContractAuctionReservePriceUpdatedIterator, error) {

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionReservePriceUpdatedIterator{contract: _SoulContract.contract, event: "AuctionReservePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionReservePriceUpdated is a free log subscription operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_SoulContract *SoulContractFilterer) WatchAuctionReservePriceUpdated(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionReservePriceUpdated) (event.Subscription, error) {

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionReservePriceUpdated)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
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

// ParseAuctionReservePriceUpdated is a log parse operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_SoulContract *SoulContractFilterer) ParseAuctionReservePriceUpdated(log types.Log) (*SoulContractAuctionReservePriceUpdated, error) {
	event := new(SoulContractAuctionReservePriceUpdated)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the SoulContract contract.
type SoulContractAuctionSettledIterator struct {
	Event *SoulContractAuctionSettled // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionSettled)
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
		it.Event = new(SoulContractAuctionSettled)
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
func (it *SoulContractAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionSettled represents a AuctionSettled event raised by the SoulContract contract.
type SoulContractAuctionSettled struct {
	TokenId *big.Int
	Winner  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed tokenId, address winner, uint256 amount)
func (_SoulContract *SoulContractFilterer) FilterAuctionSettled(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulContractAuctionSettledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionSettled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionSettledIterator{contract: _SoulContract.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed tokenId, address winner, uint256 amount)
func (_SoulContract *SoulContractFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionSettled, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionSettled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionSettled)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed tokenId, address winner, uint256 amount)
func (_SoulContract *SoulContractFilterer) ParseAuctionSettled(log types.Log) (*SoulContractAuctionSettled, error) {
	event := new(SoulContractAuctionSettled)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractAuctionTimeBufferUpdatedIterator is returned from FilterAuctionTimeBufferUpdated and is used to iterate over the raw logs and unpacked data for AuctionTimeBufferUpdated events raised by the SoulContract contract.
type SoulContractAuctionTimeBufferUpdatedIterator struct {
	Event *SoulContractAuctionTimeBufferUpdated // Event containing the contract specifics and raw log

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
func (it *SoulContractAuctionTimeBufferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractAuctionTimeBufferUpdated)
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
		it.Event = new(SoulContractAuctionTimeBufferUpdated)
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
func (it *SoulContractAuctionTimeBufferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractAuctionTimeBufferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractAuctionTimeBufferUpdated represents a AuctionTimeBufferUpdated event raised by the SoulContract contract.
type SoulContractAuctionTimeBufferUpdated struct {
	TimeBuffer *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionTimeBufferUpdated is a free log retrieval operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_SoulContract *SoulContractFilterer) FilterAuctionTimeBufferUpdated(opts *bind.FilterOpts) (*SoulContractAuctionTimeBufferUpdatedIterator, error) {

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return &SoulContractAuctionTimeBufferUpdatedIterator{contract: _SoulContract.contract, event: "AuctionTimeBufferUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionTimeBufferUpdated is a free log subscription operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_SoulContract *SoulContractFilterer) WatchAuctionTimeBufferUpdated(opts *bind.WatchOpts, sink chan<- *SoulContractAuctionTimeBufferUpdated) (event.Subscription, error) {

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractAuctionTimeBufferUpdated)
				if err := _SoulContract.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
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

// ParseAuctionTimeBufferUpdated is a log parse operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_SoulContract *SoulContractFilterer) ParseAuctionTimeBufferUpdated(log types.Log) (*SoulContractAuctionTimeBufferUpdated, error) {
	event := new(SoulContractAuctionTimeBufferUpdated)
	if err := _SoulContract.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the SoulContract contract.
type SoulContractClaimIterator struct {
	Event *SoulContractClaim // Event containing the contract specifics and raw log

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
func (it *SoulContractClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractClaim)
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
		it.Event = new(SoulContractClaim)
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
func (it *SoulContractClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractClaim represents a Claim event raised by the SoulContract contract.
type SoulContractClaim struct {
	Reserver    common.Address
	TokenId     *big.Int
	Owner       common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0xad03f837a9207e368d73ec028e1f54428184da8cfea258cc116da2225f3ac5eb.
//
// Solidity: event Claim(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_SoulContract *SoulContractFilterer) FilterClaim(opts *bind.FilterOpts, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (*SoulContractClaimIterator, error) {

	var reserverRule []interface{}
	for _, reserverItem := range reserver {
		reserverRule = append(reserverRule, reserverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Claim", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractClaimIterator{contract: _SoulContract.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0xad03f837a9207e368d73ec028e1f54428184da8cfea258cc116da2225f3ac5eb.
//
// Solidity: event Claim(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_SoulContract *SoulContractFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *SoulContractClaim, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var reserverRule []interface{}
	for _, reserverItem := range reserver {
		reserverRule = append(reserverRule, reserverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Claim", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractClaim)
				if err := _SoulContract.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0xad03f837a9207e368d73ec028e1f54428184da8cfea258cc116da2225f3ac5eb.
//
// Solidity: event Claim(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_SoulContract *SoulContractFilterer) ParseClaim(log types.Log) (*SoulContractClaim, error) {
	event := new(SoulContractClaim)
	if err := _SoulContract.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SoulContract contract.
type SoulContractInitializedIterator struct {
	Event *SoulContractInitialized // Event containing the contract specifics and raw log

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
func (it *SoulContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractInitialized)
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
		it.Event = new(SoulContractInitialized)
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
func (it *SoulContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractInitialized represents a Initialized event raised by the SoulContract contract.
type SoulContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SoulContract *SoulContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*SoulContractInitializedIterator, error) {

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SoulContractInitializedIterator{contract: _SoulContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SoulContract *SoulContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SoulContractInitialized) (event.Subscription, error) {

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractInitialized)
				if err := _SoulContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SoulContract *SoulContractFilterer) ParseInitialized(log types.Log) (*SoulContractInitialized, error) {
	event := new(SoulContractInitialized)
	if err := _SoulContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SoulContract contract.
type SoulContractOwnershipTransferredIterator struct {
	Event *SoulContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SoulContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractOwnershipTransferred)
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
		it.Event = new(SoulContractOwnershipTransferred)
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
func (it *SoulContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractOwnershipTransferred represents a OwnershipTransferred event raised by the SoulContract contract.
type SoulContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoulContract *SoulContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SoulContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractOwnershipTransferredIterator{contract: _SoulContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoulContract *SoulContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SoulContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractOwnershipTransferred)
				if err := _SoulContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SoulContract *SoulContractFilterer) ParseOwnershipTransferred(log types.Log) (*SoulContractOwnershipTransferred, error) {
	event := new(SoulContractOwnershipTransferred)
	if err := _SoulContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SoulContract contract.
type SoulContractPausedIterator struct {
	Event *SoulContractPaused // Event containing the contract specifics and raw log

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
func (it *SoulContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractPaused)
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
		it.Event = new(SoulContractPaused)
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
func (it *SoulContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractPaused represents a Paused event raised by the SoulContract contract.
type SoulContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SoulContract *SoulContractFilterer) FilterPaused(opts *bind.FilterOpts) (*SoulContractPausedIterator, error) {

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SoulContractPausedIterator{contract: _SoulContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SoulContract *SoulContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SoulContractPaused) (event.Subscription, error) {

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractPaused)
				if err := _SoulContract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SoulContract *SoulContractFilterer) ParsePaused(log types.Log) (*SoulContractPaused, error) {
	event := new(SoulContractPaused)
	if err := _SoulContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractReserveIterator is returned from FilterReserve and is used to iterate over the raw logs and unpacked data for Reserve events raised by the SoulContract contract.
type SoulContractReserveIterator struct {
	Event *SoulContractReserve // Event containing the contract specifics and raw log

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
func (it *SoulContractReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractReserve)
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
		it.Event = new(SoulContractReserve)
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
func (it *SoulContractReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractReserve represents a Reserve event raised by the SoulContract contract.
type SoulContractReserve struct {
	Reserver    common.Address
	TokenId     *big.Int
	Owner       common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReserve is a free log retrieval operation binding the contract event 0x6f8d4fbb9604038077481649590dd5ddb6c97e7236b4b644b44084733d5f8560.
//
// Solidity: event Reserve(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_SoulContract *SoulContractFilterer) FilterReserve(opts *bind.FilterOpts, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (*SoulContractReserveIterator, error) {

	var reserverRule []interface{}
	for _, reserverItem := range reserver {
		reserverRule = append(reserverRule, reserverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Reserve", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractReserveIterator{contract: _SoulContract.contract, event: "Reserve", logs: logs, sub: sub}, nil
}

// WatchReserve is a free log subscription operation binding the contract event 0x6f8d4fbb9604038077481649590dd5ddb6c97e7236b4b644b44084733d5f8560.
//
// Solidity: event Reserve(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_SoulContract *SoulContractFilterer) WatchReserve(opts *bind.WatchOpts, sink chan<- *SoulContractReserve, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var reserverRule []interface{}
	for _, reserverItem := range reserver {
		reserverRule = append(reserverRule, reserverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Reserve", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractReserve)
				if err := _SoulContract.contract.UnpackLog(event, "Reserve", log); err != nil {
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

// ParseReserve is a log parse operation binding the contract event 0x6f8d4fbb9604038077481649590dd5ddb6c97e7236b4b644b44084733d5f8560.
//
// Solidity: event Reserve(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_SoulContract *SoulContractFilterer) ParseReserve(log types.Log) (*SoulContractReserve, error) {
	event := new(SoulContractReserve)
	if err := _SoulContract.contract.UnpackLog(event, "Reserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SoulContract contract.
type SoulContractTransferIterator struct {
	Event *SoulContractTransfer // Event containing the contract specifics and raw log

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
func (it *SoulContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractTransfer)
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
		it.Event = new(SoulContractTransfer)
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
func (it *SoulContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractTransfer represents a Transfer event raised by the SoulContract contract.
type SoulContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SoulContractTransferIterator, error) {

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

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulContractTransferIterator{contract: _SoulContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SoulContract *SoulContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SoulContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractTransfer)
				if err := _SoulContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SoulContract *SoulContractFilterer) ParseTransfer(log types.Log) (*SoulContractTransfer, error) {
	event := new(SoulContractTransfer)
	if err := _SoulContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SoulContract contract.
type SoulContractUnpausedIterator struct {
	Event *SoulContractUnpaused // Event containing the contract specifics and raw log

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
func (it *SoulContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulContractUnpaused)
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
		it.Event = new(SoulContractUnpaused)
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
func (it *SoulContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulContractUnpaused represents a Unpaused event raised by the SoulContract contract.
type SoulContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SoulContract *SoulContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SoulContractUnpausedIterator, error) {

	logs, sub, err := _SoulContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SoulContractUnpausedIterator{contract: _SoulContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SoulContract *SoulContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SoulContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _SoulContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulContractUnpaused)
				if err := _SoulContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SoulContract *SoulContractFilterer) ParseUnpaused(log types.Log) (*SoulContractUnpaused, error) {
	event := new(SoulContractUnpaused)
	if err := _SoulContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
