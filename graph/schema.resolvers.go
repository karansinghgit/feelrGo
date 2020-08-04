package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/karansinghgit/feelrGo/db"
	"github.com/karansinghgit/feelrGo/doa"
	"github.com/karansinghgit/feelrGo/graph/generated"
	"github.com/karansinghgit/feelrGo/graph/model"
)

func (r *mutationResolver) CreateFeelr(ctx context.Context, question string, topic string) (*model.Feelr, error) {
	f, err := doa.AddFeelr(ctx, client, index, question, topic)

	if err != nil {
		fmt.Println("Could not add Feelr to db")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return f, nil
}

func (r *mutationResolver) SendTextMessage(ctx context.Context, chatID string, senderID string, text string) (*model.Message, error) {
	m, err := doa.AddTextMessage(ctx, client, index, chatID, senderID, text)

	if err != nil {
		fmt.Println("Could not add Text Message to db")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return m, nil
}

func (r *mutationResolver) SendFeelrMessage(ctx context.Context, chatID string, feelrID string, senderID string, answer string) (*model.Message, error) {
	docID, _ := doa.GetFeelrMessage(ctx, client, index, chatID, feelrID)

	var m *model.Message
	var err error

	if docID != "" {
		m, err = doa.SendMessageResponse(ctx, client, index, docID, answer)
	} else {
		m, err = doa.AddFeelrMessage(ctx, client, index, chatID, senderID, feelrID, answer)
	}

	if err != nil {
		fmt.Println("Could not add Feelr Message Answer to db")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return m, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, senderID string, receiverID string) (*model.Chat, error) {
	c, err := doa.AddChat(ctx, client, index, senderID, receiverID)

	if err != nil {
		fmt.Println("Could not add Couple to db")
		return nil, err
	}

	fmt.Println("Insertion Successful")
	return c, nil
}

func (r *queryResolver) GetTopFeelrs(ctx context.Context, top *int) ([]*model.Feelr, error) {
	feelrs, err := doa.GetFeelrs(ctx, client, index, *top)

	if err != nil {
		fmt.Println("Could not fetch feelrs from db")
		return nil, err
	}

	fmt.Println("[ELASTIC] Fetch Successful")
	return feelrs, nil
}

func (r *queryResolver) GetMessages(ctx context.Context, chatID string, last *int) ([]*model.Message, error) {
	messages, err := doa.GetChatMessages(ctx, client, index, chatID, *last)

	if err != nil {
		fmt.Println("Could not fetch messages from db")
		return nil, err
	}

	fmt.Println("[ELASTIC] Fetch Successful")
	return messages, nil
}

func (r *queryResolver) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	user, err := doa.GetUser(ctx, client, index, userID)

	if err != nil {
		fmt.Println("Could not fetch messages from db")
		return nil, err
	}

	fmt.Println("[ELASTIC] Fetch Successful")
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var client = db.GetNewClient()
var index = "app"
