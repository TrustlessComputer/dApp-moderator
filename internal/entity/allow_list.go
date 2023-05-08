package entity

import "dapp-moderator/utils"

type AllowList struct {
	BaseEntity       `bson:",inline"`
	Name             string `bson:"name" json:"name"`                           //token address name, exp: meme, pepe ...
	Address          string `bson:"address" json:"address"`                     //token address (ERC720)
	Threshold        string `bson:"threshold" json:"threshold"`                 // if user wallet address has tokens >= Threshold, he will be got reward, default: 10
	ThresholdDecimal int    `bson:"threshold_decimal" json:"threshold_decimal"` //default 18
	Reward           string `bson:"reward" json:"reward"`                       //To TC token // default: 0.1
	RewardDecimal    int    `bson:"reward_decimal" json:"reward_decimal"`       //default 18

}

type InsertedAllowList struct {
	BaseEntity    `bson:",inline"`
	WalletAddress string `bson:"wallet_address"` //each user is only inserted into this DB once.
	TokenAddress  string `bson:"token_address"`
	Decimals      int    `bson:"decimals"`
	Balance       string `bson:"balance"`
}

func (u AllowList) CollectionName() string {
	return utils.COLLECTION_ALLOW_LIST
}

func (u InsertedAllowList) CollectionName() string {
	return utils.COLLECTION_INSERTED_ALLOW_LIST
}
