package db

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

//try not to have any duplicate names
var mappinng = `{
	"mappings": {
	  "properties": {
		"feelra": {
		  "properties": {
			"id": {
			  "type": "text"
			},
			"question": {
			  "type": "text"
			},
			"timestamp": {
			  "type": "text"
			},
			"topic": {
			  "type": "text"
			}
		  }
		},
		"message": {
		  "properties": {
			"chat": {
			  "type": "text"
			},
			"sender": {
			  "type": "text"
			},
			"text": {
			  "type": "text"
			},
			"feelr": {
			  "type": "text"
			},
			"senderAnswer": {
			  "type": "text"
			},
			"receiverAnswer": {
			  "type": "text"
			},
			"timestamp": {
			  "type": "text"
			}
		  }
		},
		"user": {
		  "properties": {
			"id": {
			  "type": "text"
			},
			"name": {
		      "type": "text"
			}
		  }
		},
		"chata": {
		  "properties": {
			"id": {
			  "type": "text"
			},
			"sender": {
			  "type": "text"
			},
			"receiver": {
			  "type": "text"
			}
		  }
		}
	  }
	}
  }`

//GetNewClient creates and returns a new client
func GetNewClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	ctx := context.Background()

	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	exists, err := client.IndexExists("app").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(exists)

	if !exists {
		fmt.Println("not found feelr index")
		createIndex, err := client.CreateIndex("app").BodyString(mappinng).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("Index could not be ack")
		}
	}
	return client
}
