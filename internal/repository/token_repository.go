package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindToken(ctx context.Context, filter entity.TokenFilter) (*entity.Token, error) {

	query := bson.D{}
	if filter.Address != "" {
		query = append(query, bson.E{Key: "address", Value: filter.Address})
	}
	var token entity.Token
	err := r.DB.Collection(utils.COLLECTION_TOKENS).FindOne(ctx, query).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
