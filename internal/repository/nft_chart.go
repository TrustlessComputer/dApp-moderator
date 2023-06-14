package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func (r *Repository) GetCollectionChart(filter entity.FilterCollectionChart) ([]*entity.CollectionChart, error) {
	f2 := bson.A{
		bson.D{
			{"$match", bson.D{
				{"contract", strings.ToLower(*filter.ContractAddress)},
				{"volume_created_at", bson.M{"$gte": filter.FromDate}},
				{"volume_created_at", bson.M{"$lte": filter.ToDate}},
			}},
		},
		bson.D{{"$sort", bson.D{{"volume_created_at", -1}}}},
	}

	groupedNfts := []*entity.CollectionChart{}
	cursor, err := r.DB.Collection(entity.CollectionChart{}.CollectionName()).Aggregate(context.TODO(), f2)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cursor.All((context.TODO()), &groupedNfts); err != nil {
		return nil, errors.WithStack(err)
	}

	return groupedNfts, nil
}
