package request

type BytescodeRequest struct {
	Bytescode string `json:"bytescode"`
}

type CreateSignatureRequest struct {
	WalletAddress string `json:"wallet_address"`
}
