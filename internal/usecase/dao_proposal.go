package usecase

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/contracts/generative_dao"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"math/big"
	"os"
	"strings"
)

// Crontab
func (u *Usecase) DaoProposalCreated(eventData interface{}, chainLog types.Log) error {
	input := eventData.(*generative_dao.GenerativeDaoProposalCreated)

	logKey := fmt.Sprintf("DaoProposalCreated - %s", input.ProposalId)

	targets := []string{}
	for _, target := range input.Targets {
		targets = append(targets, strings.ToLower(target.String()))
	}

	values := []int64{}
	for _, value := range input.Values {
		values = append(values, value.Int64())
	}

	blockNumber := chainLog.BlockNumber
	txHash := chainLog.TxHash.String()
	createdProposal := &entity.Proposal{
		ProposalID:      strings.ToLower(input.ProposalId.String()),
		Proposer:        strings.ToLower(input.Proposer.String()),
		StartBlock:      input.StartBlock.Int64(),
		EndBlock:        input.EndBlock.Int64(),
		Title:           input.Description,
		Targets:         targets,
		Values:          values,
		Signatures:      input.Signatures,
		CallData:        input.Calldatas,
		Amount:          "0",
		TokenType:       "NATIVE",
		ReceiverAddress: strings.ToLower(input.Proposer.String()),
		BlockNumber:     blockNumber,
		TxHash:          strings.ToLower(txHash),
	}

	stB, err := u.TCPublicNode.GetBlockByNumber(*big.NewInt(createdProposal.StartBlock))
	if err == nil {
		createdProposal.StartBlockTime = helpers.ParseUintToUnixTime(stB.Time())
	} else {
		logger.AtLog.Logger.Error(logKey, zap.Error(err))
	}

	eBB, err := u.TCPublicNode.GetBlockByNumber(*big.NewInt(createdProposal.EndBlock))
	if err == nil {
		createdProposal.EndBlockTime = helpers.ParseUintToUnixTime(eBB.Time())
	} else {
		logger.AtLog.Logger.Error(logKey, zap.Error(err))
	}

	//Create proposal
	created, err := u.Repo.InsertOne(createdProposal)
	if err != nil {
		logger.AtLog.Logger.Error(logKey, zap.Error(err))
		return err
	}
	logger.AtLog.Logger.Info(logKey, zap.Any("createdProposal", created))
	return nil
}

func (u *Usecase) DaoProposalCastVoted(eventData interface{}, chainLog types.Log) error {
	parsedCastVote := eventData.(*generative_dao.GenerativeDaoVoteCast)
	logKey := fmt.Sprintf("DaoProposalCastVoted - prosal: %s, voter: %s", parsedCastVote.ProposalId, parsedCastVote.Voter)

	blockNumber := chainLog.BlockNumber
	txHash := chainLog.TxHash.String()
	obj := &entity.ProposalVotes{
		ProposalID:      strings.ToLower(parsedCastVote.ProposalId.String()),
		ContractAddress: strings.ToLower(chainLog.Address.String()),
		Voter:           strings.ToLower(parsedCastVote.Voter.String()),
		Support:         int(parsedCastVote.Support),
		Weight:          parsedCastVote.Weight.String(),
		Reason:          parsedCastVote.Reason,
		BlockNumber:     blockNumber,
		TxHash:          txHash,
	}

	created, err := u.Repo.InsertOne(obj)
	if err != nil {
		logger.AtLog.Logger.Error(logKey, zap.Error(err))
		return err
	}

	logger.AtLog.Logger.Info(logKey, zap.Any("obj", obj), zap.Any("created", created))
	return nil
}

func (u *Usecase) ParseDao(chainLog types.Log, eventType entity.TokenActivityType, daoEventName string) (func(eventData interface{}, chainLog types.Log) error, interface{}, *types.Block, error) {
	logKey := fmt.Sprintf("ParseDao - txHash: %s - logIndex: %d", chainLog.TxHash, chainLog.Index)

	daoContract, err := generative_dao.NewGenerativeDao(chainLog.Address, u.TCPublicNode.GetClient())
	if err != nil {
		logger.AtLog.Logger.Error(logKey, zap.Error(err))
		return nil, nil, nil, err
	}

	bn := big.NewInt(int64(chainLog.BlockNumber))
	blockInfo, err := u.TCPublicNode.GetBlockByNumber(*bn)
	if err != nil {
		logger.AtLog.Logger.Error(logKey, zap.Error(err))
		return nil, nil, nil, err
	}

	switch daoEventName {
	case strings.ToLower(os.Getenv("DAO_PROPOSAL_CREATED")):
		pFunction := u.DaoProposalCreated

		parsedProposal, err := daoContract.ParseProposalCreated(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error(logKey, zap.Error(err))
			return nil, nil, nil, err
		}

		return pFunction, parsedProposal, blockInfo, nil

	case strings.ToLower(os.Getenv("DAO_PROPOSAL_CAST_VOTE")):
		pFunction := u.DaoProposalCastVoted

		parsedProposal, err := daoContract.ParseVoteCast(chainLog)
		if err != nil {
			logger.AtLog.Logger.Error(logKey, zap.Error(err))
			return nil, nil, nil, err
		}

		return pFunction, parsedProposal, blockInfo, nil
	}

	err = errors.New(fmt.Sprintf("Cannot detect event log - %d - txHash: %s, topics %s ", eventType, chainLog.TxHash, chainLog.Topics[0].String()))
	logger.AtLog.Logger.Error(logKey, zap.Error(err))
	return nil, nil, nil, err
}

//end crontab

// API
func (u Usecase) CreateDraftProposal(req *entity.ProposalDetail) (*entity.ProposalDetail, error) {
	_, err := u.Repo.InsertOne(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (u Usecase) MapOffToOnChainProposal(ID string, proposalID string) (*entity.ProposalDetail, error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	f := bson.D{{
		"_id", id,
	}}

	update := bson.M{"$set": bson.D{
		{"proposal_id", strings.ToLower(proposalID)},
	}}

	_, err = u.Repo.UpdateOne(entity.ProposalDetail{}.CollectionName(), f, update)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u Usecase) GetProposals(req entity.FilterProposal) ([]*entity.Proposal, error) {
	proposals, err := u.Repo.FilterProposals(req)
	if err != nil {
		return nil, err
	}
	return proposals, nil
}

func (u Usecase) GetProposal(proposalID string) (*entity.Proposal, error) {
	return u.Repo.GetProposal(proposalID)
}

func (u Usecase) GetProposalVotes(filter entity.FilterProposalVotes) ([]*entity.ProposalVotes, error) {
	return u.Repo.FilterProposalVotes(filter)
}

//End API
