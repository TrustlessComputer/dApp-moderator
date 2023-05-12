package entity

import (
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transactions struct {
	BaseEntity         `bson:",inline"`
	TransactionType    string               `json:"transaction_type" bson:"transaction_type"`
	TransactionStatus  string               `json:"transaction_status" bson:"transaction_status"`
	Method             string               `json:"method" bson:"method"` // Owner of a contract (contract address)
	Hash               string               `json:"hash" bson:"hash"`
	FromAddress        string               `json:"from_address" bson:"from_address"`
	ToAddress          string               `json:"to_address" bson:"to_address"`
	ToName             string               `json:"to_name" bson:"to_name"`
	Amount             primitive.Decimal128 `json:"amount" bson:"amount"`
	Fee                primitive.Decimal128 `json:"fee" bson:"fee"`
	TransferFrom1      string               `json:"transfer_from1" bson:"transfer_from1"`
	TransferTo1        string               `json:"transfer_to1" bson:"transfer_to1"`
	TransferToken1     string               `json:"transfer_token1" bson:"transfer_token1"`
	TransferTokenName1 string               `json:"transfer_token_name1" bson:"transfer_token_name1"`
	TransferAmount1    primitive.Decimal128 `json:"transfer_amount1" bson:"transfer_amount1"`
	TransferTokenID1   string               `json:"transfer_token_id1" bson:"transfer_token_id1"`
	TransferFrom2      string               `json:"transfer_from2" bson:"transfer_from2"`
	TransferTo2        string               `json:"transfer_to2" bson:"transfer_to2"`
	TransferToken2     string               `json:"transfer_token2" bson:"transfer_token2"`
	TransferTokenName2 string               `json:"transfer_token_name2" bson:"transfer_token_name2"`
	TransferAmount2    primitive.Decimal128 `json:"transfer_amount2" bson:"transfer_amount2"`
	TransferTokenID2   string               `json:"transfer_token_id2" bson:"transfer_token_id2"`
	Block              uint                 `json:"block" bson:"block"`
	BlockTime          time.Time            `json:"block_time" bson:"block_time"`
}

func (t *Transactions) CollectionName() string {
	return utils.TRANSACTIONS
}
