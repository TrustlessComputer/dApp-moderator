package response

import "dapp-moderator/internal/entity"

type AuctionDetailResponse struct {
	Available     bool                 `json:"available"`
	AuctionStatus entity.AuctionStatus `json:"auction_status"`
	HighestBid    string               `json:"highest_bid"`
	EndTime       string               `json:"end_time"`
}

type AuctionBidResponseItem struct {
}
