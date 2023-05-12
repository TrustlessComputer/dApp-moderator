package generative_respository

import (
	"dapp-moderator/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (g *GenerativeRepository) InsertFaucet(object *entity.Faucet) (*mongo.InsertOneResult, error) {
	object.ID = primitive.NewObjectID()
	object.UUID = object.ID.Hex()
	return g.InsertOne(object)
}
