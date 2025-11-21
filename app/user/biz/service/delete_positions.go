package service

import (
	"context"
	"gen/kitex_gen/base"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/position"
)

type DeletePositionsService struct {
	ctx context.Context
} // NewDeletePositionsService new DeletePositionsService
func NewDeletePositionsService(ctx context.Context) *DeletePositionsService {
	return &DeletePositionsService{ctx: ctx}
}

// Run create note info
func (s *DeletePositionsService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	_, err = db.Client.Position.Delete().Where(position.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
