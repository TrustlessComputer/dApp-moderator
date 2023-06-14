package structure

import (
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/internal/entity"
)

type BaseFilters struct {
	Limit  *string
	Page   *string
	SortBy *string
	Cursor *string
	Sort   *int
}

type UpdateCollection struct {
	Cover       *string `json:"cover"`
	Thumbnail   *string `json:"thumbnail"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Social      Social  `json:"social"`
}

type Social struct {
	Website   *string `json:"website"`
	DisCord   *string `json:"discord"`
	Twitter   *string `json:"twitter"`
	Telegram  *string `json:"telegram"`
	Medium    *string `json:"medium"`
	Instagram *string `json:"instagram"`
}

type UpdateUploadedFileTxHash struct {
	TxHash        string `json:"-"`
	FileID        string `json:"-"`
	WalletAddress string `json:"wallet_address"`
	TokenID       string `json:"token_id"`
}

type TxHashInfo struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      interface{} `json:"id"`
	Result  Result      `json:"result"`
}

type Result struct {
	Nonce            string `json:"nonce"`
	GasPrice         string `json:"gasPrice"`
	Gas              string `json:"gas"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Input            string `json:"input"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
	Hash             string `json:"hash"`
	From             string `json:"from"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
}

type NftsResp struct {
	nft_explorer.NftsResp
	BnsData  []*entity.FilteredBNS `json:"bns_data,omitempty"`
	FileSize int                   `json:"file_size"` //bytes
}

type CompressedFile struct {
	OriginalSize   int `json:"original_size"`
	CompressedSize int `json:"compressed_size"`
}

type BnsRespChan struct {
	Bns *entity.Bns
	Nft entity.Nfts
	Err error
}

type MarketplaceCollectionAttributeResp struct {
	TraitName       string                                `json:"trait_name"`
	TraitValuesStat []MarketplaceCollectionAttributeValue `json:"trait_values_stat"`
}

type MarketplaceCollectionAttributeValue struct {
	Value  string  `json:"value"`
	Rarity float64 `json:"rarity"`
}
type GMDepositItemResponse struct {
	Ens             string  `json:"ens"`
	Avatar          string  `json:"avatar"`
	From            string  `json:"from"`
	To              string  `json:"to"`
	Value           string  `json:"value"`
	UsdtValue       float64 `json:"usdt_value"`
	UsdtValueExtra  int     `json:"usdt_value_extra"`
	ExtraPercent    int     `json:"extra_percent"`
	Percent         float64 `json:"percent"`
	GmReceive       float64 `json:"gm_receive"`
	GmReceiveString string  `json:"gm_receive_string"`
	Currency        string  `json:"currency"`
}

type GMDepositResponse struct {
	Error  interface{} `json:"error"`
	Status bool        `json:"status"`
	Data   struct {
		Value          string                  `json:"value"`
		Currency       string                  `json:"currency"`
		CurrencyRate   int                     `json:"currencyRate"`
		UsdtValue      float64                 `json:"usdtValue"`
		UsdtExtra      float64                 `json:"usdtExtra"`
		TotalGMReceive float64                 `json:"totalGMReceive"`
		Items          []GMDepositItemResponse `json:"items"`
		MapItems       struct {
		} `json:"MapItems"`
		MapTokensDeposit struct {
		} `json:"map_tokens_deposit"`
	} `json:"data"`
}

type CreateSignatureResp struct {
	Signature   string `json:"signature"`
	MessageHash string `json:"message_hash"`
	Deadline    string `json:"-"`
	GM          string `json:"gm"`
}
