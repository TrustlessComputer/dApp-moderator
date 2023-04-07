package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func (r *Repository) FindToken(ctx context.Context, filter entity.TokenFilter) (*entity.Token, error) {
	var token entity.Token
	err := r.DB.Collection(utils.COLLECTION_TOKENS).FindOne(ctx, r.parseTokenFilter(filter)).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *Repository) parseTokenFilter(filter entity.TokenFilter) bson.D {
	query := bson.D{}

	andCond := make([]bson.E, 0)
	orCond := make([]bson.E, 0)
	// Define your OR query
	if filter.Address != "" {
		andCond = append(andCond, bson.E{Key: "address", Value: filter.Address})
	}

	if filter.Key != "" {
		orCond = append(orCond,
			bson.E{Key: "slug", Value: bson.M{"$regex": strings.ToLower(filter.Key)}},
			bson.E{Key: "name", Value: bson.M{"$regex": filter.Key}},
		)
	}

	if filter.CreatedBy != "" {
		andCond = append(andCond, bson.E{Key: "owner", Value: filter.CreatedBy})
	}

	if len(andCond) > 0 {
		query = append(query, bson.E{Key: "$and", Value: andCond})
	}
	return query
}

func (r *Repository) FindTokens(ctx context.Context, filter entity.TokenFilter) ([]entity.Token, error) {
	var tokens []entity.Token

	// pagination
	page := filter.Page
	if page == 0 {
		page = 1
	}

	numToSkip := (filter.Page - 1) * filter.Limit
	// Set the options for the query
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_TOKENS).Find(ctx, r.parseTokenFilter(filter), options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var token entity.Token
		err = cursor.Decode(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
