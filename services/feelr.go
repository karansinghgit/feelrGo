package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karansinghgit/feelrGo/dao"
	"github.com/karansinghgit/feelrGo/graphql/model"

	"github.com/karansinghgit/feelrGo/utils"
	"github.com/olivere/elastic/v7"
)

//CreateFeelr creates a new Feelr
func CreateFeelr(ctx context.Context, client *elastic.Client, index string, question string, topic string, createdBy string) (*model.Feelr, error) {
	f := &model.Feelr{
		FeelrID:   uuid.New().String(),
		Question:  question,
		Topic:     topic,
		Timestamp: time.Now(),
		CreatedBy: createdBy,
	}

	s, err := utils.ParseToString(f)

	if err != nil {
		return nil, err
	}

	err = dao.AddFeelr(ctx, client, index, s)

	if err != nil {
		fmt.Println("[ELASTIC] Error Storing the Feelr")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return f, nil
}

//GetTopFeelrs fetches specified number of feelrs
func GetTopFeelrs(ctx context.Context, client *elastic.Client, index string, top *int) ([]*model.Feelr, error) {
	feelrs, err := dao.GetFeelrs(ctx, client, index, *top)

	if err != nil {
		return nil, err
	}

	return feelrs, nil
}
