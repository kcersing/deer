package service

import (
	"context"
	contents "gen/kitex_gen/contents"
)

type CreateArticleService struct {
	ctx context.Context
}

// NewCreateArticleService new CreateArticleService
func NewCreateArticleService(ctx context.Context) *CreateArticleService {
	return &CreateArticleService{ctx: ctx}
}

// Run create note info
func (s *CreateArticleService) Run(req *contents.CreateArticleReq) (resp *contents.ArticleResp, err error) {
	// Finish your business logic.

	return
}
