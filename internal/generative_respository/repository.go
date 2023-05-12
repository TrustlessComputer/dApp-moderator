package generative_respository

import (
	"dapp-moderator/internal/repository"
	"dapp-moderator/utils/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenerativeRepository struct {
	repository.Repository
}

func NewGenerativeRepository(g *global.Global) (*GenerativeRepository, error) {

	clientOption := &options.ClientOptions{}
	opt := &options.DatabaseOptions{
		ReadConcern:    clientOption.ReadConcern,
		WriteConcern:   clientOption.WriteConcern,
		ReadPreference: clientOption.ReadPreference,
		Registry:       clientOption.Registry,
	}

	r := new(GenerativeRepository)
	connection := g.GenerativeDBConnection.GetType()
	r.Connection = connection.(*mongo.Client)
	r.DB = r.Connection.Database(g.Conf.Databases.GenerativeMongo.Name, opt)
	return r, nil
}
