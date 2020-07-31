package graph

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

//InitClient method initializes the ElastiClient on a port
func (r *Resolver) InitClient() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	fmt.Println("ES initialized!")
	r.Client = client
}

//Resolver Struct
type Resolver struct {
	Client *elastic.Client
}
