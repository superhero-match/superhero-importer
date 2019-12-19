package main

import (
	"fmt"
	"github.com/superhero-importer/cmd/importer/importer"

	"github.com/superhero-importer/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	_, err = importer.NewImporter(cfg)
	if err != nil {
		panic(err)
	}

	//err = impr.ES.DeleteIndex()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = impr.ES.CreateIndex()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("Everything works!!!")
}