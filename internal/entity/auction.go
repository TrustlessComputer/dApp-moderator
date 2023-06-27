package entity

import (
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuctionStatus int

const (
	AuctionStatusInProgress AuctionStatus = 1
	AuctionStatusEnded      AuctionStatus = 2
	AuctionStatusSettled    AuctionStatus = 3
)

func (v AuctionStatus) Ordinal() int {
	return int(v)
}

type Auction struct {
	BaseEntity        `bson:",inline"`
	CollectionAddress string `json:"collection_address" bson:"collection_address"`
	TokenID           string `json:"token_id" bson:"token_id"`
	TokenIDInt        uint64 `json:"token_id_int" bson:"token_id_int"`
	AuctionID         string `json:"auction_id" bson:"auction_id"`
	StartTimeBlock    string `json:"start_time_block" bson:"start_time_block"`
	EndTimeBlock      string `json:"end_time_block" bson:"end_time_block"`
	BlockNumber       string `json:"block_number" bson:"block_number"`

	Status      AuctionStatus `json:"status" bson:"status"`
	TotalAmount string        `json:"total_amount" bson:"total_amount"`

	Winner *string `json:"winner,omitempty" bson:"winner,omitempty"`
}

func (Auction) CollectionName() string {
	return utils.COLLECTION_AUCTION
}

type AuctionBid struct { // auctionBid History
	BaseEntity        `bson:",inline"`
	DBAuctionID       primitive.ObjectID `json:"db_auction_id" bson:"db_auction_id"`
	ChainAuctionID    string             `json:"chain_auction_id" bson:"chain_auction_id"`
	TokenID           string             `json:"token_id" bson:"token_id"`
	CollectionAddress string             `json:"collection_address" bson:"collection_address"`
	Amount            string             `json:"amount" bson:"amount"`
	Sender            string             `json:"sender" bson:"sender"`
	BlockNumber       string             `json:"block_number" bson:"block_number"`
}

func (AuctionBid) CollectionName() string {
	return utils.COLLECTION_AUCTION_BID
}

type AuctionBidSummary struct { // summary total amount per auction per user
	BaseEntity        `bson:",inline"`
	DBAuctionID       primitive.ObjectID `json:"db_auction_id" bson:"db_auction_id"`
	ChainAuctionID    string             `json:"chain_auction_id" bson:"chain_auction_id"`
	TokenID           string             `json:"token_id" bson:"token_id"`
	CollectionAddress string             `json:"collection_address" bson:"collection_address"`
	TotalAmount       string             `json:"total_amount" bson:"total_amount"`
	Sender            string             `json:"sender" bson:"sender"`
}

func (AuctionBidSummary) CollectionName() string {
	return utils.COLLECTION_AUCTION_BID_SUMMARY
}

type AuctionClaim struct {
	BaseEntity        `bson:",inline"`
	DBAuctionID       primitive.ObjectID `json:"db_auction_id" bson:"db_auction_id"`
	ChainAuctionID    string             `json:"chain_auction_id" bson:"chain_auction_id"`
	TokenID           string             `json:"token_id" bson:"token_id"`
	CollectionAddress string             `json:"collection_address" bson:"collection_address"`
	Claimer           string             `json:"claimer" bson:"claimer"`
	Amount            string             `json:"amount" bson:"amount"`
}

func (AuctionClaim) CollectionName() string {
	return utils.COLLECTION_AUCTION_CLAIM
}
