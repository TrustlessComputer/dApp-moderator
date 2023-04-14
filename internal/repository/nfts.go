package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) CollectionsByNfts(ownerAddress string) ([]entity.GroupedCollection, error) {
	f2 := bson.A{
		bson.D{{"$match", bson.D{{"owner", primitive.Regex{Pattern: ownerAddress, Options: "i"}}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", bson.D{{"collection_address", "$collection_address"}}},
					{"tokens", bson.D{{"$sum", 1}}},
				},
			},
		},
		bson.D{{"$project", bson.D{{"tokens", 1}}}},
	}

	groupedNfts := []entity.GroupedCollection{}
	cursor, err := r.DB.Collection(entity.Nfts{}.CollectionName()).Aggregate(context.TODO(), f2)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cursor.All((context.TODO()), &groupedNfts); err != nil {
		return nil, errors.WithStack(err)
	}

	return groupedNfts, nil
}

func (r *Repository) CreateNftHistories(histories *entity.NftHistories) (*mongo.InsertOneResult, error) {
	inserted, err := r.InsertOne(histories)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) GetNft(contract string, tokenID string) (*entity.Nfts, error) {
	nftResp, err := r.FindOne(entity.Nfts{}.CollectionName(), bson.D{
		{"collection_address", contract},
		{"token_id", tokenID},
	})

	if err != nil {
		return nil, err
	}

	nft := &entity.Nfts{}
	err = nftResp.Decode(nft)

	if err != nil {
		return nil, err
	}

	return nft, nil

}

func (r *Repository) UpdateNftOwner(contract string, tokenID string, owner string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"collection_address", contract},
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"owner": strings.ToLower(owner)}}

	updated, err := r.UpdateOne(entity.Nfts{}.CollectionName(), f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil

}
