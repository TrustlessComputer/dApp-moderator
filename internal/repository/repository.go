package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/global"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Repository struct {
	Connection *mongo.Client
	DB         *mongo.Database
}

func NewRepository(g *global.Global) (*Repository, error) {

	clientOption := &options.ClientOptions{}
	opt := &options.DatabaseOptions{
		ReadConcern:    clientOption.ReadConcern,
		WriteConcern:   clientOption.WriteConcern,
		ReadPreference: clientOption.ReadPreference,
		Registry:       clientOption.Registry,
	}

	r := new(Repository)
	connection := g.DBConnection.GetType()
	r.Connection = connection.(*mongo.Client)
	r.DB = r.Connection.Database(g.Conf.Databases.Mongo.Name, opt)
	return r, nil
}

func (r Repository) InsertOne(data entity.IEntity) (*mongo.InsertOneResult, error) {
	inserted, err := r.DB.Collection(data.CollectionName()).InsertOne(context.TODO(), &data)
	if err != nil {
		return nil, err
	}
	return inserted,  nil
}

func (r Repository) InsertMany(data []entity.IEntity)  (*mongo.InsertManyResult, error){
	// if len(data) <= 0 {
	// 	return nil, errors.New("Insert data is empty")
	// }

	// inserted, err := r.DB.Collection(data[0].CollectionName()).InsertMany(context.TODO(), data)
	// if err != nil {
	// 	return nil, err
	// }
	// return inserted,  nil

	//TODO - implement me
	return nil,  nil
}

func (r Repository) UpdateOne(collectionName string, filter bson.D, updatedData bson.D) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, updatedData)
	if err != nil {
		return nil, err
	}
	return inserted,  nil
}

func (r Repository) UpdateMany(collectionName string, filter bson.D, updatedData bson.D) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateMany(context.TODO(), filter, updatedData)
	if err != nil {
		return nil, err
	}
	return inserted,  nil
}

func (r Repository) ReplaceOne(filter bson.D,data entity.IEntity) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(data.CollectionName()).ReplaceOne(context.TODO(), filter, &data)
	if err != nil {
		return nil, err
	}
	return inserted,  nil
}