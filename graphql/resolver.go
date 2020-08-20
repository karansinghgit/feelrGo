package graphql

import (
	"github.com/olivere/elastic/v7"
)

// Resolver implementation will be done here
type Resolver struct {
	client *elastic.Client
	index  string
}
