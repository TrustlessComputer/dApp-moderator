package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) FindSwapWallet(ctx context.Context, filter entity.SwapWalletAddressFilter) (*entity.SwapWalletAddress, error) {
	var swapConfigs entity.SwapWalletAddress
	err := r.DB.Collection(utils.COLLECTION_SWAP_WALLET_ADDRESS).FindOne(ctx, r.parseSwapWalletFilter(filter)).Decode(&swapConfigs)
	if err != nil {
		return nil, err
	}
	return &swapConfigs, nil
}

func (r *Repository) FindSwapWalletByAddress(ctx context.Context, address string) (*entity.SwapWalletAddress, error) {
	var swapConfigs entity.SwapWalletAddress
	filter := entity.SwapWalletAddressFilter{
		Address: address,
	}

	err := r.DB.Collection(utils.COLLECTION_SWAP_WALLET_ADDRESS).FindOne(ctx, r.parseSwapWalletFilter(filter)).Decode(&swapConfigs)
	if err != nil {
		return nil, err
	}
	return &swapConfigs, nil
}

func (r *Repository) parseSwapWalletFilter(filter entity.SwapWalletAddressFilter) bson.M {
	andCond := make([]bson.M, 0)
	if filter.Address != "" {
		andCond = append(andCond, bson.M{"address": filter.Address})
	}

	if len(andCond) == 0 {
		return bson.M{}
	}
	return bson.M{"$and": andCond}
}

func (r *Repository) UpdateWallet(ctx context.Context, wallet *entity.SwapWalletAddress) error {
	collectionName := wallet.CollectionName()
	result, err := r.DB.Collection(collectionName).UpdateOne(ctx, bson.M{"_id": wallet.ID}, bson.M{"$set": wallet})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
