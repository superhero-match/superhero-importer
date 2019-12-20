package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
	"github.com/jmoiron/sqlx"
	"github.com/superhero-importer/internal/config"
)

// DB holds the database connection.
type DB struct {
	DB                     *sqlx.DB
	Limit                  int
	stmtGetSuperheros      *sqlx.Stmt
	stmtGetProfilePictures *sqlx.Stmt
}

// NewDB returns database.
func NewDB(cfg *config.Config) (dbs *DB, err error) {
	db, err := sqlx.Open(
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

	stmtGetSuperheros, err := db.Preparex(`call get_superheros(?,?)`)
	if err != nil {
		return nil, err
	}

	stmtGetProfilePictures, err := db.Preparex(`call get_profile_pictures(?)`)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:                     db,
		Limit:                  cfg.DB.Limit,
		stmtGetSuperheros:      stmtGetSuperheros,
		stmtGetProfilePictures: stmtGetProfilePictures,
	}, nil
}
