package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/karansinghgit/feelrGo/db"
	"github.com/karansinghgit/feelrGo/graphql/generated"
	"github.com/karansinghgit/feelrGo/graphql/model"
	"github.com/karansinghgit/feelrGo/services"
)

func (r *mutationResolver) CreateFeelr(ctx context.Context, question string, topic string, createdBy string) (*model.Feelr, error) {
	f, err := services.CreateFeelr(ctx, r.client, r.index, question, topic, createdBy)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *mutationResolver) SendTextMessage(ctx context.Context, chatID string, senderID string, text string) (*model.Message, error) {
	m, err := services.SendTextMessage(ctx, r.client, r.index, chatID, senderID, text)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *mutationResolver) SendFeelrMessage(ctx context.Context, chatID string, senderID string, feelrID string, answer string) (*model.Message, error) {
	m, err := services.SendFeelrMessage(ctx, r.client, r.index, chatID, senderID, feelrID, answer)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, senderID string, receiverID string) (*model.Chat, error) {
	c, err := services.CreateChat(ctx, r.client, r.index, senderID, receiverID)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, username string, userToken *string, partnerID *string) (*model.User, error) {
	u, err := services.CreateUser(ctx, r.client, r.index, username, userToken, partnerID)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *queryResolver) GetTopFeelrs(ctx context.Context, top *int) ([]*model.Feelr, error) {
	f, err := services.GetTopFeelrs(ctx, r.client, r.index, top)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *queryResolver) GetMessages(ctx context.Context, chatID string, last *int) ([]*model.Message, error) {
	m, err := services.GetMessages(ctx, r.client, r.index, chatID, last)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *queryResolver) GetUserInfo(ctx context.Context, userID string) (*model.User, error) {
	u, err := services.GetUserInfo(ctx, r.client, r.index, userID)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *queryResolver) CheckUsername(ctx context.Context, username string) (string, error) {
	u, err := services.CheckUsername(ctx, r.client, r.index, username)

	if err != nil {
		return "", err
	}

	return u, nil
}

func (r *queryResolver) CheckPartnerState(ctx context.Context, userID string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, chatID string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	r.client = db.GetNewClient()
	r.index = "app"
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	r.client = db.GetNewClient()
	r.index = "app"
	return &queryResolver{r}
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
