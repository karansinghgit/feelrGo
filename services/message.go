package services

import (
	"context"
	"fmt"
	"time"

	"github.com/karansinghgit/feelrGo/dao"
	"github.com/karansinghgit/feelrGo/graphql/model"

	"github.com/karansinghgit/feelrGo/utils"
	"github.com/olivere/elastic/v7"
)

//SendTextMessage sends a new Text Message
func SendTextMessage(ctx context.Context, client *elastic.Client, index string, chatID string, senderID string, text string) (*model.Message, error) {
	m := &model.Message{
		ChatID:    chatID,
		SenderID:  senderID,
		Timestamp: time.Now(),
		Text:      &text,
	}

	s, err := utils.ParseToString(m)

	if err != nil {
		return nil, err
	}

	err = dao.AddTextMessage(ctx, client, index, s)

	if err != nil {
		fmt.Println("[ELASTIC] Error Storing the Message")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return m, nil
}

//SendFeelrMessage sends a new Feelr Message
func SendFeelrMessage(ctx context.Context, client *elastic.Client, index string, chatID string, senderID string, feelrID string, answer string) (*model.Message, error) {
	docID, _ := dao.GetFeelrMessage(ctx, client, index, chatID, feelrID)

	var m *model.Message
	var err error

	if docID != "" {
		m, err = dao.SendMessageResponse(ctx, client, index, docID, answer)
	} else {
		m := &model.Message{
			ChatID:    chatID,
			SenderID:  senderID,
			FeelrID:   &feelrID,
			Timestamp: time.Now(),
			Text:      &answer,
		}

		s, err := utils.ParseToString(m)

		if err != nil {
			return nil, err
		}

		err = dao.AddFeelrMessage(ctx, client, index, s)
	}

	if err != nil {
		fmt.Println("[ELASTIC] Error Storing the Message")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return m, nil
}
