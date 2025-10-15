package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"user/biz/dal/db"
	"user/biz/dal/db/ent/user"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *Base.IdReq) (resp *Base.NilResponse, err error) {
	_, err = db.Client.User.Delete().Where(user.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
