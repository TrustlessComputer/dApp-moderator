package repository

import (
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func (r *Repository) UpdateBnsResolver(tokenID string, resolver string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"resolver": strings.ToLower(resolver)}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *Repository) UpdateBnsPfp(tokenID string, pfp string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"pfp": strings.ToLower(pfp)}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}
