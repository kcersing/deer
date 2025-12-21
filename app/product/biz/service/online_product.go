package service

import (
	"context"
	base "gen/kitex_gen/base"
	"product/biz/dal/db"
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

	_, err = db.Client.Product.UpdateOneID(req.GetId()).
		SetStatus(1).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
