package repository

import (
	"context"
	"dapp-moderator/internal/entity"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)


func (r *Repository) CollectionsByNfts(ownerAddress string) ([]entity.GroupedCollection,  error) {
	f2 := bson.A{
		bson.D{{"$match", bson.D{{"owner",  ownerAddress}}}},
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