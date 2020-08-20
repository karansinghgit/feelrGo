package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/karansinghgit/feelrGo/dao"
	"github.com/karansinghgit/feelrGo/graphql/model"

	"github.com/karansinghgit/feelrGo/utils"
	"github.com/olivere/elastic/v7"
)

//CreateChat creates a new Chat
func CreateChat(ctx context.Context, client *elastic.Client, index string, senderID string, receiverID string) (*model.Chat, error) {
	c := &model.Chat{
		ChatID:     uuid.New().String(),
		SenderID:   senderID,
		ReceiverID: receiverID,
	}

	s, err := utils.ParseToString(c)

	if err != nil {
		return nil, err
	}

	err = dao.AddChat(ctx, client, index, s)

	if err != nil {
		fmt.Println("[ELASTIC] Error Storing the Chat")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return c, nil
}

//GetMessages gets all the messages related to a Chat
func GetMessages(ctx context.Context, client *elastic.Client, index string, chatID string, last *int) ([]*model.Message, error) {
	m, err := dao.GetChatMessages(ctx, client, index, chatID, *last)

	if err != nil {
		fmt.Println("[ELASTIC] Error fetching the User Info")
		return nil, err
	}

	fmt.Println("[ELASTIC] Fetch Successful")
	return m, nil
}
