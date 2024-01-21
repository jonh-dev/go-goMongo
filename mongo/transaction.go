package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// TransactionSession encapsula uma sessão de transação do MongoDB.
type TransactionSession struct {
	sessionContext mongo.SessionContext
}

// GetSessionContext retorna o contexto da sessão.
func (ts *TransactionSession) GetSessionContext() mongo.SessionContext {
	return ts.sessionContext
}

// SetSessionContext define o contexto da sessão.
func (ts *TransactionSession) SetSessionContext(sessionContext mongo.SessionContext) {
	ts.sessionContext = sessionContext
}

// StartTransaction inicia uma nova transação.
func (ts *TransactionSession) StartTransaction() error {
	return ts.sessionContext.StartTransaction()
}

// CommitTransaction confirma a transação atual.
func (ts *TransactionSession) CommitTransaction() error {
	return ts.sessionContext.CommitTransaction(context.Background())
}

// AbortTransaction aborta a transação atual.
func (ts *TransactionSession) AbortTransaction() error {
	return ts.sessionContext.AbortTransaction(context.Background())
}
