package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func (r *Repository) ListenedDaoContract() ([]entity.ProposalDaoContract, error) {
	result := []entity.ProposalDaoContract{}

	f := bson.A{}

	cursor, err := r.DB.Collection(entity.ProposalDaoContract{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) AllOpenProposal() ([]entity.Proposal, error) {
	result := []entity.Proposal{}

	f := bson.A{
		bson.D{
			{
				"$match", bson.D{{"state", bson.E{"$in", bson.A{entity.PStateActivate, entity.PStatePending}}}},
			},
			{"$project",
				bson.D{
					{"contract", 1},
					{"deployed_at_block", 1},
					{"contract_type", 1},
				},
			},
		},
	}

	cursor, err := r.DB.Collection(entity.Proposal{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) FilterProposals(filter entity.FilterProposal) ([]*entity.Proposal, error) {
	result := []*entity.Proposal{}

	match := bson.D{}
	if filter.Proposer != nil && *filter.Proposer != "" {
		match = append(match, bson.E{"proposer", *filter.Proposer})
	}

	if filter.ProposalID != nil && *filter.ProposalID != "" {
		match = append(match, bson.E{"proposal_id", *filter.ProposalID})
	}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"contract_address", *filter.ContractAddress})
	}

	if len(filter.State) > 0 {
		match = append(match, bson.E{"state", bson.M{"$in": filter.State}})
	}

	f := bson.A{}

	if len(match) > 0 {
		f = append(f, bson.D{{
			"$match", match,
		}})
	}

	f = append(f, bson.D{
		{"$sort", bson.D{
			{filter.SortBy, filter.Sort},
		}},
	})
	f = append(f, bson.D{
		{"$skip", filter.Offset},
	})
	f = append(f, bson.D{
		{"$limit", filter.Limit},
	})

	cursor, err := r.DB.Collection(entity.Proposal{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetProposal(proposalID string) (*entity.Proposal, error) {
	result := &entity.Proposal{}

	f := bson.D{
		{"proposal_id", strings.ToLower(proposalID)},
	}

	cursor := r.DB.Collection(entity.Proposal{}.CollectionName()).FindOne(context.TODO(), f)
	if err := cursor.Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) FilterProposalVotes(filter entity.FilterProposalVotes) ([]*entity.ProposalVotes, error) {
	result := []*entity.ProposalVotes{}

	match := bson.D{}
	if filter.Voter != nil && *filter.Voter != "" {
		match = append(match, bson.E{"voter", *filter.Voter})
	}

	if filter.ProposalID != nil && *filter.ProposalID != "" {
		match = append(match, bson.E{"proposal_id", *filter.ProposalID})
	}

	if filter.ContractAddress != nil && *filter.ContractAddress != "" {
		match = append(match, bson.E{"contract_address", *filter.ContractAddress})
	}

	f := bson.A{}

	if len(match) > 0 {
		f = append(f, bson.D{{
			"$match", match,
		}})
	}

	f = append(f, bson.D{
		{"$sort", bson.D{
			{filter.SortBy, filter.Sort},
		}},
	})
	f = append(f, bson.D{
		{"$skip", filter.Offset},
	})
	f = append(f, bson.D{
		{"$limit", filter.Limit},
	})

	cursor, err := r.DB.Collection(entity.ProposalVotes{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}
