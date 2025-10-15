package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
)

type DeleteMemberService struct {
	ctx context.Context
} // NewDeleteMemberService new DeleteMemberService
func NewDeleteMemberService(ctx context.Context) *DeleteMemberService {
	return &DeleteMemberService{ctx: ctx}
}

// Run create note info
func (s *DeleteMemberService) Run(req *Base.IdReq) (resp *Base.NilResponse, err error) {
	_, err = db.Client.Member.Delete().Where(member.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
