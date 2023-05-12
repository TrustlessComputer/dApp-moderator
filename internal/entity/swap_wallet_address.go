package entity

import (
	"dapp-moderator/utils"
)

type SwapWalletAddress struct {
	BaseEntity `bson:",inline"`
	Address    string `json:"address"  bson:"address,omitempty"`
	Prk        string `json:"prk"  bson:"prk,omitempty"`
}

func (t *SwapWalletAddress) CollectionName() string {
	return utils.COLLECTION_SWAP_WALLET_ADDRESS
}

type SwapWalletAddressFilter struct {
	BaseFilters
	Address string
}
