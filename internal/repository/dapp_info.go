package repository

import (
	"context"
	"dapp-moderator/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) InsertDappInfo(data *entity.DappInfo) error {
	_, err := r.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListDappInfo() ([]*entity.DappInfo, error) {
	resp := []*entity.DappInfo{}
	filter := bson.M{}

	cursor, err := r.DB.Collection(entity.DappInfo{}.CollectionName()).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
