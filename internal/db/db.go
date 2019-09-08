package db

import (
	"database/sql"
	"fmt"

	"github.com/superhero-importer/internal/config"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
)

// DB holds the database connection.
type DB struct {
	DB *sql.DB
}

// NewDB returns database.
func NewDB(cfg *config.Config) (dbs *DB, err error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Name,
		),
	)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: db,
	}, nil
}