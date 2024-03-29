package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) FindToken(ctx context.Context, filter entity.TokenFilter) (*entity.Token, error) {
	var token entity.Token
	err := r.DB.Collection(utils.COLLECTION_TOKENS).FindOne(ctx, r.parseTokenFilter(filter)).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *Repository) FindTokensByContracts(ctx context.Context, contracts []string) ([]*entity.Token, error) {
	tokens := []*entity.Token{}

	f := bson.D{{"address", bson.M{"$in": contracts}}}
	cursor, err := r.DB.Collection(utils.COLLECTION_TOKENS).Find(ctx, f)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		token := &entity.Token{}
		err = cursor.Decode(token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (r *Repository) parseTokenFilter(filter entity.TokenFilter) bson.M {

	andCond := make([]bson.M, 0)
	// Define your OR query
	if filter.Address != "" {
		andCond = append(andCond, bson.M{"address": filter.Address})
	}

	if filter.Key != "" {
		andCond = append(andCond, bson.M{"$or": []bson.M{
			{"slug": bson.M{"$regex": strings.ToLower(filter.Key)}},
			{"name": bson.M{"$regex": filter.Key}},
			{"owner": primitive.Regex{Pattern: filter.Key, Options: "i"}},
		}})
	}

	if filter.CreatedBy != "" {
		andCond = append(andCond, bson.M{"owner": filter.CreatedBy})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) FindTokens(ctx context.Context, filter entity.TokenFilter) ([]entity.Token, error) {
	var tokens []entity.Token

	// pagination
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

func (r *Repository) UpdateToken(ctx context.Context, token *entity.Token) error {
	collectionName := token.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"address": token.Address}, bson.M{"$set": token})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *Repository) FindListTokens(ctx context.Context, filter entity.TokenFilter) ([]*entity.Token, error) {
	var tokens []*entity.Token

	// pagination
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
		tokens = append(tokens, &token)
	}
	return tokens, nil
}

func (r *Repository) FindBlackListTokens(ctx context.Context, filter entity.SwapBlackListokenFilter) ([]*entity.SwapBlackListToken, error) {
	tokens := []*entity.SwapBlackListToken{}

	// pagination
	numToSkip := (filter.Page - 1) * filter.Limit
	// Set the options for the query
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(filter.Limit)

	cursor, err := r.DB.Collection(utils.COLLECTION_SWAP_BLACKLIST_TOKENS).Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var token entity.SwapBlackListToken
		err = cursor.Decode(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}

func (r *Repository) UpdateBaseSymbolToken(ctx context.Context, token *entity.Token) error {
	collectionName := token.CollectionName()
	mapBaseTk := map[string]string{
		"token":             token.Address,
		"base_token_symbol": token.BaseTokenSymbol,
	}
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"address": token.Address},
		bson.M{"$set": bson.M{
			"base_token_symbol":     token.BaseTokenSymbol,
			"base_token_symbol_obj": mapBaseTk,
		}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *Repository) AggregateAllTokens(ctx context.Context, filter entity.TokenFilter) ([]entity.Token, error) {
	var tokens []entity.Token
	f := bson.A{
		bson.D{
			{"$project",
				bson.D{
					{"address", 1},
					{"symbol", 1},
					{"decimal", 1},
				},
			},
		},
	}

	cursor, err := r.DB.Collection(utils.COLLECTION_TOKENS).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
