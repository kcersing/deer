package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type OnlineProductService struct {
	ctx context.Context
} // NewOnlineProductService new OnlineProductService
func NewOnlineProductService(ctx context.Context) *OnlineProductService {
	return &OnlineProductService{ctx: ctx}
}

// Run create note info
func (s *OnlineProductService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
