package moralis

import "time"

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

type HashResp struct {
	Hash                     string        `json:"hash"`
	Nonce                    string        `json:"nonce"`
	TransactionIndex         string        `json:"transaction_index"`
	FromAddress              string        `json:"from_address"`
	ToAddress                string        `json:"to_address"`
	Value                    string        `json:"value"`
	Gas                      string        `json:"gas"`
	GasPrice                 string        `json:"gas_price"`
	Input                    string        `json:"input"`
	ReceiptCumulativeGasUsed string        `json:"receipt_cumulative_gas_used"`
	ReceiptGasUsed           string        `json:"receipt_gas_used"`
	ReceiptContractAddress   interface{}   `json:"receipt_contract_address"`
	ReceiptRoot              interface{}   `json:"receipt_root"`
	ReceiptStatus            string        `json:"receipt_status"`
	BlockTimestamp           time.Time     `json:"block_timestamp"`
	BlockNumber              string        `json:"block_number"`
	BlockHash                string        `json:"block_hash"`
	TransferIndex            interface{}   `json:"transfer_index"`
	Logs                     []interface{} `json:"logs"`
}
