package generative_respository

import (
	"dapp-moderator/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (g *GenerativeRepository) InsertFaucet(object *entity.Faucet) (*mongo.InsertOneResult, error) {
	return g.InsertOne(object)
}
