package service

import (
	"common/pkg/errno"
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent"
)

type GetPositionsService struct {
	ctx context.Context
} // NewGetPositionsService new GetPositionsService
func NewGetPositionsService(ctx context.Context) *GetPositionsService {
	return &GetPositionsService{ctx: ctx}
}

// Run create note info
func (s *GetPositionsService) Run(req *base.IdReq) (resp *user.PositionsResp, err error) {
	// Finish your business logic.

	if req == nil {
		return nil, errno.InvalidParameterErr
	}

	entity, err := db.Client.Position.Get(s.ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errno.NotFound
		}
		return nil, errno.QueryFailed
	}
	dataResp := convert.EntToPosition(entity)
	resp = &user.PositionsResp{
		Data: dataResp,
	}
	return
}
