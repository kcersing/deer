package service

import (
	"context"
	base "gen/kitex_gen/base"
)

type OfflineProductService struct {
	ctx context.Context
} // NewOfflineProductService new OfflineProductService
func NewOfflineProductService(ctx context.Context) *OfflineProductService {
	return &OfflineProductService{ctx: ctx}
}

// Run create note info
func (s *OfflineProductService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
