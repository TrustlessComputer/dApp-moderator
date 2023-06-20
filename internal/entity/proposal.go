package entity

import (
	"dapp-moderator/utils"
	"time"
)

type DAOActivityType int

const (
	DAOProposalCreated  DAOActivityType = iota
	DAOProposalCastVote DAOActivityType = 1
)

const (
	PStatePending = iota
	PStateActivate
	PStateCanceled
	PStateDefeated
	PStateSuccesseded
	PStateQueued
	PStateExpired
	PStateExecuted
)

type FilterProposal struct {
	BaseFilters
	Proposer        *string
	ContractAddress *string
	ProposalID      *string
	State           []int
}

type CreateDraftProposal struct {
	Proposal       Proposal
	ProposalDetail ProposalDetail
}

type Proposal struct {
	BaseEntity       `bson:",inline"`
	ContractAddress  string     `bson:"contract_address" json:"contract_address"`
	ProposalID       string     `bson:"proposal_id" json:"proposal_id"`
	Proposer         string     `bson:"proposer" json:"proposer"`
	ReceiverAddress  string     `bson:"receiver_address" json:"receiver_address"`
	Targets          []string   `bson:"targets" json:"targets"`
	Values           []int64    `bson:"values" json:"values"`
	Signatures       []string   `bson:"signatures" json:"signatures"`
	CallData         [][]byte   `bson:"call_data" json:"call_data"`
	StartBlock       int64      `bson:"start_block" json:"start_block"`
	StartBlockTime   *time.Time `bson:"start_block_time" json:"start_block_time"`
	CurrentBlock     int64      `bson:"current_block" json:"current_block"`
	CurrentBlockTime *time.Time `bson:"current_block_time" json:"current_block_time"`
	EndBlock         int64      `bson:"end_block" json:"end_block"`
	EndBlockTime     *time.Time `bson:"end_block_time" json:"end_block_time"`
	Title            string     `bson:"title" json:"title"`
	Description      string     `bson:"description" json:"description"`
	Amount           string     `bson:"amount" json:"amount"`
	TokenType        string     `bson:"token_type" json:"token_type"`
	State            uint8      `bson:"state" json:"state"`
	TxHash           string     `bson:"tx_hash" json:"tx_hash"`
	BlockNumber      uint64     `bson:"block_number" json:"block_number"`
}

func (u Proposal) CollectionName() string {
	return utils.COLLECTION_DAO_PROPOSAL
}
