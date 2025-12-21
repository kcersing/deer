package service

import (
	"context"
	base "gen/kitex_gen/base"
	"product/biz/dal/db"
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

	_, err = db.Client.Product.UpdateOneID(req.GetId()).
		SetStatus(0).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
