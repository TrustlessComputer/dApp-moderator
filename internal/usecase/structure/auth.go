package structure

import "time"

type GenerateMessage struct {
	Address    string
	WalletType string
}

type VerifyMessage struct {
	ETHSignature     string  `json:"-"`
	Signature        string  `json:"signature"`
	Address          string  `json:"address"`
	AddressBTC       *string `json:"-"` // taproot address
	AddressBTCSegwit *string `json:"-"`
	MessagePrefix    *string `json:"-"`
	AddressPayment   string  `json:"-"`
}

type VerifyResponse struct {
	IsVerified   bool   `json:"is_verified"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ProfileResponse struct {
	ID               string     `json:"id"`
	WalletAddress    string     `json:"wallet_address"`
	DisplayName      string     `json:"display_name"`
	Bio              string     `json:"bio"`
	Avatar           string     `json:"avatar"`
	CreatedAt        *time.Time `json:"created_at"`
	WalletAddressBTC string     `json:"wallet_address_btc"`
}

type CreateHistoryMessage struct {
	TxHash         string     `json:"tx_hash"`
	DappTypeTxHash string     `json:"dapp_type"`
	FromAddress    string     `json:"from_address"`
	ToAddress      string     `json:"to_address"`
	Value          string     `json:"value"`
	Decimal        int        `json:"decimal"`
	Currency       string     `json:"currency"`
	Time           *time.Time `json:"time"`
	WalletAddress  string     `json:"-"`
	Status         string     `json:"-"`
	BTCTxHash      string     `json:"btc_tx_hash"`
}
