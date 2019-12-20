package es

import (
	"context"
	"fmt"

	"github.com/superhero-importer/internal/es/model"

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
