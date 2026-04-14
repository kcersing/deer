package service

import (
	"common/pkg/errno"
	"contents/biz/convert"
	"contents/biz/dal/db"
	"contents/biz/dal/db/ent"
	"context"
	base "gen/kitex_gen/base"
	contents "gen/kitex_gen/contents"
)

type GetArticleService struct {
	ctx context.Context
}

// NewGetArticleService new GetArticleService
func NewGetArticleService(ctx context.Context) *GetArticleService {
	return &GetArticleService{ctx: ctx}
}

// Run create note info
func (s *GetArticleService) Run(req *base.IdReq) (resp *contents.ArticleResp, err error) {
	// Finish your business logic.

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Article.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToArticle(entity)
	resp = &contents.ArticleResp{
		Data: dataResp,
	}
	return
}
