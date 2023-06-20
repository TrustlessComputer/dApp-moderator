package entity

import "dapp-moderator/utils"

type FilterProposalVotes struct {
	BaseFilters
	ProposalID      *string `bson:"proposal_id" json:"proposal_id"`
	ContractAddress *string `bson:"contract_address" json:"contract_address"`
	Voter           *string `bson:"voter" json:"voter"`
}

type ProposalVotes struct {
	BaseEntity      `bson:",inline"`
	ProposalID      string `bson:"proposal_id" json:"proposal_id"`
	ContractAddress string `bson:"contract_address" json:"contract_address"`
	Voter           string `bson:"voter" json:"voter"`
	Support         int    `bson:"support" json:"support"`
	Weight          string `bson:"weight" json:"weight"`
	Reason          string `bson:"reason" json:"reason"`
	TxHash          string `bson:"tx_hash" json:"tx_hash"`
	BlockNumber     uint64 `bson:"block_number" json:"block_number"`
}

func (u ProposalVotes) CollectionName() string {
	return utils.COLLECTION_DAO_PROPOSAL_VOTES
}
