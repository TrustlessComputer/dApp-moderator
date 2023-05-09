package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindTransactionByHash(ctx context.Context, hash string) (*entity.Transactions, error) {
	var transaction entity.Transactions
	mg := r.DB.Collection(utils.TRANSACTIONS).FindOne(ctx, r.parseTransactionFilter(hash))
	if mg != nil {
		err := mg.Decode(&transaction)
		if err != nil {
			return nil, err
		}
	}

	return &transaction, nil
}

func (r *Repository) parseTransactionFilter(hash string) bson.M {
	andCond := make([]bson.M, 0)
	andCond = append(andCond, bson.M{"hash": hash})
	return bson.M{"$and": andCond}
}
