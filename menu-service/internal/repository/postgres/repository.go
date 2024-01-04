package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	statements map[string]*sqlx.NamedStmt
	DB         *sqlx.DB
}

func (r *Repository) Statement(query string) (*sqlx.NamedStmt, error) {
	stmt, ok := r.statements[query]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", query)
	}

	return stmt, nil
}

func prepareStatements(db *sqlx.DB, queries map[string]string) (map[string]*sqlx.NamedStmt, error) {
	statements := make(map[string]*sqlx.NamedStmt, len(queries))

	for queryName, query := range queries {
		stmt, err := db.PrepareNamed(query)
		if err != nil {
			log.Printf("error preparing statement %s: %v", queryName, err)
			return statements, err
		}
		statements[queryName] = stmt
	}

	return statements, nil
}
