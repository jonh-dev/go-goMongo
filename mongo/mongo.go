package mongo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jonh-dev/go-logger/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoORM é uma implementação concreta de IMongoORM
type MongoORM struct {
	client   *mongo.Client
	uri      string
	database string // Adicione este campo
}

// NewMongoORM cria uma nova instância de IMongoORM
func NewMongoORM(database string) (IMongoORM, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		logger.Error("MONGO_URI não definida")
		return nil, errors.New("MONGO_URI não definida")
	}
	return &MongoORM{
		uri:      uri,
		database: database, // Inicialize este campo
	}, nil
}

// Connect implementa a função Connect da interface IMongoORM
func (m *MongoORM) Connect() error {
	// Define um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Cria uma nova conexão com o MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	if err != nil {
		logger.Error("erro ao conectar ao MongoDB: %s" + err.Error())
		return fmt.Errorf("erro ao conectar ao MongoDB: %s", err.Error())
	}

	// Verifica a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("erro ao verificar a conexão com o MongoDB: %s" + err.Error())
		return fmt.Errorf("erro ao verificar a conexão com o MongoDB: %s", err.Error())
	}

	logger.Success("conexão com o MongoDB estabelecida com sucesso")

	// Armazena o cliente MongoDB na estrutura MongoORM
	m.client = client

	return nil
}

// StartTransaction inicia uma nova sessão de transação e retorna uma instância de SessionContext.
func (m *MongoORM) StartTransaction() (mongo.SessionContext, error) {
	session, err := m.client.StartSession()
	if err != nil {
		return nil, err
	}

	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}

	ctx := mongo.NewSessionContext(context.Background(), session)

	return ctx, nil
}

// ExecuteInTransaction executa uma operação dentro de uma transação.
func (m *MongoORM) ExecuteInTransaction(ts *TransactionSession, collection string, document interface{}, operation func(mongo.SessionContext, string, interface{}) error) error {
	err := operation(ts.GetSessionContext(), collection, document)
	if err != nil {
		ts.AbortTransaction()
		return err
	}

	err = ts.CommitTransaction()
	if err != nil {
		ts.AbortTransaction()
		return err
	}

	return nil
}

// GetClient implementa a função GetClient da interface IMongoORM
func (m *MongoORM) GetClient() *mongo.Client {
	return m.client
}
