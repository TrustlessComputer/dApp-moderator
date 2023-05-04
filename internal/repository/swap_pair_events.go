package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindSwapPairEvents(ctx context.Context, filter entity.SwapPairEventFilter) (*entity.SwapPairEvents, error) {
	var swapPairEvents entity.SwapPairEvents
	err := r.DB.Collection(utils.COLLECTION_SWAP_PAIR_EVENTS).FindOne(ctx, r.parseSwapPairEventsFilter(filter)).Decode(&swapPairEvents)
	if err != nil {
		return nil, err
	}
	return &swapPairEvents, nil
}

func (r *Repository) parseSwapPairEventsFilter(filter entity.SwapPairEventFilter) bson.M {
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
