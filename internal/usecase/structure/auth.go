package structure

import "time"


type GenerateMessage struct {
	Address    string
	WalletType string
}

type VerifyMessage struct {
	ETHSignature     string
	Signature        string
	Address          string
	AddressBTC       *string // taproot address
	AddressBTCSegwit *string
	MessagePrefix    *string
	AddressPayment   string
}

type VerifyResponse struct {
	IsVerified   bool
	Token        string
	RefreshToken string
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
