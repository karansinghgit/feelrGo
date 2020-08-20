package dao

import (
	"context"
	"encoding/json"

	"github.com/karansinghgit/feelrGo/graphql/model"
	"github.com/olivere/elastic/v7"
)

//AddChat will create a new chat (partner relationship) to the DB
func AddChat(ctx context.Context, client *elastic.Client, index string, s string) error {
	_, err := client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		return err
	}
	return nil
}

//GetChatMessages will fetch chat messages from specified chatID from the DB
func GetChatMessages(ctx context.Context, client *elastic.Client, index string, chatID string, last int) ([]*model.Message, error) {
	chatQuery := elastic.NewMatchQuery("chatID", chatID)
	searchResult, err := client.Search().
		Index(index).
		Query(chatQuery).
		Sort("timestamp", false).
		Size(last).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	var messages []*model.Message

	for _, hit := range searchResult.Hits.Hits {
		var message model.Message
		err := json.Unmarshal(hit.Source, &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}
