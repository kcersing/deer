package service

import (
	base "gen/kitex_gen/base"

	"context"
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

	return
}
