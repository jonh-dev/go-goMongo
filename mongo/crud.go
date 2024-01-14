package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create implementa a função Create da interface IMongoORM
func (m *MongoORM) Create(collection string, document interface{}) (primitive.ObjectID, error) {
	// Verifica se o cliente MongoDB está conectado
	if m.client == nil {
		return primitive.NilObjectID, errors.New("cliente MongoDB não está conectado")
	}

	// Obtém a coleção do MongoDB
	coll := m.client.Database(m.database).Collection(collection) // Use o campo database aqui

	// Insere o documento na coleção
	res, err := coll.InsertOne(context.Background(), document)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("erro ao inserir documento: %s", err.Error())
	}

	// Retorna o ID do documento inserido
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("não foi possível obter o ID do documento inserido")
	}

	return id, nil
}

// TODO: Implementar as operações Read, Update e Delete
