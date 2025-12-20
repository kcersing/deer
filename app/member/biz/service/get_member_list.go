package service

import (
	"context"
	Base "gen/kitex_gen/base"

	Member "gen/kitex_gen/member"
	"member/biz/dal/db"
	"member/biz/dal/db/ent"

	"member/biz/convert"
	"member/biz/dal/db/ent/member"
	"member/biz/dal/db/ent/predicate"

	"github.com/cloudwego/kitex/pkg/klog"
)

type GetMemberListService struct {
	ctx context.Context
} // NewGetMemberListService new GetMemberListService
func NewGetMemberListService(ctx context.Context) *GetMemberListService {
	return &GetMemberListService{ctx: ctx}
}

// Run Update
func (s *GetMemberListService) Run(req *Member.GetMemberListReq) (resp *Member.MemberListResp, err error) {
	klog.Info("GetMemberListService.Run req: %v", req)
	var (
		dataResp []*Base.Member
	)
	var predicates []predicate.Member
	if req.GetKeyword() != "" {
		predicates = append(predicates, member.Or(
			member.NameContains(req.GetKeyword()),
		))
	}

	all, err := db.Client.Member.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(member.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range all {
		profile, err := v.QueryMemberProfile().First(s.ctx)
		if err != nil {
			return nil, err
		}
		dataResp = append(dataResp, convert.EntToMember(v, profile))
	}
	return &Member.MemberListResp{
		Data: dataResp,
	}, nil
}
