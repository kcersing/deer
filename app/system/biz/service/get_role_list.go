package service

import (
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"

	"github.com/cloudwego/kitex/pkg/klog"
)

type GetRoleListService struct {
	ctx context.Context
} // NewGetRoleListService new GetRoleListService
func NewGetRoleListService(ctx context.Context) *GetRoleListService {
	return &GetRoleListService{ctx: ctx}
}

// Run create note info
func (s *GetRoleListService) Run(req *system.GetRoleListReq) (resp *system.RoleListResp, err error) {
	// Finish your business logic.
	var (
		predicates []predicate.Role
		dataResp   []*base.Role
	)
	all, err := db.Client.Role.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return resp, err
	}
	for _, v := range all {
		klog.Info("role: %v", v)
		dataResp = append(dataResp, convert.EntToRole(v))
	}
	resp = &system.RoleListResp{
		Data: dataResp,
	}
	return
}
