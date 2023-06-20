package entity

import (
	"dapp-moderator/utils"
)

type ProposalDetail struct {
	BaseEntity      `bson:",inline"`
	ProposalID      string `bson:"proposal_id" json:"proposal_id"` //proposalID from chain
	ContractAddress string `bson:"contract_address" json:"contract_address"`
	ReceiverAddress string `bson:"receiver_address" json:"receiver_address"`
	Title           string `bson:"title" json:"title"`
	Description     string `bson:"description" json:"description"`
	Amount          string `bson:"amount" json:"amount"`
	TokenType       string `bson:"token_type" json:"token_type"`
	IsDraft         bool   `bson:"is_draft" json:"is_draft"`
}

func (u ProposalDetail) CollectionName() string {
	return utils.COLLECTION_DAO_PROPOSAL_DETAIL
}
