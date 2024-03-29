package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/global"
	"dapp-moderator/utils/helpers"
	"errors"

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

func (r *Repository) InsertOne(data entity.IEntity) (*mongo.InsertOneResult, error) {
	data.SetID()
	data.SetCreatedAt()
	insertedData, err := helpers.ToDoc(data)
	if err != nil {
		return nil, err
	}

	collectionName := data.CollectionName()
	inserted, err := r.DB.Collection(collectionName).InsertOne(context.TODO(), *insertedData)
	if err != nil {
		return nil, err
	}

	return inserted, nil
}

func (r *Repository) InsertOneWithContext(ctx context.Context, data entity.IEntity) (*mongo.InsertOneResult, error) {
	data.SetID()
	data.SetCreatedAt()
	insertedData, err := helpers.ToDoc(data)
	if err != nil {
		return nil, err
	}

	collectionName := data.CollectionName()
	inserted, err := r.DB.Collection(collectionName).InsertOne(ctx, *insertedData)
	if err != nil {
		return nil, err
	}

	return inserted, nil
}

func (r *Repository) WithTransaction(ctx context.Context, callback func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	session, err := r.DB.Client().StartSession()
	defer session.EndSession(ctx)
	if err != nil {
		return nil, err
	}
	return session.WithTransaction(ctx, callback, opts...)
}

func (r *Repository) InsertMany(data []entity.IEntity) (*mongo.InsertManyResult, error) {
	if len(data) <= 0 {
		return nil, errors.New("Insert data is empty")
	}
	insertedData := make([]interface{}, 0)
	for _, item := range data {
		item.SetID()
		item.SetCreatedAt()
		tmp, err := helpers.ToDoc(item)
		if err != nil {
			return nil, err
		}
		insertedData = append(insertedData, *tmp)
	}

	opts := options.InsertMany().SetOrdered(false)
	inserted, err := r.DB.Collection(data[0].CollectionName()).InsertMany(context.TODO(), insertedData, opts)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) UpdateOne(collectionName string, filter bson.D, updatedData bson.M) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, updatedData)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) UpdateOneWithOptions(collectionName string, filter bson.D, updatedData bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, updatedData, opts...)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) UpdateMany(collectionName string, filter bson.D, updatedData bson.M) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateMany(context.TODO(), filter, updatedData)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) UpdateManyWithOptions(collectionName string, filter bson.D, updatedData bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	inserted, err := r.DB.Collection(collectionName).UpdateMany(context.TODO(), filter, updatedData, opts...)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) ReplaceOne(filter bson.D, data entity.IEntity) (*mongo.UpdateResult, error) {
	bsonData, err := helpers.ToDoc(data)
	if err != nil {
		return nil, err
	}

	inserted, err := r.DB.Collection(data.CollectionName()).ReplaceOne(context.TODO(), filter, bsonData)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (r *Repository) DeleteOne(collectionName string, filter bson.D) (*mongo.DeleteResult, error) {
	deleted, err := r.DB.Collection(collectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}

func (r *Repository) DeleteMany(collectionName string, filter bson.D) (*mongo.DeleteResult, error) {
	deleted, err := r.DB.Collection(collectionName).DeleteMany(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}

func (r *Repository) CountDocuments(collectionName string, filter bson.D) (*int64, *int64, error) {
	estCount, estCountErr := r.DB.Collection(collectionName).EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		return nil, nil, estCountErr
	}
	count, err := r.DB.Collection(collectionName).CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, nil, err
	}

	return &count, &estCount, nil
}

func (r *Repository) FindOne(collectionName string, filter bson.D) (*mongo.SingleResult, error) {

	sr := r.DB.Collection(collectionName).FindOne(context.TODO(), filter)
	if sr.Err() != nil {
		return nil, sr.Err()
	}

	return sr, nil
}

func (r *Repository) FindOneWithResult(collectionName string, filter bson.M, result interface{}, opts ...*options.FindOneOptions) error {
	sr := r.DB.Collection(collectionName).FindOne(context.TODO(), filter, opts...)
	if sr.Err() != nil {
		return sr.Err()
	}

	if err := sr.Decode(result); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Find(collectionName string, filter bson.D, limit int64, offset int64, result interface{}, sort bson.D) error {
	opts := &options.FindOptions{}
	opts.Limit = &limit
	opts.Skip = &offset
	opts.Sort = sort

	cursor, err := r.DB.Collection(collectionName).Find(context.TODO(), filter, opts)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindWithProjections(collectionName string, filter bson.D, limit int64, offset int64, result interface{}, sort bson.D, project bson.D) error {
	opts := &options.FindOptions{}
	opts.Limit = &limit
	opts.Skip = &offset
	opts.Sort = sort
	opts.Projection = project

	cursor, err := r.DB.Collection(collectionName).Find(context.TODO(), filter, opts)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := cursor.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (r *Repository) AllItems(collectionName string, filter bson.D) (int64, error) {
	count, err := r.DB.Collection(collectionName).CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Count before add sort and skip
func (r *Repository) CountTotalFromPipeline(collectionName string, pipelines bson.A) (int32, error) {
	var countPipeline = bson.A(make([]interface{}, len(pipelines), len(pipelines)))
	copy(countPipeline, pipelines) // để ko change trên pipeline truyền vào
	countPipeline = append(countPipeline, bson.M{"$count": "total"})
	totalCur, err := r.DB.Collection(collectionName).Aggregate(context.TODO(), countPipeline)
	if err != nil {
		return 0, err
	}
	var totalResult []bson.M
	if err := totalCur.All(context.TODO(), &totalResult); err != nil {
		return 0, err
	}
	if len(totalResult) > 0 {
		if _val, ok := totalResult[0]["total"].(int32); ok {
			return _val, nil
		}
	}

	return 0, nil
}
