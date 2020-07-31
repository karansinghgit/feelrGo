package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karansinghgit/feelrGo/cmd/utils"
	db "github.com/karansinghgit/feelrGo/db"
	"github.com/karansinghgit/feelrGo/graph/generated"
	"github.com/karansinghgit/feelrGo/graph/model"
	"github.com/olivere/elastic/v7"
)

var client = db.GetNewClient()

var index = "app"

//Works
func (r *mutationResolver) CreateFeelr(ctx context.Context, question string, topic string) (*model.Feelr, error) {
	f := &model.Feelr{
		ID:        uuid.New().String(),
		Question:  question,
		Topic:     topic,
		Timestamp: time.Now(),
	}

	s, err := utils.ParseToString(f)

	if err != nil {
		fmt.Println("fa")
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

	fmt.Println("Insertion Successful")
	return f, nil
}

//Works
func (r *mutationResolver) SendTextMessage(ctx context.Context, chatID string, sender string, text string) (*model.Message, error) {
	m := &model.Message{
		Chat:      chatID,
		Sender:    sender,
		Text:      &text,
		Timestamp: time.Now(),
	}

	s, err := utils.ParseToString(m)

	if err != nil {
		fmt.Println("fa")
		return nil, err
	}

	_, err = client.Index().
		Index(index).
		BodyString(s).
		Do(ctx)

	if err != nil {
		fmt.Println("Error Storing the Feelr", err)
		return nil, err
	}

	fmt.Println("Insertion Successful")
	return m, nil
}

//Works
func (r *mutationResolver) SendFeelrMessage(ctx context.Context, chatID string, feelrID string, sender string, answer string) (*model.Message, error) {
	chatQuery := elastic.NewMatchQuery("chat", chatID)
	feelrQuery := elastic.NewMatchQuery("feelr", feelrID)

	query := elastic.NewBoolQuery().Must(chatQuery, feelrQuery)
	searchResult, err := client.Search().
		Index(index).
		Query(query).
		Do(ctx)

	if err != nil {
		return nil, err
	}
	var m *model.Message

	if searchResult.Hits.TotalHits.Value > 0 {
		res, err := client.Update().Index("feelr").Id(searchResult.Hits.Hits[0].Id).Doc(map[string]interface{}{"receiverAnswer": answer}).Do(ctx)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(res.GetResult.Source, &m)
		if err != nil {
			fmt.Println("Error initializing : ", err)
			return nil, err
		}
	} else {
		m = &model.Message{
			Chat:         chatID,
			Sender:       sender,
			Feelr:        &feelrID,
			SenderAnswer: &answer,
			Timestamp:    time.Now(),
		}
		dataJSON, err := json.Marshal(m)
		js := string(dataJSON)
		_, err = client.Index().
			Index(index).
			BodyJson(js).
			Do(ctx)

		if err != nil {
			return nil, err
		}
		fmt.Println("[Elastic]Insertion Successful")
	}
	return m, nil
}

//Works
func (r *queryResolver) GetTopFeelrs(ctx context.Context, top *int) ([]*model.Feelr, error) {
	existsQuery := elastic.NewExistsQuery("question")
	fmt.Println(existsQuery)
	searchResult, err := client.Search().
		Index(index).
		Query(existsQuery).
		Sort("timestamp", false).
		Size(*top).
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

//Doesnt Work. Something to do with NULL Values
func (r *queryResolver) GetMessages(ctx context.Context, chatID string, last *int) ([]*model.Message, error) {
	chatQuery := elastic.NewMatchQuery("chat", chatID)
	searchResult, err := client.Search().
		Index("feelr").
		Query(chatQuery).
		Sort("timestamp", false).
		Size(*last).
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

//Works
func (r *queryResolver) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	userQuery := elastic.NewMatchQuery("id", userID)
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

func (r *subscriptionResolver) MessageAdded(ctx context.Context, chatID string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
