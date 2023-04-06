package structure

type InscriptionByOutput struct {
	Outpoint     string   `json:"outpoint"`
	Chain        string   `json:"chain"`
	Inscriptions []string `json:"inscriptions"`
}

type WalletAddressBalanceResp struct {
	Version  int `json:"version"`
	Height   int64  `json:"height"`
	Script   string `json:"script"`
	Address  string  `json:"address"`
	Coinbase bool   `json:"coinbase"`
	Hash     string `json:"hash"`
	Index    int    `json:"index"`
	Value    uint64 `json:"value"`
	IsOrdinal bool  `json:"isOrdinal"`
}