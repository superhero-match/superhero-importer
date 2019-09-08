package main

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

