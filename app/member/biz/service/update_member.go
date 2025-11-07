package service

import (
	"common/pkg/errno"
	"context"
	Member "gen/kitex_gen/member"
	"member/biz/convert"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
	"time"
)

type UpdateMemberService struct {
	ctx context.Context
} // NewUpdateMemberService new UpdateMemberService
func NewUpdateMemberService(ctx context.Context) *UpdateMemberService {
	return &UpdateMemberService{ctx: ctx}
}

// Run Update note info
func (s *UpdateMemberService) Run(req *Member.UpdateMemberReq) (resp *Member.MemberResp, err error) {

	ok, _ := db.Client.Member.Query().Where(member.MobileEQ(req.GetMobile()), member.IDNEQ(req.GetId())).Exist(s.ctx)
	if ok {
		return nil, errno.MemberMobileExistErr
	}

	birthday, err := time.Parse(time.DateOnly, req.GetBirthday())
	if err != nil {
		return nil, errno.TimeFormatErr
	}
	_, err = db.Client.Member.UpdateOneID(req.GetId())).
		SetAvatar(req.GetAvatar()).
		SetMobile(req.GetMobile()).
		SetName(req.GetName()).
		SetStatus(req.GetStatus()).
		SetGender(req.GetGender()).
		SetBirthday(birthday).
		SetDetail(req.GetDetail()).
		//SetRoleID(req.GetRoleId()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &Member.MemberResp{
		Data: convert.EntToMember(only),
	}
	return
}
