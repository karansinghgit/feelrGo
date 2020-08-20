package dao

import (
	"context"
	"encoding/json"

	"github.com/karansinghgit/feelrGo/graphql/model"
	"github.com/olivere/elastic/v7"
)

//AddFeelr will add a feelr to the DB
func AddFeelr(ctx context.Context, client *elastic.Client, index string, s string) error {
	_, err := client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		return err
	}
	return nil
}

//GetFeelrs will fetch feelrs from the DB
func GetFeelrs(ctx context.Context, client *elastic.Client, index string, count int) ([]*model.Feelr, error) {
	existsQuery := elastic.NewExistsQuery("question")
	searchResult, err := client.Search().
		Index(index).
		Query(existsQuery).
		Sort("timestamp", false).
		Size(count).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	var feelrs []*model.Feelr
	for _, hit := range searchResult.Hits.Hits {
		var feelr model.Feelr
		err := json.Unmarshal(hit.Source, &feelr)
		if err != nil {
			return nil, err
		}
		feelrs = append(feelrs, &feelr)
	}

	return feelrs, nil
}
