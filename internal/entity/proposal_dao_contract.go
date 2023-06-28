package entity

import (
	"dapp-moderator/utils"
)

type ProposalDaoContract struct {
	BaseEntity      `bson:",inline"`
	ContractAddress string `bson:"contract_address" json:"contract_address"`
}

func (u ProposalDaoContract) CollectionName() string {
	return utils.COLLECTION_DAO_PROPOSAL_CONTRACTS
}
