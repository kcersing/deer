package service

import (
	"context"
	User "gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent"
	"user/biz/dal/db/ent/predicate"
	"user/biz/dal/db/ent/user"
)

type GetUserListService struct {
	ctx context.Context
} // NewGetUserListService new GetUserListService
func NewGetUserListService(ctx context.Context) *GetUserListService {
	return &GetUserListService{ctx: ctx}
}

// Run Update
func (s *GetUserListService) Run(req *User.GetUserListReq) (resp *User.UserListResp, err error) {

	var (
		userResp []*User.User
	)

	var predicates []predicate.User
	if req.Mobile != "" {
		predicates = append(predicates, user.MobileEQ(req.Mobile))
	}

	if req.Name != "" {
		predicates = append(predicates, user.Name(req.Name))
	}

	users, err := db.Client.User.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(user.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	for _, v := range users {

		userResp = append(userResp, convert.EntToUser(v))
	}
	if err != nil {
		return nil, err
	}
	return &User.UserListResp{
		Data: userResp,
	}, nil
}
