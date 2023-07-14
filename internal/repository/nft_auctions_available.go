package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) InsertAuction(data *entity.NftAuctionsAvailable) error {
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

func (r *Repository) FindAuction(contractAddress string, tokenID string) (*entity.NftAuctionsAvailable, error) {

	filter := bson.D{
		{"collection_address", strings.ToLower(contractAddress)},
		{"token_id", tokenID},
	}

	resp := r.DB.Collection(utils.VIEW_NFT_AUCTION_AVAILABLE).FindOne(context.TODO(), filter)

	data := &entity.NftAuctionsAvailable{}
	err := resp.Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) NftWithoutCapturedImage(contractAddress string, offset int, limit int) ([]entity.Nfts, error) {
	resp := []entity.Nfts{}

	f := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"collection_address", strings.ToLower(contractAddress)},
					{"image_capture_at", bson.D{{"$in", bson.A{
						"",
						primitive.Null{},
					},
					}}},
				},
			},
		},
		bson.D{{"$skip", offset}},
		bson.D{{"$limit", limit}},
	}

	cursor, err := r.DB.Collection(entity.Nfts{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) NftCapturedImageHistories(contractAddress string, offset int, limit int, specialNfts []string) ([]entity.Nfts, error) {
	resp := []entity.Nfts{}
	f := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"collection_address", strings.ToLower(contractAddress)},
					//{"token_id", bson.M{"$in": specialNfts}},
				},
			},
		},
		bson.D{{"$skip", offset}},
		bson.D{{"$limit", limit}},
	}

	cursor, err := r.DB.Collection(entity.Nfts{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
