package service

import (
	"common/pkg/errno"
	"context"
	Member "gen/kitex_gen/member"

	"github.com/pkg/errors"
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
		return nil, errno.UserMobileExistErr
	}

	birthday, err := time.Parse(time.DateOnly, req.GetBirthday())
	if err != nil {
		return nil, errno.TimeFormatErr
	}
	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "starting a transaction:")
	}
	save, err := tx.Member.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetStatus(req.Status).
		SetAvatar(req.GetAvatar()).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "update user failed"))
	}
	profile, err := save.QueryMemberProfile().First(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "query profile nil"))
	}
	profile, err = tx.MemberProfile.UpdateOneID(req.GetId()).
		SetCreatedID(req.CreatedId).
		SetBirthday(birthday).
		SetIntention(req.GetIntention()).
		SetGender(req.GetGender()).
		SetIntention(req.GetIntention()).
		//SetSource(req.GetSource()).
		//SetEmail(req.GetEmail()).
		//SetWecom(req.GetWecom()).
		//SetRelationMid(req.GetRelationMid()).
		//SetRelationMame(req.GetRelationMame()).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "update Member Profile failed"))
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	resp = &Member.MemberResp{
		Data: convert.EntToMember(save, profile),
	}
	return
}
