package db

import (
	"github.com/superhero-importer/internal/db/model"
)

// GetSuperheros fetches a batch of superheros.
func (db *DB) GetSuperheros(offset int64) (superheros map[string]model.Superhero, err error) {
	sups := make([]model.Superhero, 0)

	err = db.stmtGetSuperheros.Select(&sups, offset, db.Limit)
	if err != nil {
		return nil, err
	}

	superheros = make(map[string]model.Superhero)

	for _, superhero := range sups {
		superheros[superhero.ID] = superhero
	}

	return superheros, nil
} 
