package repository

import (
	"dapp-moderator/utils/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

const MongoIdGenCollectionName = "id-gen"

type Repository struct {
	Connection *mongo.Client
	Logger     logger.Ilogger

	DB         *mongo.Database
}
