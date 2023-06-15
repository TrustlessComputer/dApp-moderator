package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuctionStatus int

const (
	AuctionStatusInProgress AuctionStatus = 1
)

type Auction struct {
	BaseEntity        `bson:",inline"`
	CollectionAddress string `json:"collection_address" bson:"collection_address"`
	TokenID           string `json:"token_id" bson:"token_id"`
	TokenIDInt        uint64 `json:"token_id_int" bson:"token_id_int"`
	AuctionID         uint64 `json:"auction_id" bson:"auction_id"`
	StartTimeBlock    uint64 `json:"start_time_block" bson:"start_time_block"`
	EndTimeBlock      uint64 `json:"end_time_block" bson:"end_time_block"`

	Status AuctionStatus `json:"status" bson:"status"`
}

func (Auction) CollectionName() string {
	return "auction"
}

type AuctionBid struct {
	BaseEntity        `bson:",inline"`
	DBAuctionID       primitive.ObjectID `json:"db_auction_id" bson:"db_auction_id"`
	ChainAuctionID    uint64             `json:"chain_auction_id" bson:"chain_auction_id"`
	TokenID           string             `json:"token_id" bson:"token_id"`
	CollectionAddress string             `json:"collection_address" bson:"collection_address"`
	Amount            string             `json:"amount" bson:"amount"`
	Sender            string             `json:"sender" bson:"sender"`
}

func (AuctionBid) CollectionName() string {
	return "auction_bid"
}
