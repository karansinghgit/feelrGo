package doa

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/karansinghgit/feelrGo/graph/model"
	"github.com/karansinghgit/feelrGo/utils"
	"github.com/olivere/elastic/v7"
)

func AddChat(ctx context.Context, client *elastic.Client, index string, senderID string, receiverID string) (*model.Chat, error) {
	c := &model.Chat{
		ChatID:     uuid.New().String(),
		SenderID:   senderID,
		ReceiverID: receiverID,
	}

	s, err := utils.ParseToString(c)

	if err != nil {
		return nil, err
	}

	_, err = client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		fmt.Println("Error Storing the Chat")
		return nil, err
	}
	return c, nil
}

func GetChatMessages(ctx context.Context, client *elastic.Client, index string, chatID string, last int) ([]*model.Message, error) {
	chatQuery := elastic.NewMatchQuery("chatID", chatID)
	searchResult, err := client.Search().
		Index("feelr").
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
