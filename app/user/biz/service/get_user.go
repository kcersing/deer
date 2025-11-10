package service

import (
	"context"
	Base "gen/kitex_gen/base"
	User "gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run Get note info
func (s *GetUserService) Run(req *Base.IdReq) (resp *User.UserResp, err error) {

	var (
		dataResp *Base.User
	)

	eg, ctx := errgroup.WithContext(s.ctx)
	eg.Go(func() error {

		only, err := db.Client.User.Query().Where(user.IDEQ(req.GetId())).Only(ctx)
		if err != nil {
			klog.CtxErrorf(ctx, "call details error: %s", err.Error())
			return err
		}
		dataResp = convert.EntToUser(only)
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
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &User.UserResp{
		Data: dataResp,
	}

	return
}
