package service

import (
	"contents/biz/dal/db"
	"contents/biz/dal/db/ent/article"
	"context"
	"gen/kitex_gen/base"
)

type DeleteArticleService struct {
	ctx context.Context
}

// NewDeleteArticleService new DeleteArticleService
func NewDeleteArticleService(ctx context.Context) *DeleteArticleService {
	return &DeleteArticleService{ctx: ctx}
}

// Run create note info
func (s *DeleteArticleService) Run(req *base.IdReq) (resp *base.BaseResp, err error) {
	// Finish your business logic.
	_, err = db.Client.Article.Delete().Where(article.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
