package doa

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/karansinghgit/feelrGo/graph/model"
	"github.com/karansinghgit/feelrGo/utils"
	"github.com/olivere/elastic/v7"
)

func AddTextMessage(ctx context.Context, client *elastic.Client, index string, chatID string, senderID string, text string) (*model.Message, error) {
	m := &model.Message{
		ChatID:    chatID,
		SenderID:  senderID,
		Text:      &text,
		Timestamp: time.Now(),
	}

	s, err := utils.ParseToString(m)

	if err != nil {
		return nil, err
	}

	_, err = client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		fmt.Println("Error Storing the Text Message")
		return nil, err
	}
	return m, nil
}

func AddFeelrMessage(ctx context.Context, client *elastic.Client, index string, chatID string, senderID string, feelrID string, answer string) (*model.Message, error) {
	m := &model.Message{
		ChatID:       chatID,
		SenderID:     senderID,
		FeelrID:      &feelrID,
		SenderAnswer: &answer,
		Timestamp:    time.Now(),
	}

	s, err := utils.ParseToString(m)

	if err != nil {
		return nil, err
	}

	_, err = client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		fmt.Println("Error Storing the Text Message")
		return nil, err
	}
	return m, nil
}

func GetFeelrMessage(ctx context.Context, client *elastic.Client, index string, chatID string, feelrID string) (string, error) {
	chatQuery := elastic.NewMatchQuery("chatId", chatID)
	feelrQuery := elastic.NewMatchQuery("feelrId", feelrID)

	query := elastic.NewBoolQuery().Must(chatQuery, feelrQuery)
	searchResult, err := client.Search().
		Index(index).
		Query(query).
		Do(ctx)

	if err != nil {
		fmt.Println("Something bad happened with the Combination")
		return "", err
	}

	if searchResult.Hits.TotalHits.Value > 0 {
		return searchResult.Hits.Hits[0].Id, nil
	}

	return "", err
}

func SendMessageResponse(ctx context.Context, client *elastic.Client, index string, docId string, answer string) (*model.Message, error) {
	var m model.Message

	_, err := client.Update().Index(index).Id(docId).Doc(map[string]interface{}{"receiverAnswer": answer}).Do(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.Get().Index(index).Id(docId).Do(ctx)
	if err != nil {
		fmt.Println("Error initializing : ", err)
		return nil, err
	}

	err = json.Unmarshal(res.Source, &m)
	if err != nil {
		fmt.Println("Error unmarshaling : ", err)
		return nil, err
	}
	return &m, nil
}
