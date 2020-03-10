/*
  Copyright (C) 2019 - 2020 MWSOFT
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
	"github.com/superhero-match/superhero-importer/internal/db/model"
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
