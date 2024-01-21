// crud_test.go
package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jonh-dev/go-goMongo/mongo"
	"github.com/jonh-dev/go-logger/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mongoURI = "mongodb://root:root@localhost:27017/admin?retryWrites=true&loadBalanced=false&serverSelectionTimeoutMS=2000&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-256&directConnection=true"

func TestCreate(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", mongoURI)

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

// crud_test.go
func TestRead(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", mongoURI)

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

	// Define um filtro de teste
	filter := bson.M{"nome": "Teste"}

	// Testa a função Read
	docs, err := mongoORM.Read("testCollection", filter)
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}
	logger.Info(fmt.Sprintf("Documentos retornados: %v\n", docs))

	// Verifica se algum documento foi retornado
	if len(docs) == 0 {
		t.Fatalf("esperava pelo menos um documento, obteve %v", len(docs))
	}
}

// crud_test.go
func TestUpdate(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", mongoURI)

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

	// Define um filtro de teste
	filter := bson.M{"nome": "Teste"}

	// Define uma atualização de teste
	update := bson.M{"$set": bson.M{"email": "teste_atualizado_02@teste.com"}}

	// Testa a função Update
	count, err := mongoORM.Update("testCollection", filter, update)
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}
	logger.Info(fmt.Sprintf("Número de documentos atualizados: %v\n", count))

	// Verifica se algum documento foi atualizado
	if count == 0 {
		t.Fatalf("esperava pelo menos um documento atualizado, obteve %v", count)
	}
}

// crud_test.go
func TestDelete(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", mongoURI)

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

	// Define um filtro de teste
	filter := bson.M{"nome": "Teste"}

	// Testa a função Delete
	count, err := mongoORM.Delete("testCollection", filter)
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}
	logger.Info(fmt.Sprintf("Número de documentos excluídos: %v\n", count))

	// Verifica se algum documento foi excluído
	if count == 0 {
		t.Fatalf("esperava pelo menos um documento excluído, obteve %v", count)
	}
}
