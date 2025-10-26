package service

import (
	"context"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/convert"
	"system/biz/dal/db"
	"time"
)

type VerifyRoleAuthService struct {
	ctx context.Context
} // NewVerifyRoleAuthService new VerifyRoleAuthService
func NewVerifyRoleAuthService(ctx context.Context) *VerifyRoleAuthService {
	return &VerifyRoleAuthService{ctx: ctx}
}

// Run create note info
func (s *VerifyRoleAuthService) Run(req *system.VerifyRoleAuthReq) (resp *system.RoleResp, err error) {
	// Finish your business logic.

	save, err := db.Client.Role.UpdateOneID(req.GetId()).
		SetName(req.GetName()).
		SetValue(req.GetValue()).
		SetDefaultRouter(req.GetDefaultRouter()).
		//SetStatus(req.Status).
		SetRemark(req.GetRemark()).
		//SetOrderNo(req.OrderNo).
		SetUpdatedAt(time.Now()).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role failed")
		return nil, err
	}
	resp.Data = convert.EntToRole(save)
	return
}
