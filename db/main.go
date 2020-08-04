package db

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

var index = "app"

//try not to have any duplicate names
var mapping = Mapping

//GetNewClient creates and returns a new client
func GetNewClient() *elastic.Client {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	if err != nil {
		fmt.Println("Error initializing a new ES Client:: ", err)
		panic("Client fail ")
	}

	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		fmt.Println("Error pinging the ES Client:: ", err)
		panic(err)
	}

	fmt.Printf("[SUCCESS] Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		fmt.Printf("Index [%v] wasn't found. Initializing index [%v]\n", index, index)
		createIndex, err := client.CreateIndex("app").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("Index Created but couldn't be checked.")
		}
	} else {
		fmt.Printf("Continuing with index [%v]\n", index)
	}
	return client
}
