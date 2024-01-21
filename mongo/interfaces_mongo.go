package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoORM interface {
	Connect() error
	GetClient() *mongo.Client
	StartTransaction() (mongo.SessionContext, error)
	ExecuteInTransaction(ts *TransactionSession, collection string, document interface{}, operation func(mongo.SessionContext, string, interface{}) error) error
	Create(collection string, document interface{}) (primitive.ObjectID, error)
	Read(collection string, filter interface{}) ([]bson.M, error)
	Update(collection string, filter interface{}, update interface{}) (int64, error) // Adicionado aqui
	Delete(collection string, filter interface{}) (int64, error)                     // Adicionado aqui
}
