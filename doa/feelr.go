package doa

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karansinghgit/feelrGo/graph/model"
	"github.com/karansinghgit/feelrGo/utils"
	"github.com/olivere/elastic/v7"
)

func AddFeelr(ctx context.Context, client *elastic.Client, index string, question string, topic string) (*model.Feelr, error) {
	f := &model.Feelr{
		FeelrID:   uuid.New().String(),
		Question:  question,
		Topic:     topic,
		Timestamp: time.Now(),
	}

	s, err := utils.ParseToString(f)

	if err != nil {
		return nil, err
	}

	_, err = client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		fmt.Println("Error Storing the Feelr")
		return nil, err
	}
	return f, err
}

func GetFeelrs(ctx context.Context, client *elastic.Client, index string, count int) ([]*model.Feelr, error) {
	existsQuery := elastic.NewExistsQuery("senderID")
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
