package response

import "dapp-moderator/internal/entity"

type AuctionDetailResponse struct {
	Available     bool                 `json:"available"`
	AuctionStatus entity.AuctionStatus `json:"auction_status"`
	HighestBid    string               `json:"highest_bid"`
	EndTime       string               `json:"end_time"`

	DBAuctionID    string `json:"db_auction_id"`
	ChainAuctionID string `json:"chain_auction_id"`
}

type AuctionBidResponseItem struct {
}
