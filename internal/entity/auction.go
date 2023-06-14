package entity

type AuctionStatus int

const (
	AuctionStatusInProgress AuctionStatus = 1
)

type Auction struct {
	BaseEntity        `bson:",inline"`
	CollectionAddress string `json:"collection_address" bson:"collection_address"`
	TokenID           string `json:"token_id" bson:"token_id"`
	StartTimeBlock    uint64 `json:"start_time_block" bson:"start_time_block"`
	EndTimeBlock      uint64 `json:"end_time_block" bson:"end_time_block"`

	Status AuctionStatus `json:"status" bson:"status"`
}

func (Auction) CollectionName() string {
	return "auction"
}
