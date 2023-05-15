package generative_respository

import (
	"dapp-moderator/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (g *GenerativeRepository) InsertFaucet(object *entity.Faucet) (*mongo.InsertOneResult, error) {
	object.ID = primitive.NewObjectID()
	object.UUID = object.ID.Hex()
	return g.InsertOne(object)
}

func (g *GenerativeRepository) FindDappFaucet(userAddress string, source string) (*entity.Faucet, error) {
	f := bson.D{
		{"source", source},
		{"address", userAddress},
	}
	r, err := g.FindOne(entity.Faucet{}.CollectionName(), f)
	if err != nil {
		return nil, err
	}

	faucet := &entity.Faucet{}
	err = r.Decode(faucet)
	if err != nil {
		return nil, err
	}

	return faucet, nil

}
