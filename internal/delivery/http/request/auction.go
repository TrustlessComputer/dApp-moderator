package request

type FilterAuctionBid struct {
	PaginationReq
	DBAuctionID *string
	Sender      *string
}
