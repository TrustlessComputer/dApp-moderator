// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package soul

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

// AuctionHouseAuction is an auto generated low-level Go binding around an user-defined struct.
type AuctionHouseAuction struct {
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
	AuctionId                 [32]byte
}

// SoulMetaData contains all meta data concerning the Soul contract.
var SoulMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structAuctionHouse.Auction\",\"name\":\"auction\",\"type\":\"tuple\"}],\"name\":\"HandleAuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionClaimBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structAuctionHouse.Auction\",\"name\":\"auction\",\"type\":\"tuple\"}],\"name\":\"HandleAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structAuctionHouse.Auction\",\"name\":\"auction\",\"type\":\"tuple\"}],\"name\":\"AuctionExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"}],\"name\":\"AuctionMinBidIncrementPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"AuctionReservePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structAuctionHouse.Auction\",\"name\":\"auction\",\"type\":\"tuple\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"}],\"name\":\"AuctionTimeBufferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"Reserve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_auctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_auctionsList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bfs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_bidderAuctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_coreTeamTreasury\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gmToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_mintAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_minted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_paramsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_randomizerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_script\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_signerMint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"batchMint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"biddable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdm\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBfs\",\"type\":\"address\"}],\"name\":\"changeBfs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBrc20\",\"type\":\"address\"}],\"name\":\"changeBrc20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"changeParamAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"changeRandomizerAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newScript\",\"type\":\"string\"}],\"name\":\"changeScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdd\",\"type\":\"address\"}],\"name\":\"changeSignerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"claimBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalGM\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paramsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"randomizerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bfs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signerMint\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalGM\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"p5jsScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"projectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenHTML\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"variableScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"web3Script\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SoulABI is the input ABI used to generate the binding from.
// Deprecated: Use SoulMetaData.ABI instead.
var SoulABI = SoulMetaData.ABI

// Soul is an auto generated Go binding around an Ethereum contract.
type Soul struct {
	SoulCaller     // Read-only binding to the contract
	SoulTransactor // Write-only binding to the contract
	SoulFilterer   // Log filterer for contract events
}

