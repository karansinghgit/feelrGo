package dao

import (
	"context"
	"encoding/json"

	"github.com/karansinghgit/feelrGo/graphql/model"
	"github.com/olivere/elastic/v7"
)

//AddTextMessage will add a text message to the DB
func AddTextMessage(ctx context.Context, client *elastic.Client, index string, s string) error {
	_, err := client.
		Index().
		BodyString(s).
		Do(ctx)

	if err != nil {
		return err
	}
	return err
}

//AddFeelrMessage will add a text message to the DB
func AddFeelrMessage(ctx context.Context, client *elastic.Client, index string, s string) error {
	_, err := client.
		Index().
		BodyString(s).
		Do(ctx)

	if err != nil {
		return err
	}
	return err
}

//GetFeelrMessage will fetch a feelr associated with a chat from the DB
func GetFeelrMessage(ctx context.Context, client *elastic.Client, index string, chatID string, feelrID string) (string, error) {
	chatQuery := elastic.NewMatchQuery("chatId", chatID)
	feelrQuery := elastic.NewMatchQuery("feelrId", feelrID)

	query := elastic.NewBoolQuery().Must(chatQuery, feelrQuery)
	searchResult, err := client.Search().
		Index(index).
		Query(query).
		Do(ctx)

	if err != nil {
		return "", err
	}

	if searchResult.Hits.TotalHits.Value > 0 {
		return searchResult.Hits.Hits[0].Id, nil
	}

	return "", err
}

//SendMessageResponse will add the partners response of a feelr to the DB
func SendMessageResponse(ctx context.Context, client *elastic.Client, index string, docID string, answer string) (*model.Message, error) {
	var m model.Message

	_, err := client.Update().Index(index).Id(docID).Doc(map[string]interface{}{"receiverAnswer": answer}).Do(ctx)
	if err != nil {
		return nil, err
	}

	res, err := client.Get().Index(index).Id(docID).Do(ctx)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Source, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
