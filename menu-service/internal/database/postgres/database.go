package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"menu-service/internal/config"
)

func Connect(cfg *config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password)

	db, err := sqlx.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
