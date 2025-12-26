package service

import (
	"context"
	base "gen/kitex_gen/base"
	"product/biz/dal/db"
	"product/biz/dal/db/ent/productitem"
)

type DeleteItemService struct {
	ctx context.Context
} // NewDeleteItemService new DeleteItemService
func NewDeleteItemService(ctx context.Context) *DeleteItemService {
	return &DeleteItemService{ctx: ctx}
}

// Run create note info
func (s *DeleteItemService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	_, err = db.Client.ProductItem.Delete().Where(productitem.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return

}
