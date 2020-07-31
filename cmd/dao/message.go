package dao

import (
	"context"

	"github.com/karansinghgit/feelrGo/cmd/graph/model"
	"github.com/karansinghgit/feelrGo/cmd/utils"
	"github.com/olivere/elastic"
)

func AddMessage(ctx context.Context, c *elastic.Client, m *model.Message) error {
	s, err := utils.ParseToString(m)

	if err != nil {
		return err
	}
	_, err = c.Index().
		Index("feelr").
		BodyJson(s).
		Do(ctx)
	return nil
}
