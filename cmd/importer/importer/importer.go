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
package importer

import (
	"github.com/superhero-importer/internal/config"
	"github.com/superhero-importer/internal/db"
	"github.com/superhero-importer/internal/es"
)

// Importer holds all the data relevant.
type Importer struct {
	DB       *db.DB
	ES       *es.ES
}

// NewImporter configures Importer.
func NewImporter(cfg *config.Config) (im *Importer, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	return &Importer{
		DB:       dbs,
		ES:       e,
	}, nil
}

