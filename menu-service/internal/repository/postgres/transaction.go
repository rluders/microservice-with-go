package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

type PostgreSQLTransaction struct {
	db          *sqlx.DB
	sqlTx       *sql.Tx
	isCompleted bool
}

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

func (t *PostgreSQLTransaction) Begin() (domain.Transaction, error) {
	return t, nil
}

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
