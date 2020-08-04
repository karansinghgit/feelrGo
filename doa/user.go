package doa

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/karansinghgit/feelrGo/graph/model"
	"github.com/olivere/elastic/v7"
)

func GetUser(ctx context.Context, client *elastic.Client, index string, userID string) (*model.User, error) {
	userQuery := elastic.NewMatchQuery("userID", userID)
	searchResult, err := client.Search().
		Index("feelr").
		Query(userQuery).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	if searchResult.Hits.TotalHits.Value == 0 {
		fmt.Println("The user doesn't exist!")
		return nil, err
	}
	var user *model.User
	json.Unmarshal(searchResult.Hits.Hits[0].Source, &user)
	return user, nil
}
