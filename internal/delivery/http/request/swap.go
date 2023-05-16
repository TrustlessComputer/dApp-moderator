package request

import "time"

type IdoRequest struct {
	ID                string    `json:"id"`
	Signature         string    `json:"signature"`
	TokenAddress      string    `json:"token_address"`
	UserWalletAddress string    `json:"user_wallet_address"`
	StartAt           time.Time `json:"start_at"`
	Price             string    `json:"price"`
	Link              string    `json:"link"`
	Website           string    `json:"website"`
	Twitter           string    `json:"twitter"`
	WhitePaper        string    `json:"white_papper"`
	Discord           string    `json:"discord"`
}

type SwapWalletAddressRequest struct {
	WalletAddress           string `json:"address"`
	WalletAddressPrivateKey string `json:"prk_key"`
}

type SwapBotConfigRequest struct {
	Pair     string  `json:"pair"`
	Address  string  `json:"address"`
	MinValue float64 `json:"min_value"`
	MaxValue float64 `json:"max_value"`
}
