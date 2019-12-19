package es

import (
	"context"
	"fmt"
)

// DeleteIndex deletes the whole index.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! DO NOT CALL THIS METHOD THIS IS ONLY FOR LOCAL DEVELOPMENT PURPOSES !!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func(es *ES) DeleteIndex() error {
	deleteIndex, err := es.Client.DeleteIndex(es.Index).Do(context.Background())
	if err != nil {
		fmt.Println("deleteIndex")
		fmt.Println(err)
		return err
	}

	fmt.Printf("%+v", deleteIndex)

	return nil
}