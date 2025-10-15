package service

import (
	"context"
	Base "gen/kitex_gen/base"
	Member "gen/kitex_gen/member"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
	"member/biz/dal/db"
	"member/biz/dal/db/ent/member"
)

type GetMemberService struct {
	ctx context.Context
} // NewGetMemberService new GetMemberService
func NewGetMemberService(ctx context.Context) *GetMemberService {
	return &GetMemberService{ctx: ctx}
}

// Run Get note info
func (s *GetMemberService) Run(req *Base.IdReq) (resp *Member.MemberResp, err error) {

	var (
		memberResp *Member.Member
	)

	eg, ctx := errgroup.WithContext(s.ctx)
	eg.Go(func() error {

		only, err := db.Client.Member.Query().Where(member.IDEQ(req.GetId())).Only(ctx)
		if err != nil {
			klog.CtxErrorf(ctx, "call details error: %s", err.Error())
			return err
		}
		memberResp = &Member.Member{
			Id:   &only.ID,
			Name: &only.Name,
		}
		return nil
	})
	//eg.Go(func() error {
	//	res, err := h.detailsClient.GetProduct(ctx, &details.GetProductReq{ID: productID})
	//	if err != nil {
	//		klog.CtxErrorf(ctx, "call details error: %s", err.Error())
	//		return err
	//	}
	//
	//	detailsResp = res
	//	return nil
	//})
	if err := eg.Wait(); err != nil {
		return
	}

	resp = &Member.MemberResp{
		Data: memberResp,
	}

	return
}
