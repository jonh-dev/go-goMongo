package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jonh-dev/go-goMongo/mongo"
	"github.com/jonh-dev/go-logger/logger"
	"go.mongodb.org/mongo-driver/bson"

	mongodb "go.mongodb.org/mongo-driver/mongo"
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

func TestTransaction(t *testing.T) {
	// Define a URI do MongoDB para o teste
	os.Setenv("MONGO_URI", mongoURI)

	// Cria uma nova instância de MongoORM
	mongoORM, err := mongo.NewMongoORM("test")
	if err != nil {
		t.Fatalf("esperava nil, obteve %v", err)
	}

	// Estabelece a conexão com o MongoDB
	err = mongoORM.Connect()
	if err != nil {
		t.Fatalf("erro ao conectar ao MongoDB: %v", err)
	}

	// Inicia uma nova transação
	sessionContext, err := mongoORM.StartTransaction()
	if err != nil {
		t.Fatalf("erro ao iniciar a transação: %v", err)
	}

	// Cria uma nova TransactionSession e define o contexto da sessão
	ts := &mongo.TransactionSession{}
	ts.SetSessionContext(sessionContext)

	// Executa uma operação dentro da transação
	err = mongoORM.ExecuteInTransaction(ts, "testCollection", bson.M{"nome": "teste"}, func(sessionContext mongodb.SessionContext, collection string, document interface{}) error {
		_, err := mongoORM.Create(collection, document)
		logger.Info(fmt.Sprintf("documento inserido: %v", document))
		return err
	})
	if err != nil {
		t.Fatalf("erro ao executar a operação: %v", err)
	}

	// Verifica se a operação foi aplicada corretamente
	// Por exemplo, você pode verificar se o documento foi inserido corretamente na coleção
	docs, err := mongoORM.Read("testcollection", bson.M{"nome": "teste"})
	if err != nil {
		t.Fatalf("erro ao buscar o documento: %v", err)
	}
	if len(docs) == 0 || docs[0]["nome"] != "teste" {
		t.Fatalf("o documento inserido não é o esperado: %v", docs)
	}
}
