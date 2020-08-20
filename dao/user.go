package dao

import (
	"context"
	"encoding/json"

	"github.com/karansinghgit/feelrGo/graphql/model"
	"github.com/olivere/elastic/v7"
)

//AddUser will add a user to the DB
func AddUser(ctx context.Context, client *elastic.Client, index string, s string) error {
	_, err := client.
		Index().
		BodyString(s).
		Do(ctx)

	if err != nil {
		return err
	}
	return err
}

//GetUser will fetch a user from the DB
func GetUser(ctx context.Context, client *elastic.Client, index string, userID string) (*model.User, error) {
	userQuery := elastic.NewMatchQuery("userID", userID)
	searchResult, err := client.Search().
		Index(index).
		Query(userQuery).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	if searchResult.Hits.TotalHits.Value == 0 {
		return nil, err
	}

	var user *model.User
	json.Unmarshal(searchResult.Hits.Hits[0].Source, &user)

	return user, nil
}

//CheckUsername checks if the username exists in DB
func CheckUsername(ctx context.Context, client *elastic.Client, index string, username string) (string, error) {
	userQuery := elastic.NewMatchQuery("username", username)
	searchResult, err := client.Search().
		Index(index).
		Query(userQuery).
		Do(ctx)

	if err != nil {
		return "false", err
	}

	if searchResult.Hits.TotalHits.Value == 0 {
		return "false", nil
	}

	return "true", nil
}
