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

//CreateUser creates a new User
func CreateUser(ctx context.Context, client *elastic.Client, index string, username string, userToken *string, partnerID *string) (*model.User, error) {
	var relState string
	if *partnerID == "" {
		relState = "available"
	} else {
		relState = "engaged"
	}

	u := &model.User{
		UserID:    uuid.New().String(),
		Username:  username,
		UserToken: userToken,
		PartnerID: partnerID,
		RelState:  relState,
	}

	s, err := utils.ParseToString(u)

	if err != nil {
		return nil, err
	}

	err = dao.AddUser(ctx, client, index, s)

	if err != nil {
		fmt.Println("[ELASTIC] Error Storing the User")
		return nil, err
	}

	fmt.Println("[ELASTIC] Insertion Successful")
	return u, nil
}

//GetUserInfo gets the info of a user
func GetUserInfo(ctx context.Context, client *elastic.Client, index string, userID string) (*model.User, error) {
	u, err := dao.GetUser(ctx, client, index, userID)

	if err != nil {
		fmt.Println("[ELASTIC] Error fetching the User Info")
		return nil, err
	}

	fmt.Println("[ELASTIC] Fetch Successful")
	return u, nil
}

//CheckUsername checks for a valid username
func CheckUsername(ctx context.Context, client *elastic.Client, index string, userID string) (string, error) {
	isPresent, err := dao.CheckUsername(ctx, client, index, userID)

	if err != nil {
		fmt.Println("[ELASTIC] Error fetching the User Info")
		return "false", err
	}

	fmt.Println("[ELASTIC] Fetch Successful")
	return isPresent, nil
}
