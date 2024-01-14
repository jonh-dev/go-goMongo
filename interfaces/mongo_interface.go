package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoORM interface {
	Connect() error
	GetClient() *mongo.Client
	Create(collection string, document interface{}) (primitive.ObjectID, error)
}
