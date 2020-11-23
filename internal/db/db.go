/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
	"github.com/jmoiron/sqlx"
	"github.com/superhero-match/superhero-importer/internal/config"
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
