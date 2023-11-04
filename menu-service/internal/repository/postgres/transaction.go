// Package postgres provides an implementation of the domain.Transaction interface
// using PostgreSQL database transactions.
package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

// PostgreSQLTransaction represents a PostgreSQL database transaction.
type PostgreSQLTransaction struct {
	db          *sqlx.DB
	sqlTx       *sql.Tx
	isCompleted bool
}

// NewPostgreSQLTransaction creates a new PostgreSQLTransaction for the given sqlx.DB.
func NewPostgreSQLTransaction(db *sqlx.DB) (*PostgreSQLTransaction, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return &PostgreSQLTransaction{
		db:    db,
		sqlTx: tx,
	}, nil
}

// Begin returns the current transaction.
func (t *PostgreSQLTransaction) Begin() (domain.Transaction, error) {
	return t, nil
}

// Commit commits the current transaction. If the transaction is already completed,
// it returns nil. If an error occurs during commit, it rolls back the transaction.
func (t *PostgreSQLTransaction) Commit() error {
	if t.isCompleted {
		return nil
	}

	if err := t.sqlTx.Commit(); err != nil {
		t.Rollback()
		return err
	}

	t.isCompleted = true
	return nil
}

// Rollback rolls back the current transaction. If the transaction is already completed,
// it returns nil. If an error occurs during rollback, it returns the error.
func (t *PostgreSQLTransaction) Rollback() error {
	if t.isCompleted {
		return nil
	}

	if err := t.sqlTx.Rollback(); err != nil {
		return err
	}

	t.isCompleted = true
	return nil
}
