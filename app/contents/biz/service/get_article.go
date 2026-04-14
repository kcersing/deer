package service

import (
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

	return
}
