package repository

import (
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//ETH
func (r Repository) FindUserByWalletAddress(walletAddress string) (*entity.Users, error) {
	resp := &entity.Users{}
	f := bson.D{{utils.KEY_WALLET_ADDRESS, primitive.Regex{Pattern: walletAddress, Options: "i"}}}

	usr, err := r.FindOne(utils.COLLECTION_USERS, f)
	if err != nil {
		return nil, err
	}

	err = usr.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r Repository) FindUserByBTCWalletAddress(walletAddress string) (*entity.Users, error) {
	resp := &entity.Users{}
	f := bson.D{{utils.KEY_WALLET_ADDRESS_BTC, primitive.Regex{Pattern: walletAddress, Options: "i"}}}

	usr, err := r.FindOne(utils.COLLECTION_USERS, f)
	if err != nil {
		return nil, err
	}

	err = usr.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r Repository) FindUserByBTCTaprootWalletAddress(walletAddress string) (*entity.Users, error) {
	resp := &entity.Users{}
	f := bson.D{{utils.KEY_WALLET_ADDRESS_BTC_TAPROOT, primitive.Regex{Pattern: walletAddress, Options: "i"}}}

	usr, err := r.FindOne(utils.COLLECTION_USERS, f)
	if err != nil {
		return nil, err
	}

	err = usr.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r Repository) UpdateUserMessage(walletAddress string, message string) (*mongo.UpdateResult, error) {
	now := time.Now().UTC()
	f := bson.D{{utils.KEY_WALLET_ADDRESS, primitive.Regex{Pattern: walletAddress, Options: "i"}}}
	data := bson.M{"$set": bson.M{
		"message": message,
		"updated_at" :now,
	}}

	updated, err := r.UpdateOne(utils.COLLECTION_USERS, f, data)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r Repository) UpdateUserLastLoggedIn(walletAddress string) (*mongo.UpdateResult, error) {
	now := time.Now().UTC()
	f := bson.D{{utils.KEY_WALLET_ADDRESS, primitive.Regex{Pattern: walletAddress, Options: "i"}}}
	data := bson.M{"$set": bson.M{
		"updated_at" :now,
		"last_loggedin_at": now,
	}}

	updated, err := r.UpdateOne(utils.COLLECTION_USERS, f, data)
	if err != nil {
		return nil, err
	}

	return updated, nil
}