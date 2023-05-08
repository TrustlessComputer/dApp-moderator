package moralis

type RequestData struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type Erc20BalanceResp struct {
	TokenAddress string  `json:"token_address"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Logo         *string `json:"logo"`
	Thumbnail    *string `json:"thumbnail"`
	Decimals     int     `json:"decimals"`
	Balance      string  `json:"balance"`
	PossibleSpam bool    `json:"possible_spam"`
}
