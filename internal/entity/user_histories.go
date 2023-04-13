package entity

import (
	"dapp-moderator/utils"
	"time"
)

const (
	HISTORY_PENDING   string = "pending"
	HISTORY_CONFIRMED string = "confirmed"
)

type UserHistories struct {
	BaseEntity     `bson:",inline"`
	WalletAddress  string `bson:"wallet_address" json:"wallet_address"` // eth wallet define user in platform by connect wallet and sign
	TxHash         string `bson:"tx_hash" json:"tx_hash"`
	DappTypeTxHash string `bson:"tx_hash_type" json:"tx_hash_type"`
	Status         string `bson:"status" json:"status"`
	FromAddress    string `bson:"from_address" json:"from_address"`
	ToAddress      string `bson:"to_address" json:"to_address"`
	Value          string `bson:"value" json:"value"`
	Decimal        int    `bson:"decimal" json:"decimal"`
	Currency       string `bson:"currency" json:"currency"`

	Time *time.Time `bson:"time" json:"time"`
}

func (t *UserHistories) CollectionName() string {
	return utils.COLLECTION_USER_HISTORIES
}
