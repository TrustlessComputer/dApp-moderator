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

type AuctionListBidResponse struct {
	Items []*AuctionListBidResponseItem `json:"items"`
	Total int64                         `json:"total"`
}

type AuctionListBidResponseItem struct {
	Amount string    `json:"amount"`
	Sender string    `json:"sender"`
	Avatar string    `json:"avatar"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
}