// SoulCaller is an auto generated read-only Go binding around an Ethereum contract.
type SoulCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SoulTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SoulFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoulSession struct {
	Contract     *Soul             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SoulCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SoulCallerSession struct {
	Contract *SoulCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SoulTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SoulTransactorSession struct {
	Contract     *SoulTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SoulRaw is an auto generated low-level Go binding around an Ethereum contract.
type SoulRaw struct {
	Contract *Soul // Generic contract binding to access the raw methods on
}

// SoulCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SoulCallerRaw struct {
	Contract *SoulCaller // Generic read-only contract binding to access the raw methods on
}

// SoulTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SoulTransactorRaw struct {
	Contract *SoulTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSoul creates a new instance of Soul, bound to a specific deployed contract.
func NewSoul(address common.Address, backend bind.ContractBackend) (*Soul, error) {
	contract, err := bindSoul(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Soul{SoulCaller: SoulCaller{contract: contract}, SoulTransactor: SoulTransactor{contract: contract}, SoulFilterer: SoulFilterer{contract: contract}}, nil
}

// NewSoulCaller creates a new read-only instance of Soul, bound to a specific deployed contract.
func NewSoulCaller(address common.Address, caller bind.ContractCaller) (*SoulCaller, error) {
	contract, err := bindSoul(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SoulCaller{contract: contract}, nil
}

// NewSoulTransactor creates a new write-only instance of Soul, bound to a specific deployed contract.
func NewSoulTransactor(address common.Address, transactor bind.ContractTransactor) (*SoulTransactor, error) {
	contract, err := bindSoul(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SoulTransactor{contract: contract}, nil
}

// NewSoulFilterer creates a new log filterer instance of Soul, bound to a specific deployed contract.
func NewSoulFilterer(address common.Address, filterer bind.ContractFilterer) (*SoulFilterer, error) {
	contract, err := bindSoul(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SoulFilterer{contract: contract}, nil
}

// bindSoul binds a generic wrapper to an already deployed contract.
func bindSoul(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SoulMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Soul *SoulRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Soul.Contract.SoulCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Soul *SoulRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Soul.Contract.SoulTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Soul *SoulRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Soul.Contract.SoulTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Soul *SoulCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Soul.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Soul *SoulTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Soul.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Soul *SoulTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Soul.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x01bc45c9.
//
// Solidity: function _admin() view returns(address)
func (_Soul *SoulCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x01bc45c9.
//
// Solidity: function _admin() view returns(address)
func (_Soul *SoulSession) Admin() (common.Address, error) {
	return _Soul.Contract.Admin(&_Soul.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0x01bc45c9.
//
// Solidity: function _admin() view returns(address)
func (_Soul *SoulCallerSession) Admin() (common.Address, error) {
	return _Soul.Contract.Admin(&_Soul.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x17dc6bf6.
//
// Solidity: function _auctions(uint256 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage, bytes32 auctionId)
func (_Soul *SoulCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	AuctionId                 [32]byte
}, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_auctions", arg0)

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
		AuctionId                 [32]byte
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
	outstruct.AuctionId = *abi.ConvertType(out[10], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x17dc6bf6.
//
// Solidity: function _auctions(uint256 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage, bytes32 auctionId)
func (_Soul *SoulSession) Auctions(arg0 *big.Int) (struct {
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
	AuctionId                 [32]byte
}, error) {
	return _Soul.Contract.Auctions(&_Soul.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x17dc6bf6.
//
// Solidity: function _auctions(uint256 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage, bytes32 auctionId)
func (_Soul *SoulCallerSession) Auctions(arg0 *big.Int) (struct {
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
	AuctionId                 [32]byte
}, error) {
	return _Soul.Contract.Auctions(&_Soul.CallOpts, arg0)
}

// AuctionsList is a free data retrieval call binding the contract method 0x9a247cb3.
//
// Solidity: function _auctionsList(bytes32 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage, bytes32 auctionId)
func (_Soul *SoulCaller) AuctionsList(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	AuctionId                 [32]byte
}, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_auctionsList", arg0)

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
		AuctionId                 [32]byte
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
	outstruct.AuctionId = *abi.ConvertType(out[10], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// AuctionsList is a free data retrieval call binding the contract method 0x9a247cb3.
//
// Solidity: function _auctionsList(bytes32 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage, bytes32 auctionId)
func (_Soul *SoulSession) AuctionsList(arg0 [32]byte) (struct {
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
	AuctionId                 [32]byte
}, error) {
	return _Soul.Contract.AuctionsList(&_Soul.CallOpts, arg0)
}

// AuctionsList is a free data retrieval call binding the contract method 0x9a247cb3.
//
// Solidity: function _auctionsList(bytes32 ) view returns(uint256 tokenId, address erc20Token, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled, uint256 timeBuffer, uint256 reservePrice, uint256 minBidIncrementPercentage, bytes32 auctionId)
func (_Soul *SoulCallerSession) AuctionsList(arg0 [32]byte) (struct {
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
	AuctionId                 [32]byte
}, error) {
	return _Soul.Contract.AuctionsList(&_Soul.CallOpts, arg0)
}

// Bfs is a free data retrieval call binding the contract method 0x33348979.
//
// Solidity: function _bfs() view returns(address)
func (_Soul *SoulCaller) Bfs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_bfs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bfs is a free data retrieval call binding the contract method 0x33348979.
//
// Solidity: function _bfs() view returns(address)
func (_Soul *SoulSession) Bfs() (common.Address, error) {
	return _Soul.Contract.Bfs(&_Soul.CallOpts)
}

// Bfs is a free data retrieval call binding the contract method 0x33348979.
//
// Solidity: function _bfs() view returns(address)
func (_Soul *SoulCallerSession) Bfs() (common.Address, error) {
	return _Soul.Contract.Bfs(&_Soul.CallOpts)
}

// BidderAuctions is a free data retrieval call binding the contract method 0x55cd6326.
//
// Solidity: function _bidderAuctions(uint256 , bytes32 , address ) view returns(uint256)
func (_Soul *SoulCaller) BidderAuctions(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_bidderAuctions", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidderAuctions is a free data retrieval call binding the contract method 0x55cd6326.
//
// Solidity: function _bidderAuctions(uint256 , bytes32 , address ) view returns(uint256)
func (_Soul *SoulSession) BidderAuctions(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _Soul.Contract.BidderAuctions(&_Soul.CallOpts, arg0, arg1, arg2)
}

// BidderAuctions is a free data retrieval call binding the contract method 0x55cd6326.
//
// Solidity: function _bidderAuctions(uint256 , bytes32 , address ) view returns(uint256)
func (_Soul *SoulCallerSession) BidderAuctions(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _Soul.Contract.BidderAuctions(&_Soul.CallOpts, arg0, arg1, arg2)
}

// CoreTeamTreasury is a free data retrieval call binding the contract method 0x2824d162.
//
// Solidity: function _coreTeamTreasury(address ) view returns(uint256)
func (_Soul *SoulCaller) CoreTeamTreasury(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_coreTeamTreasury", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoreTeamTreasury is a free data retrieval call binding the contract method 0x2824d162.
//
// Solidity: function _coreTeamTreasury(address ) view returns(uint256)
func (_Soul *SoulSession) CoreTeamTreasury(arg0 common.Address) (*big.Int, error) {
	return _Soul.Contract.CoreTeamTreasury(&_Soul.CallOpts, arg0)
}

// CoreTeamTreasury is a free data retrieval call binding the contract method 0x2824d162.
//
// Solidity: function _coreTeamTreasury(address ) view returns(uint256)
func (_Soul *SoulCallerSession) CoreTeamTreasury(arg0 common.Address) (*big.Int, error) {
	return _Soul.Contract.CoreTeamTreasury(&_Soul.CallOpts, arg0)
}

// GmToken is a free data retrieval call binding the contract method 0x9c6f35f9.
//
// Solidity: function _gmToken() view returns(address)
func (_Soul *SoulCaller) GmToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_gmToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmToken is a free data retrieval call binding the contract method 0x9c6f35f9.
//
// Solidity: function _gmToken() view returns(address)
func (_Soul *SoulSession) GmToken() (common.Address, error) {
	return _Soul.Contract.GmToken(&_Soul.CallOpts)
}

// GmToken is a free data retrieval call binding the contract method 0x9c6f35f9.
//
// Solidity: function _gmToken() view returns(address)
func (_Soul *SoulCallerSession) GmToken() (common.Address, error) {
	return _Soul.Contract.GmToken(&_Soul.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0x22f4596f.
//
// Solidity: function _maxSupply() view returns(uint256)
func (_Soul *SoulCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0x22f4596f.
//
// Solidity: function _maxSupply() view returns(uint256)
func (_Soul *SoulSession) MaxSupply() (*big.Int, error) {
	return _Soul.Contract.MaxSupply(&_Soul.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0x22f4596f.
//
// Solidity: function _maxSupply() view returns(uint256)
func (_Soul *SoulCallerSession) MaxSupply() (*big.Int, error) {
	return _Soul.Contract.MaxSupply(&_Soul.CallOpts)
}

// MintAt is a free data retrieval call binding the contract method 0xcc22e004.
//
// Solidity: function _mintAt(uint256 ) view returns(uint256)
func (_Soul *SoulCaller) MintAt(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_mintAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintAt is a free data retrieval call binding the contract method 0xcc22e004.
//
// Solidity: function _mintAt(uint256 ) view returns(uint256)
func (_Soul *SoulSession) MintAt(arg0 *big.Int) (*big.Int, error) {
	return _Soul.Contract.MintAt(&_Soul.CallOpts, arg0)
}

// MintAt is a free data retrieval call binding the contract method 0xcc22e004.
//
// Solidity: function _mintAt(uint256 ) view returns(uint256)
func (_Soul *SoulCallerSession) MintAt(arg0 *big.Int) (*big.Int, error) {
	return _Soul.Contract.MintAt(&_Soul.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x7de77ecc.
//
// Solidity: function _minted(address ) view returns(uint256)
func (_Soul *SoulCaller) Minted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_minted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x7de77ecc.
//
// Solidity: function _minted(address ) view returns(uint256)
func (_Soul *SoulSession) Minted(arg0 common.Address) (*big.Int, error) {
	return _Soul.Contract.Minted(&_Soul.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x7de77ecc.
//
// Solidity: function _minted(address ) view returns(uint256)
func (_Soul *SoulCallerSession) Minted(arg0 common.Address) (*big.Int, error) {
	return _Soul.Contract.Minted(&_Soul.CallOpts, arg0)
}

// ParamsAddress is a free data retrieval call binding the contract method 0xadfc7dae.
//
// Solidity: function _paramsAddress() view returns(address)
func (_Soul *SoulCaller) ParamsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_paramsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParamsAddress is a free data retrieval call binding the contract method 0xadfc7dae.
//
// Solidity: function _paramsAddress() view returns(address)
func (_Soul *SoulSession) ParamsAddress() (common.Address, error) {
	return _Soul.Contract.ParamsAddress(&_Soul.CallOpts)
}

// ParamsAddress is a free data retrieval call binding the contract method 0xadfc7dae.
//
// Solidity: function _paramsAddress() view returns(address)
func (_Soul *SoulCallerSession) ParamsAddress() (common.Address, error) {
	return _Soul.Contract.ParamsAddress(&_Soul.CallOpts)
}

// RandomizerAddr is a free data retrieval call binding the contract method 0x66215eb4.
//
// Solidity: function _randomizerAddr() view returns(address)
func (_Soul *SoulCaller) RandomizerAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_randomizerAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomizerAddr is a free data retrieval call binding the contract method 0x66215eb4.
//
// Solidity: function _randomizerAddr() view returns(address)
func (_Soul *SoulSession) RandomizerAddr() (common.Address, error) {
	return _Soul.Contract.RandomizerAddr(&_Soul.CallOpts)
}

// RandomizerAddr is a free data retrieval call binding the contract method 0x66215eb4.
//
// Solidity: function _randomizerAddr() view returns(address)
func (_Soul *SoulCallerSession) RandomizerAddr() (common.Address, error) {
	return _Soul.Contract.RandomizerAddr(&_Soul.CallOpts)
}

// Script is a free data retrieval call binding the contract method 0x8a016249.
//
// Solidity: function _script() view returns(string)
func (_Soul *SoulCaller) Script(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_script")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Script is a free data retrieval call binding the contract method 0x8a016249.
//
// Solidity: function _script() view returns(string)
func (_Soul *SoulSession) Script() (string, error) {
	return _Soul.Contract.Script(&_Soul.CallOpts)
}

// Script is a free data retrieval call binding the contract method 0x8a016249.
//
// Solidity: function _script() view returns(string)
func (_Soul *SoulCallerSession) Script() (string, error) {
	return _Soul.Contract.Script(&_Soul.CallOpts)
}

// SignerMint is a free data retrieval call binding the contract method 0xbd896843.
//
// Solidity: function _signerMint() view returns(address)
func (_Soul *SoulCaller) SignerMint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "_signerMint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerMint is a free data retrieval call binding the contract method 0xbd896843.
//
// Solidity: function _signerMint() view returns(address)
func (_Soul *SoulSession) SignerMint() (common.Address, error) {
	return _Soul.Contract.SignerMint(&_Soul.CallOpts)
}

// SignerMint is a free data retrieval call binding the contract method 0xbd896843.
//
// Solidity: function _signerMint() view returns(address)
func (_Soul *SoulCallerSession) SignerMint() (common.Address, error) {
	return _Soul.Contract.SignerMint(&_Soul.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_Soul *SoulCaller) Available(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "available", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_Soul *SoulSession) Available(tokenId *big.Int) (bool, error) {
	return _Soul.Contract.Available(&_Soul.CallOpts, tokenId)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_Soul *SoulCallerSession) Available(tokenId *big.Int) (bool, error) {
	return _Soul.Contract.Available(&_Soul.CallOpts, tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Soul *SoulCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Soul *SoulSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Soul.Contract.BalanceOf(&_Soul.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Soul *SoulCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Soul.Contract.BalanceOf(&_Soul.CallOpts, owner)
}

// Biddable is a free data retrieval call binding the contract method 0x68a17bd8.
//
// Solidity: function biddable(uint256 tokenId) view returns(bool)
func (_Soul *SoulCaller) Biddable(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "biddable", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Biddable is a free data retrieval call binding the contract method 0x68a17bd8.
//
// Solidity: function biddable(uint256 tokenId) view returns(bool)
func (_Soul *SoulSession) Biddable(tokenId *big.Int) (bool, error) {
	return _Soul.Contract.Biddable(&_Soul.CallOpts, tokenId)
}

// Biddable is a free data retrieval call binding the contract method 0x68a17bd8.
//
// Solidity: function biddable(uint256 tokenId) view returns(bool)
func (_Soul *SoulCallerSession) Biddable(tokenId *big.Int) (bool, error) {
	return _Soul.Contract.Biddable(&_Soul.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Soul *SoulCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Soul *SoulSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Soul.Contract.GetApproved(&_Soul.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Soul *SoulCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Soul.Contract.GetApproved(&_Soul.CallOpts, tokenId)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_Soul *SoulCaller) GetMessageHash(opts *bind.CallOpts, user common.Address, totalGM *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "getMessageHash", user, totalGM)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_Soul *SoulSession) GetMessageHash(user common.Address, totalGM *big.Int) ([32]byte, error) {
	return _Soul.Contract.GetMessageHash(&_Soul.CallOpts, user, totalGM)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x512c91df.
//
// Solidity: function getMessageHash(address user, uint256 totalGM) view returns(bytes32)
func (_Soul *SoulCallerSession) GetMessageHash(user common.Address, totalGM *big.Int) ([32]byte, error) {
	return _Soul.Contract.GetMessageHash(&_Soul.CallOpts, user, totalGM)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Soul *SoulCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Soul *SoulSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Soul.Contract.IsApprovedForAll(&_Soul.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Soul *SoulCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Soul.Contract.IsApprovedForAll(&_Soul.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Soul *SoulCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Soul *SoulSession) Name() (string, error) {
	return _Soul.Contract.Name(&_Soul.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Soul *SoulCallerSession) Name() (string, error) {
	return _Soul.Contract.Name(&_Soul.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Soul *SoulCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Soul *SoulSession) Owner() (common.Address, error) {
	return _Soul.Contract.Owner(&_Soul.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Soul *SoulCallerSession) Owner() (common.Address, error) {
	return _Soul.Contract.Owner(&_Soul.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Soul *SoulCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Soul *SoulSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Soul.Contract.OwnerOf(&_Soul.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Soul *SoulCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Soul.Contract.OwnerOf(&_Soul.CallOpts, tokenId)
}

// P5jsScript is a free data retrieval call binding the contract method 0x74973a41.
//
// Solidity: function p5jsScript() view returns(string result)
func (_Soul *SoulCaller) P5jsScript(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "p5jsScript")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// P5jsScript is a free data retrieval call binding the contract method 0x74973a41.
//
// Solidity: function p5jsScript() view returns(string result)
func (_Soul *SoulSession) P5jsScript() (string, error) {
	return _Soul.Contract.P5jsScript(&_Soul.CallOpts)
}

// P5jsScript is a free data retrieval call binding the contract method 0x74973a41.
//
// Solidity: function p5jsScript() view returns(string result)
func (_Soul *SoulCallerSession) P5jsScript() (string, error) {
	return _Soul.Contract.P5jsScript(&_Soul.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Soul *SoulCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Soul *SoulSession) Paused() (bool, error) {
	return _Soul.Contract.Paused(&_Soul.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Soul *SoulCallerSession) Paused() (bool, error) {
	return _Soul.Contract.Paused(&_Soul.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 projectId, uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Soul *SoulCaller) RoyaltyInfo(opts *bind.CallOpts, projectId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "royaltyInfo", projectId, _salePrice)

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
func (_Soul *SoulSession) RoyaltyInfo(projectId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Soul.Contract.RoyaltyInfo(&_Soul.CallOpts, projectId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 projectId, uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Soul *SoulCallerSession) RoyaltyInfo(projectId *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Soul.Contract.RoyaltyInfo(&_Soul.CallOpts, projectId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Soul *SoulCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Soul *SoulSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Soul.Contract.SupportsInterface(&_Soul.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Soul *SoulCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Soul.Contract.SupportsInterface(&_Soul.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Soul *SoulCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Soul *SoulSession) Symbol() (string, error) {
	return _Soul.Contract.Symbol(&_Soul.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Soul *SoulCallerSession) Symbol() (string, error) {
	return _Soul.Contract.Symbol(&_Soul.CallOpts)
}

// TokenHTML is a free data retrieval call binding the contract method 0xeb29c232.
//
// Solidity: function tokenHTML(bytes32 seed, uint256 tokenId) view returns(string result)
func (_Soul *SoulCaller) TokenHTML(opts *bind.CallOpts, seed [32]byte, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "tokenHTML", seed, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenHTML is a free data retrieval call binding the contract method 0xeb29c232.
//
// Solidity: function tokenHTML(bytes32 seed, uint256 tokenId) view returns(string result)
func (_Soul *SoulSession) TokenHTML(seed [32]byte, tokenId *big.Int) (string, error) {
	return _Soul.Contract.TokenHTML(&_Soul.CallOpts, seed, tokenId)
}

// TokenHTML is a free data retrieval call binding the contract method 0xeb29c232.
//
// Solidity: function tokenHTML(bytes32 seed, uint256 tokenId) view returns(string result)
func (_Soul *SoulCallerSession) TokenHTML(seed [32]byte, tokenId *big.Int) (string, error) {
	return _Soul.Contract.TokenHTML(&_Soul.CallOpts, seed, tokenId)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 tokenId) view returns(bytes32)
func (_Soul *SoulCaller) TokenIdToHash(opts *bind.CallOpts, tokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "tokenIdToHash", tokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 tokenId) view returns(bytes32)
func (_Soul *SoulSession) TokenIdToHash(tokenId *big.Int) ([32]byte, error) {
	return _Soul.Contract.TokenIdToHash(&_Soul.CallOpts, tokenId)
}

// TokenIdToHash is a free data retrieval call binding the contract method 0x621a1f74.
//
// Solidity: function tokenIdToHash(uint256 tokenId) view returns(bytes32)
func (_Soul *SoulCallerSession) TokenIdToHash(tokenId *big.Int) ([32]byte, error) {
	return _Soul.Contract.TokenIdToHash(&_Soul.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_Soul *SoulCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_Soul *SoulSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Soul.Contract.TokenURI(&_Soul.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string result)
func (_Soul *SoulCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Soul.Contract.TokenURI(&_Soul.CallOpts, tokenId)
}

// VariableScript is a free data retrieval call binding the contract method 0x32e32326.
//
// Solidity: function variableScript(bytes32 seed, uint256 tokenId) view returns(string result)
func (_Soul *SoulCaller) VariableScript(opts *bind.CallOpts, seed [32]byte, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "variableScript", seed, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VariableScript is a free data retrieval call binding the contract method 0x32e32326.
//
// Solidity: function variableScript(bytes32 seed, uint256 tokenId) view returns(string result)
func (_Soul *SoulSession) VariableScript(seed [32]byte, tokenId *big.Int) (string, error) {
	return _Soul.Contract.VariableScript(&_Soul.CallOpts, seed, tokenId)
}

// VariableScript is a free data retrieval call binding the contract method 0x32e32326.
//
// Solidity: function variableScript(bytes32 seed, uint256 tokenId) view returns(string result)
func (_Soul *SoulCallerSession) VariableScript(seed [32]byte, tokenId *big.Int) (string, error) {
	return _Soul.Contract.VariableScript(&_Soul.CallOpts, seed, tokenId)
}

// Web3Script is a free data retrieval call binding the contract method 0xa0d41e11.
//
// Solidity: function web3Script() view returns(string result)
func (_Soul *SoulCaller) Web3Script(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Soul.contract.Call(opts, &out, "web3Script")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Web3Script is a free data retrieval call binding the contract method 0xa0d41e11.
//
// Solidity: function web3Script() view returns(string result)
func (_Soul *SoulSession) Web3Script() (string, error) {
	return _Soul.Contract.Web3Script(&_Soul.CallOpts)
}

// Web3Script is a free data retrieval call binding the contract method 0xa0d41e11.
//
// Solidity: function web3Script() view returns(string result)
func (_Soul *SoulCallerSession) Web3Script() (string, error) {
	return _Soul.Contract.Web3Script(&_Soul.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Soul *SoulTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Soul *SoulSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.Approve(&_Soul.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Soul *SoulTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.Approve(&_Soul.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x624de2a9.
//
// Solidity: function batchMint(address to, uint256 n, bytes signatures) payable returns(uint256[])
func (_Soul *SoulTransactor) BatchMint(opts *bind.TransactOpts, to common.Address, n *big.Int, signatures []byte) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "batchMint", to, n, signatures)
}

// BatchMint is a paid mutator transaction binding the contract method 0x624de2a9.
//
// Solidity: function batchMint(address to, uint256 n, bytes signatures) payable returns(uint256[])
func (_Soul *SoulSession) BatchMint(to common.Address, n *big.Int, signatures []byte) (*types.Transaction, error) {
	return _Soul.Contract.BatchMint(&_Soul.TransactOpts, to, n, signatures)
}

// BatchMint is a paid mutator transaction binding the contract method 0x624de2a9.
//
// Solidity: function batchMint(address to, uint256 n, bytes signatures) payable returns(uint256[])
func (_Soul *SoulTransactorSession) BatchMint(to common.Address, n *big.Int, signatures []byte) (*types.Transaction, error) {
	return _Soul.Contract.BatchMint(&_Soul.TransactOpts, to, n, signatures)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdm) returns()
func (_Soul *SoulTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdm common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeAdmin", newAdm)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdm) returns()
func (_Soul *SoulSession) ChangeAdmin(newAdm common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeAdmin(&_Soul.TransactOpts, newAdm)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdm) returns()
func (_Soul *SoulTransactorSession) ChangeAdmin(newAdm common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeAdmin(&_Soul.TransactOpts, newAdm)
}

// ChangeBfs is a paid mutator transaction binding the contract method 0x2b7d13de.
//
// Solidity: function changeBfs(address newBfs) returns()
func (_Soul *SoulTransactor) ChangeBfs(opts *bind.TransactOpts, newBfs common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeBfs", newBfs)
}

// ChangeBfs is a paid mutator transaction binding the contract method 0x2b7d13de.
//
// Solidity: function changeBfs(address newBfs) returns()
func (_Soul *SoulSession) ChangeBfs(newBfs common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeBfs(&_Soul.TransactOpts, newBfs)
}

// ChangeBfs is a paid mutator transaction binding the contract method 0x2b7d13de.
//
// Solidity: function changeBfs(address newBfs) returns()
func (_Soul *SoulTransactorSession) ChangeBfs(newBfs common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeBfs(&_Soul.TransactOpts, newBfs)
}

// ChangeBrc20Token is a paid mutator transaction binding the contract method 0x93c6cdc0.
//
// Solidity: function changeBrc20Token(address newBrc20) returns()
func (_Soul *SoulTransactor) ChangeBrc20Token(opts *bind.TransactOpts, newBrc20 common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeBrc20Token", newBrc20)
}

// ChangeBrc20Token is a paid mutator transaction binding the contract method 0x93c6cdc0.
//
// Solidity: function changeBrc20Token(address newBrc20) returns()
func (_Soul *SoulSession) ChangeBrc20Token(newBrc20 common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeBrc20Token(&_Soul.TransactOpts, newBrc20)
}

// ChangeBrc20Token is a paid mutator transaction binding the contract method 0x93c6cdc0.
//
// Solidity: function changeBrc20Token(address newBrc20) returns()
func (_Soul *SoulTransactorSession) ChangeBrc20Token(newBrc20 common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeBrc20Token(&_Soul.TransactOpts, newBrc20)
}

// ChangeParamAddr is a paid mutator transaction binding the contract method 0xfebfd6c3.
//
// Solidity: function changeParamAddr(address newAddr) returns()
func (_Soul *SoulTransactor) ChangeParamAddr(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeParamAddr", newAddr)
}

// ChangeParamAddr is a paid mutator transaction binding the contract method 0xfebfd6c3.
//
// Solidity: function changeParamAddr(address newAddr) returns()
func (_Soul *SoulSession) ChangeParamAddr(newAddr common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeParamAddr(&_Soul.TransactOpts, newAddr)
}

// ChangeParamAddr is a paid mutator transaction binding the contract method 0xfebfd6c3.
//
// Solidity: function changeParamAddr(address newAddr) returns()
func (_Soul *SoulTransactorSession) ChangeParamAddr(newAddr common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeParamAddr(&_Soul.TransactOpts, newAddr)
}

// ChangeRandomizerAddr is a paid mutator transaction binding the contract method 0x1ca9741b.
//
// Solidity: function changeRandomizerAddr(address newAddr) returns()
func (_Soul *SoulTransactor) ChangeRandomizerAddr(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeRandomizerAddr", newAddr)
}

// ChangeRandomizerAddr is a paid mutator transaction binding the contract method 0x1ca9741b.
//
// Solidity: function changeRandomizerAddr(address newAddr) returns()
func (_Soul *SoulSession) ChangeRandomizerAddr(newAddr common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeRandomizerAddr(&_Soul.TransactOpts, newAddr)
}

// ChangeRandomizerAddr is a paid mutator transaction binding the contract method 0x1ca9741b.
//
// Solidity: function changeRandomizerAddr(address newAddr) returns()
func (_Soul *SoulTransactorSession) ChangeRandomizerAddr(newAddr common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeRandomizerAddr(&_Soul.TransactOpts, newAddr)
}

// ChangeScript is a paid mutator transaction binding the contract method 0x140f0e07.
//
// Solidity: function changeScript(string newScript) returns()
func (_Soul *SoulTransactor) ChangeScript(opts *bind.TransactOpts, newScript string) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeScript", newScript)
}

// ChangeScript is a paid mutator transaction binding the contract method 0x140f0e07.
//
// Solidity: function changeScript(string newScript) returns()
func (_Soul *SoulSession) ChangeScript(newScript string) (*types.Transaction, error) {
	return _Soul.Contract.ChangeScript(&_Soul.TransactOpts, newScript)
}

// ChangeScript is a paid mutator transaction binding the contract method 0x140f0e07.
//
// Solidity: function changeScript(string newScript) returns()
func (_Soul *SoulTransactorSession) ChangeScript(newScript string) (*types.Transaction, error) {
	return _Soul.Contract.ChangeScript(&_Soul.TransactOpts, newScript)
}

// ChangeSignerMint is a paid mutator transaction binding the contract method 0xf3f83ddd.
//
// Solidity: function changeSignerMint(address newAdd) returns()
func (_Soul *SoulTransactor) ChangeSignerMint(opts *bind.TransactOpts, newAdd common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "changeSignerMint", newAdd)
}

// ChangeSignerMint is a paid mutator transaction binding the contract method 0xf3f83ddd.
//
// Solidity: function changeSignerMint(address newAdd) returns()
func (_Soul *SoulSession) ChangeSignerMint(newAdd common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeSignerMint(&_Soul.TransactOpts, newAdd)
}

// ChangeSignerMint is a paid mutator transaction binding the contract method 0xf3f83ddd.
//
// Solidity: function changeSignerMint(address newAdd) returns()
func (_Soul *SoulTransactorSession) ChangeSignerMint(newAdd common.Address) (*types.Transaction, error) {
	return _Soul.Contract.ChangeSignerMint(&_Soul.TransactOpts, newAdd)
}

// ClaimBid is a paid mutator transaction binding the contract method 0xf50cbc8b.
//
// Solidity: function claimBid(uint256 tokenId, bytes32 auctionId) returns()
func (_Soul *SoulTransactor) ClaimBid(opts *bind.TransactOpts, tokenId *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "claimBid", tokenId, auctionId)
}

// ClaimBid is a paid mutator transaction binding the contract method 0xf50cbc8b.
//
// Solidity: function claimBid(uint256 tokenId, bytes32 auctionId) returns()
func (_Soul *SoulSession) ClaimBid(tokenId *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _Soul.Contract.ClaimBid(&_Soul.TransactOpts, tokenId, auctionId)
}

// ClaimBid is a paid mutator transaction binding the contract method 0xf50cbc8b.
//
// Solidity: function claimBid(uint256 tokenId, bytes32 auctionId) returns()
func (_Soul *SoulTransactorSession) ClaimBid(tokenId *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _Soul.Contract.ClaimBid(&_Soul.TransactOpts, tokenId, auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xd5563f31.
//
// Solidity: function createAuction(uint256 tokenId) returns()
func (_Soul *SoulTransactor) CreateAuction(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "createAuction", tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xd5563f31.
//
// Solidity: function createAuction(uint256 tokenId) returns()
func (_Soul *SoulSession) CreateAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.CreateAuction(&_Soul.TransactOpts, tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xd5563f31.
//
// Solidity: function createAuction(uint256 tokenId) returns()
func (_Soul *SoulTransactorSession) CreateAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.CreateAuction(&_Soul.TransactOpts, tokenId)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 tokenId, uint256 amount) payable returns()
func (_Soul *SoulTransactor) CreateBid(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "createBid", tokenId, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 tokenId, uint256 amount) payable returns()
func (_Soul *SoulSession) CreateBid(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.CreateBid(&_Soul.TransactOpts, tokenId, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xb7751c71.
//
// Solidity: function createBid(uint256 tokenId, uint256 amount) payable returns()
func (_Soul *SoulTransactorSession) CreateBid(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.CreateBid(&_Soul.TransactOpts, tokenId, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Soul *SoulTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Soul *SoulSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.Deposit(&_Soul.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Soul *SoulTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.Deposit(&_Soul.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x94f5a66e.
//
// Solidity: function initialize(string name, string symbol, address admin, address paramsAddress, address randomizerAddr, address gmToken, address bfs, address signerMint) returns()
func (_Soul *SoulTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, admin common.Address, paramsAddress common.Address, randomizerAddr common.Address, gmToken common.Address, bfs common.Address, signerMint common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "initialize", name, symbol, admin, paramsAddress, randomizerAddr, gmToken, bfs, signerMint)
}

// Initialize is a paid mutator transaction binding the contract method 0x94f5a66e.
//
// Solidity: function initialize(string name, string symbol, address admin, address paramsAddress, address randomizerAddr, address gmToken, address bfs, address signerMint) returns()
func (_Soul *SoulSession) Initialize(name string, symbol string, admin common.Address, paramsAddress common.Address, randomizerAddr common.Address, gmToken common.Address, bfs common.Address, signerMint common.Address) (*types.Transaction, error) {
	return _Soul.Contract.Initialize(&_Soul.TransactOpts, name, symbol, admin, paramsAddress, randomizerAddr, gmToken, bfs, signerMint)
}

// Initialize is a paid mutator transaction binding the contract method 0x94f5a66e.
//
// Solidity: function initialize(string name, string symbol, address admin, address paramsAddress, address randomizerAddr, address gmToken, address bfs, address signerMint) returns()
func (_Soul *SoulTransactorSession) Initialize(name string, symbol string, admin common.Address, paramsAddress common.Address, randomizerAddr common.Address, gmToken common.Address, bfs common.Address, signerMint common.Address) (*types.Transaction, error) {
	return _Soul.Contract.Initialize(&_Soul.TransactOpts, name, symbol, admin, paramsAddress, randomizerAddr, gmToken, bfs, signerMint)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address to, uint256 totalGM, bytes signature) payable returns(uint256 tokenId)
func (_Soul *SoulTransactor) Mint(opts *bind.TransactOpts, to common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "mint", to, totalGM, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address to, uint256 totalGM, bytes signature) payable returns(uint256 tokenId)
func (_Soul *SoulSession) Mint(to common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _Soul.Contract.Mint(&_Soul.TransactOpts, to, totalGM, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address to, uint256 totalGM, bytes signature) payable returns(uint256 tokenId)
func (_Soul *SoulTransactorSession) Mint(to common.Address, totalGM *big.Int, signature []byte) (*types.Transaction, error) {
	return _Soul.Contract.Mint(&_Soul.TransactOpts, to, totalGM, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Soul *SoulTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Soul *SoulSession) RenounceOwnership() (*types.Transaction, error) {
	return _Soul.Contract.RenounceOwnership(&_Soul.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Soul *SoulTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Soul.Contract.RenounceOwnership(&_Soul.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Soul *SoulTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Soul *SoulSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.SafeTransferFrom(&_Soul.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Soul *SoulTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.SafeTransferFrom(&_Soul.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Soul *SoulTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Soul *SoulSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Soul.Contract.SafeTransferFrom0(&_Soul.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Soul *SoulTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Soul.Contract.SafeTransferFrom0(&_Soul.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Soul *SoulTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Soul *SoulSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Soul.Contract.SetApprovalForAll(&_Soul.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Soul *SoulTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Soul.Contract.SetApprovalForAll(&_Soul.TransactOpts, operator, approved)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x2e993611.
//
// Solidity: function settleAuction(uint256 tokenId) returns()
func (_Soul *SoulTransactor) SettleAuction(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "settleAuction", tokenId)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x2e993611.
//
// Solidity: function settleAuction(uint256 tokenId) returns()
func (_Soul *SoulSession) SettleAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.SettleAuction(&_Soul.TransactOpts, tokenId)
}

// SettleAuction is a paid mutator transaction binding the contract method 0x2e993611.
//
// Solidity: function settleAuction(uint256 tokenId) returns()
func (_Soul *SoulTransactorSession) SettleAuction(tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.SettleAuction(&_Soul.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Soul *SoulTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Soul *SoulSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.TransferFrom(&_Soul.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Soul *SoulTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Soul.Contract.TransferFrom(&_Soul.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Soul *SoulTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Soul *SoulSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Soul.Contract.TransferOwnership(&_Soul.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Soul *SoulTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Soul.Contract.TransferOwnership(&_Soul.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address erc20Addr) returns()
func (_Soul *SoulTransactor) Withdraw(opts *bind.TransactOpts, erc20Addr common.Address) (*types.Transaction, error) {
	return _Soul.contract.Transact(opts, "withdraw", erc20Addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address erc20Addr) returns()
func (_Soul *SoulSession) Withdraw(erc20Addr common.Address) (*types.Transaction, error) {
	return _Soul.Contract.Withdraw(&_Soul.TransactOpts, erc20Addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address erc20Addr) returns()
func (_Soul *SoulTransactorSession) Withdraw(erc20Addr common.Address) (*types.Transaction, error) {
	return _Soul.Contract.Withdraw(&_Soul.TransactOpts, erc20Addr)
}

// SoulApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Soul contract.
type SoulApprovalIterator struct {
	Event *SoulApproval // Event containing the contract specifics and raw log

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
func (it *SoulApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulApproval)
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
		it.Event = new(SoulApproval)
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
func (it *SoulApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulApproval represents a Approval event raised by the Soul contract.
type SoulApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Soul *SoulFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SoulApprovalIterator, error) {

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

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulApprovalIterator{contract: _Soul.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Soul *SoulFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SoulApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulApproval)
				if err := _Soul.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Soul *SoulFilterer) ParseApproval(log types.Log) (*SoulApproval, error) {
	event := new(SoulApproval)
	if err := _Soul.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Soul contract.
type SoulApprovalForAllIterator struct {
	Event *SoulApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SoulApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulApprovalForAll)
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
		it.Event = new(SoulApprovalForAll)
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
func (it *SoulApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulApprovalForAll represents a ApprovalForAll event raised by the Soul contract.
type SoulApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Soul *SoulFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SoulApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SoulApprovalForAllIterator{contract: _Soul.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Soul *SoulFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SoulApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulApprovalForAll)
				if err := _Soul.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Soul *SoulFilterer) ParseApprovalForAll(log types.Log) (*SoulApprovalForAll, error) {
	event := new(SoulApprovalForAll)
	if err := _Soul.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for HandleAuctionBid events raised by the Soul contract.
type SoulAuctionBidIterator struct {
	Event *SoulAuctionBid // Event containing the contract specifics and raw log

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
func (it *SoulAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionBid)
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
		it.Event = new(SoulAuctionBid)
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
func (it *SoulAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionBid represents a HandleAuctionBid event raised by the Soul contract.
type SoulAuctionBid struct {
	TokenId  *big.Int
	Sender   common.Address
	Value    *big.Int
	Extended bool
	Auction  AuctionHouseAuction
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0xfd09e259e5c48be16f7dd63468e19266c684a9f3dce3b31881dbc1f909fc0d7c.
//
// Solidity: event HandleAuctionBid(uint256 indexed tokenId, address sender, uint256 value, bool extended, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) FilterAuctionBid(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulAuctionBidIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "HandleAuctionBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulAuctionBidIterator{contract: _Soul.contract, event: "HandleAuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0xfd09e259e5c48be16f7dd63468e19266c684a9f3dce3b31881dbc1f909fc0d7c.
//
// Solidity: event HandleAuctionBid(uint256 indexed tokenId, address sender, uint256 value, bool extended, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *SoulAuctionBid, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "HandleAuctionBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionBid)
				if err := _Soul.contract.UnpackLog(event, "HandleAuctionBid", log); err != nil {
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

// ParseAuctionBid is a log parse operation binding the contract event 0xfd09e259e5c48be16f7dd63468e19266c684a9f3dce3b31881dbc1f909fc0d7c.
//
// Solidity: event HandleAuctionBid(uint256 indexed tokenId, address sender, uint256 value, bool extended, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) ParseAuctionBid(log types.Log) (*SoulAuctionBid, error) {
	event := new(SoulAuctionBid)
	if err := _Soul.contract.UnpackLog(event, "HandleAuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionClaimBidIterator is returned from FilterAuctionClaimBid and is used to iterate over the raw logs and unpacked data for AuctionClaimBid events raised by the Soul contract.
type SoulAuctionClaimBidIterator struct {
	Event *SoulAuctionClaimBid // Event containing the contract specifics and raw log

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
func (it *SoulAuctionClaimBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionClaimBid)
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
		it.Event = new(SoulAuctionClaimBid)
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
func (it *SoulAuctionClaimBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionClaimBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionClaimBid represents a AuctionClaimBid event raised by the Soul contract.
type SoulAuctionClaimBid struct {
	TokenId   *big.Int
	Sender    common.Address
	Value     *big.Int
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionClaimBid is a free log retrieval operation binding the contract event 0x3efeb5ca9b10c05ba2df4bcac38445e4a4e8a996887f004e5036d254c14a9864.
//
// Solidity: event AuctionClaimBid(uint256 indexed tokenId, address sender, uint256 value, bytes32 auctionId)
func (_Soul *SoulFilterer) FilterAuctionClaimBid(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulAuctionClaimBidIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionClaimBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulAuctionClaimBidIterator{contract: _Soul.contract, event: "AuctionClaimBid", logs: logs, sub: sub}, nil
}

// WatchAuctionClaimBid is a free log subscription operation binding the contract event 0x3efeb5ca9b10c05ba2df4bcac38445e4a4e8a996887f004e5036d254c14a9864.
//
// Solidity: event AuctionClaimBid(uint256 indexed tokenId, address sender, uint256 value, bytes32 auctionId)
func (_Soul *SoulFilterer) WatchAuctionClaimBid(opts *bind.WatchOpts, sink chan<- *SoulAuctionClaimBid, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionClaimBid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionClaimBid)
				if err := _Soul.contract.UnpackLog(event, "AuctionClaimBid", log); err != nil {
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

// ParseAuctionClaimBid is a log parse operation binding the contract event 0x3efeb5ca9b10c05ba2df4bcac38445e4a4e8a996887f004e5036d254c14a9864.
//
// Solidity: event AuctionClaimBid(uint256 indexed tokenId, address sender, uint256 value, bytes32 auctionId)
func (_Soul *SoulFilterer) ParseAuctionClaimBid(log types.Log) (*SoulAuctionClaimBid, error) {
	event := new(SoulAuctionClaimBid)
	if err := _Soul.contract.UnpackLog(event, "AuctionClaimBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionClosedIterator is returned from FilterAuctionClosed and is used to iterate over the raw logs and unpacked data for AuctionClosed events raised by the Soul contract.
type SoulAuctionClosedIterator struct {
	Event *SoulAuctionClosed // Event containing the contract specifics and raw log

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
func (it *SoulAuctionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionClosed)
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
		it.Event = new(SoulAuctionClosed)
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
func (it *SoulAuctionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionClosed represents a AuctionClosed event raised by the Soul contract.
type SoulAuctionClosed struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionClosed is a free log retrieval operation binding the contract event 0xac4a907ec29adcc56774b757ecb1e1b4d597374fc9386107d05e2670259df7d3.
//
// Solidity: event AuctionClosed(uint256 indexed tokenId)
func (_Soul *SoulFilterer) FilterAuctionClosed(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulAuctionClosedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionClosed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulAuctionClosedIterator{contract: _Soul.contract, event: "AuctionClosed", logs: logs, sub: sub}, nil
}

// WatchAuctionClosed is a free log subscription operation binding the contract event 0xac4a907ec29adcc56774b757ecb1e1b4d597374fc9386107d05e2670259df7d3.
//
// Solidity: event AuctionClosed(uint256 indexed tokenId)
func (_Soul *SoulFilterer) WatchAuctionClosed(opts *bind.WatchOpts, sink chan<- *SoulAuctionClosed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionClosed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionClosed)
				if err := _Soul.contract.UnpackLog(event, "AuctionClosed", log); err != nil {
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
func (_Soul *SoulFilterer) ParseAuctionClosed(log types.Log) (*SoulAuctionClosed, error) {
	event := new(SoulAuctionClosed)
	if err := _Soul.contract.UnpackLog(event, "AuctionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for HandleAuctionCreated events raised by the Soul contract.
type SoulAuctionCreatedIterator struct {
	Event *SoulAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SoulAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionCreated)
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
		it.Event = new(SoulAuctionCreated)
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
func (it *SoulAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionCreated represents a HandleAuctionCreated event raised by the Soul contract.
type SoulAuctionCreated struct {
	TokenId   *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Sender    common.Address
	AuctionId [32]byte
	Auction   AuctionHouseAuction
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xb404e07dbfe9071da655c32567071b2ee4e88867784a1b39f01a8457787e504f.
//
// Solidity: event HandleAuctionCreated(uint256 indexed tokenId, uint256 startTime, uint256 endTime, address sender, bytes32 auctionId, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) FilterAuctionCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulAuctionCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "HandleAuctionCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulAuctionCreatedIterator{contract: _Soul.contract, event: "HandleAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xb404e07dbfe9071da655c32567071b2ee4e88867784a1b39f01a8457787e504f.
//
// Solidity: event HandleAuctionCreated(uint256 indexed tokenId, uint256 startTime, uint256 endTime, address sender, bytes32 auctionId, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SoulAuctionCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "HandleAuctionCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionCreated)
				if err := _Soul.contract.UnpackLog(event, "HandleAuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xb404e07dbfe9071da655c32567071b2ee4e88867784a1b39f01a8457787e504f.
//
// Solidity: event HandleAuctionCreated(uint256 indexed tokenId, uint256 startTime, uint256 endTime, address sender, bytes32 auctionId, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) ParseAuctionCreated(log types.Log) (*SoulAuctionCreated, error) {
	event := new(SoulAuctionCreated)
	if err := _Soul.contract.UnpackLog(event, "HandleAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionExtendedIterator is returned from FilterAuctionExtended and is used to iterate over the raw logs and unpacked data for AuctionExtended events raised by the Soul contract.
type SoulAuctionExtendedIterator struct {
	Event *SoulAuctionExtended // Event containing the contract specifics and raw log

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
func (it *SoulAuctionExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionExtended)
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
		it.Event = new(SoulAuctionExtended)
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
func (it *SoulAuctionExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionExtended represents a AuctionExtended event raised by the Soul contract.
type SoulAuctionExtended struct {
	TokenId *big.Int
	EndTime *big.Int
	Auction AuctionHouseAuction
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionExtended is a free log retrieval operation binding the contract event 0x4c1666cb87ed4fd7706cb19b1e3160612601aca7235c307fe64a11bc8db0f41f.
//
// Solidity: event AuctionExtended(uint256 indexed tokenId, uint256 endTime, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) FilterAuctionExtended(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulAuctionExtendedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionExtended", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulAuctionExtendedIterator{contract: _Soul.contract, event: "AuctionExtended", logs: logs, sub: sub}, nil
}

// WatchAuctionExtended is a free log subscription operation binding the contract event 0x4c1666cb87ed4fd7706cb19b1e3160612601aca7235c307fe64a11bc8db0f41f.
//
// Solidity: event AuctionExtended(uint256 indexed tokenId, uint256 endTime, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) WatchAuctionExtended(opts *bind.WatchOpts, sink chan<- *SoulAuctionExtended, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionExtended", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionExtended)
				if err := _Soul.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
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

// ParseAuctionExtended is a log parse operation binding the contract event 0x4c1666cb87ed4fd7706cb19b1e3160612601aca7235c307fe64a11bc8db0f41f.
//
// Solidity: event AuctionExtended(uint256 indexed tokenId, uint256 endTime, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) ParseAuctionExtended(log types.Log) (*SoulAuctionExtended, error) {
	event := new(SoulAuctionExtended)
	if err := _Soul.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionMinBidIncrementPercentageUpdatedIterator is returned from FilterAuctionMinBidIncrementPercentageUpdated and is used to iterate over the raw logs and unpacked data for AuctionMinBidIncrementPercentageUpdated events raised by the Soul contract.
type SoulAuctionMinBidIncrementPercentageUpdatedIterator struct {
	Event *SoulAuctionMinBidIncrementPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *SoulAuctionMinBidIncrementPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionMinBidIncrementPercentageUpdated)
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
		it.Event = new(SoulAuctionMinBidIncrementPercentageUpdated)
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
func (it *SoulAuctionMinBidIncrementPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionMinBidIncrementPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionMinBidIncrementPercentageUpdated represents a AuctionMinBidIncrementPercentageUpdated event raised by the Soul contract.
type SoulAuctionMinBidIncrementPercentageUpdated struct {
	MinBidIncrementPercentage *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterAuctionMinBidIncrementPercentageUpdated is a free log retrieval operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_Soul *SoulFilterer) FilterAuctionMinBidIncrementPercentageUpdated(opts *bind.FilterOpts) (*SoulAuctionMinBidIncrementPercentageUpdatedIterator, error) {

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &SoulAuctionMinBidIncrementPercentageUpdatedIterator{contract: _Soul.contract, event: "AuctionMinBidIncrementPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionMinBidIncrementPercentageUpdated is a free log subscription operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_Soul *SoulFilterer) WatchAuctionMinBidIncrementPercentageUpdated(opts *bind.WatchOpts, sink chan<- *SoulAuctionMinBidIncrementPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionMinBidIncrementPercentageUpdated)
				if err := _Soul.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
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
func (_Soul *SoulFilterer) ParseAuctionMinBidIncrementPercentageUpdated(log types.Log) (*SoulAuctionMinBidIncrementPercentageUpdated, error) {
	event := new(SoulAuctionMinBidIncrementPercentageUpdated)
	if err := _Soul.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionReservePriceUpdatedIterator is returned from FilterAuctionReservePriceUpdated and is used to iterate over the raw logs and unpacked data for AuctionReservePriceUpdated events raised by the Soul contract.
type SoulAuctionReservePriceUpdatedIterator struct {
	Event *SoulAuctionReservePriceUpdated // Event containing the contract specifics and raw log

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
func (it *SoulAuctionReservePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionReservePriceUpdated)
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
		it.Event = new(SoulAuctionReservePriceUpdated)
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
func (it *SoulAuctionReservePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionReservePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionReservePriceUpdated represents a AuctionReservePriceUpdated event raised by the Soul contract.
type SoulAuctionReservePriceUpdated struct {
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionReservePriceUpdated is a free log retrieval operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_Soul *SoulFilterer) FilterAuctionReservePriceUpdated(opts *bind.FilterOpts) (*SoulAuctionReservePriceUpdatedIterator, error) {

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return &SoulAuctionReservePriceUpdatedIterator{contract: _Soul.contract, event: "AuctionReservePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionReservePriceUpdated is a free log subscription operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_Soul *SoulFilterer) WatchAuctionReservePriceUpdated(opts *bind.WatchOpts, sink chan<- *SoulAuctionReservePriceUpdated) (event.Subscription, error) {

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionReservePriceUpdated)
				if err := _Soul.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
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
func (_Soul *SoulFilterer) ParseAuctionReservePriceUpdated(log types.Log) (*SoulAuctionReservePriceUpdated, error) {
	event := new(SoulAuctionReservePriceUpdated)
	if err := _Soul.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the Soul contract.
type SoulAuctionSettledIterator struct {
	Event *SoulAuctionSettled // Event containing the contract specifics and raw log

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
func (it *SoulAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionSettled)
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
		it.Event = new(SoulAuctionSettled)
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
func (it *SoulAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionSettled represents a AuctionSettled event raised by the Soul contract.
type SoulAuctionSettled struct {
	TokenId *big.Int
	Winner  common.Address
	Amount  *big.Int
	Auction AuctionHouseAuction
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0x727cff623966c8adc61ac68940e8c898c04ff7ced5d69d9735d51f0589daf532.
//
// Solidity: event AuctionSettled(uint256 indexed tokenId, address winner, uint256 amount, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) FilterAuctionSettled(opts *bind.FilterOpts, tokenId []*big.Int) (*SoulAuctionSettledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionSettled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulAuctionSettledIterator{contract: _Soul.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0x727cff623966c8adc61ac68940e8c898c04ff7ced5d69d9735d51f0589daf532.
//
// Solidity: event AuctionSettled(uint256 indexed tokenId, address winner, uint256 amount, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *SoulAuctionSettled, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionSettled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionSettled)
				if err := _Soul.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0x727cff623966c8adc61ac68940e8c898c04ff7ced5d69d9735d51f0589daf532.
//
// Solidity: event AuctionSettled(uint256 indexed tokenId, address winner, uint256 amount, (uint256,address,uint256,uint256,uint256,address,bool,uint256,uint256,uint256,bytes32) auction)
func (_Soul *SoulFilterer) ParseAuctionSettled(log types.Log) (*SoulAuctionSettled, error) {
	event := new(SoulAuctionSettled)
	if err := _Soul.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulAuctionTimeBufferUpdatedIterator is returned from FilterAuctionTimeBufferUpdated and is used to iterate over the raw logs and unpacked data for AuctionTimeBufferUpdated events raised by the Soul contract.
type SoulAuctionTimeBufferUpdatedIterator struct {
	Event *SoulAuctionTimeBufferUpdated // Event containing the contract specifics and raw log

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
func (it *SoulAuctionTimeBufferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulAuctionTimeBufferUpdated)
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
		it.Event = new(SoulAuctionTimeBufferUpdated)
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
func (it *SoulAuctionTimeBufferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulAuctionTimeBufferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulAuctionTimeBufferUpdated represents a AuctionTimeBufferUpdated event raised by the Soul contract.
type SoulAuctionTimeBufferUpdated struct {
	TimeBuffer *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionTimeBufferUpdated is a free log retrieval operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_Soul *SoulFilterer) FilterAuctionTimeBufferUpdated(opts *bind.FilterOpts) (*SoulAuctionTimeBufferUpdatedIterator, error) {

	logs, sub, err := _Soul.contract.FilterLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return &SoulAuctionTimeBufferUpdatedIterator{contract: _Soul.contract, event: "AuctionTimeBufferUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionTimeBufferUpdated is a free log subscription operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_Soul *SoulFilterer) WatchAuctionTimeBufferUpdated(opts *bind.WatchOpts, sink chan<- *SoulAuctionTimeBufferUpdated) (event.Subscription, error) {

	logs, sub, err := _Soul.contract.WatchLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulAuctionTimeBufferUpdated)
				if err := _Soul.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
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
func (_Soul *SoulFilterer) ParseAuctionTimeBufferUpdated(log types.Log) (*SoulAuctionTimeBufferUpdated, error) {
	event := new(SoulAuctionTimeBufferUpdated)
	if err := _Soul.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Soul contract.
type SoulClaimIterator struct {
	Event *SoulClaim // Event containing the contract specifics and raw log

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
func (it *SoulClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulClaim)
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
		it.Event = new(SoulClaim)
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
func (it *SoulClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulClaim represents a Claim event raised by the Soul contract.
type SoulClaim struct {
	Reserver    common.Address
	TokenId     *big.Int
	Owner       common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0xad03f837a9207e368d73ec028e1f54428184da8cfea258cc116da2225f3ac5eb.
//
// Solidity: event Claim(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_Soul *SoulFilterer) FilterClaim(opts *bind.FilterOpts, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (*SoulClaimIterator, error) {

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

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Claim", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SoulClaimIterator{contract: _Soul.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0xad03f837a9207e368d73ec028e1f54428184da8cfea258cc116da2225f3ac5eb.
//
// Solidity: event Claim(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_Soul *SoulFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *SoulClaim, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Claim", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulClaim)
				if err := _Soul.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_Soul *SoulFilterer) ParseClaim(log types.Log) (*SoulClaim, error) {
	event := new(SoulClaim)
	if err := _Soul.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Soul contract.
type SoulInitializedIterator struct {
	Event *SoulInitialized // Event containing the contract specifics and raw log

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
func (it *SoulInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulInitialized)
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
		it.Event = new(SoulInitialized)
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
func (it *SoulInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulInitialized represents a Initialized event raised by the Soul contract.
type SoulInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Soul *SoulFilterer) FilterInitialized(opts *bind.FilterOpts) (*SoulInitializedIterator, error) {

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SoulInitializedIterator{contract: _Soul.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Soul *SoulFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SoulInitialized) (event.Subscription, error) {

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulInitialized)
				if err := _Soul.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Soul *SoulFilterer) ParseInitialized(log types.Log) (*SoulInitialized, error) {
	event := new(SoulInitialized)
	if err := _Soul.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Soul contract.
type SoulOwnershipTransferredIterator struct {
	Event *SoulOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SoulOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulOwnershipTransferred)
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
		it.Event = new(SoulOwnershipTransferred)
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
func (it *SoulOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulOwnershipTransferred represents a OwnershipTransferred event raised by the Soul contract.
type SoulOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Soul *SoulFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SoulOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Soul.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SoulOwnershipTransferredIterator{contract: _Soul.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Soul *SoulFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SoulOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Soul.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulOwnershipTransferred)
				if err := _Soul.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Soul *SoulFilterer) ParseOwnershipTransferred(log types.Log) (*SoulOwnershipTransferred, error) {
	event := new(SoulOwnershipTransferred)
	if err := _Soul.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Soul contract.
type SoulPausedIterator struct {
	Event *SoulPaused // Event containing the contract specifics and raw log

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
func (it *SoulPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulPaused)
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
		it.Event = new(SoulPaused)
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
func (it *SoulPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulPaused represents a Paused event raised by the Soul contract.
type SoulPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Soul *SoulFilterer) FilterPaused(opts *bind.FilterOpts) (*SoulPausedIterator, error) {

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SoulPausedIterator{contract: _Soul.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Soul *SoulFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SoulPaused) (event.Subscription, error) {

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulPaused)
				if err := _Soul.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Soul *SoulFilterer) ParsePaused(log types.Log) (*SoulPaused, error) {
	event := new(SoulPaused)
	if err := _Soul.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulReserveIterator is returned from FilterReserve and is used to iterate over the raw logs and unpacked data for Reserve events raised by the Soul contract.
type SoulReserveIterator struct {
	Event *SoulReserve // Event containing the contract specifics and raw log

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
func (it *SoulReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulReserve)
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
		it.Event = new(SoulReserve)
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
func (it *SoulReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulReserve represents a Reserve event raised by the Soul contract.
type SoulReserve struct {
	Reserver    common.Address
	TokenId     *big.Int
	Owner       common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReserve is a free log retrieval operation binding the contract event 0x6f8d4fbb9604038077481649590dd5ddb6c97e7236b4b644b44084733d5f8560.
//
// Solidity: event Reserve(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_Soul *SoulFilterer) FilterReserve(opts *bind.FilterOpts, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (*SoulReserveIterator, error) {

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

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Reserve", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SoulReserveIterator{contract: _Soul.contract, event: "Reserve", logs: logs, sub: sub}, nil
}

// WatchReserve is a free log subscription operation binding the contract event 0x6f8d4fbb9604038077481649590dd5ddb6c97e7236b4b644b44084733d5f8560.
//
// Solidity: event Reserve(address indexed reserver, uint256 indexed tokenId, address indexed owner, uint256 blockNumber)
func (_Soul *SoulFilterer) WatchReserve(opts *bind.WatchOpts, sink chan<- *SoulReserve, reserver []common.Address, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Reserve", reserverRule, tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulReserve)
				if err := _Soul.contract.UnpackLog(event, "Reserve", log); err != nil {
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
func (_Soul *SoulFilterer) ParseReserve(log types.Log) (*SoulReserve, error) {
	event := new(SoulReserve)
	if err := _Soul.contract.UnpackLog(event, "Reserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Soul contract.
type SoulTransferIterator struct {
	Event *SoulTransfer // Event containing the contract specifics and raw log

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
func (it *SoulTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulTransfer)
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
		it.Event = new(SoulTransfer)
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
func (it *SoulTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulTransfer represents a Transfer event raised by the Soul contract.
type SoulTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Soul *SoulFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SoulTransferIterator, error) {

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

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SoulTransferIterator{contract: _Soul.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Soul *SoulFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SoulTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulTransfer)
				if err := _Soul.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Soul *SoulFilterer) ParseTransfer(log types.Log) (*SoulTransfer, error) {
	event := new(SoulTransfer)
	if err := _Soul.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Soul contract.
type SoulUnpausedIterator struct {
	Event *SoulUnpaused // Event containing the contract specifics and raw log

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
func (it *SoulUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulUnpaused)
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
		it.Event = new(SoulUnpaused)
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
func (it *SoulUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulUnpaused represents a Unpaused event raised by the Soul contract.
type SoulUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Soul *SoulFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SoulUnpausedIterator, error) {

	logs, sub, err := _Soul.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SoulUnpausedIterator{contract: _Soul.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Soul *SoulFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SoulUnpaused) (event.Subscription, error) {

	logs, sub, err := _Soul.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulUnpaused)
				if err := _Soul.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Soul *SoulFilterer) ParseUnpaused(log types.Log) (*SoulUnpaused, error) {
	event := new(SoulUnpaused)
	if err := _Soul.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
