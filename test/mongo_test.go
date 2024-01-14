package test

import (
	"os"
	"testing"

	"github.com/jonh-dev/go-goMongo/mongo"
)

func TestConnect(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")

	// Cria uma nova instância de MongoORM
	mongoORM, err := mongo.NewMongoORM("test")
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}

	// Testa a conexão com o MongoDB
	err = mongoORM.Connect()
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}

	// Verifica se o cliente MongoDB foi definido
	if mongoORM.GetClient() == nil {
		t.Fatalf("esperava *mongo.Client, obteve nil")
	}
}
