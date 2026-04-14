package service

import (
	"context"
	contents "gen/kitex_gen/contents"
)

type UpdateArticleService struct {
	ctx context.Context
}

// NewUpdateArticleService new UpdateArticleService
func NewUpdateArticleService(ctx context.Context) *UpdateArticleService {
	return &UpdateArticleService{ctx: ctx}
}

// Run create note info
func (s *UpdateArticleService) Run(req *contents.UpdateArticleReq) (resp *contents.ArticleResp, err error) {
	// Finish your business logic.

	return
}
