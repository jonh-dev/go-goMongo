// crud_test.go
package test

import (
	"os"
	"testing"

	"github.com/jonh-dev/go-goMongo/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", "mongodb://root:root@localhost:27017/admin?retryWrites=true&loadBalanced=false&serverSelectionTimeoutMS=2000&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-256&directConnection=true")

	// Cria uma nova instância de MongoORM
	mongoORM, err := mongo.NewMongoORM("test") // Adicione o nome do banco de dados aqui
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}

	// Testa a conexão com o MongoDB
	err = mongoORM.Connect()
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}

	// Define um documento de teste
	doc := map[string]string{
		"nome":  "Teste",
		"email": "teste@teste.com",
	}

	// Testa a função Create
	id, err := mongoORM.Create("testCollection", doc)
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}

	// Verifica se o ID retornado é válido
	if id == primitive.NilObjectID {
		t.Fatalf("esperava um ID válido, obteve %v", id)
	}
}
