package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

func (r *Repository) InsertAuction(data *entity.NftAuctions) error {
	now := time.Now().UTC()
	data.UpdatedAt = &now
	//data.SetID()

	filter := bson.D{
		{"contract", strings.ToLower(data.ContractAddress)},
		{"token_id", strings.ToLower(data.TokenID)},
	}

	options := &options2.UpdateOptions{}
	options.SetUpsert(true)

	updated, err := r.DB.Collection(utils.COLLECTION_NFT_AUCTIONS).UpdateOne(context.TODO(), filter, bson.M{"$set": data}, options)
	if err != nil {
		return err
	}

	if updated.UpsertedID != nil {
		createdAt := time.Now().UTC()

		_, err := r.DB.Collection(utils.COLLECTION_NFT_AUCTIONS).UpdateOne(context.TODO(), filter, bson.M{"$set": bson.M{"created_at": createdAt}}, options)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) FindAuction(contractAddress string, tokenID string) (*entity.NftAuctions, error) {

	filter := bson.D{
		{"collection_address", strings.ToLower(contractAddress)},
		{"token_id", tokenID},
	}

	resp := r.DB.Collection(utils.VIEW_NFT_AUCTION).FindOne(context.TODO(), filter)

	data := &entity.NftAuctions{}
	err := resp.Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
