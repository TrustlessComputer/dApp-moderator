package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindSwapPairSync(ctx context.Context, filter entity.SwapPairSyncFilter) (*entity.SwapPairSync, error) {
	var swapPairSync entity.SwapPairSync
	err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR_SYNC).FindOne(ctx, r.parseSwapPairSyncFilter(filter)).Decode(&swapPairSync)
	if err != nil {
		return nil, err
	}
	return &swapPairSync, nil
}

func (r *Repository) parseSwapPairSyncFilter(filter entity.SwapPairSyncFilter) bson.M {
	andCond := make([]bson.M, 0)
	// Define your OR query
	if filter.ContractAddress != "" {
		andCond = append(andCond, bson.M{"contract_address": filter.ContractAddress})
	}
	if filter.TxHash != "" {
		andCond = append(andCond, bson.M{"tx_hash": filter.TxHash})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}
