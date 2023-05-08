package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) GetAllAllowList() ([]entity.AllowList, error) {

	allowList := []entity.AllowList{}
	cursor, err := r.DB.Collection(entity.AllowList{}.CollectionName()).Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &allowList); err != nil {
		return nil, err
	}

	return allowList, nil
}

func (r *Repository) GetInsertedAllowList(walletAddress string) (*entity.InsertedAllowList, error) {

	insertedAllowList := &entity.InsertedAllowList{}
	usr, err := r.FindOne(utils.COLLECTION_INSERTED_ALLOW_LIST, bson.D{{"wallet_address", walletAddress}})
	if err != nil {
		return nil, err
	}

	err = usr.Decode(insertedAllowList)
	if err != nil {
		return nil, err
	}

	return insertedAllowList, nil
}

func (r *Repository) CreateInsertedAllowList(obj *entity.InsertedAllowList) (*mongo.InsertOneResult, error) {

	res, err := r.InsertOne(obj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
