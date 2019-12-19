package db

import (
	"github.com/superhero-importer/internal/db/model"
)

// GetSuperheros fetches a batch of superheros.
func (db *DB) GetSuperheros(offset int64) (superheros map[string]model.Superhero, err error) {
	rows, err := db.stmtGetSuperheros.Query(offset, db.Limit)
	if err != nil {
		return nil, err
	}

	superheros = make(map[string]model.Superhero)

	for rows.Next() {
		var superhero model.Superhero

		err = rows.Scan(&superhero)
		if err != nil {
			return nil, err
		}

		superheros[superhero.ID] = superhero
	}

	return superheros, nil
} 
