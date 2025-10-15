package service

import (
	"common/pkg/encrypt"
	"context"
	Base "gen/kitex_gen/base"
	Member "gen/kitex_gen/member"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
)

type ChangePasswordService struct {
	ctx context.Context
} // NewChangePasswordService new ChangePasswordService
func NewChangePasswordService(ctx context.Context) *ChangePasswordService {
	return &ChangePasswordService{ctx: ctx}
}

// Run create note info
func (s *ChangePasswordService) Run(req *Member.ChangePasswordReq) (resp *Base.NilResponse, err error) {
	password, _ := encrypt.Crypt(req.GetPassword())
	_, err = db.Client.Member.Update().Where(member.IDEQ(req.GetId())).SetPassword(password).Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
