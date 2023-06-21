package response

import (
	"dapp-moderator/internal/entity"
	"time"
)

type AuctionDetailResponse struct {
	Available     bool                 `json:"available"`
	AuctionStatus entity.AuctionStatus `json:"auction_status"`
	HighestBid    string               `json:"highest_bid"`
	EndTime       string               `json:"end_time"`

	DBAuctionID    string `json:"db_auction_id"`
	ChainAuctionID string `json:"chain_auction_id"`
}

type AuctionBidResponseItem struct {
	Items []*AuctionBidItemResponse `json:"items"`
	Total int64                     `json:"total"`
}

type AuctionBidItemResponse struct {
	Amount string    `json:"amount"`
	Sender string    `json:"sender"`
	Time   time.Time `json:"time"`
}
