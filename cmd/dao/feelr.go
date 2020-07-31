package dao

import (
	"context"

	"github.com/karansinghgit/feelrGo/cmd/graph/model"
	"github.com/karansinghgit/feelrGo/cmd/utils"
	"github.com/olivere/elastic"
)

func AddFeelr(ctx context.Context, c *elastic.Client, f *model.Feelr) error {
	s, err := utils.ParseToString(f)

	if err != nil {
		return err
	}
	_, err = c.Index().
		Index("feelr").
		BodyJson(s).
		Do(ctx)
	return nil
}
