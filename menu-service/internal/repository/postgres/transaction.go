// Package postgre provides an implementation of the domain.Transaction interface
// using PostgreSQL database transactions.
package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

// Transaction represents a PostgreSQL database transaction.
type Transaction struct {
	db          *sqlx.DB
	sqlTx       *sql.Tx
	isCompleted bool
}

// NewTransaction creates a new Transaction for the given sqlx.DB.
func NewTransaction(db *sqlx.DB) (*Transaction, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return &Transaction{
		db:    db,
		sqlTx: tx,
	}, nil
}

// Begin returns the current transaction.
func (t *Transaction) Begin() (domain.Transaction, error) {
	return t, nil
}

// Commit commits the current transaction. If the transaction is already completed,
// it returns nil. If an error occurs during commit, it rolls back the transaction.
func (t *Transaction) Commit() error {
	if t.isCompleted {
		return nil
	}

	if err := t.sqlTx.Commit(); err != nil {
		rberr := t.Rollback()
		if rberr != nil {
			return rberr
		}
		return err
	}

	t.isCompleted = true
	return nil
}

// Rollback rolls back the current transaction. If the transaction is already completed,
// it returns nil. If an error occurs during rollback, it returns the error.
func (t *Transaction) Rollback() error {
	if t.isCompleted {
		return nil
	}

	if err := t.sqlTx.Rollback(); err != nil {
		return err
	}

	t.isCompleted = true
	return nil
}
