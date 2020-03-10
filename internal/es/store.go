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
package es

import (
	"context"
	"fmt"

	"github.com/superhero-match/superhero-importer/internal/es/model"

	"gopkg.in/olivere/elastic.v7"
)

// StoreSuperheros saves existing superheros in Elasticsearch.
func (es *ES) StoreSuperheros(superheros []model.Superhero) error {
	bulk := es.Client.Bulk()

	for _, superhero := range superheros {
		req := elastic.NewBulkIndexRequest().
			OpType("index").
			Index(es.Index).
			Doc(superhero)

		bulk = bulk.Add(req)
	}

	fmt.Println("NewBulkIndexRequest().NumberOfActions():", bulk.NumberOfActions())

	bulkResp, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)

		return err
	}

	indexed := bulkResp.Indexed()
	for _, item := range indexed {
		fmt.Printf("Indexed Superhero %s to index %s, type %s\n", item.Id, item.Index, item.Type)
	}

	return nil
}
