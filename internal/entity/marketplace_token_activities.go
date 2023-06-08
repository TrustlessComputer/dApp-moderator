package entity

import (
	"dapp-moderator/utils"
	"time"
)

type TokenActivityType int

const (
	TokenMint          TokenActivityType = 0
	TokenListing       TokenActivityType = 1
	TokenCancelListing TokenActivityType = 2
	TokenMatched       TokenActivityType = 3
	TokenTransfer      TokenActivityType = 4
	TokenMakeOffer     TokenActivityType = 5
	TokenCancelOffer   TokenActivityType = 6
	TokenAcceptOffer   TokenActivityType = 7
	TokenPurchase      TokenActivityType = 8
	BNSResolverUpdated TokenActivityType = 9
	BNSResolverCreated TokenActivityType = 10
	BNSPfpUpdated      TokenActivityType = 11
)

var TokenActivityName = map[TokenActivityType]string{
	TokenMint:          "mint",
	TokenListing:       "listing",
	TokenCancelListing: "cancel listing",
	TokenMatched:       "market place: token match",
	TokenTransfer:      "transfer",
	TokenMakeOffer:     "make offer",
	TokenCancelOffer:   "cancel offer",
	TokenAcceptOffer:   "accept offer",
	TokenPurchase:      "purchase",
	BNSResolverUpdated: "BNS Resolver updated",
	BNSResolverCreated: "BNS registered",
	BNSPfpUpdated:      "BNS pfp updated",
}

type MarketplaceTokenActivity struct {
	BaseEntity         `bson:",inline" json:"base_entity"`
	Type               TokenActivityType `bson:"type" json:"type"`
	Title              string            `bson:"title" json:"title"`
	UserAAddress       string            `bson:"user_a_address" json:"user_a_address"`
	UserBAddress       string            `bson:"user_b_address" json:"user_b_address"`
	Amount             int64             `bson:"amount" json:"-"`
	AmountStr          string            `bson:"-" json:"amount"`
	Erc20Address       string            `bson:"erc_20_address" json:"erc_20_address"`
	Time               *time.Time        `bson:"time" json:"time"`
	InscriptionID      string            `bson:"inscription_id" json:"inscription_id"`
	CollectionContract string            `bson:"collection_contract" json:"collection_contract"`
	OfferingID         string            `bson:"offering_id" json:"offering_id"`
	BlockNumber        uint64            `bson:"block_number" json:"block_number"`
	TxHash             string            `json:"tx_hash" bson:"tx_hash"`
	LogIndex           uint              `json:"log_index" bson:"log_index"`
}

func (u MarketplaceTokenActivity) CollectionName() string {
	return utils.COLLECTION_MARKETPLACE_TOKEN_ACTIVITY
}
