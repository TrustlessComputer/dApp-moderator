package request

type UpdateTokenReq struct {
	Symbol      string `json:"symbol" `
	Name        string `json:"name" `
	Thumbnail   string `json:"thumbnail" `
	Description string `json:"description" `
	Social      struct {
		Website   string `json:"website" `
		DisCord   string `json:"discord" `
		Twitter   string `json:"twitter" `
		Telegram  string `json:"telegram" `
		Medium    string `json:"medium" `
		Instagram string `json:"instagram" `
	} `json:"social" `
}
