package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent"
	"user/biz/dal/db/ent/position"
	"user/biz/dal/db/ent/predicate"
)

type GetPositionsListService struct {
	ctx context.Context
} // NewGetPositionsListService new GetPositionsListService
func NewGetPositionsListService(ctx context.Context) *GetPositionsListService {
	return &GetPositionsListService{ctx: ctx}
}

// Run create note info
func (s *GetPositionsListService) Run(req *user.GetPositionsListReq) (resp *user.PositionsListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*Base.Positions
	)

	var predicates []predicate.Position
	if req.GetKeyword() != "" {
		predicates = append(predicates, position.NameContains(req.GetKeyword()))
	}

	all, err := db.Client.Position.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(position.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)

	if err != nil {
		return nil, err
	}

	dataResp = convert.FindPositionChildren(all, 1)

	return &user.PositionsListResp{
		Data: dataResp,
	}, nil
}
