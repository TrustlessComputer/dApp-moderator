package request

type UpsertRedisRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListtokenIDsReq struct {
	InscriptionID []string `json:"inscriptionIDs"`

	SellOrdAddress string            `json:"seller_ord_address"`
	SellerAddress  string            `json:"seller_address"`
	Price          string            `json:"price"`
	PayType        map[string]string `bson:"payType"`
}
