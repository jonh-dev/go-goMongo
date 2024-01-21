package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

// Read implementa a função Read da interface IMongoORM
func (m *MongoORM) Read(collection string, filter interface{}) ([]bson.M, error) {
	// Verifica se o cliente MongoDB está conectado
	if m.client == nil {
		return nil, errors.New("cliente MongoDB não está conectado")
	}

	// Obtém a coleção do MongoDB
	coll := m.client.Database(m.database).Collection(collection)

	// Busca documentos na coleção que correspondem ao filtro
	cur, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar documentos: %s", err.Error())
	}
	defer cur.Close(context.Background())

	// Decodifica os documentos retornados
	var results []bson.M
	if err := cur.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("erro ao decodificar documentos: %s", err.Error())
	}

	return results, nil
}

// Update implementa a função Update da interface IMongoORM
func (m *MongoORM) Update(collection string, filter interface{}, update interface{}) (int64, error) {
	// Verifica se o cliente MongoDB está conectado
	if m.client == nil {
		return 0, errors.New("cliente MongoDB não está conectado")
	}

	// Obtém a coleção do MongoDB
	coll := m.client.Database(m.database).Collection(collection)

	// Atualiza documentos na coleção que correspondem ao filtro
	res, err := coll.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return 0, fmt.Errorf("erro ao atualizar documentos: %s", err.Error())
	}

	// Retorna o número de documentos atualizados
	return res.ModifiedCount, nil
}

// Delete implementa a função Delete da interface IMongoORM
func (m *MongoORM) Delete(collection string, filter interface{}) (int64, error) {
	// Verifica se o cliente MongoDB está conectado
	if m.client == nil {
		return 0, errors.New("cliente MongoDB não está conectado")
	}

	// Obtém a coleção do MongoDB
	coll := m.client.Database(m.database).Collection(collection)

	// Exclui documentos na coleção que correspondem ao filtro
	res, err := coll.DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, fmt.Errorf("erro ao excluir documentos: %s", err.Error())
	}

	// Retorna o número de documentos excluídos
	return res.DeletedCount, nil
}
