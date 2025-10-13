package main

import (
	"context"
	base "gen/kitex_gen/base"
	user "gen/kitex_gen/user"
	base "kitex_gen/base"
	user "kitex_gen/user"
	"user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.UserResp, err error) {
	resp, err = service.NewCreateUserService(ctx).Run(req)

	return resp, err
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *base.IdReq) (resp *user.UserResp, err error) {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		res, err := h.reviewsClient.ReviewProduct(ctx, &reviews.ReviewReq{ProductID: productID})
		if err != nil {
			klog.CtxErrorf(ctx, "call reviews error: %s", err.Error())
			return err
		}
		reviewsResp = res
		return nil
	})
	eg.Go(func() error {
		res, err := h.detailsClient.GetProduct(ctx, &details.GetProductReq{ID: productID})
		if err != nil {
			klog.CtxErrorf(ctx, "call details error: %s", err.Error())
			return err
		}

		detailsResp = res
		return nil
	})
	if err := eg.Wait(); err != nil {
		c.JSON(http.StatusInternalServerError, &base.BaseResp{
			StatusMessage: "internal error",
			StatusCode:    http.StatusInternalServerError,
			Extra:         nil,
		})
		return
	}
	p := detailsResp.GetProduct()
	resp := &product.Product{
		ID:          productID,
		Title:       p.GetTitle(),
		Author:      p.GetAuthor(),
		Description: p.GetDescription(),
		Rating:      reviewsResp.GetReview().GetRating(),
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserList(ctx context.Context, req *user.GetUserListReq) (resp *user.UserListResp, err error) {
	// TODO: Your code here...
	return
}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *base.CheckAccountReq) (resp *user.UserResp, err error) {
	// TODO: Your code here...
	return
}
