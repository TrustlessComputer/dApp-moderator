package request

import "time"

type IdoRequest struct {
	ID                string    `json:"id"`
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
